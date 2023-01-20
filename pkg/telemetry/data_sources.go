package telemetry

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"

	"github.com/wundergraph/wundergraph/pkg/loadvariable"
	"github.com/wundergraph/wundergraph/pkg/wgpb"
)

const (
	urlHashTag = "urlHash"
)

func isPrivateHost(host string) bool {
	privateIPBlocks := []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	}
	addrs, _ := net.LookupHost(host)
	for _, addr := range addrs {
		ip := net.ParseIP(addr)
		if ip == nil {
			continue
		}
		if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
			return true
		}
		for _, blockAddr := range privateIPBlocks {
			_, block, err := net.ParseCIDR(blockAddr)
			if err == nil && block.Contains(ip) {
				return true
			}
		}
	}
	return false
}

func dataSourceMetric(dataSourceTag string, urlVariable *wgpb.ConfigurationVariable) (*Metric, error) {
	metric := NewDataSourceMetric(dataSourceTag)
	if urlVariable != nil {
		urlValue := loadvariable.String(urlVariable)
		if urlValue != "" {
			urlHash, err := Hash(urlValue)
			if err != nil {
				return nil, err
			}
			if err := metric.AddTag(urlHashTag, urlHash); err != nil {
				return nil, err
			}
			if u, err := url.Parse(urlValue); err == nil {
				// Try to parse as host:port, fallback to host
				host, _, _ := net.SplitHostPort(u.Host)
				if host == "" {
					host = u.Host
				}
				if isPrivateHost(host) {
					if err := metric.AddTag("isPrivate", "true"); err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return metric, nil
}

func DataSourceMetrics(wunderGraphDir string) ([]*Metric, error) {
	configFile := filepath.Join(wunderGraphDir, "generated", "wundergraph.config.json")
	f, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var wgConfig wgpb.WunderGraphConfiguration
	if err := dec.Decode(&wgConfig); err != nil {
		return nil, err
	}
	var metrics []*Metric
	for _, ds := range wgConfig.GetApi().GetEngineConfiguration().GetDatasourceConfigurations() {
		var dataSourceTag string
		var urlVariable *wgpb.ConfigurationVariable
		switch ds.Kind {
		case wgpb.DataSourceKind_STATIC:
			// Not reported
		case wgpb.DataSourceKind_REST:
			dataSourceTag = ds.Kind.String()
			urlVariable = ds.GetCustomRest().GetFetch().GetUrl()
		case wgpb.DataSourceKind_GRAPHQL:
			dataSourceTag = ds.Kind.String()
			urlVariable = ds.GetCustomGraphql().GetFetch().GetUrl()
		case wgpb.DataSourceKind_POSTGRESQL,
			wgpb.DataSourceKind_MYSQL,
			wgpb.DataSourceKind_SQLSERVER,
			wgpb.DataSourceKind_MONGODB,
			wgpb.DataSourceKind_SQLITE:

			dataSourceTag = ds.Kind.String()
			urlVariable = ds.CustomDatabase.GetDatabaseURL()
		default:
			if err != nil {
				return nil, fmt.Errorf("unhandled data source kind %v", ds.Kind)
			}
		}
		if dataSourceTag == "" {
			continue
		}
		metric, err := dataSourceMetric(dataSourceTag, urlVariable)
		if err != nil {
			return nil, err
		}
		// Same underlying data source can be included multiple times, deduplicate
		found := false
		for _, m := range metrics {
			if m.Equal(metric) {
				found = true
				break
			}
		}
		if !found {
			metrics = append(metrics, metric)
		}
	}

	return metrics, nil
}
