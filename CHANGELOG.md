## [0.0.10](https://github.com/pascaliske/magicmirror/compare/v0.0.9...v0.0.10) (2022-07-05)


### Bug Fixes

* **dockerfile:** pin alpine image to latest version ([4df57df](https://github.com/pascaliske/magicmirror/commit/4df57df6598a98d9c3555b9b1234bdb7cdc94ad3))
* **server:** defer metric update after socket client disconnects ([4e29b3e](https://github.com/pascaliske/magicmirror/commit/4e29b3ee88a42daa56dab1621a66db93dc1a671f))
* **server:** handle client write message errors ([5c0c8a4](https://github.com/pascaliske/magicmirror/commit/5c0c8a425e47233fb82178b2c912f1db6ba3281c))


### Features

* initial docs setup ([2a84b98](https://github.com/pascaliske/magicmirror/commit/2a84b98ac0633405aa14f66b9bb978715535fffb))
* **news-card:** implement link behind news headlines ([e165e1d](https://github.com/pascaliske/magicmirror/commit/e165e1d1616fd1dc4aea0b4d04905fd5b99c53f5))



## [0.0.9](https://github.com/pascaliske/magicmirror/compare/v0.0.8...v0.0.9) (2022-04-30)


### Bug Fixes

* **server:** use logger for socket errors ([f9d0002](https://github.com/pascaliske/magicmirror/commit/f9d00022af69ed1085b020894385554cecfd7207))
* **socket:** replace deprecated rxjs operator ([762c95d](https://github.com/pascaliske/magicmirror/commit/762c95dd946976bf454e1462ba4acec700ca8c2d))


### Features

* **server:** add config file reload metrics ([007134b](https://github.com/pascaliske/magicmirror/commit/007134b10b790ee4e70ddc75c3e39c151efc24a1))
* **server:** add socket client metric ([2782829](https://github.com/pascaliske/magicmirror/commit/27828291fe6969b4197fe4b9e74de7888da0d188))
* **server:** add uptime metric ([b9e83b0](https://github.com/pascaliske/magicmirror/commit/b9e83b03ed6f17bac5521c4647f6517b0778375b))



## [0.0.8](https://github.com/pascaliske/magicmirror/compare/v0.0.7...v0.0.8) (2022-04-27)


### Features

* implement build time variable ([8245d92](https://github.com/pascaliske/magicmirror/commit/8245d9276910e07c0efa8eea3fe03176ca0ba37c))
* **server:** implement build info metric ([4deb76c](https://github.com/pascaliske/magicmirror/commit/4deb76c9d54a522fb9d4cd441526f04b9eae2a15))



## [0.0.7](https://github.com/pascaliske/magicmirror/compare/v0.0.6...v0.0.7) (2022-03-21)


### Bug Fixes

* **weather-card:** ensure width of temperature works with baseline alignment in safari ([3875795](https://github.com/pascaliske/magicmirror/commit/38757952337b69fedf200d7207f9251b4b984c9a))


### Features

* **socket:** retry socket connection every 5 seconds after server close event ([ee974b2](https://github.com/pascaliske/magicmirror/commit/ee974b24258bbea4712a229ebf62ee8d6cd554e6))



## [0.0.6](https://github.com/pascaliske/magicmirror/compare/v0.0.5...v0.0.6) (2022-02-25)


### Bug Fixes

* properly align top cards ([7e7b86a](https://github.com/pascaliske/magicmirror/commit/7e7b86ae046d15da19252026e21bbf6b328dfa31))
* **server:** remove unused feed struct ([49f18a5](https://github.com/pascaliske/magicmirror/commit/49f18a527d24068d0ddb0e332f05e4f7d49b700d))


### Features

* **ansible:** add tailscale provisioning role ([f92c636](https://github.com/pascaliske/magicmirror/commit/f92c6367c6b0175c126ed52ac8f429ab6a3bdaeb))
* re-write sass using new module system ([c0a6b1c](https://github.com/pascaliske/magicmirror/commit/c0a6b1c03e4097ae0294b1b6c7979fd29179e375))
* **server:** exit  programm correctly on errors ([48f7ce7](https://github.com/pascaliske/magicmirror/commit/48f7ce71db4c2e3174aafc0c7dff1057fb933ae6))



## [0.0.5](https://github.com/pascaliske/magicmirror/compare/v0.0.4...v0.0.5) (2022-02-09)



## [0.0.4](https://github.com/pascaliske/magicmirror/compare/v0.0.3...v0.0.4) (2022-02-09)


### Features

* **cards:** automatically refresh weather and news cards after 10 minutes ([c7a5c84](https://github.com/pascaliske/magicmirror/commit/c7a5c84bbd6819406402dd603e7f2cdf49036b82))
* **server:** add log line for health endpoint ([56ed06b](https://github.com/pascaliske/magicmirror/commit/56ed06b803328b17cee9dead33e9e598302e6b11))
* **server:** add log line for socket endpoint ([3881d59](https://github.com/pascaliske/magicmirror/commit/3881d590877b426816b3fb310b218a6176d368a4))



## [0.0.3](https://github.com/pascaliske/magicmirror/compare/v0.0.2...v0.0.3) (2022-02-09)


### Bug Fixes

* **icon:** use newer icon type for font awesome ([a96e313](https://github.com/pascaliske/magicmirror/commit/a96e31339ca482358d339a02cce1fd37a3e58fe4))
* **server:** fix config reload handling for multiple clients ([82b0d4f](https://github.com/pascaliske/magicmirror/commit/82b0d4f3c6a14d6f6a84218f256304edcadf2625))


### Features

* **server:** implement logs with configurable levels ([de88b91](https://github.com/pascaliske/magicmirror/commit/de88b91ad55af517031300d6e3e70c01b3caea30))



## [0.0.2](https://github.com/pascaliske/magicmirror/compare/v0.0.1...v0.0.2) (2022-02-07)


### Features

* reload app on config file changes ([60d4435](https://github.com/pascaliske/magicmirror/commit/60d4435433881d9a1adfa5120d959df383256410))
* **server:** add initial http metrics ([eeb8f13](https://github.com/pascaliske/magicmirror/commit/eeb8f13be721975d483a1c2c1d7d44146685d2a1))
* **server:** inject build information into server binary ([e94cfa4](https://github.com/pascaliske/magicmirror/commit/e94cfa40a410570fa9b32e66cadbc794d390013c))



## 0.0.1 (2022-01-25)


### Bug Fixes

* **news-card:** properly cache multiple news feed requests ([0bd9fc7](https://github.com/pascaliske/magicmirror/commit/0bd9fc7f4045234b2d0230dd6326d83a53e6e7cf))
* **weather-card:** ensure card is no-op if api key is missing ([41f0f3c](https://github.com/pascaliske/magicmirror/commit/41f0f3c8f155cd0ef249e4e8e8026f45b5f570c6))


### Features

* check health of application periodically ([9610f03](https://github.com/pascaliske/magicmirror/commit/9610f03159f2a2fe5d426cad630634707241c017))
* implement translation setup ([32ded95](https://github.com/pascaliske/magicmirror/commit/32ded95e5a01c2d5a51cf7ed39ba52f8b6b76959))
* **news-card:** add initial news card implementation ([ea746b4](https://github.com/pascaliske/magicmirror/commit/ea746b4c47fbbee2bb50eec093fe00880f9a53ca))
* **server:** always search for config file in default directory ([6af7dea](https://github.com/pascaliske/magicmirror/commit/6af7dea4416893d9688d746435de37fc26f9fb4a))
* **time-card:** add initial time card implementation ([afee2cc](https://github.com/pascaliske/magicmirror/commit/afee2cc882a1b4100e194b54075ed781074dc9e4))
* **weather-card:** add initial weather card implementation ([d13457e](https://github.com/pascaliske/magicmirror/commit/d13457e718729768efc24a9124ea908c3430b925))




