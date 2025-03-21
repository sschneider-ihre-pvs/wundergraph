# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

## [0.132.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.132.0...@wundergraph/sdk@0.132.1) (2023-01-23)

### Bug Fixes

* mark json-schema-to-typescript as dependency instead of devDependency ([#584](https://github.com/wundergraph/wundergraph/issues/584)) ([1c890d1](https://github.com/wundergraph/wundergraph/commit/1c890d11e9ea40fe322eef272818af578016c9f5)) (@fiam)
* use operationPath in client operation metadata ([#583](https://github.com/wundergraph/wundergraph/issues/583)) ([fb2c322](https://github.com/wundergraph/wundergraph/commit/fb2c3224a240a3b73e453cfe6874ef5a066f4d4a)) (@spetrunin)

## [0.132.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.131.2...@wundergraph/sdk@0.132.0) (2023-01-20)

### Features

* add upload profiles for S3 providers ([#476](https://github.com/wundergraph/wundergraph/issues/476)) ([6144545](https://github.com/wundergraph/wundergraph/commit/614454539133c7f235aea6aa72ade36059f41c97)) (@fiam)
* migrate to .graphqlrc ([#537](https://github.com/wundergraph/wundergraph/issues/537)) ([f99a046](https://github.com/wundergraph/wundergraph/commit/f99a046c6cccb00e9c57224304be0c3dccc9c909)) (@fiam)

## [0.131.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.131.1...@wundergraph/sdk@0.131.2) (2023-01-20)

### Bug Fixes

* operations loading and normalization ([#568](https://github.com/wundergraph/wundergraph/issues/568)) ([f34fc98](https://github.com/wundergraph/wundergraph/commit/f34fc98227e9a988499fd388e3883a07128dcdd7)) (@spetrunin)

## [0.131.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.131.0...@wundergraph/sdk@0.131.1) (2023-01-19)

### Bug Fixes

* handle consecutive hyphens or underscores when formatting field … ([#566](https://github.com/wundergraph/wundergraph/issues/566)) ([810b31c](https://github.com/wundergraph/wundergraph/commit/810b31cfbc135a93b65ffaeed20f64d003646548)) (@Aenimus)
* operation metadata type imports ([#501](https://github.com/wundergraph/wundergraph/issues/501)) ([1dcb746](https://github.com/wundergraph/wundergraph/commit/1dcb746110a9dc518d02fe82184bedf67e2a5d51)) (@rwest202)
* path formatting for resolveFieldName ([#565](https://github.com/wundergraph/wundergraph/issues/565)) ([7416b61](https://github.com/wundergraph/wundergraph/commit/7416b61e99037aef15ce6e2bb3812a36c2ca3bed)) (@Aenimus)

## [0.131.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.130.2...@wundergraph/sdk@0.131.0) (2023-01-17)

### Features

* add support for disconnecting from Auth0 ([#525](https://github.com/wundergraph/wundergraph/issues/525)) ([7465fb2](https://github.com/wundergraph/wundergraph/commit/7465fb21a3618924c7dfb59a6a2f94c7d740f0f8)) (@fiam)
* make input required if there are required variables ([#551](https://github.com/wundergraph/wundergraph/issues/551)) ([00256bf](https://github.com/wundergraph/wundergraph/commit/00256bf456f0c733beb45a5cdbc258f84631975e)) (@Pagebakers)
* upgrade to nextjs 13 ([#504](https://github.com/wundergraph/wundergraph/issues/504)) ([45bc431](https://github.com/wundergraph/wundergraph/commit/45bc431243cc61765c2712b03e89818a1bb3d14a)) (@Pagebakers)

### Bug Fixes

* allow to treat a subgraph as regular graphql api ([#496](https://github.com/wundergraph/wundergraph/issues/496)) ([cac56f1](https://github.com/wundergraph/wundergraph/commit/cac56f1ffc7d16701abd1e921820d558121e94f1)) (@spetrunin)
* json scalar types rendering not supported for graphql api ([#516](https://github.com/wundergraph/wundergraph/issues/516)) ([76ad844](https://github.com/wundergraph/wundergraph/commit/76ad84425ad0851c7217e87ac40b2ba89c90fc88)) (@spetrunin)
* prevent swallowing of input validation errors ([#529](https://github.com/wundergraph/wundergraph/issues/529)) ([39ea3fc](https://github.com/wundergraph/wundergraph/commit/39ea3fc3fbf96916f1165228194b8c915882b133)) (@Aenimus)

## [0.130.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.130.1...@wundergraph/sdk@0.130.2) (2022-12-31)

**Note:** Version bump only for package @wundergraph/sdk

## [0.130.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.130.0...@wundergraph/sdk@0.130.1) (2022-12-31)

**Note:** Version bump only for package @wundergraph/sdk

## [0.130.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.129.0...@wundergraph/sdk@0.130.0) (2022-12-29)

### ⚠ BREAKING CHANGES

* **node:** restructure server imports (#497)

### Features

* **node:** restructure server imports ([#497](https://github.com/wundergraph/wundergraph/issues/497)) ([ac277de](https://github.com/wundergraph/wundergraph/commit/ac277dec5c06bb761d6acb026248dedd3d1f59c0)) (@StarpTech)

### Bug Fixes

* use 0.0.0.0 for test server ([#499](https://github.com/wundergraph/wundergraph/issues/499)) ([126778b](https://github.com/wundergraph/wundergraph/commit/126778b88d3e546be1870e4bfb7c8e9d429f3fe1)) (@spetrunin)

## [0.129.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.128.0...@wundergraph/sdk@0.129.0) (2022-12-28)

### Features

* add [@remove](https://github.com/remove)NullVariables directives ([#477](https://github.com/wundergraph/wundergraph/issues/477)) ([0f4398b](https://github.com/wundergraph/wundergraph/commit/0f4398b1509b5939e0c4b5824ae2b64c0646e101)) (@spetrunin)

## [0.128.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.127.0...@wundergraph/sdk@0.128.0) (2022-12-21)

### Features

* better error messages when resolving undefined variables ([#419](https://github.com/wundergraph/wundergraph/issues/419)) ([aa0dd65](https://github.com/wundergraph/wundergraph/commit/aa0dd65cb6fc837cf4de962e915bb1c541d18418)), closes [#262](https://github.com/wundergraph/wundergraph/issues/262) (@fiam)

## [0.127.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.126.0...@wundergraph/sdk@0.127.0) (2022-12-20)

### Features

* **ci:** add CI workflow for examples ([#428](https://github.com/wundergraph/wundergraph/issues/428)) ([72c1616](https://github.com/wundergraph/wundergraph/commit/72c16160205cd2e61ffcf493c9eb488214ff42cb)) (@fiam)

### Bug Fixes

* csrf and authenticated uploads ([#456](https://github.com/wundergraph/wundergraph/issues/456)) ([6138a98](https://github.com/wundergraph/wundergraph/commit/6138a98c20286e9693ab9df0245006c2d73043ab)) (@thisisnithin)
* detect server ready when the hooks server health check is disabled ([#462](https://github.com/wundergraph/wundergraph/issues/462)) ([e324f1c](https://github.com/wundergraph/wundergraph/commit/e324f1c681d83658ce0300e2e938ffbb73a3c779)) (@fiam)

## [0.126.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.125.0...@wundergraph/sdk@0.126.0) (2022-12-15)

### Features

* listen on both IPv4 and IPv6 by default ([#451](https://github.com/wundergraph/wundergraph/issues/451)) ([b1fd332](https://github.com/wundergraph/wundergraph/commit/b1fd332dcaf8c15209b8329f0c96cb1d4c4972ab)) (@fiam)

## [0.125.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.124.1...@wundergraph/sdk@0.125.0) (2022-12-14)

### Features

* add WunderGraph testing library ([#345](https://github.com/wundergraph/wundergraph/issues/345)) ([13d47b5](https://github.com/wundergraph/wundergraph/commit/13d47b50e7f54dba1c6f7188204c285d8de523c4)) (@fiam)
* implement telemetry ([#424](https://github.com/wundergraph/wundergraph/issues/424)) ([8cd3d29](https://github.com/wundergraph/wundergraph/commit/8cd3d299e39afcd0bce6c1e6c2459e268a09af7b)) (@StarpTech)

### Bug Fixes

* nextjs-swr example ([#434](https://github.com/wundergraph/wundergraph/issues/434)) ([0cae074](https://github.com/wundergraph/wundergraph/commit/0cae0746a0dc36a8a5f8514e4c078363afab3b94)) (@fiam)

## [0.124.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.124.0...@wundergraph/sdk@0.124.1) (2022-12-08)

**Note:** Version bump only for package @wundergraph/sdk

## [0.124.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.123.2...@wundergraph/sdk@0.124.0) (2022-12-06)

### Features

* make detecting invalid operations easier ([#374](https://github.com/wundergraph/wundergraph/issues/374)) ([45639db](https://github.com/wundergraph/wundergraph/commit/45639db0ae3adb8cac4f62b623b04061118f7bf7)) (@fiam)
* react query client ([#340](https://github.com/wundergraph/wundergraph/issues/340)) ([c5769a4](https://github.com/wundergraph/wundergraph/commit/c5769a422970e8eaf626dd767cf09252789bcd1f)) (@Pagebakers)

### Bug Fixes

* generate all required templates ([#360](https://github.com/wundergraph/wundergraph/issues/360)) ([ce2f130](https://github.com/wundergraph/wundergraph/commit/ce2f130829dfc4e4c7e468fcdc364bc7adef00d2)) (@JivusAyrus)
* remove colors from configurePublishWunderGraphAPI() ([#371](https://github.com/wundergraph/wundergraph/issues/371)) ([371ed96](https://github.com/wundergraph/wundergraph/commit/371ed96e759f2bf0daa774c5f26705504b79f11d)) (@fiam)
* replace graphql-weather-api.herokuapp.com with our own copy ([#390](https://github.com/wundergraph/wundergraph/issues/390)) ([1b99525](https://github.com/wundergraph/wundergraph/commit/1b995252af8bf7970cb1e0740b0b17412760de13)) (@fiam)
* stop wunderctl on normalization errors, exit with non-zero ([#370](https://github.com/wundergraph/wundergraph/issues/370)) ([a0b9b9e](https://github.com/wundergraph/wundergraph/commit/a0b9b9ef446aee6031b07e9bbbf2d7ed471fd066)) (@fiam)

## [0.123.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.123.0...@wundergraph/sdk@0.123.2) (2022-11-20)

**Note:** Version bump only for package @wundergraph/sdk

## [0.123.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.123.0...@wundergraph/sdk@0.123.1) (2022-11-20)

**Note:** Version bump only for package @wundergraph/sdk

## [0.123.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.122.0...@wundergraph/sdk@0.123.0) (2022-11-18)

### Features

* remove app/main from URL structure ([#335](https://github.com/wundergraph/wundergraph/issues/335)) ([e49e585](https://github.com/wundergraph/wundergraph/commit/e49e585528297126b93958105e80b68f1943d781)), closes [#333](https://github.com/wundergraph/wundergraph/issues/333) (@fiam)
* sdk, schema extension fields config ([#336](https://github.com/wundergraph/wundergraph/issues/336)) ([ca09b3c](https://github.com/wundergraph/wundergraph/commit/ca09b3cf2002763b7ea53a7d50f6dce50d08c120)) (@YuriBuerov)

## [0.122.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.121.0...@wundergraph/sdk@0.122.0) (2022-11-17)

### Features

* **breaking:** the Next.js client now uses SWR ([#319](https://github.com/wundergraph/wundergraph/issues/319)) ([020cd74](https://github.com/wundergraph/wundergraph/commit/020cd74091517faedcd97071d48a19395cbcd9bf)) (@Pagebakers)
* logging, request id ([#323](https://github.com/wundergraph/wundergraph/issues/323)) ([ed39136](https://github.com/wundergraph/wundergraph/commit/ed3913693b7233ee58ce8423d1ee0ab833c1d161)) (@YuriBuerov)

## [0.121.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.120.0...@wundergraph/sdk@0.121.0) (2022-11-11)

### Features

* allow static string with placeholder OR environment variable for OAS baseURL ([41b7ce6](https://github.com/wundergraph/wundergraph/commit/41b7ce6a02b5159fbdff522a881a61da37221cd0)) (@jensneuse)
* improved SWR hooks api, useMutation result no longer cached and updated to SWR 2.0 ([#305](https://github.com/wundergraph/wundergraph/issues/305)) ([fc4848b](https://github.com/wundergraph/wundergraph/commit/fc4848b3f15b447a487b31bd3d152f134c6f3aeb)) (@Pagebakers)

### Bug Fixes

* errors when wundergraph.server.ts does not exist ([#327](https://github.com/wundergraph/wundergraph/issues/327)) ([5df223d](https://github.com/wundergraph/wundergraph/commit/5df223d9f7428b36e7d7f95632007d966f624e10)) (@fiam)
* make baseURL optional ([#330](https://github.com/wundergraph/wundergraph/issues/330)) ([fcb7133](https://github.com/wundergraph/wundergraph/commit/fcb7133c62fe7ae8057584860857e858e6e2a8d9)) (@Pagebakers)
* **oas:** use primitive schema references as the field schema ([#326](https://github.com/wundergraph/wundergraph/issues/326)) ([3a6d636](https://github.com/wundergraph/wundergraph/commit/3a6d6367a81dd7aed91b84950897258c775790dd)) (@acdn-sglanzer)
* operation metadata in react provider ([#293](https://github.com/wundergraph/wundergraph/issues/293)) ([ad02a27](https://github.com/wundergraph/wundergraph/commit/ad02a27af5dddbcf8e1126e5bf32c949005630b2)) (@rwest202)
* set upload type when s3 is not configured ([#324](https://github.com/wundergraph/wundergraph/issues/324)) ([68b34d7](https://github.com/wundergraph/wundergraph/commit/68b34d7b5436dfe8d88c7e3a27e0975a263e5034)) (@Pagebakers)

## [0.120.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.119.0...@wundergraph/sdk@0.120.0) (2022-11-08)

### Features

* createClient now returns typesafe client ([#307](https://github.com/wundergraph/wundergraph/issues/307)) ([803ebd5](https://github.com/wundergraph/wundergraph/commit/803ebd5d799688586a8a8abdde8cf5d7b2ea9557)) (@Pagebakers)

## [0.119.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.118.0...@wundergraph/sdk@0.119.0) (2022-11-04)

### Features

* eng 640 compose subgraphs without apollo dependencies ([#315](https://github.com/wundergraph/wundergraph/issues/315)) ([628a183](https://github.com/wundergraph/wundergraph/commit/628a18303acb47df5a10118b17a4e88916b2903a)) (@jensneuse)
* sdk, support schema extension ([#302](https://github.com/wundergraph/wundergraph/issues/302)) ([e952af6](https://github.com/wundergraph/wundergraph/commit/e952af61428d0592876362bc19155d45fec626f1)) (@YuriBuerov)

## [0.118.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.117.4...@wundergraph/sdk@0.118.0) (2022-11-03)

### Features

* improve server graceful shutdown ([#314](https://github.com/wundergraph/wundergraph/issues/314)) ([14da07d](https://github.com/wundergraph/wundergraph/commit/14da07d6ca1b6c6cd1571e8322338c1684f92ff8)) (@StarpTech)

## [0.117.4](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.117.3...@wundergraph/sdk@0.117.4) (2022-11-03)

**Note:** Version bump only for package @wundergraph/sdk

## [0.117.3](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.117.2...@wundergraph/sdk@0.117.3) (2022-11-03)

**Note:** Version bump only for package @wundergraph/sdk

## [0.117.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.117.1...@wundergraph/sdk@0.117.2) (2022-11-03)

### Bug Fixes

* skip introspection cache when OpenAPI source is a file ([#309](https://github.com/wundergraph/wundergraph/issues/309)) ([911a551](https://github.com/wundergraph/wundergraph/commit/911a551ced23b5d455d0321f1dc2eeebe979d65d)) (@fiam)

## [0.117.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.117.0...@wundergraph/sdk@0.117.1) (2022-10-23)

**Note:** Version bump only for package @wundergraph/sdk

## [0.117.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.116.1...@wundergraph/sdk@0.117.0) (2022-10-20)

### Features

* healthcheck, node and hooks server ([#274](https://github.com/wundergraph/wundergraph/issues/274)) ([b7f7ecc](https://github.com/wundergraph/wundergraph/commit/b7f7eccef038cf03913740cb43360ca1d38dc016)) (@YuriBuerov)
* make WunderGraphConfigApplicationConfig.cors optional ([#270](https://github.com/wundergraph/wundergraph/issues/270)) ([4bb2658](https://github.com/wundergraph/wundergraph/commit/4bb26586b795b4b7682d4dc43574f9d765d6be86)) (@fiam)
* subscriptions, operation hooks ([#260](https://github.com/wundergraph/wundergraph/issues/260)) ([3199931](https://github.com/wundergraph/wundergraph/commit/31999317c0d95098cefaef7b19b51aa9c0248353)) (@YuriBuerov)

### Bug Fixes

* exit node process when parent exited ([#271](https://github.com/wundergraph/wundergraph/issues/271)) ([fbc8608](https://github.com/wundergraph/wundergraph/commit/fbc8608a96ed5aa40afe17fa2ad658bd650d6257)) (@spetrunin)

## [0.116.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.116.0...@wundergraph/sdk@0.116.1) (2022-10-18)

**Note:** Version bump only for package @wundergraph/sdk

## [0.116.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.115.0...@wundergraph/sdk@0.116.0) (2022-10-18)

### Features

* add configurable per-source timeouts ([#232](https://github.com/wundergraph/wundergraph/issues/232)) ([bb3b6bd](https://github.com/wundergraph/wundergraph/commit/bb3b6bd31250b402fe0c9a099b0dad993976cf39)) (@fiam)
* align logging format hooks server and ([#240](https://github.com/wundergraph/wundergraph/issues/240)) ([e601d4c](https://github.com/wundergraph/wundergraph/commit/e601d4c1597a949c2c564a5b953b4424dae5ee7c)) (@spetrunin)
* subscriptions, ws connection init hook ([#243](https://github.com/wundergraph/wundergraph/issues/243)) ([2e4ae85](https://github.com/wundergraph/wundergraph/commit/2e4ae8506558165a9bf3ada4b8f45cee55a9f18d)) (@YuriBuerov)

### Bug Fixes

* fix mutations are possibly undefined ([#265](https://github.com/wundergraph/wundergraph/issues/265)) ([cea8607](https://github.com/wundergraph/wundergraph/commit/cea860703a1f76a63dae292770514218624ce3a1)) (@JivusAyrus)

## [0.115.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.6...@wundergraph/sdk@0.115.0) (2022-10-12)

### Features

* refactor cache handling, move storage out of generated ([#238](https://github.com/wundergraph/wundergraph/issues/238)) ([6351e35](https://github.com/wundergraph/wundergraph/commit/6351e35419215b4bc63bce4a80e16e20d8e9d2d0)) (@fiam)

### Bug Fixes

* correctly generate typenames for nested array objects ([#257](https://github.com/wundergraph/wundergraph/issues/257)) ([ca61274](https://github.com/wundergraph/wundergraph/commit/ca612747122195a826b9362e24765b66562e06cf)) (@jensneuse)
* live query support in generated web client ([#248](https://github.com/wundergraph/wundergraph/issues/248)) ([a1bc5f5](https://github.com/wundergraph/wundergraph/commit/a1bc5f5fe5182d87759ea895257e1f643e472e4a)), closes [#78](https://github.com/wundergraph/wundergraph/issues/78) (@xzyfer)

## [0.114.6](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.5...@wundergraph/sdk@0.114.6) (2022-10-10)

**Note:** Version bump only for package @wundergraph/sdk

## [0.114.5](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.4...@wundergraph/sdk@0.114.5) (2022-10-07)

### Bug Fixes

* automatically rename subscription object fields for oas ([#246](https://github.com/wundergraph/wundergraph/issues/246)) ([aaf018e](https://github.com/wundergraph/wundergraph/commit/aaf018e038758dcf5f79c9ee42641a9c562d5a70)) (@jensneuse)

## [0.114.4](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.3...@wundergraph/sdk@0.114.4) (2022-10-07)

### Bug Fixes

* **sdk/client:** include extra headers from constructor ([#244](https://github.com/wundergraph/wundergraph/issues/244)) ([d59e3ab](https://github.com/wundergraph/wundergraph/commit/d59e3ab1d90c82a2971253e6afdc793d6f6e3f46)) (@chronotc)
* **sdk:** add fastify route only if the hook is configured ([#204](https://github.com/wundergraph/wundergraph/issues/204)) ([4d744a3](https://github.com/wundergraph/wundergraph/commit/4d744a3c3923b9c0db926ea393cbba821b4b8b74)) (@JivusAyrus)

## [0.114.3](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.2...@wundergraph/sdk@0.114.3) (2022-10-04)

### Bug Fixes

* introspection, infinite loop in file changes ([#230](https://github.com/wundergraph/wundergraph/issues/230)) ([4fdf635](https://github.com/wundergraph/wundergraph/commit/4fdf6352bddd63125be673f55808f4d8493299b8)) (@YuriBuerov)

## [0.114.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.1...@wundergraph/sdk@0.114.2) (2022-09-30)

### Bug Fixes

* re-enable to specify wundergraph dir via arg ([#226](https://github.com/wundergraph/wundergraph/issues/226)) ([50cb5f2](https://github.com/wundergraph/wundergraph/commit/50cb5f22468fa481089caeba9935ee65e9dfe1f3)) (@jensneuse)

## [0.114.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.114.0...@wundergraph/sdk@0.114.1) (2022-09-29)

**Note:** Version bump only for package @wundergraph/sdk

## [0.114.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.113.2...@wundergraph/sdk@0.114.0) (2022-09-29)

### Features

*  add graphql introspection middleware ([#214](https://github.com/wundergraph/wundergraph/issues/214)) ([33e840c](https://github.com/wundergraph/wundergraph/commit/33e840c995c5d482e6e755d19bbd76b006d19f3c)) (@spetrunin)
* sse subscriptions configuration ([#217](https://github.com/wundergraph/wundergraph/issues/217)) ([7bf72ef](https://github.com/wundergraph/wundergraph/commit/7bf72efd16a8bac422db32fe957e102395d7357c)) (@YuriBuerov)

### Bug Fixes

* allow to configure internal and public node urls separately ([#207](https://github.com/wundergraph/wundergraph/issues/207)) ([c01bd9a](https://github.com/wundergraph/wundergraph/commit/c01bd9a1cedefbb5fd0ecde83f3b96b3dfee6f41)) (@spetrunin)
* synchronize prisma install for multiple prisma datasources ([#222](https://github.com/wundergraph/wundergraph/issues/222)) ([9ca1fc9](https://github.com/wundergraph/wundergraph/commit/9ca1fc9d3f75ce381bafd4f895c8fb803547a932)) (@jensneuse)

## [0.113.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.113.1...@wundergraph/sdk@0.113.2) (2022-09-26)

### Bug Fixes

* **cache:** dont print dev messages in prod mode ([#208](https://github.com/wundergraph/wundergraph/issues/208)) ([b22c32a](https://github.com/wundergraph/wundergraph/commit/b22c32ab7f2705b19ad537f76dfef43d589fd026)) (@StarpTech)
* **codegen:** resolve template dependencies recursively ([#209](https://github.com/wundergraph/wundergraph/issues/209)) ([283a9c8](https://github.com/wundergraph/wundergraph/commit/283a9c8db4c28f8fb1177fd68aa6a55acc98e8a0)) (@StarpTech)

## [0.113.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.113.0...@wundergraph/sdk@0.113.1) (2022-09-22)

**Note:** Version bump only for package @wundergraph/sdk

## [0.113.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.112.0...@wundergraph/sdk@0.113.0) (2022-09-21)

### Features

* add auth example for swr ([#200](https://github.com/wundergraph/wundergraph/issues/200)) ([7a5f34e](https://github.com/wundergraph/wundergraph/commit/7a5f34e8c80141f400a9b6ed195d04a22a86c92b)) (@StarpTech)

## [0.112.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.111.0...@wundergraph/sdk@0.112.0) (2022-09-21)

### Features

* add post logout hook ([#196](https://github.com/wundergraph/wundergraph/issues/196)) ([6870130](https://github.com/wundergraph/wundergraph/commit/6870130b0c4e6fc269d81160994384c1d1cf6e59)) (@jensneuse)

## [0.111.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.110.0...@wundergraph/sdk@0.111.0) (2022-09-20)

### Features

* add swr hooks for auth and file upload ([#195](https://github.com/wundergraph/wundergraph/issues/195)) ([a0b1bae](https://github.com/wundergraph/wundergraph/commit/a0b1bae6a18e30fcefc35bdde2f72326f3de1392)) (@StarpTech)

## [0.110.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.109.0...@wundergraph/sdk@0.110.0) (2022-09-19)

### Features

* make typescript client typesafe ([#179](https://github.com/wundergraph/wundergraph/issues/179)) ([942b278](https://github.com/wundergraph/wundergraph/commit/942b2782255de3b9e6374500c7a8047de074e4ff)), closes [#188](https://github.com/wundergraph/wundergraph/issues/188) (@Pagebakers)

## [0.109.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.108.0...@wundergraph/sdk@0.109.0) (2022-09-19)

### Features

* implement config first approach ([#151](https://github.com/wundergraph/wundergraph/issues/151)) ([803da0e](https://github.com/wundergraph/wundergraph/commit/803da0e51beb3a7b23ee826dfde835eccfa1c2dd)), closes [#190](https://github.com/wundergraph/wundergraph/issues/190) (@spetrunin)

### Bug Fixes

* update User interface on client to parse json correctly ([#176](https://github.com/wundergraph/wundergraph/issues/176)) ([f80e410](https://github.com/wundergraph/wundergraph/commit/f80e410530433c6c9b1c290abca2e51b1e7ea907)) (@thisisnithin)

## [0.108.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.107.1...@wundergraph/sdk@0.108.0) (2022-09-15)

### Features

* allow templates in baseURL for openapi ([#183](https://github.com/wundergraph/wundergraph/issues/183)) ([cb279e7](https://github.com/wundergraph/wundergraph/commit/cb279e7829e704d75a1bcea5a87c42b09331c624)) (@jensneuse)
* custom stable hash function for js values ([#182](https://github.com/wundergraph/wundergraph/issues/182)) ([c16c635](https://github.com/wundergraph/wundergraph/commit/c16c635dd6c6ee5a9784e83b4063941b8e3f3435)) (@StarpTech)

## [0.107.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.107.0...@wundergraph/sdk@0.107.1) (2022-09-15)

**Note:** Version bump only for package @wundergraph/sdk

## [0.107.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.106.0...@wundergraph/sdk@0.107.0) (2022-09-14)

### Features

* oidc,  allow passing additional query parameters to the IDP ([#178](https://github.com/wundergraph/wundergraph/issues/178)) ([d015bb1](https://github.com/wundergraph/wundergraph/commit/d015bb150762cba7a46865e66f3de633e731de07)) (@YuriBuerov)

## [0.106.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.105.1...@wundergraph/sdk@0.106.0) (2022-09-09)

### Features

* add golang client ([#170](https://github.com/wundergraph/wundergraph/issues/170)) ([a26bc32](https://github.com/wundergraph/wundergraph/commit/a26bc32d4a178134f893012ca4e2648460b4e7f8)) (@jensneuse)

## [0.105.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.105.0...@wundergraph/sdk@0.105.1) (2022-09-08)

**Note:** Version bump only for package @wundergraph/sdk

## [0.105.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.104.0...@wundergraph/sdk@0.105.0) (2022-09-08)

### Features

* improve webhook types ([#173](https://github.com/wundergraph/wundergraph/issues/173)) ([8cb6d97](https://github.com/wundergraph/wundergraph/commit/8cb6d9750f6e14e2fdd3d87ad97ca0cf3236f95b)) (@StarpTech)

## [0.104.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.103.1...@wundergraph/sdk@0.104.0) (2022-09-06)

### Features

* ensure that user is always set in auth hooks, disable introspection cache in `wunderctl generate` ([#167](https://github.com/wundergraph/wundergraph/issues/167)) ([4b40572](https://github.com/wundergraph/wundergraph/commit/4b40572dd993be1c84e421f1796eb8a2913ecf69)) (@StarpTech)

### Bug Fixes

* sync wunderctl ([06dfe11](https://github.com/wundergraph/wundergraph/commit/06dfe11e885acafe48b6d7e8776cb763f0c75a66)) (@StarpTech)

## [0.103.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.103.0...@wundergraph/sdk@0.103.1) (2022-09-05)

### Bug Fixes

* fixes open api errors for int based enums and json fields ([#164](https://github.com/wundergraph/wundergraph/issues/164)) ([a79fe3e](https://github.com/wundergraph/wundergraph/commit/a79fe3ebc8c0b8d863123d565edfe0942f03048d)) (@rpeterson)

## [0.103.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.102.0...@wundergraph/sdk@0.103.0) (2022-09-04)

### Features

* don't exit, throw error and handle at root ([#161](https://github.com/wundergraph/wundergraph/issues/161)) ([5495d27](https://github.com/wundergraph/wundergraph/commit/5495d27c181f12a96655fae0f403ffaedda50816)) (@StarpTech)

### Bug Fixes

* subscription url config ([#162](https://github.com/wundergraph/wundergraph/issues/162)) ([c503400](https://github.com/wundergraph/wundergraph/commit/c503400061a33243702c8e7be753e14d863e5d98)) (@jensneuse)

## [0.102.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.101.0...@wundergraph/sdk@0.102.0) (2022-09-02)

### Features

* refactor introspection runner ([#158](https://github.com/wundergraph/wundergraph/issues/158)) ([49d0ab3](https://github.com/wundergraph/wundergraph/commit/49d0ab3be58552e017d5120feb7f09d0f48b4aae)) (@StarpTech)

## [0.101.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.100.0...@wundergraph/sdk@0.101.0) (2022-09-01)

### Features

* implement userId for fromClaim directive ([#152](https://github.com/wundergraph/wundergraph/issues/152)) ([51df6e5](https://github.com/wundergraph/wundergraph/commit/51df6e50244bee9f5f8d579ff6f604e1a1c853d9)) (@jensneuse)

## [0.100.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.99.0...@wundergraph/sdk@0.100.0) (2022-08-30)

### Features

* **server:** introduce pino base logger ([#146](https://github.com/wundergraph/wundergraph/issues/146)) ([d261b8f](https://github.com/wundergraph/wundergraph/commit/d261b8fe5d8fa6e21058468c2e70b45defa0601a)) (@StarpTech)

### Bug Fixes

* internal directive breaks code generation ([#148](https://github.com/wundergraph/wundergraph/issues/148)) ([a9cb48c](https://github.com/wundergraph/wundergraph/commit/a9cb48cbfd840862cd38f17b9c185407acad7772)) (@jensneuse)
* **types:** make webhooks optional ([#149](https://github.com/wundergraph/wundergraph/issues/149)) ([fa0d243](https://github.com/wundergraph/wundergraph/commit/fa0d243e3bd0bfbf62448d1348709375df404cac)) (@StarpTech)

## [0.99.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.98.3...@wundergraph/sdk@0.99.0) (2022-08-29)

### Features

* use web std header implementation ([#145](https://github.com/wundergraph/wundergraph/issues/145)) ([7c0359b](https://github.com/wundergraph/wundergraph/commit/7c0359bdc3efac0a8c11ceb23cd49172a991fbd3))(@StarpTech)

## [0.98.3](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.98.2...@wundergraph/sdk@0.98.3) (2022-08-25)

**Note:** Version bump only for package @wundergraph/sdk

## [0.98.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.98.1...@wundergraph/sdk@0.98.2) (2022-08-25)

**Note:** Version bump only for package @wundergraph/sdk

## [0.98.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.98.0...@wundergraph/sdk@0.98.1) (2022-08-18)

**Note:** Version bump only for package @wundergraph/sdk

## [0.98.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.97.0...@wundergraph/sdk@0.98.0) (2022-08-18)

### Features

* native webhooks support ([#126](https://github.com/wundergraph/wundergraph/issues/126)) ([a0b38bd](https://github.com/wundergraph/wundergraph/commit/a0b38bd54b88198db6cc176432d577dab0931245))

### Bug Fixes

* issue with unhandled hyphens in input names ([#123](https://github.com/wundergraph/wundergraph/issues/123)) ([b01caaa](https://github.com/wundergraph/wundergraph/commit/b01caaa8c4036afbeb579dbbf14a52d82971b116))

## [0.97.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.96.1...@wundergraph/sdk@0.97.0) (2022-08-09)

### Features

* add fragment support to sdk ([5839e35](https://github.com/wundergraph/wundergraph/commit/5839e35ad4ab00f9174e8e18a54375580dd1c6a0))
* extract typescript client from nextjs ([#72](https://github.com/wundergraph/wundergraph/issues/72)) ([282797d](https://github.com/wundergraph/wundergraph/commit/282797dd4d28dce922cca8a3d5092d68c508f5bd))
* replace the legacy client with the new implementation ([#78](https://github.com/wundergraph/wundergraph/issues/78)) ([e2468c8](https://github.com/wundergraph/wundergraph/commit/e2468c8856e02a7d1d89dc1c08c1731871bc19f3))

### Bug Fixes

* update tsdoc for hooks config ([8f5d916](https://github.com/wundergraph/wundergraph/commit/8f5d9161383981e5abae2be5c66587cf2b5fb547))

## [0.96.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.96.0...@wundergraph/sdk@0.96.1) (2022-07-18)

**Note:** Version bump only for package @wundergraph/sdk

## [0.96.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.95.0...@wundergraph/sdk@0.96.0) (2022-07-13)

### Features

- use headersobject for transport hooks ([#75](https://github.com/wundergraph/wundergraph/issues/75)) ([82059cf](https://github.com/wundergraph/wundergraph/commit/82059cfb87292b3baadc8d618732314a532b5ed6))

### Bug Fixes

- **auth:** pass raw access token to hook ([#76](https://github.com/wundergraph/wundergraph/issues/76)) ([c31644d](https://github.com/wundergraph/wundergraph/commit/c31644ddcb29dcb74063ef20d80d7ef71aa3c88f))

## [0.95.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.94.4...@wundergraph/sdk@0.95.0) (2022-07-07)

### Features

- add introspection caching & DataSource polling ([#63](https://github.com/wundergraph/wundergraph/issues/63)) ([ec6226e](https://github.com/wundergraph/wundergraph/commit/ec6226e19f930d53e0a621c9a831d2ac5deea913))

### Bug Fixes

- restart hooks server, ensure \_\_wg exists ([#68](https://github.com/wundergraph/wundergraph/issues/68)) ([55435df](https://github.com/wundergraph/wundergraph/commit/55435dfcf9d03187385266bc6d6a3cc9c6606edf))

## [0.94.4](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.94.3...@wundergraph/sdk@0.94.4) (2022-07-05)

### Bug Fixes

- **codegen:** detect internal input correctly ([#64](https://github.com/wundergraph/wundergraph/issues/64)) ([7c36904](https://github.com/wundergraph/wundergraph/commit/7c36904e2d5d5a5a8c36b9c31a6f98844aa34081))

## [0.94.3](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.94.2...@wundergraph/sdk@0.94.3) (2022-06-30)

**Note:** Version bump only for package @wundergraph/sdk

## [0.94.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.94.1...@wundergraph/sdk@0.94.2) (2022-06-30)

### Bug Fixes

- call mutation hooks ([#58](https://github.com/wundergraph/wundergraph/issues/58)) ([8ff5f75](https://github.com/wundergraph/wundergraph/commit/8ff5f75ee50483b150a0f1b7512f9e2a2cbcba2d))

## [0.94.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.94.0...@wundergraph/sdk@0.94.1) (2022-06-30)

**Note:** Version bump only for package @wundergraph/sdk

## [0.94.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.93.2...@wundergraph/sdk@0.94.0) (2022-06-29)

### Features

- improve error message when graphql introspection fails ([#54](https://github.com/wundergraph/wundergraph/issues/54)) ([b774e73](https://github.com/wundergraph/wundergraph/commit/b774e7341bff0da2343e959854d58deab8dbf580))

### Bug Fixes

- set correct client request, remove inflights checks in client ([06df8dc](https://github.com/wundergraph/wundergraph/commit/06df8dc779702dc257545d000f0d60eb4d99a3da))

## [0.93.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.93.1...@wundergraph/sdk@0.93.2) (2022-06-29)

### Bug Fixes

- openapi introspection defect ([#53](https://github.com/wundergraph/wundergraph/issues/53)) ([9da07df](https://github.com/wundergraph/wundergraph/commit/9da07df6b84301ade07bbecd741aa643e06a11d4))

## [0.93.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.93.0...@wundergraph/sdk@0.93.1) (2022-06-23)

**Note:** Version bump only for package @wundergraph/sdk

## [0.93.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.6...@wundergraph/sdk@0.93.0) (2022-06-20)

### ⚠ BREAKING CHANGES

- **hooks:** refactor hooks interface (#40)

### Code Refactoring

- **hooks:** refactor hooks interface ([#40](https://github.com/wundergraph/wundergraph/issues/40)) ([9e58149](https://github.com/wundergraph/wundergraph/commit/9e581498899f3251cd41d6e33c784c4960979c51))

## [0.92.6](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.5...@wundergraph/sdk@0.92.6) (2022-06-12)

### Bug Fixes

- **hooks:** pass response correctly, pass input arg when available ([#38](https://github.com/wundergraph/wundergraph/issues/38)) ([5e4fe75](https://github.com/wundergraph/wundergraph/commit/5e4fe755a3c46446eaefbb3b5c8e581d55d608d8))

## [0.92.5](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.4...@wundergraph/sdk@0.92.5) (2022-06-11)

**Note:** Version bump only for package @wundergraph/sdk

## [0.92.4](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.3...@wundergraph/sdk@0.92.4) (2022-06-11)

**Note:** Version bump only for package @wundergraph/sdk

## [0.92.3](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.2...@wundergraph/sdk@0.92.3) (2022-06-11)

**Note:** Version bump only for package @wundergraph/sdk

## [0.92.2](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.1...@wundergraph/sdk@0.92.2) (2022-06-11)

**Note:** Version bump only for package @wundergraph/sdk

## [0.92.1](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.92.0...@wundergraph/sdk@0.92.1) (2022-06-10)

**Note:** Version bump only for package @wundergraph/sdk

## [0.92.0](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.91.5...@wundergraph/sdk@0.92.0) (2022-06-10)

### Features

- **server:** reimplement bundling, watcher, script runner ([#32](https://github.com/wundergraph/wundergraph/issues/32)) ([594af1d](https://github.com/wundergraph/wundergraph/commit/594af1d3b53c1e9b12dd21bd79a4cc8a51784c3a))

### Bug Fixes

- add test for schema merge conflict, improve error message ([#27](https://github.com/wundergraph/wundergraph/issues/27)) ([7f41a65](https://github.com/wundergraph/wundergraph/commit/7f41a651eb0975c92fb2b8fbe345fe7062c35824))

## [0.91.5](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.91.4...@wundergraph/sdk@0.91.5) (2022-06-04)

**Note:** Version bump only for package @wundergraph/sdk

## [0.91.4](https://github.com/wundergraph/wundergraph/compare/@wundergraph/sdk@0.91.3...@wundergraph/sdk@0.91.4) (2022-06-02)

**Note:** Version bump only for package @wundergraph/sdk

## 0.91.3

### Patch Changes

- [`e507ffd`](https://github.com/wundergraph/wundergraph/commit/e507ffd05916d089957b44084d8f3c5387320ef3) Thanks [@StarpTech](https://github.com/StarpTech)! - sync wunderctl version

## 0.91.2

### Patch Changes

- [`b5dbe92`](https://github.com/wundergraph/wundergraph/commit/b5dbe92e6c9d3160bfba3713c43594790cb2effd) Thanks [@StarpTech](https://github.com/StarpTech)! - sync wunderctl version

## 0.91.1

### Patch Changes

- [`d122589`](https://github.com/wundergraph/wundergraph/commit/d122589b501dfa2f6565630de4005e1bc83cc729) Thanks [@StarpTech](https://github.com/StarpTech)! - update wunderctl

## 0.91.0

### Minor Changes

- [#18](https://github.com/wundergraph/wundergraph/pull/18) [`afea237`](https://github.com/wundergraph/wundergraph/commit/afea23771191e049aab5ce56ce775775389e8770) Thanks [@StarpTech](https://github.com/StarpTech)! - move wunderctl / go binary into local node_modules

### Patch Changes

- [#20](https://github.com/wundergraph/wundergraph/pull/20) [`93cf9e1`](https://github.com/wundergraph/wundergraph/commit/93cf9e1cd3b2a30ad79d434f13f84596dd0b3607) Thanks [@jensneuse](https://github.com/jensneuse)! - fix typescript codegen for list of enum

- Updated dependencies [[`afea237`](https://github.com/wundergraph/wundergraph/commit/afea23771191e049aab5ce56ce775775389e8770)]:
  - @wundergraph/protobuf@0.91.0

## 0.90.6

### Patch Changes

- [`ad2e3a6`](https://github.com/wundergraph/wundergraph/commit/ad2e3a6fa97441b49ab477e47463a9ce2d561049) Thanks [@jensneuse](https://github.com/jensneuse)! - fix openapi transformation when array contained ref

## 0.90.5

### Patch Changes

- [`e68d101`](https://github.com/wundergraph/wundergraph/commit/e68d101f5af39918fc810c68ec0cde606009d40c) Thanks [@jensneuse](https://github.com/jensneuse)! - fix oas path resolving

## 0.90.4

### Patch Changes

- [`a53fc56`](https://github.com/wundergraph/wundergraph/commit/a53fc56a054d4b4dc68de53a8d178e3d5341ef58) Thanks [@jensneuse](https://github.com/jensneuse)! - trigger ci

## 0.90.3

### Patch Changes

- [`a27e9f5`](https://github.com/wundergraph/wundergraph/commit/a27e9f50093f7b4775f4d6a1d2f03a56398bc169) Thanks [@jensneuse](https://github.com/jensneuse)! - trigger ci

## 0.90.2

### Patch Changes

- [`7c6b3ae`](https://github.com/wundergraph/wundergraph/commit/7c6b3ae3f86bbe7ee3402556668ce81052f192f4) Thanks [@jensneuse](https://github.com/jensneuse)! - fix database introspection

## 0.90.1

### Patch Changes

- Updated dependencies [[`0857db3`](https://github.com/wundergraph/wundergraph/commit/0857db3d55209fb878fe6326629b125c6f2d2315)]:
  - @wundergraph/protobuf@0.90.1
