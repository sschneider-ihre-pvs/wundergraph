import {
	authProviders,
	configureWunderGraphApplication,
	cors,
	EnvironmentVariable,
	introspect,
	templates,
} from '@wundergraph/sdk';
import { NextJsTemplate } from '@wundergraph/nextjs/dist/template';
import server from './wundergraph.server';
import operations from './wundergraph.operations';

const federatedApi = introspect.federation({
	apiNamespace: 'federated',
	upstreams: [
		{
			url: 'http://localhost:4001/graphql',
		},
		{
			url: 'http://localhost:4002/graphql',
		},
		{
			url: 'http://localhost:4003/graphql',
		},
		{
			url: 'http://localhost:4004/graphql',
		},
	],
});

// configureWunderGraph emits the configuration
configureWunderGraphApplication({
	apis: [federatedApi],
	server,
	operations,
	codeGenerators: [
		{
			templates: [...templates.typescript.all],
		},
		{
			templates: [new NextJsTemplate()],
			path: '../components/generated',
		},
	],
	cors: {
		...cors.allowAll,
		allowedOrigins: process.env.NODE_ENV === 'production' ? ['http://localhost:3000'] : ['http://localhost:3000'],
	},
	authentication: {
		cookieBased: {
			providers: [authProviders.demo()],
			authorizedRedirectUris: ['http://localhost:3000'],
			secureCookieHashKey: new EnvironmentVariable('WUNDERGRAPH_SECURE_COOKIE_HASH_KEY'), // must be of length 32
			secureCookieBlockKey: new EnvironmentVariable('WUNDERGRAPH_SECURE_COOKIE_BLOCK_KEY'), // must be of length 32
			csrfTokenSecret: new EnvironmentVariable('WUNDERGRAPH_CSRF_TOKEN_SECRET'), // must be of length 11
		},
	},
	security: {
		enableGraphQLEndpoint: process.env.NODE_ENV !== 'production',
	},
});
