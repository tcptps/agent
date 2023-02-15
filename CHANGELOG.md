# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v3.44.0](https://github.com/buildkite/agent/tree/v3.44.0) (2023-02-15)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.43.1...v3.44.0)

### Fixed
- Restore old tini path on Ubuntu 20.04 [#1934](https://github.com/buildkite/agent/pull/1934) (@triarius)
- When the `ansi-timestamps` experiment is enabled, timestamps are now computed at the end of each line [#1940](https://github.com/buildkite/agent/pull/1940) (@DrJosh9000)

### Added
- New flag (`git-checkout-flags`) and environment variable (`BUILDKITE_GIT_CHECKOUT_FLAGS`) for passing extra flags to `git checkout` [#1891](https://github.com/buildkite/agent/pull/1891) (@jmelahman)

### Changed
- Upstart is no longer supported [#1946](https://github.com/buildkite/agent/pull/1946) (@sj26)
- Better errors when config loading fails [#1937](https://github.com/buildkite/agent/pull/1937) (@moskyb)
- Pipelines are now parsed with gopkg.in/yaml.v3. This change should be invisible, but involved a non-trivial amount of new code. [#1930](https://github.com/buildkite/agent/pull/1930) (@DrJosh9000)
- Many dependency updates, notably Go v1.20.1 [#1955](https://github.com/buildkite/agent/pull/1955).
- Several minor improvements and clean-ups (@sj26, @triarius, @jonahbull, @DrJosh9000)

## [3.43.1](https://github.com/buildkite/agent/tree/3.43.1) (2023-01-20)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.43.0...3.43.1)

### Fixed
- An issue introduced in v3.43.0 where agents running in acquire mode would exit after ~4.5 minutes, failing the job they were running [#1923](https://github.com/buildkite/agent/pull/1923) (@leathekd)

## [3.43.0](https://github.com/buildkite/agent/tree/3.43.0) (2023-01-18)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.42.0...3.43.0)

### Fixed
- A nil pointer dereference introduced in 3.42.0 due to missing error handling after calling `user.Current` [#1910](https://github.com/buildkite/agent/pull/1910) (@DrJosh9000)

### Added
- A flag to allow empty results with doing an artifact search [#1887](https://github.com/buildkite/agent/pull/1887) (@MatthewDolan)
- Docker Images for linux/arm64 [#1901](https://github.com/buildkite/agent/pull/1901) (@triarius)
- Agent tags are added from ECS container metadata [#1870](https://github.com/buildkite/agent/pull/1870) (@francoiscampbell)

### Changed
- The `env` subcommand is now `env dump` [#1920](https://github.com/buildkite/agent/pull/1920) (@pda)
- AcquireJob now retries while the job is locked [#1894](https://github.com/buildkite/agent/pull/1894) (@triarius)
- Various miscellaneous updates and improvements (@moskyb, @triarius, @mitchbne, @dependabot[bot])

## [v3.42.0](https://github.com/buildkite/agent/tree/v3.42.0) (2023-01-05)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.41.0...v3.42.0)

 ### Added
 - Add an in-built hierarchical status page [#1873](https://github.com/buildkite/agent/pull/1873) (@DrJosh9000)
 - Add an `agent-startup` hook that fires at the same time as the `agent-shutdown` hook is registered [#1778](https://github.com/buildkite/agent/pull/1778) (@donalmacc)

 ### Changed
- Enforce a timeout on `finishJob` and `onUploadChunk` [#1854](https://github.com/buildkite/agent/pull/1854) (@DrJosh9000)
- A variety of dependency updates, documentation, and code cleanups! (@dependabot[bot], @DrJosh9000, @moskyb)
- Flakey test fixes and test suite enhancements (@triarius, @DrJosh9000)

 ### Fixed
 - Ensure that unrecoverable errors for Heartbeat and Ping stop the agent [#1855](https://github.com/buildkite/agent/pull/1855) (@moskyb)

 ### Security
 - Update `x/crypto/ssh` to `0.3.0`, patching CVE-2020-9283 [#1857](https://github.com/buildkite/agent/pull/1857) (@moskyb)


## [v3.41.0](https://github.com/buildkite/agent/tree/v3.41.0) (2022-11-24)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.40.0...v3.41.0)

### Added
- Experimental `buildkite-agent oidc request-token` command [#1827](https://github.com/buildkite/agent/pull/1827) (@triarius)
- Option to set the service name for tracing [#1779](https://github.com/buildkite/agent/pull/1779) (@goodspark)

### Changed

- Update windows install script to detect arm64 systems [#1768](https://github.com/buildkite/agent/pull/1768) (@yob)
- Install docker compose v2 plugin in agent alpine and ubuntu docker images [#1841](https://github.com/buildkite/agent/pull/1841) (@ajoneil) (@triarius)
- 🧹 A variety of dependency updates, documentation, and cleanups! (@dependabot[bot]) (@DrJosh9000)


## [v3.40.0](https://github.com/buildkite/agent/tree/v3.40.0) (2022-11-08)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.39.0...v3.40.0)

### Added

- Agent binaries for windows/arm64 [#1767](https://github.com/buildkite/agent/pull/1767) (@yob)
- Alpine k8s image [#1771](https://github.com/buildkite/agent/pull/1771) (@dabarrell)

### Security

- (Fixed in 3.39.1) A security issue in environment handling between buildkite-agent and Bash 5.2 [#1781](https://github.com/buildkite/agent/pull/1781) (@moskyb)
- Secret redaction now handles secrets containing UTF-8 code points greater than 255 [#1809](https://github.com/buildkite/agent/pull/1809) (@DrJosh9000)
- The update to Go 1.19.3 fixes two Go security issues (particularly on Windows):
   - The current directory (`.`) in `$PATH` is now ignored for finding executables - see https://go.dev/blog/path-security
   - Environment variable values containing null bytes are now sanitised - see https://github.com/golang/go/issues/56284

### Changed

- 5xx responses are now retried when attempting to start a job [#1777](https://github.com/buildkite/agent/pull/1777) (@jonahbull)
- 🧹 A variety of dependency updates and cleanups!

## [v3.39.0](https://github.com/buildkite/agent/tree/v3.39.0) (2022-09-08)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.38.0...v3.39.0)

### Added
- gcp:instance-name and tweak GCP labels fetching [#1742](https://github.com/buildkite/agent/pull/1742) (@pda)
- Support for not-yet-released per-job agent tokens [#1745](https://github.com/buildkite/agent/pull/1745) (@moskyb)

### Changed
- Retry Disconnect API calls [#1761](https://github.com/buildkite/agent/pull/1761) (@pda)
- Only search for finished artifacts [#1728](https://github.com/buildkite/agent/pull/1728) (@moskyb)
- Cache S3 clients between artifact downloads [#1732](https://github.com/buildkite/agent/pull/1732) (@moskyb)
- Document label edge case [#1718](https://github.com/buildkite/agent/pull/1718) (@plaindocs)

### Fixed
- Docker: run /sbin/tini without -g for graceful termination [#1763](https://github.com/buildkite/agent/pull/1763) (@pda)
- Fix multiple-nested plugin repos on gitlab [#1746](https://github.com/buildkite/agent/pull/1746) (@moskyb)
- Fix unowned plugin reference [#1733](https://github.com/buildkite/agent/pull/1733) (@moskyb)
- Fix order of level names for logger.Level.String() [#1722](https://github.com/buildkite/agent/pull/1722) (@moskyb)
- Fix warning log level [#1721](https://github.com/buildkite/agent/pull/1721) (@ChrisBr)

## [v3.38.0](https://github.com/buildkite/agent/tree/v3.38.0) (2022-07-20)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.37.0...v3.38.0)

### Changed
- Include a list of enabled features in the register request [#1706](https://github.com/buildkite/agent/pull/1706) (@moskyb)
- Promote opentelemetry tracing to mainline feature status [#1702](https://github.com/buildkite/agent/pull/1702) (@moskyb)
- Improve opentelemetry implementation [#1699](https://github.com/buildkite/agent/pull/1699) [#1705](https://github.com/buildkite/agent/pull/1705) (@moskyb)

## [v3.37.0](https://github.com/buildkite/tree/v3.37.0) (2022-07-06)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.36.1...v3.37.0)

### Added

* Agent metadata includes `arch` (e.g. `arch=amd64`) alongside `hostname` and `os` [#1691](https://github.com/buildkite/agent/pull/1691) ([moskyb](https://github.com/moskyb))
* Allow forcing clean checkout of plugins [#1636](https://github.com/buildkite/agent/pull/1636) ([toothbrush](https://github.com/toothbrush))

### Fixed

* Environment modification in hooks that set bash arrays [#1692](https://github.com/buildkite/agent/pull/1692) ([moskyb](https://github.com/moskyb))
* Unescape backticks when parsing env from export -p output [#1687](https://github.com/buildkite/agent/pull/1687) ([moskyb](https://github.com/moskyb))
* Log “Using flock-file-locks experiment 🧪” when enabled [#1688](https://github.com/buildkite/agent/pull/1688) ([lox](https://github.com/lox))
* flock-file-locks experiment: errors logging [#1689](https://github.com/buildkite/agent/pull/1689) ([KevinGreen](https://github.com/KevinGreen))
* Remove potentially-corrupted mirror dir if clone fails [#1671](https://github.com/buildkite/agent/pull/1671) ([lox](https://github.com/lox))
* Improve log-level flag usage description [#1676](https://github.com/buildkite/agent/pull/1676) ([pzeballos](https://github.com/pzeballos))

### Changed

* datadog-go major version upgrade to v5.1.1 [#1666](https://github.com/buildkite/agent/pull/1666) ([moskyb](https://github.com/moskyb))
* Revert to delegating directory creation permissions to system umask [#1667](https://github.com/buildkite/agent/pull/1667) ([moskyb](https://github.com/moskyb))
* Replace retry code with [roko](https://github.com/buildkite/roko) [#1675](https://github.com/buildkite/agent/pull/1675) ([moskyb](https://github.com/moskyb))
* bootstrap/shell: round command durations to 5 significant digits [#1651](https://github.com/buildkite/agent/pull/1651) ([kevinburkesegment](https://github.com/kevinburkesegment))


## [v3.36.1](https://github.com/buildkite/agent/tree/v3.36.1) (2022-05-27)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.36.0...v3.36.1)

### Fixed
- Fix nil pointer deref when using --log-format json [#1653](https://github.com/buildkite/agent/pull/1653) (@moskyb)

## [v3.36.0](https://github.com/buildkite/agent/tree/v3.36.0) (2022-05-17)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.35.2...v3.36.0)

### Added

- Add experiment to use kernel-based flocks instead of lockfiles [#1624](https://github.com/buildkite/agent/pull/1624) (@KevinGreen)
- Add option to enable temporary job log file [#1564](https://github.com/buildkite/agent/pull/1564) (@albertywu)
- Add experimental OpenTelemetry Tracing Support [#1631](https://github.com/buildkite/agent/pull/1631) + [#1632](https://github.com/buildkite/agent/pull/1632) (@moskyb)
- Add `--log-level` flag to all commands [#1635](https://github.com/buildkite/agent/pull/1635) (@moskyb)

### Fixed

- The `no-plugins` option now works correctly when set in the config file [#1579](https://github.com/buildkite/agent/pull/1579) (@elruwen)
- Clear up usage instructions around `--disconnect-after-idle-timeout` and `--disconnect-after-job` [#1599](https://github.com/buildkite/agent/pull/1599) (@moskyb)

### Changed
- Refactor retry machinery to allow the use of exponential backoff [#1588](https://github.com/buildkite/agent/pull/1588) (@moskyb)
- Create all directories with 0775 permissions [#1616](https://github.com/buildkite/agent/pull/1616) (@moskyb)
- Dependency Updates:
  - github.com/urfave/cli: 1.22.4 -> 1.22.9 [#1619](https://github.com/buildkite/agent/pull/1619) + [#1638](https://github.com/buildkite/agent/pull/1638) (@dependabot[bot])
  - Golang: 1.17.6 -> 1.18.1 (yay, generics!) [#1603](https://github.com/buildkite/agent/pull/1603) + [#1627](https://github.com/buildkite/agent/pull/1627) (@dependabot[bot])
  - Alpine Build Images: 3.15.0 -> 3.15.4 [#1626](https://github.com/buildkite/agent/pull/1626) (@dependabot[bot])
  - Alpine Release Images: 3.12 -> 3.15.4 [#1628](https://github.com/buildkite/agent/pull/1628) (@moskyb)

## [v3.35.2](https://github.com/buildkite/agent/tree/v3.35.2) (2022-04-13)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.35.1...v3.35.2)

### Fixed
- Fix race condition in bootstrap.go [#1606](https://github.com/buildkite/agent/pull/1606) (@moskyb)

### Changed
- Bump some dependency versions - thanks @dependabot!
  - github.com/stretchr/testify: 1.5.1 -> 1.7.1 [#1608](https://github.com/buildkite/agent/pull/1608)
  - github.com/mitchellh/go-homedir: 1.0.0 -> 1.1.0 [#1576](https://github.com/buildkite/agent/pull/1576)

## [v3.35.1](https://github.com/buildkite/agent/tree/v3.35.1) (2022-04-05)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.35.0...v3.35.1)

### Fixed

- Revert file permission changes made in [#1580](https://github.com/buildkite/agent/pull/1580). They were creating issues with docker-based workflows [#1601](https://github.com/buildkite/agent/pull/1601) (@pda + @moskyb)

## [v3.35.0](https://github.com/buildkite/agent/tree/v3.35.0) (2022-03-23)
[Full Changelog](https://github.com/buildkite/agent/compare/v3.34.0...v3.35.0)

### Changed

- Make `go fmt` mandatory in this repo [#1587](https://github.com/buildkite/agent/pull/1587) (@moskyb)
- Only search for finished artifact uploads when using `buildkite-agent artifact download` and `artifact shasum` [#1584](https://github.com/buildkite/agent/pull/1584) (@pda)
- Improve help/usage/errors for `buildkite-agent artifact shasum` [#1581](https://github.com/buildkite/agent/pull/1581) (@pda)
- Make the agent look for work immediately after completing a job, rather than waiting the ping interval [#1567](https://github.com/buildkite/agent/pull/1567) (@extemporalgenome)
- Update github.com/aws/aws-sdk-go to the latest v1 release [#1573](https://github.com/buildkite/agent/pull/1573) (@yob)
- Enable dependabot for go.mod [#1574](https://github.com/buildkite/agent/pull/1574) (@yob)
- Use build matrix feature to simplify CI pipeline [#1566](https://github.com/buildkite/agent/pull/1566) (@ticky)
  - Interested in using Build Matrices yourself? Check out [our docs!](https://buildkite.com/docs/pipelines/build-matrix)
- Buildkite pipeline adjustments [#1597](https://github.com/buildkite/agent/pull/1597) (@moskyb)

### Fixed

- Use `net.JoinHostPort()` to join host/port combos, rather than `fmt.Sprintf()` [#1585](https://github.com/buildkite/agent/pull/1585) (@pda)
- Fix minor typo in help text for `buildkite-agent pipeline upload [#1595](https://github.com/buildkite/agent/pull/1595) (@moskyb)

### Added

- Add option to skip updating the mirror when using git mirrors. Useful when git is mounted from an external volume, NFS mount etc [#1552](https://github.com/buildkite/agent/pull/1552) (@fatmcgav)
- Use the more secure SHA256 hashing algorithm alongside SHA1 when working with artifacts [#1582](https://github.com/buildkite/agent/pull/1582) [#1583](https://github.com/buildkite/agent/pull/1583) [#1584](https://github.com/buildkite/agent/pull/1584) (@pda)

### Security

- When creating directories, make them only accessible by current user and group [#1580](https://github.com/buildkite/agent/pull/1580) (@pda)

## [v3.34.1](https://github.com/buildkite/agent/compare/v3.34.0...v3.34.1) (2022-03-23)

### Fixed

- Make secret value rejection on pipeline upload optional. **This undoes a breaking change accidentally included in v3.34.0** [#1589](https://github.com/buildkite/agent/pull/1589) (@moskyb)

## [v3.34.0](https://github.com/buildkite/agent/compare/v3.33.3...v3.34.0) (2022-03-01)

### Added

* Introduce `spawn-with-priority` option [#1530](https://github.com/buildkite/agent/pull/1530) ([sema](https://github.com/sema))

### Fixed

* Retry 500 responses when batch creating artifacts [#1568](https://github.com/buildkite/agent/pull/1568) ([moskyb](https://github.com/moskyb))
* Report OS versions when running on AIX and Solaris [#1559](https://github.com/buildkite/agent/pull/1559) ([yob](https://github.com/yob))
* Support multiple commands on Windows [#1543](https://github.com/buildkite/agent/pull/1543) ([keithduncan](https://github.com/keithduncan))
* Allow `BUILDKITE_S3_DEFAULT_REGION` to be used for unconditional bucket region [#1535](https://github.com/buildkite/agent/pull/1535) ([keithduncan](https://github.com/keithduncan))

### Changed

* Go version upgraded from 1.16 to 1.17 [#1557](https://github.com/buildkite/agent/pull/1557) [#1549](https://github.com/buildkite/agent/pull/1549)
* Remove the CentOS (end-of-life) docker image [#1561](https://github.com/buildkite/agent/pull/1561) ([tessereth](https://github.com/tessereth))
* Plugin `git clone` is retried up to 3 times [#1539](https://github.com/buildkite/agent/pull/1539) ([pzeballos](https://github.com/pzeballos))
* Docker image alpine upgraded from 3.14.2 to 3.15.0 [#1541](https://github.com/buildkite/agent/pull/1541)

### Security

* Lock down file permissions on windows [#1562](https://github.com/buildkite/agent/pull/1562) ([tessereth](https://github.com/tessereth))
* Reject pipeline uploads containing redacted vars [#1523](https://github.com/buildkite/agent/pull/1523) ([keithduncan](https://github.com/keithduncan))

## [v3.33.3](https://github.com/buildkite/agent/compare/v3.33.2...v3.33.3) (2021-09-29)

### Fixed

* Fix erroneous working directory change for hooks that early exit [#1520](https://github.com/buildkite/agent/pull/1520)

## [v3.33.2](https://github.com/buildkite/agent/compare/v3.33.1...v3.33.2) (2021-09-29)

### Fixed

* Non backwards compatible change to artifact download path handling [#1518](https://github.com/buildkite/agent/pull/1518)

## [v3.33.1](https://github.com/buildkite/agent/compare/v3.33.0...v3.33.1) (2021-09-28)

### Fixed

* A crash in `buildkite-agent bootstrap` when command hooks early exit [#1516](https://github.com/buildkite/agent/pull/1516)

## [v3.33.0](https://github.com/buildkite/agent/compare/v3.32.3...v3.33.0) (2021-09-27)

### Added

* Support for `unset` environment variables in Job Lifecycle Hooks [#1488](https://github.com/buildkite/agent/pull/1488)

### Changed

* Remove retry handling when deleting annotations that are already deleted [#1507](https://github.com/buildkite/agent/pull/1507) ([@lox](https://github.com/lox))
* Alpine base image from 3.14.0 to 3.14.2 [#1499](https://github.com/buildkite/agent/pull/1499)

### Fixed

* Support for trailing slash path behaviour in artifact download [#1504](https://github.com/buildkite/agent/pull/1504) ([@jonathan-brand](https://github.com/jonathan-brand))

## [v3.32.3](https://github.com/buildkite/agent/compare/v3.32.2...v3.32.3) (2021-09-01)

### Fixed

* PowerShell hooks on Windows [#1497](https://github.com/buildkite/agent/pull/1497)

## [v3.32.2](https://github.com/buildkite/agent/compare/v3.32.1...v3.32.2) (2021-08-31)

### Added

* Improved error logging around AWS Credentials [#1490](https://github.com/buildkite/agent/pull/1490)
* Logging to the artifact upload command to say where artifacts are being sent [#1486](https://github.com/buildkite/agent/pull/1486)
* Support for cross-region artifact buckets [#1495](https://github.com/buildkite/agent/pull/1495)

### Changed

* artifact_paths failures no longer mask a command error [#1487](https://github.com/buildkite/agent/pull/1487)

### Fixed

* Failed plug-in checkouts using the default branch instead of the requested version [#1493](https://github.com/buildkite/agent/pull/1493)
* Missing quote in the PowerShell hook wrapper [#1494](https://github.com/buildkite/agent/pull/1494)

## [v3.32.1](https://github.com/buildkite/agent/compare/v3.32.0...v3.32.1) (2021-08-06)

### Fixed

* A panic in the log redactor when processing certain bytes [#1478](https://github.com/buildkite/agent/issues/1478) ([scv119](https://github.com/scv119))

## [v3.32.0](https://github.com/buildkite/agent/compare/v3.31.0...v3.32.0) (2021-07-30)

### Added

* A new pre-bootstrap hook which can accept or reject jobs before environment variables are loaded [#1456](https://github.com/buildkite/agent/pull/1456)
* `ppc64` and `ppc64le` architecture binaries to the DEB and RPM packages [#1474](https://github.com/buildkite/agent/pull/1474) [#1473](https://github.com/buildkite/agent/pull/1473) ([staticfloat](https://github.com/staticfloat))
* Use text/yaml mime type for .yml and .yaml artifacts [#1470](https://github.com/buildkite/agent/pull/1470)

### Changed

* Add BUILDKITE_BIN_PATH to end, not start, of PATH [#1465](https://github.com/buildkite/agent/pull/1465) ([DavidSpickett](https://github.com/DavidSpickett))

## [v3.31.0](https://github.com/buildkite/agent/compare/v3.30.0...v3.31.0) (2021-07-02)

### Added

* Output secret redaction is now on by default [#1452](https://github.com/buildkite/agent/pull/1452)
* Improved CLI docs for `buildkite-agent artifact download` [#1446](https://github.com/buildkite/agent/pull/1446)

### Changed

* Build using golang 1.16.5 [#1460](https://github.com/buildkite/agent/pull/1460)

### Fixed

* Discovery of the `buildkite-agent` binary path in more situations [#1444](https://github.com/buildkite/agent/pull/1444) [#1457](https://github.com/buildkite/agent/pull/1457)

## [v3.30.0](https://github.com/buildkite/agent/compare/v3.29.0...v3.30.0) (2021-05-28)

### Added
* Send queue metrics to Datadog when job received [#1442](https://github.com/buildkite/agent/pull/1442) ([keithduncan](https://github.com/keithduncan))
* Add flag to send Datadog Metrics as Distributions [#1433](https://github.com/buildkite/agent/pull/1433) ([amukherjeetwilio](https://github.com/amukherjeetwilio))
* Ubuntu 18.04 based Docker image [#1441](https://github.com/buildkite/agent/pull/1441) ([keithduncan](https://github.com/keithduncan))
* Build binaries for `netbsd` and `s390x` [#1432](https://github.com/buildkite/agent/pull/1432), [#1421](https://github.com/buildkite/agent/pull/1421) ([yob](https://github.com/yob))
* Add `wait-for-ec2-meta-data-timeout` config variable [#1425](https://github.com/buildkite/agent/pull/1425) ([OliverKoo](https://github.com/OliverKoo))

### Changed
* Build using golang 1.16.4 [#1429](https://github.com/buildkite/agent/pull/1429)
* Replace kr/pty with creack/pty and upgrade from 1.1.2 to 1.1.12 [#1431](https://github.com/buildkite/agent/pull/1431) ([ibuclaw](https://github.com/ibuclaw))

### Fixed
* Trim trailing slash from `buildkite-agent artifact upload` when using custom S3 bucket paths [#1427](https://github.com/buildkite/agent/pull/1427) ([shevaun](https://github.com/shevaun))
* Use /usr/pkg/bin/bash as default shell on NetBSD [#1430](https://github.com/buildkite/agent/pull/1430) ([ibuclaw](https://github.com/ibuclaw))

## [v3.29.0](https://github.com/buildkite/agent/compare/v3.28.1...v3.29.0) (2021-04-21)

### Changed
* Support mips64le architecture target. [#1379](https://github.com/buildkite/agent/pull/1379) ([houfangdong](https://github.com/houfangdong))
* Search the path for bash when running bootstrap scripts [#1404](https://github.com/buildkite/agent/pull/1404) ([yob](https://github.com/yob))
* Output-redactor: redact shell logger, including changed env vars [#1401](https://github.com/buildkite/agent/pull/1401) ([pda](https://github.com/pda))
* Add *_ACCESS_KEY & *_SECRET_KEY to default redactor-var [#1405](https://github.com/buildkite/agent/pull/1405) ([pda](https://github.com/pda))
* Build with Golang 1.16.3 [#1412](https://github.com/buildkite/agent/pull/1412) ([dependabot[bot]](https://github.com/apps/dependabot))
* Update [Buildkite CLI](https://github.com/buildkite/cli) release from 1.0.0 to 1.2.0 [#1403](https://github.com/buildkite/agent/pull/1403) ([yob](https://github.com/yob))

### Fixed
* Avoid occasional failure to run jobs when working directory is missing [#1402](https://github.com/buildkite/agent/pull/1402) ([yob](https://github.com/yob))
* Avoid a rare panic when running `buildkite-agent pipeline upload` [#1406](https://github.com/buildkite/agent/pull/1406) ([yob](https://github.com/yob))

## [v3.28.1](https://github.com/buildkite/agent/compare/v3.27.0...v3.28.1)

### Added

* collect instance-life-cycle as a default tag on EC2 instances [#1374](https://github.com/buildkite/agent/pull/1374) [yob](https://github.com/yob))
* Expose plugin config in two new instance variables, `BUILDKITE_PLUGIN_NAME` and `BUILDKITE_PLUGIN_CONFIGURATION` [#1382](https://github.com/buildkite/agent/pull/1382) [moensch](https://github.com/moensch)
* Add `buildkite-agent annotation remove` command [#1364](https://github.com/buildkite/agent/pull/1364/) [ticky](https://github.com/ticky)
* Allow customizing the signal bootstrap sends to processes on cancel  [#1390](https://github.com/buildkite/agent/pull/1390/) [brentleyjones](https://github.com/brentleyjones)

### Changed

* On new installs the default agent name has changed from `%hostname-%n` to `%hostname-%spawn` [#1389](https://github.com/buildkite/agent/pull/1389) [pda](https://github.com/pda)

### Fixed

* Fixed --no-pty flag [#1394][https://github.com/buildkite/agent/pull/1394] [pda](https://github.com/pda)

## v3.28.0

* Skipped due to a versioning error

## [v3.27.0](https://github.com/buildkite/agent/compare/v3.26.0...v3.27.0)

### Added
* Add support for agent tracing using Datadog APM [#1273](https://github.com/buildkite/agent/pull/1273) ([goodspark](https://github.com/goodspark), [Sam Schlegel](https://github.com/samschlegel))
* Improvements to ARM64 support (Apple Silicon/M1) [#1346](https://github.com/buildkite/agent/pull/1346), [#1354](https://github.com/buildkite/agent/pull/1354), [#1343](https://github.com/buildkite/agent/pull/1343) ([ticky](https://github.com/ticky))
* Add a Linux ppc64 build to the pipeline [#1362](https://github.com/buildkite/agent/pull/1362) ([ticky](https://github.com/ticky))
* Agent can now upload artifacts using AssumedRoles using `BUILDKITE_S3_SESSION_TOKEN` [#1359](https://github.com/buildkite/agent/pull/1359) ([grahamc](https://github.com/grahamc))
* Agent name `%spawn` interpolation to deprecate/replace `%n` [#1377](https://github.com/buildkite/agent/pull/1377) ([ticky](https://github.com/ticky))

### Changed
* Compile the darwin/arm64 binary using go 1.16rc1 [#1352](https://github.com/buildkite/agent/pull/1352) ([yob](https://github.com/yob)) [#1369](https://github.com/buildkite/agent/pull/1369) ([chloeruka](https://github.com/chloeruka))
* Use Docker CLI packages, update Docker Compose, and update centos to 8.x [#1351](https://github.com/buildkite/agent/pull/1351) ([RemcodM](https://github.com/RemcodM))

## Fixed
* Fixed an issue in #1314 that broke pull requests with git-mirrors [#1347](https://github.com/buildkite/agent/pull/1347) ([ticky](https://github.com/ticky))

## [v3.26.0](https://github.com/buildkite/agent/compare/v3.25.0...v3.26.0) (2020-12-03)

### Added

* Compile an experimental native executable for Apple Silicon [#1339](https://github.com/buildkite/agent/pull/1339) ([yob](https://github.com/yob))
  * Using a pre-release version of go, we'll switch to compiling with go 1.16 once it's released

### Changed

* Install script: use the arm64 binary for aarch64 machines [#1340](https://github.com/buildkite/agent/pull/1340) ([gc-plp](https://github.com/gc-plp))
* Build with golang 1.15 [#1334](https://github.com/buildkite/agent/pull/1334) ([yob](https://github.com/yob))
* Bump alpine docker image from alpine 3.8 to 3.12 [#1333](https://github.com/buildkite/agent/pull/1333) ([yob](https://github.com/yob))
* Upgrade docker ubuntu to 20.04 focal [#1312](https://github.com/buildkite/agent/pull/1312) ([sj26](https://github.com/sj26))

## [v3.25.0](https://github.com/buildkite/agent/compare/v3.24.0...v3.25.0) (2020-10-21)

### Added
* Add --mirror flag by default for mirror clones [#1328](https://github.com/buildkite/agent/pull/1328) ([chrislloyd](https://github.com/chrislloyd))
* Add an agent-wide shutdown hook [#1275](https://github.com/buildkite/agent/pull/1275) ([goodspark](https://github.com/goodspark)) [#1322](https://github.com/buildkite/agent/pull/1322) ([pda](https://github.com/pda))

### Fixed
* Improve windows telemetry so that we report the version accurately in-platform [#1330](https://github.com/buildkite/agent/pull/1330) ([yob](https://github.com/yob)) [#1316](https://github.com/buildkite/agent/pull/1316) ([yob](https://github.com/yob))
* Ensure no orphaned processes when Windows jobs complete [#1329](https://github.com/buildkite/agent/pull/1329) ([yob](https://github.com/yob))
* Log error messages when canceling a running job fails [#1317](https://github.com/buildkite/agent/pull/1317) ([yob](https://github.com/yob))
* gitCheckout() validates branch, plus unit tests [#1315](https://github.com/buildkite/agent/pull/1315) ([pda](https://github.com/pda))
* gitFetch() terminates options with -- before repo/refspecs [#1314](https://github.com/buildkite/agent/pull/1314) ([pda](https://github.com/pda))

## [v3.24.0](https://github.com/buildkite/agent/compare/v3.23.1...v3.24.0) (2020-09-29)

### Fixed
* Fix build script [#1300](https://github.com/buildkite/agent/pull/1300) ([goodspark](https://github.com/goodspark))
* Fix typos [#1297](https://github.com/buildkite/agent/pull/1297) ([JuanitoFatas](https://github.com/JuanitoFatas))
* Fix flaky tests: experiments and parallel tests don't mix [#1295](https://github.com/buildkite/agent/pull/1295) ([pda](https://github.com/pda))
* artifact_uploader_test fixed for Windows. [#1288](https://github.com/buildkite/agent/pull/1288) ([pda](https://github.com/pda))
* Windows integration tests: git file path quirk fix [#1291](https://github.com/buildkite/agent/pull/1291) ([pda](https://github.com/pda))

### Changed
* git-mirrors: set BUILDKITE_REPO_MIRROR=/path/to/mirror/repo [#1311](https://github.com/buildkite/agent/pull/1311) ([pda](https://github.com/pda))
* Truncate BUILDKITE_MESSAGE to 64 KiB [#1307](https://github.com/buildkite/agent/pull/1307) ([pda](https://github.com/pda))
* CI: windows tests on queue=buildkite-agent-windows without Docker [#1294](https://github.com/buildkite/agent/pull/1294) ([pda](https://github.com/pda))
* buildkite:git:commit meta-data via stdin; avoid Argument List Too Long [#1302](https://github.com/buildkite/agent/pull/1302) ([pda](https://github.com/pda))
* Expand the CLI test step [#1293](https://github.com/buildkite/agent/pull/1293) ([ticky](https://github.com/ticky))
* Improve Apple Silicon Mac support [#1289](https://github.com/buildkite/agent/pull/1289) ([ticky](https://github.com/ticky))
* update github.com/urfave/cli to the latest v1 release [#1287](https://github.com/buildkite/agent/pull/1287) ([yob](https://github.com/yob))


## [v3.23.1](https://github.com/buildkite/agent/compare/v3.23.0...v3.23.1) (2020-09-09)

### Fixed
* Fix CLI flag/argument order sensitivity regression [#1286](https://github.com/buildkite/agent/pull/1286) ([yob](https://github.com/yob))


## [v3.23.0](https://github.com/buildkite/agent/compare/v3.22.1...v3.23.0) (2020-09-04)

### Added
* New artifact search subcommand [#1278](https://github.com/buildkite/agent/pull/1278) ([chloeruka](https://github.com/chloeruka))
![image](https://user-images.githubusercontent.com/30171259/92212159-e32bd700-eed4-11ea-9af8-2ad024eaecc1.png)
* Add sidecar agent suitable for being shared via volume in ECS or Kubernetes [#1218](https://github.com/buildkite/agent/pull/1218) ([keithduncan](https://github.com/keithduncan)) [#1263](https://github.com/buildkite/agent/pull/1263) ([yob](https://github.com/yob))
* We now fetch amd64 binaries on Apple Silicon Macs in anticipation of new macOS ARM computers [#1237](https://github.com/buildkite/agent/pull/1237) ([ticky](https://github.com/ticky))
* Opt-in experimental `resolve-commit-after-checkout` flag to resolve `BUILDKITE_COMMIT` refs (for example, "HEAD") to a commit hash [#1256](https://github.com/buildkite/agent/pull/1256) ([jayco](https://github.com/jayco))
* Experimental: Build & publish RPM ARM64 package for aarch64 [#1243](https://github.com/buildkite/agent/pull/1243) ([chloeruka](https://github.com/chloeruka)) [#1241](https://github.com/buildkite/agent/pull/1241) ([chloeruka](https://github.com/chloeruka))

### Changed
* Stop building i386 for darwin after 14 years of amd64 Mac hardware [#1238](https://github.com/buildkite/agent/pull/1238) ([pda](https://github.com/pda))
* Updated github.com/urfave/cli to v2 - this is the CLI framework we use for agent commands. [#1233](https://github.com/buildkite/agent/pull/1233) ([yob](https://github.com/yob)) [#1250](https://github.com/buildkite/agent/pull/1250) ([yob](https://github.com/yob))
* Send the reported system and machine names when fetching latest release [#1240](https://github.com/buildkite/agent/pull/1240) ([ticky](https://github.com/ticky))
* Bump build environment from [go](https://github.com/golang/go) 1.14.2 to 1.14.7 [#1235](https://github.com/buildkite/agent/pull/1235) ([yob](https://github.com/yob)) [#1262](https://github.com/buildkite/agent/pull/1262) ([yob](https://github.com/yob))
* Update [aws-sdk-go](https://github.com/aws/aws-sdk-go) to 1.32.10 [#1234](https://github.com/buildkite/agent/pull/1234) ([yob](https://github.com/yob))

### Fixed
* `git-mirrors` experiment now only fetches branch instead of a full remote update [#1112](https://github.com/buildkite/agent/pull/1112) ([lox](https://github.com/lox))
* Hooks can introduce empty environment variables [#1232](https://github.com/buildkite/agent/pull/1232) ([pda](https://github.com/pda))
* ArtifactUploader now deduplicates upload file paths [#1268](https://github.com/buildkite/agent/pull/1268) ([pda](https://github.com/pda))
* Added additional logging to artifact uploads  [#1266](https://github.com/buildkite/agent/pull/1266) ([yob](https://github.com/yob)) [#1265](https://github.com/buildkite/agent/pull/1265) ([denbeigh2000](https://github.com/denbeigh2000)) [#1255](https://github.com/buildkite/agent/pull/1255) ([yob](https://github.com/yob))
* Fail faster when uploading an artifact > 5Gb to unsupported destinations [#1264](https://github.com/buildkite/agent/pull/1264) ([yob](https://github.com/yob))
* Job should now reliably fail when process.Run() -> err [#1261](https://github.com/buildkite/agent/pull/1261) ([sj26](https://github.com/sj26))
* Fix checkout failure when there is a file called HEAD in the repository root [#1223](https://github.com/buildkite/agent/pull/1223) ([zhenyavinogradov](https://github.com/zhenyavinogradov)) [#1260](https://github.com/buildkite/agent/pull/1260) ([pda](https://github.com/pda))
* Enable `BUILDKITE_AGENT_DEBUG_HTTP` in jobs if it's enabled in the agent process [#1251](https://github.com/buildkite/agent/pull/1251) ([yob](https://github.com/yob))
* Avoid passing nils to Printf() during HTTP Debug mode [#1252](https://github.com/buildkite/agent/pull/1252) ([yob](https://github.com/yob))
* Allow `BUILDKITE_CLEAN_CHECKOUT` to be set via hooks [#1242](https://github.com/buildkite/agent/pull/1242) ([niceking](https://github.com/niceking))
* Add optional brackets to file arg documentation [#1276](https://github.com/buildkite/agent/pull/1276) ([harrietgrace](https://github.com/harrietgrace))
* Reword artifact shasum documentation [#1229](https://github.com/buildkite/agent/pull/1229) ([vineetgopal](https://github.com/vineetgopal))
* Provide example dogstatsd integration options [#1219](https://github.com/buildkite/agent/pull/1219) ([GaryPWhite](https://github.com/GaryPWhite))
* submit basic OS info when registering from a BSD system [#1239](https://github.com/buildkite/agent/pull/1239) ([yob](https://github.com/yob))
* Various typo fixes and light refactors [#1277](https://github.com/buildkite/agent/pull/1277) ([chloeruka](https://github.com/chloeruka)) [#1271](https://github.com/buildkite/agent/pull/1271) ([pda](https://github.com/pda)) [#1244](https://github.com/buildkite/agent/pull/1244) ([pda](https://github.com/pda)) [#1224](https://github.com/buildkite/agent/pull/1224) ([plaindocs](https://github.com/plaindocs))

## [v3.22.1](https://github.com/buildkite/agent/compare/v3.22.0...v3.22.1) (2020-06-18)

### Fixed

- Wrap calls for GCP metadata in a retry [#1230](https://github.com/buildkite/agent/pull/1230) ([jayco](https://github.com/jayco))
- Accept `--experiment` flags in all buildkite-agent subcommands [#1220](https://github.com/buildkite/agent/pull/1220) ([ticky](https://github.com/ticky))
- buildkite/interpolate updated; ability to use numeric default [#1217](https://github.com/buildkite/agent/pull/1217) ([pda](https://github.com/pda))

## [v3.22.0](https://github.com/buildkite/agent/tree/v3.22.0) (2020-05-15)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.21.1...v3.22.0)

### Changed

- Experiment: `normalised-upload-paths` Normalise upload path to Unix/URI form on Windows [#1211](https://github.com/buildkite/agent/pull/1211) (@ticky)
- Improve some outputs for error timers [#1212](https://github.com/buildkite/agent/pull/1212) (@ticky)

## [v3.21.1](https://github.com/buildkite/agent/tree/v3.21.1) (2020-05-05)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.21.0...v3.21.1)

### Fixed

- Rebuild with golang 1.14.2 to fix panic on some linux kernels [#1213](https://github.com/buildkite/agent/pull/1213) (@zifnab06)

## [v3.21.0](https://github.com/buildkite/agent/tree/v3.21.0) (2020-05-05)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.20.0...v3.21.0)

### Fixed

- Add a retry for errors during artifact search [#1210](https://github.com/buildkite/agent/pull/1210) (@lox)
- Fix checkout dir missing and hooks failing after failed checkout retries [#1192](https://github.com/buildkite/agent/pull/1192) (@sj26)

### Changed

- Bump golang build version to 1.14 [#1197](https://github.com/buildkite/agent/pull/1197) (@yob)
- Added 'spawn=1' with to all .cfg templates [#1175](https://github.com/buildkite/agent/pull/1175) (@drnic)
- Send more signal information back to Buildkite [#899](https://github.com/buildkite/agent/pull/899) (@lox)
- Updated artifact --help documentation [#1183](https://github.com/buildkite/agent/pull/1183) (@pda)
- Remove vendor in favor of go modules [#1117](https://github.com/buildkite/agent/pull/1117) (@lox)
- Update crypto [#1194](https://github.com/buildkite/agent/pull/1194) (@gavinelder)

## [v3.20.0](https://github.com/buildkite/agent/tree/v3.20.0) (2020-02-10)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.19.0...v3.20.0)

### Changed

- Multiple plugins can provide checkout & command hooks [#1161](https://github.com/buildkite/agent/pull/1161) (@pda)

## [v3.19.0](https://github.com/buildkite/agent/tree/v3.19.0) (2020-01-30)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.18.0...v3.19.0)

### Fixed

- Fix plugin execution being skipped with duplicate hook warning [#1156](https://github.com/buildkite/agent/pull/1156) (@toolmantim)
- minor changes [#1151](https://github.com/buildkite/agent/pull/1151) [#1154](https://github.com/buildkite/agent/pull/1154) [#1149](https://github.com/buildkite/agent/pull/1149)

## [v3.18.0](https://github.com/buildkite/agent/tree/v3.18.0) (2020-01-21)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.17.0...v3.18.0)

### Added

- Hooks can be written in PowerShell [#1122](https://github.com/buildkite/agent/pull/1122) (@pdemirdjian)

### Changed

- Ignore multiple checkout plugin hooks [#1135](https://github.com/buildkite/agent/pull/1135) (@toolmantim)
- clicommand/annotate: demote success log from Info to Debug [#1141](https://github.com/buildkite/agent/pull/1141) (@pda)

### Fixed

- Fix AgentPool to disconnect if AgentWorker.Start fails [#1146](https://github.com/buildkite/agent/pull/1146) (@keithduncan)
- Fix run-parts usage for CentOS docker entrypoint [#1139](https://github.com/buildkite/agent/pull/1139) (@moensch)

## [v3.17.0](https://github.com/buildkite/agent/tree/v3.17.0) (2019-12-11)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.16.0...v3.17.0)

### Added

- CentOS 7.x Docker image [#1137](https://github.com/buildkite/agent/pull/1137) (@moensch)
- Added --acquire-job for optionally accepting a specific job [#1138](https://github.com/buildkite/agent/pull/1138) (@keithpitt)
- Add filter to remove passwords, etc from job output [#1109](https://github.com/buildkite/agent/pull/1109) (@dbaggerman)
- Allow fetching arbitrary tag=suffix pairs from GCP/EC2 meta-data [#1067](https://github.com/buildkite/agent/pull/1067) (@plasticine)

### Fixed

- Propagate signals in intermediate bash shells [#1116](https://github.com/buildkite/agent/pull/1116) (@lox)
- Detect ansi clear lines and add ansi timestamps in ansi-timestamps experiments [#1128](https://github.com/buildkite/agent/pull/1128) (@lox)
- Added v3 for better go module support [#1115](https://github.com/buildkite/agent/pull/1115) (@sayboras)
- Convert windows paths to unix ones on artifact download [#1113](https://github.com/buildkite/agent/pull/1113) (@lox)

## [v3.16.0](https://github.com/buildkite/agent/tree/v3.16.0) (2019-10-14)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.15.2...v3.16.0)

### Added

- Add ANSI timestamp output experiment [#1103](https://github.com/buildkite/agent/pull/1103) (@lox)

### Changed

- Bump golang build version to 1.13 [#1107](https://github.com/buildkite/agent/pull/1107) (@lox)
- Drop support for setting process title [#1106](https://github.com/buildkite/agent/pull/1106) (@lox)

### Fixed

- Avoid destroying the checkout on specific git errors [#1104](https://github.com/buildkite/agent/pull/1104) (@lox)

## [v3.15.2](https://github.com/buildkite/agent/tree/v3.15.2) (2019-10-10)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.15.1...v3.15.2)

### Added

- Support GS credentials via BUILDKITE_GS_APPLICATION_CREDENTIALS [#1093](https://github.com/buildkite/agent/pull/1093) (@GaryPWhite)
- Add --include-retried-jobs to artifact download/shasum [#1101](https://github.com/buildkite/agent/pull/1101) (@toolmantim)

## [v3.15.1](https://github.com/buildkite/agent/tree/v3.15.1) (2019-09-30)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.15.0...v3.15.1)

### Fixed

- Fix a race condition that causes panics on job accept [#1095](https://github.com/buildkite/agent/pull/1095) (@lox)

## [v3.15.0](https://github.com/buildkite/agent/tree/v3.15.0) (2019-09-17)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.14.0...v3.15.0)

### Changed

- Let the agent serve a status page via HTTP. [#1066](https://github.com/buildkite/agent/pull/1066) (@philwo)
- Only execute the "command" hook once. [#1055](https://github.com/buildkite/agent/pull/1055) (@philwo)
- Fix goroutine leak and memory leak after job finishes [#1084](https://github.com/buildkite/agent/pull/1084) (@lox)
- Allow gs_downloader to use GS_APPLICATION_CREDENTIALS [#1086](https://github.com/buildkite/agent/pull/1086) (@GaryPWhite)
- Updates to `step update` and added `step get` [#1083](https://github.com/buildkite/agent/pull/1083) (@keithpitt)

## [v3.14.0](https://github.com/buildkite/agent/tree/v3.14.0) (2019-08-16)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.13.2...v3.14.0)

### Fixed

- Fix progress group id in debug output [#1074](https://github.com/buildkite/agent/pull/1074) (@lox)
- Avoid os.Exit in pipeline upload command [#1070](https://github.com/buildkite/agent/pull/1070) (@lox)
- Reword the cancel-grace-timeout config option [#1071](https://github.com/buildkite/agent/pull/1071) (@toolmantim)
- Correctly log last successful heartbeat time. [#1065](https://github.com/buildkite/agent/pull/1065) (@philwo)
- Add a test that BUILDKITE_GIT_SUBMODULES is handled [#1054](https://github.com/buildkite/agent/pull/1054) (@lox)

### Changed

- Added feature to enable encryption at rest for artifacts in S3. [#1072](https://github.com/buildkite/agent/pull/1072) (@wolfeidau)
- If commit is HEAD, always use FETCH_HEAD in agent checkout [#1064](https://github.com/buildkite/agent/pull/1064) (@matthewd)
- Updated `--help` output in the README and include more stuff in the "Development" section [#1061](https://github.com/buildkite/agent/pull/1061) (@keithpitt)
- Allow signal agent sends to bootstrap on cancel to be customized [#1041](https://github.com/buildkite/agent/pull/1041) (@lox)
- buildkite/pipeline.yaml etc in pipeline upload default search [#1051](https://github.com/buildkite/agent/pull/1051) (@pda)
- Move plugin tests to table-driven tests [#1048](https://github.com/buildkite/agent/pull/1048) (@lox)

## [v3.13.2](https://github.com/buildkite/agent/tree/v3.13.2) (2019-07-20)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.13.1...v3.13.2)

### Changed

- Fix panic on incorrect token [#1046](https://github.com/buildkite/agent/pull/1046) (@lox)
- Add artifactory vars to artifact upload --help output [#1042](https://github.com/buildkite/agent/pull/1042) (@harrietgrace)
- Fix buildkite-agent upload with absolute path (regression in v3.11.1) [#1044](https://github.com/buildkite/agent/pull/1044) (@petercgrant)
- Don't show vendored plugin header if none are present [#984](https://github.com/buildkite/agent/pull/984) (@lox)

## [v3.13.1](https://github.com/buildkite/agent/tree/v3.13.1) (2019-07-08)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.13.0...v3.13.1)

### Changed

- Add meta-data keys command [#1039](https://github.com/buildkite/agent/pull/1039) (@lox)
- Fix bug where file upload hangs and add a test [#1036](https://github.com/buildkite/agent/pull/1036) (@lox)
- Fix memory leak in artifact uploading with FormUploader [#1033](https://github.com/buildkite/agent/pull/1033) (@lox)
- Add profile option to all cli commands [#1032](https://github.com/buildkite/agent/pull/1032) (@lox)

## [v3.13.0](https://github.com/buildkite/agent/tree/v3.13.0) (2019-06-12)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.12.0...v3.13.0)

### Changed

- Quote command to git submodule foreach to fix error with git 2.20.0 [#1029](https://github.com/buildkite/agent/pull/1029) (@lox)
- Refactor api package to an interface [#1020](https://github.com/buildkite/agent/pull/1020) (@lox)

## [v3.12.0](https://github.com/buildkite/agent/tree/v3.12.0) (2019-05-22)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.11.5...v3.12.0)

### Added

- Add checksums for artifactory uploaded artifacts [#961](https://github.com/buildkite/agent/pull/961) (@lox)
- Add BUILDKITE_GCS_PATH_PREFIX for the path of GCS artifacts [#1000](https://github.com/buildkite/agent/pull/1000) (@DoomGerbil)

### Changed

- Don't force set the executable bit on scripts to be set [#1001](https://github.com/buildkite/agent/pull/1001) (@kuroneko)
- Deprecate disconnect-after-job-timeout [#1009](https://github.com/buildkite/agent/pull/1009) (@lox)
- Update Ubuntu docker image to docker-compose 1.24 [#1005](https://github.com/buildkite/agent/pull/1005) (@pecigonzalo)
- Update Artifactory path parsing to support windows [#1013](https://github.com/buildkite/agent/pull/1013) (@GaryPWhite)
- Refactor: Move signal handling out of AgentPool and into start command [#1012](https://github.com/buildkite/agent/pull/1012) (@lox)
- Refactor: Simplify how we handle idle timeouts [#1010](https://github.com/buildkite/agent/pull/1010) (@lox)
- Remove api socket proxy experiment [#1019](https://github.com/buildkite/agent/pull/1019) (@lox)
- Remove msgpack experiment [#1015](https://github.com/buildkite/agent/pull/1015) (@lox)

## [v3.11.5](https://github.com/buildkite/agent/tree/v3.11.5) (2019-05-13)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.11.4...v3.11.5)

### Fixed

- Fix broken signal handling [#1011](https://github.com/buildkite/agent/pull/1011) (@lox)

### Changed

- Update Ubuntu docker image to docker-compose 1.24 [#1005](https://github.com/buildkite/agent/pull/1005) (@pecigonzalo)

## [v3.11.4](https://github.com/buildkite/agent/tree/v3.11.4) (2019-05-08)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.11.3...v3.11.4)

### Changed

- Fix bug where disconnect-after-idle stopped working [#1004](https://github.com/buildkite/agent/pull/1004) (@lox)

## [v3.11.3](https://github.com/buildkite/agent/tree/v3.11.3) (2019-05-08)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.11.2...v3.11.3)

### Fixed

- Prevent host tags from overwriting aws/gcp tags [#1002](https://github.com/buildkite/agent/pull/1002) (@lox)

### Changed

- Replace signalwatcher package with os/signal [#998](https://github.com/buildkite/agent/pull/998) (@lox)
- Only trigger idle disconnect if all workers are idle [#999](https://github.com/buildkite/agent/pull/999) (@lox)

## [v3.11.2](https://github.com/buildkite/agent/tree/v3.11.2) (2019-04-20)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.11.1...v3.11.2)

### Changed

- Send logging to stderr again [#996](https://github.com/buildkite/agent/pull/996) (@lox)

## [v3.11.1](https://github.com/buildkite/agent/tree/v3.11.1) (2019-04-20)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.11.0...v3.11.1)

### Fixed

- Ensure heartbeats run until agent is stopped [#992](https://github.com/buildkite/agent/pull/992) (@lox)
- Revert "Refactor AgentConfiguration into JobRunnerConfig" to fix error accepting jobs[#993](https://github.com/buildkite/agent/pull/993) (@lox)

## [v3.11.0](https://github.com/buildkite/agent/tree/v3.11.0) (2019-04-16)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.10.4...v3.11.0)

### Fixed

- Allow applying ec2 tags when config tags are empty [#990](https://github.com/buildkite/agent/pull/990) (@vanstee)
- Upload Artifactory artifacts to correct path [#989](https://github.com/buildkite/agent/pull/989) (@GaryPWhite)

### Changed

- Combine apache and nginx sources for mime types. [#988](https://github.com/buildkite/agent/pull/988) (@blueimp)
- Support log output in json [#966](https://github.com/buildkite/agent/pull/966) (@lox)
- Add git-fetch-flags [#957](https://github.com/buildkite/agent/pull/957) (@lox)

## [v3.10.4](https://github.com/buildkite/agent/tree/v3.10.4) (2019-04-05)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.10.3...v3.10.4)

### Fixed

- Fix bug where logger was defaulting to debug [#974](https://github.com/buildkite/agent/pull/974) (@lox)
- Fix race condition between stop/cancel and register [#971](https://github.com/buildkite/agent/pull/971) (@lox)
- Fix incorrect artifactory upload url [#977](https://github.com/buildkite/agent/pull/977) (@GaryPWhite)

## [v3.10.3](https://github.com/buildkite/agent/tree/v3.10.3) (2019-04-02)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.10.2...v3.10.3)

### Fixed

- Fix bug where ec2 tags aren't added correctly [#970](https://github.com/buildkite/agent/pull/970) (@lox)
- Fix bug where host tags overwrite other tags [#969](https://github.com/buildkite/agent/pull/969) (@lox)

## [v3.10.2](https://github.com/buildkite/agent/tree/v3.10.2) (2019-03-31)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.10.1...v3.10.2)

### Fixed

- Update artifatory uploader to use the correct PUT url [#960](https://github.com/buildkite/agent/pull/960) (@GaryPWhite)

### Changed

- Refactor: Move logger.Logger to an interface [#962](https://github.com/buildkite/agent/pull/962) (@lox)
- Refactor: Move AgentWorker construction and registration out of AgentPool [#956](https://github.com/buildkite/agent/pull/956) (@lox)

## [v3.10.1](https://github.com/buildkite/agent/tree/v3.10.1) (2019-03-24)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.10.0...v3.10.1)

### Fixed

- Fix long urls for artifactory integration [#955](https://github.com/buildkite/agent/pull/955) (@GaryPWhite)

## [v3.10.0](https://github.com/buildkite/agent/tree/v3.10.0) (2019-03-12)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.9.1...v3.10.0)

### Added

- Experimental shared repositories (git mirrors) between checkouts [#936](https://github.com/buildkite/agent/pull/936) (@lox)
- Support disconnecting agent after it's been idle for a certain time period [#932](https://github.com/buildkite/agent/pull/932) (@lox)

### Changed

- Restart agents on SIGPIPE from systemd in systemd units [#945](https://github.com/buildkite/agent/pull/945) (@lox)

## [v3.9.1](https://github.com/buildkite/agent/tree/v3.9.1) (2019-03-06)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.9.0...v3.9.1)

### Changed

- Allow the Agent API to reject header times [#938](https://github.com/buildkite/agent/pull/938) (@sj26)
- Increase pipeline upload retries on 5xx errors [#937](https://github.com/buildkite/agent/pull/937) (@toolmantim)
- Pass experiment environment vars to bootstrap [#933](https://github.com/buildkite/agent/pull/933) (@lox)

## [v3.9.0](https://github.com/buildkite/agent/tree/v3.9.0) (2019-02-23)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.8.4...v3.9.0)

### Added

- Artifactory artifact support [#924](https://github.com/buildkite/agent/pull/924) (@GaryPWhite)
- Add a `--tag-from-gcp-labels` for loading agent tags from GCP [#930](https://github.com/buildkite/agent/pull/930) (@conorgil)
- Add a `--content-type` to `artifact upload` to allow specifying a content type [#912](https://github.com/buildkite/agent/pull/912) (@lox)
- Filter env used for command config out of environment [#908](https://github.com/buildkite/agent/pull/908) (@lox)
- If BUILDKITE_REPO is empty, skip checkout [#909](https://github.com/buildkite/agent/pull/909) (@lox)

### Changed

- Terminate bootstrap with unhandled signal after cancel [#890](https://github.com/buildkite/agent/pull/890) (@lox)

### Fixed

- Fix a race condition in cancellation [#928](https://github.com/buildkite/agent/pull/928) (@lox)
- Make sure checkout is removed on failure [#916](https://github.com/buildkite/agent/pull/916) (@lox)
- Ensure TempDir exists to avoid errors on windows [#915](https://github.com/buildkite/agent/pull/915) (@lox)
- Flush output immediately if timestamp-lines not on [#931](https://github.com/buildkite/agent/pull/931) (@lox)

## [v3.8.4](https://github.com/buildkite/agent/tree/v3.8.4) (2019-01-22)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.8.3...v3.8.4)

### Fixed

- Fix and test another seg fault in the artifact searcher [#901](https://github.com/buildkite/agent/pull/901) (@lox)
- Fix a seg-fault in the artifact uploader [#900](https://github.com/buildkite/agent/pull/900) (@lox)

## [v3.8.3](https://github.com/buildkite/agent/tree/v3.8.3) (2019-01-18)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.8.2...v3.8.3)

### Fixed

- Retry forever to upload job chunks [#898](https://github.com/buildkite/agent/pull/898) (@keithpitt)
- Resolve ssh hostname aliases before running ssh-keyscan [#889](https://github.com/buildkite/agent/pull/889) (@ticky)

## [v3.8.2](https://github.com/buildkite/agent/tree/v3.8.2) (2019-01-11)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.8.1...v3.8.2)

### Changed

- Fix another segfault in artifact download [#893](https://github.com/buildkite/agent/pull/893) (@lox)

## [v3.8.1](https://github.com/buildkite/agent/tree/v3.8.1) (2019-01-11)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.8.0...v3.8.1)

### Fixed

- Fixed two segfaults caused by missing loggers [#892](https://github.com/buildkite/agent/pull/892) (@lox)

## [v3.8.0](https://github.com/buildkite/agent/tree/v3.8.0) (2019-01-10)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.7.0...v3.8.0)

### Fixed

- Support absolute paths on Windows for config [#881](https://github.com/buildkite/agent/pull/881) (@petemounce)

### Changed

- Show log output colors on Windows 10+ [#885](https://github.com/buildkite/agent/pull/885) (@lox)
- Better cancel signal handling and error messages in output [#860](https://github.com/buildkite/agent/pull/860) (@lox)
- Use windows console groups for process management [#879](https://github.com/buildkite/agent/pull/879) (@lox)
- Support vendored plugins [#878](https://github.com/buildkite/agent/pull/878) (@lox)
- Show agent name in logger output [#880](https://github.com/buildkite/agent/pull/880) (@lox)
- Change git-clean-flags to cleanup submodules [#875](https://github.com/buildkite/agent/pull/875) (@lox)

## [v3.7.0](https://github.com/buildkite/agent/tree/v3.7.0) (2018-12-18)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.6.1...v3.7.0)

### Changed

- Fixed bug where submodules hosts weren't ssh keyscanned correctly [#876](https://github.com/buildkite/agent/pull/876) (@lox)
- Add a default port to metrics-datadog-host [#874](https://github.com/buildkite/agent/pull/874) (@lox)
- Hooks can now modify \$BUILDKITE_REPO before checkout to change the git clone or fetch address [#877](https://github.com/buildkite/agent/pull/877) (@sj26)
- Add a configurable cancel-grace-period [#700](https://github.com/buildkite/agent/pull/700) (@lox)
- Resolve BUILDKITE_COMMIT before pipeline upload [#871](https://github.com/buildkite/agent/pull/871) (@lox)

## [v3.6.1](https://github.com/buildkite/agent/tree/v3.6.1) (2018-12-13)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.6.0...v3.6.1)

### Added

- Add another search path for config file on Windows [#867](https://github.com/buildkite/agent/pull/867) (@petemounce)

### Fixed

- Exclude headers from timestamp-lines output [#870](https://github.com/buildkite/agent/pull/870) (@lox)

## [v3.6.0](https://github.com/buildkite/agent/tree/v3.6.0) (2018-12-04)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.5.4...v3.6.0)

### Fixed

- Fix bug that caused an extra log chunk to be sent in some cases [#845](https://github.com/buildkite/agent/pull/845) (@idledaemon)
- Don't retry checkout on build cancel [#863](https://github.com/buildkite/agent/pull/863) (@lox)
- Add buildkite-agent.cfg to docker images [#847](https://github.com/buildkite/agent/pull/847) (@lox)

### Added

- Experimental `--spawn` option to spawn multiple parallel agents [#590](https://github.com/buildkite/agent/pull/590) (@lox) - **Update:** This feature is now super stable.
- Add a linux/ppc64le build target [#859](https://github.com/buildkite/agent/pull/859) (@lox)
- Basic metrics collection for Datadog [#832](https://github.com/buildkite/agent/pull/832) (@lox)
- Added a `job update` command to make changes to a job [#833](https://github.com/buildkite/agent/pull/833) (@keithpitt)
- Remove the checkout dir if the checkout phase fails [#812](https://github.com/buildkite/agent/pull/812) (@lox)

### Changed

- Add tests around gracefully killing processes [#862](https://github.com/buildkite/agent/pull/862) (@lox)
- Removes process callbacks and moves them to job runner [#856](https://github.com/buildkite/agent/pull/856) (@lox)
- Use a channel to monitor whether process is killed [#855](https://github.com/buildkite/agent/pull/855) (@lox)
- Move to golang 1.11 [#839](https://github.com/buildkite/agent/pull/839) (@lox)
- Add a flag to disable http2 in the start command [#851](https://github.com/buildkite/agent/pull/851) (@lox)
- Use transparent for golang http2 transport [#849](https://github.com/buildkite/agent/pull/849) (@lox)

## [v3.5.4](https://github.com/buildkite/agent/tree/v3.5.4) (2018-10-24)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.5.3...v3.5.4)

### Fixed

- Prevent docker image from crashing with missing config error [#847](https://github.com/buildkite/agent/pull/847) (@lox)

## [v3.5.3](https://github.com/buildkite/agent/tree/v3.5.3) (2018-10-24)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.5.2...v3.5.3)

### Fixed

- Update to alpine to 3.8 in docker image [#842](https://github.com/buildkite/agent/pull/842) (@lox)
- Set BUILDKITE_AGENT_CONFIG in docker images to /buildkite [#834](https://github.com/buildkite/agent/pull/834) (@blakestoddard)
- Fix agent panics on ARM architecture [#831](https://github.com/buildkite/agent/pull/831) (@jhedev)

## [v3.5.2](https://github.com/buildkite/agent/tree/v3.5.2) (2018-10-09)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.5.1...v3.5.2)

### Changed

- Fix issue where pipelines with a top-level array of steps failed [#830](https://github.com/buildkite/agent/pull/830) (@lox)

## [v3.5.1](https://github.com/buildkite/agent/tree/v3.5.1) (2018-10-08)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.5.0...v3.5.1)

### Fixed

- Ensure plugin directory exists, otherwise checkout lock thrashes [#828](https://github.com/buildkite/agent/pull/828) (@lox)

## [v3.5.0](https://github.com/buildkite/agent/tree/v3.5.0) (2018-10-08)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.4.0...v3.5.0)

### Fixed

- Add plugin locking before checkout [#827](https://github.com/buildkite/agent/pull/827) (@lox)
- Ensure pipeline parser maintains map order in output [#824](https://github.com/buildkite/agent/pull/824) (@lox)
- Update aws sdk [#818](https://github.com/buildkite/agent/pull/818) (@sj26)
- Fix boostrap typo [#814](https://github.com/buildkite/agent/pull/814) (@ChefAustin)

### Changed

- `annotate` takes body as an arg, or reads from a pipe [#813](https://github.com/buildkite/agent/pull/813) (@sj26)
- Respect pre-set BUILDKITE_BUILD_CHECKOUT_PATH [#806](https://github.com/buildkite/agent/pull/806) (@lox)
- Add time since last successful heartbeat/ping [#810](https://github.com/buildkite/agent/pull/810) (@lox)
- Updating launchd templates to only restart on error [#804](https://github.com/buildkite/agent/pull/804) (@lox)
- Allow more time for systemd graceful stop [#819](https://github.com/buildkite/agent/pull/819) (@lox)

## [v3.4.0](https://github.com/buildkite/agent/tree/v3.4.0) (2018-07-18)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.3.0...v3.4.0)

### Changed

- Add basic plugin definition parsing [#748](https://github.com/buildkite/agent/pull/748) (@lox)
- Allow specifying which phases bootstrap should execute [#799](https://github.com/buildkite/agent/pull/799) (@lox)
- Warn in bootstrap when protected env are used [#796](https://github.com/buildkite/agent/pull/796) (@lox)
- Cancellation on windows kills bootstrap subprocesses [#795](https://github.com/buildkite/agent/pull/795) (@amitsaha)

## [v3.3.0](https://github.com/buildkite/agent/tree/v3.3.0) (2018-07-11)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.2.1...v3.3.0)

### Added

- Allow tags from the host to be automatically added with --add-host-tags [#791](https://github.com/buildkite/agent/pull/791) (@lox)
- Allow --no-plugins=false to force plugins on [#790](https://github.com/buildkite/agent/pull/790) (@lox)

## [v3.2.1](https://github.com/buildkite/agent/tree/v3.2.1) (2018-06-28)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.2.0...v3.2.1)

### Changed

- Remove the checkout dir when git clean fails [#786](https://github.com/buildkite/agent/pull/786) (@lox)
- Add a --dry-run to pipeline upload that dumps json [#781](https://github.com/buildkite/agent/pull/781) (@lox)
- Support PTY under OpenBSD [#785](https://github.com/buildkite/agent/pull/785) (@derekmarcotte) [#787](https://github.com/buildkite/agent/pull/787) (@derekmarcotte)
- Experiments docs and experiment cleanup [#771](https://github.com/buildkite/agent/pull/771) (@lox)

## [v3.2.0](https://github.com/buildkite/agent/tree/v3.2.0) (2018-05-25)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.1.2...v3.2.0)

### Changed

- Propagate exit code > 1 out of failing hooks [#768](https://github.com/buildkite/agent/pull/768) (@lox)
- Fix broken list parsing in cli arguments --tags and --experiments [#772](https://github.com/buildkite/agent/pull/772) (@lox)
- Add a virtual provides to the RPM package [#737](https://github.com/buildkite/agent/pull/737) (@jnewbigin)
- Clean up docker image building [#755](https://github.com/buildkite/agent/pull/755) (@lox)
- Don't trim whitespace from the annotation body [#766](https://github.com/buildkite/agent/pull/766) (@petemounce)

## [v3.1.2](https://github.com/buildkite/agent/tree/v3.1.2) (2018-05-10)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.1.1...v3.1.2)

### Changed

- Experiment: Pass jobs an authenticated unix socket rather than an access token [#759](https://github.com/buildkite/agent/pull/759) (@lox)
- Remove buildkite:git:branch meta-data [#753](https://github.com/buildkite/agent/pull/753) (@sj26)
- Set TERM and PWD for commands that get executed in shell [#751](https://github.com/buildkite/agent/pull/751) (@lox)

### Fixed

- Avoid pausing after job has finished [#764](https://github.com/buildkite/agent/pull/764) (@sj26)

## [v3.1.1](https://github.com/buildkite/agent/tree/v3.1.1) (2018-05-02)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.1.0...v3.1.1)

### Fixed

- Fix stdin detection for output redirection [#750](https://github.com/buildkite/agent/pull/750) (@lox)

## [v3.1.0](https://github.com/buildkite/agent/tree/v3.1.0) (2018-05-01)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0.1...v3.1.0)

### Changed

- Add ubuntu docker image [#749](https://github.com/buildkite/agent/pull/749) (@lox)
- Support `--no-interpolation` option in `pipeline upload` [#733](https://github.com/buildkite/agent/pull/733) (@lox)
- Bump our Docker image base to alpine v3.7 [#745](https://github.com/buildkite/agent/pull/745) (@sj26)
- Better error for multiple file args to artifact upload [#740](https://github.com/buildkite/agent/pull/740) (@toolmantim)

## [v3.0.1](https://github.com/buildkite/agent/tree/v3.0.1) (2018-04-17)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0.0...v3.0.1)

### Changed

- Don't set Content-Encoding on s3 upload [#732](@lox)

## [v3.0.0](https://github.com/buildkite/agent/tree/v3.0.0) (2018-04-03)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.44...v3.0.0)

No changes

## [v3.0-beta.44](https://github.com/buildkite/agent/tree/v3.0-beta.44) (2018-04-03)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.43...v3.0-beta.44)

### Fixed

- Normalize the `bootstrap-script` command using a new `commandpath` normalization [#714](https://github.com/buildkite/agent/pull/714) (@keithpitt)

### Changed

- Install windows binary to c:\buildkite-agent\bin [#713](https://github.com/buildkite/agent/pull/713) (@lox)

## [v3.0-beta.43](https://github.com/buildkite/agent/tree/v3.0-beta.43) (2018-04-03)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.42...v3.0-beta.43)

### Changed

- Prettier bootstrap output 💅🏻 [#708](https://github.com/buildkite/agent/pull/708) (@lox)
- Only run git submodule operations if there is a .gitmodules [#704](https://github.com/buildkite/agent/pull/704) (@lox)
- Add an agent config for no-local-hooks [#707](https://github.com/buildkite/agent/pull/707) (@lox)
- Build docker image as part of agent pipeline [#701](https://github.com/buildkite/agent/pull/701) (@lox)
- Windows install script [#699](https://github.com/buildkite/agent/pull/699) (@lox)
- Expose no-git-submodules config and arg to start [#698](https://github.com/buildkite/agent/pull/698) (@lox)

## [v3.0-beta.42](https://github.com/buildkite/agent/tree/v3.0-beta.42) (2018-03-20)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.41...v3.0-beta.42)

### Fixed

- Preserve types in pipeline.yml [#696](https://github.com/buildkite/agent/pull/696) (@lox)

## [v3.0-beta.41](https://github.com/buildkite/agent/tree/v3.0-beta.41) (2018-03-16)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.40...v3.0-beta.41)

### Added

- Retry failed checkouts [#670](https://github.com/buildkite/agent/pull/670) (@lox)

### Changed

- Write temporary batch scripts for Windows/CMD.EXE [#692](https://github.com/buildkite/agent/pull/692) (@lox)
- Enabling `no-command-eval` will also disable use of plugins [#690](https://github.com/buildkite/agent/pull/690) (@keithpitt)
- Support plugins that have a `null` config [#691](https://github.com/buildkite/agent/pull/691) (@keithpitt)
- Handle upgrading bootstrap-path from old 2.x shell script [#580](https://github.com/buildkite/agent/pull/580) (@lox)
- Show plugin commit if it's already installed [#685](https://github.com/buildkite/agent/pull/685) (@keithpitt)
- Handle windows paths in all usage of shellwords parsing [#686](https://github.com/buildkite/agent/pull/686) (@lox)
- Make NormalizeFilePath handle empty strings and windows [#688](https://github.com/buildkite/agent/pull/688) (@lox)
- Retry ssh-keyscans on error or blank output [#687](https://github.com/buildkite/agent/pull/687) (@keithpitt)
- Quote and escape env-file values [#682](https://github.com/buildkite/agent/pull/682) (@lox)
- Prevent incorrect corrupt git checkout detection on fresh checkout dir creation [#681](https://github.com/buildkite/agent/pull/681) (@lox)
- Only keyscan git/ssh urls [#675](https://github.com/buildkite/agent/pull/675) (@lox)
- Fail the job when no command is provided in the default command phase [#678](https://github.com/buildkite/agent/pull/678) (@keithpitt)
- Don't look for powershell hooks since we don't support them yet [#679](https://github.com/buildkite/agent/pull/679) (@keithpitt)
- Exit when artifacts can't be found for downloading [#676](https://github.com/buildkite/agent/pull/676) (@keithpitt)
- Run scripts via the shell, rather than invoking with exec [#673](https://github.com/buildkite/agent/pull/673) (@lox)
- Rename no-automatic-ssh-fingerprint-verification to no-ssh-keyscan [#671](https://github.com/buildkite/agent/pull/671) (@lox)

### Fixed

- Parse pipeline.yml env block in order [#668](https://github.com/buildkite/agent/pull/668) (@lox)
- Bootstrap shouldn't panic if plugin checkout fails [#672](https://github.com/buildkite/agent/pull/672) (@lox)

## [v3.0-beta.40](https://github.com/buildkite/agent/tree/v3.0-beta.40) (2018-03-07)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.39...v3.0-beta.40)

### Changed

- Commands are no longer written to temporary script files before execution [#648](https://github.com/buildkite/agent/pull/648) (@lox)
- Support more complex types in plugin config [#658](https://github.com/buildkite/agent/pull/658) (@lox)

### Added

- Write an env-file for the bootstrap [#643](https://github.com/buildkite/agent/pull/643) (@DazWorrall)
- Allow the shell interpreter to be configured [#648](https://github.com/buildkite/agent/pull/648) (@lox)

### Fixed

- Fix stdin detection on windows [#665](https://github.com/buildkite/agent/pull/665) (@lox)
- Check hook scripts get written to disk without error [#652](https://github.com/buildkite/agent/pull/652) (@sj26)

## [v3.0-beta.39](https://github.com/buildkite/agent/tree/v3.0-beta.39) (2018-01-31)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.38...v3.0-beta.39)

### Fixed

- Fix bug failing artifact upload glob would cause later globs to fail [\#620](https://github.com/buildkite/agent/pull/620) (@lox)
- Fix race condition in process management [\#618](https://github.com/buildkite/agent/pull/618) (@lox)
- Support older git versions for submodule commands [\#628](https://github.com/buildkite/agent/pull/628) (@lox)
- Lots of windows fixes and tests! [\#630](https://github.com/buildkite/agent/pull/630) [\#631](https://github.com/buildkite/agent/pull/631) [\#632](https://github.com/buildkite/agent/pull/632)

### Added

- Support for Bash for Windows for plugins and hooks! [\#636](https://github.com/buildkite/agent/pull/636) (@lox)
- Correct mimetypes for .log files [\#635](https://github.com/buildkite/agent/pull/635) (@DazWorrall)
- Usable Content-Disposition for GCE uploaded artifacts [\#640](https://github.com/buildkite/agent/pull/640) (@DazWorrall)
- Experiment for retrying checkout on failure [\#613](https://github.com/buildkite/agent/pull/613) (@lox)
- Skip local hooks when BUILDKITE_NO_LOCAL_HOOKS is set [\#622](https://github.com/buildkite/agent/pull/622) (@lox)

### Changed

- Bootstrap shell commands output stderr now [\#626](https://github.com/buildkite/agent/pull/626) (@lox)

## [v2.6.9](https://github.com/buildkite/agent/releases/tag/v2.6.9) (2018-01-18)

[Full Changelog](https://github.com/buildkite/agent/compare/v2.6.8...v2.6.9)

### Added

- Implement `BUILDKITE_CLEAN_CHECKOUT`, `BUILDKITE_GIT_CLONE_FLAGS` and `BUILDKITE_GIT_CLEAN_FLAGS` in bootstrap.bat [\#610](https://github.com/buildkite/agent/pull/610) (@solemnwarning)

### Fixed

- Fix unbounded memory usage in artifact uploads (#493)

## [v3.0-beta.38](https://github.com/buildkite/agent/tree/v3.0-beta.38) (2018-01-10)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.37...v3.0-beta.38)

### Fixed

- Fix bug where bootstrap with pty hangs on macOS [\#614](https://github.com/buildkite/agent/pull/614) (@lox)

## [v3.0-beta.37](https://github.com/buildkite/agent/tree/v3.0-beta.37) (2017-12-07)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.36...v3.0-beta.37)

### Fixed

- Fixed bug where agent uploads fail if no files match [\#600](https://github.com/buildkite/agent/pull/600) (@lox)
- Fixed bug where timestamps are incorrectly appended to header expansions [\#597](https://github.com/buildkite/agent/pull/597)

## [v3.0-beta.36](https://github.com/buildkite/agent/tree/v3.0-beta.36) (2017-11-23)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.35...v3.0-beta.36)

### Added

- Don't retry pipeline uploads on invalid pipelines [\#589](https://github.com/buildkite/agent/pull/589) (@DazWorrall)
- A vagrant box for windows testing [\#583](https://github.com/buildkite/agent/pull/583) (@lox)
- Binary is build with golang 1.9.2

### Fixed

- Fixed bug where malformed pipelines caused infinite loop [\#585](https://github.com/buildkite/agent/pull/585) (@lox)

## [v3.0-beta.35](https://github.com/buildkite/agent/tree/v3.0-beta.35) (2017-11-13)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.34...v3.0-beta.35)

### Added

- Support nested interpolated variables [\#578](https://github.com/buildkite/agent/pull/578) (@lox)
- Check for corrupt git repository before checkout [\#574](https://github.com/buildkite/agent/pull/574) (@lox)

### Fixed

- Fix bug where non-truthy bool arguments failed silently [\#582](https://github.com/buildkite/agent/pull/582) (@lox)
- Pass working directory changes between hooks [\#577](https://github.com/buildkite/agent/pull/577) (@lox)
- Kill cancelled tasks with taskkill on windows [\#575](https://github.com/buildkite/agent/pull/575) (@adill)
- Support hashed hosts in ssh known_hosts [\#579](https://github.com/buildkite/agent/pull/579) (@lox)

## [v3.0-beta.34](https://github.com/buildkite/agent/tree/v3.0-beta.34) (2017-10-19)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.33...v3.0-beta.34)

### Fixed

- Fix bug where pipeline upload doesn't get environment passed correctly [\#567](https://github.com/buildkite/agent/pull/567) (@lox)
- Only show "Running hook" if one exists [\#566](https://github.com/buildkite/agent/pull/566) (@lox)
- Fix segfault when using custom artifact bucket and EC2 instance role credentials [\#563](https://github.com/buildkite/agent/pull/563) (@sj26)
- Fix ssh keyscan of hosts with custom ports [\#565](https://github.com/buildkite/agent/pull/565) (@sj26)

## [v2.6.7](https://github.com/buildkite/agent/releases/tag/v2.6.7) (2017-11-13)

[Full Changelog](https://github.com/buildkite/agent/compare/v2.6.6...v2.6.7)

### Added

- Check for corrupt git repository before checkout [\#556](https://github.com/buildkite/agent/pull/556) (@lox)

### Fixed

- Kill cancelled tasks with taskkill on windows [\#571](https://github.com/buildkite/agent/pull/571) (@adill)

## [v2.6.6](https://github.com/buildkite/agent/releases/tag/v2.6.6) (2017-10-09)

[Full Changelog](https://github.com/buildkite/agent/compare/v2.6.5...v2.6.6)

### Fixed

- Backported new globbing library to fix "too many open files" during globbing [\#539](https://github.com/buildkite/agent/pull/539) (@sj26 & @lox)

## [v3.0-beta.33](https://github.com/buildkite/agent/tree/v3.0-beta.33) (2017-10-05)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.32...v3.0-beta.33)

### Added

- Interpolate env block before rest of pipeline.yml [\#552](https://github.com/buildkite/agent/pull/552) (@lox)

### Fixed

- Build hanging after git checkout [\#558](https://github.com/buildkite/agent/issues/558)

## [v3.0-beta.32](https://github.com/buildkite/agent/tree/v3.0-beta.32) (2017-09-25)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.31...v3.0-beta.32)

### Added

- Add --no-plugins option to agent [\#540](https://github.com/buildkite/agent/pull/540) (@lox)
- Support docker environment vars from v2 [\#545](https://github.com/buildkite/agent/pull/545) (@lox)

### Changed

- Refactored bootstrap to be more testable / maintainable [\#514](https://github.com/buildkite/agent/pull/514) [\#530](https://github.com/buildkite/agent/pull/530) [\#536](https://github.com/buildkite/agent/pull/536) [\#522](https://github.com/buildkite/agent/pull/522) (@lox)
- Add BUILDKITE_GCS_ACCESS_HOST for GCS Host choice [\#532](https://github.com/buildkite/agent/pull/532) (@jules2689)
- Prefer plugin, local, global and then default for hooks [\#549](https://github.com/buildkite/agent/pull/549) (@lox)
- Integration tests for v3 [\#548](https://github.com/buildkite/agent/pull/548) (@lox)
- Add docker integration tests [\#547](https://github.com/buildkite/agent/pull/547) (@lox)
- Use latest golang 1.9 [\#541](https://github.com/buildkite/agent/pull/541) (@lox)
- Faster globbing with go-zglob [\#539](https://github.com/buildkite/agent/pull/539) (@lox)
- Consolidate Environment into env package (@lox)

### Fixed

- Fix bug where ssh-keygen error causes agent to block [\#521](https://github.com/buildkite/agent/pull/521) (@lox)
- Pre-exit hook always fires now

## [v3.0-beta.31](https://github.com/buildkite/agent/tree/v3.0-beta.31) (2017-08-14)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.30...v3.0-beta.31)

### Fixed

- Support paths in BUILDKITE_ARTIFACT_UPLOAD_DESTINATION [\#519](https://github.com/buildkite/agent/pull/519) (@lox)

## [v3.0-beta.30](https://github.com/buildkite/agent/tree/v3.0-beta.30) (2017-08-11)

[Full Changelog](https://github.com/buildkite/agent/compare/v3.0-beta.29...v3.0-beta.30)

### Fixed

- Agent is prompted to verify remote server authenticity when cloning submodule from unkown host [\#503](https://github.com/buildkite/agent/issues/503)
- Windows agent cannot find git executable \(Environment variable/Path issue?\) [\#487](https://github.com/buildkite/agent/issues/487)
- ssh-keyscan doesn't work for submodules on a different host [\#411](https://github.com/buildkite/agent/issues/411)
- Fix boolean plugin config parsing [\#508](https://github.com/buildkite/agent/pull/508) (@toolmantim)

### Changed

- Stop making hook files executable [\#515](https://github.com/buildkite/agent/pull/515) (@yeungda-rea)
- Switch to yaml.v2 as the YAML parser [\#511](https://github.com/buildkite/agent/pull/511) (@keithpitt)
- Add submodule remotes to known_hosts [\#509](https://github.com/buildkite/agent/pull/509) (@lox)

## 3.0-beta.29 - 2017-07-18

### Added

- Added a `--timestamp-lines` option to `buildkite-agent start` that will insert RFC3339 UTC timestamps at the beginning of each log line. The timestamps are not applied to header lines. [#501](@lox)
- Ctrl-c twice will force kill the agent [#499](@lox)
- Set the content encoding on artifacts uploaded to s3 [#494] (thanks @airhorns)
- Output fetched commit sha during git fetch for pull request [#505](@sj26)

### Changed

- Migrate the aging goamz library to the latest aws-sdk [#474](@lox)

## 2.6.5 - 2017-07-18

### Added

- 🔍 Output fetched commit sha during git fetch for pull request [#505]

## 3.0-beta.28 - 2017-06-23

### Added

- 🐞 The agent will now poll the AWS EC2 Tags API until it finds some tags to apply before continuing. In some cases, the agent will start and connect to Buildkite before the tags are available. The timeout for this polling can be configured with --wait-for-ec2-tags-timeout (which defaults to 10 seconds) #492

### Fixed

- 🐛 Fixed 2 Windows bugs that caused all jobs that ran through our built-in buildkite-agent bootstrap command to fail #496

## 2.6.4 - 2017-06-16

### Added

- 🚀 The buildkite-agent upstart configuration will now source /etc/default/buildkite-agent before starting the agent process. This gives you an opportunity to configure the agent outside of the standard buildkite-agent.conf file

## 3.0-beta.27 - 2017-05-31

### Added

- Allow pipeline uploads when no-command-eval is true

### Fixed

- 🐞 Fixes to a few more edge cases when exported environment variables from hooks would include additional quotes #484
- Apt server misconfigured - `Packages` reports wrong sizes/hashes
- Rewrote `export -p` parser to support multiple line env vars

## 3.0-beta.26 - 2017-05-29

### Fixed

- 🤦‍♂️ We accidentally skipped a beta version, there's no v3.0-beta.25! Doh!
- 🐛 Fixed an issue where some environment variables exported from environment hooks would have new lines appended to the end

## 3.0-beta.24 - 2017-05-26

### Added

- 🚀 Added an --append option to buildkite-agent annotate that allows you to append to the body of an existing annotation

### Fixed

- 🐛 Fixed an issue where exporting multi-line environment variables from a hook would truncate everything but the first line

## 3.0-beta.23 - 2017-05-10

### Added

- 🚀 New command buildkite-agent annotate that gives you the power to annotate a build page with information from your pipelines. This feature is currently experimental and the CLI command API may change before an official 3.0 release

## 2.6.3 - 2017-05-04

### Added

- Added support for local and global pre-exit hooks (#466)

## 3.0-beta.22 - 2017-05-04

### Added

- Renames --meta-data to --tags (#435). --meta-data will be removed in v4, and v3 versions will now show a deprecation warning.
- Fixes multiple signals not being passed to job processes (#454)
- Adds binaries for OpenBSD (#463) and DragonflyBSD (#462)
- Adds support for local and global pre-exit hooks (#465)

## 2.6.2 - 2017-05-02

### Fixed

- Backport #381 to stable: Retries for fetching EC2 metadata and tags. #461

### Added

- Add OpenBSD builds

## 2.6.1 - 2017-04-13

### Removed

- Reverted #451 as it introduced a regression. Will re-think this fix and push it out again in another release after we do some more testing

## 3.0-beta.21 - 2017-04-13

### Removed

- Reverts the changes made in #448 as it seemed to introduce a regression. We'll rethink this change and push it out in another release.

## 2.6.0 - 2017-04-13

### Fixed

- Use /bin/sh rather than /bin/bash when executing commands. This allows use in environments that don't have bash, such as Alpine Linux.

## 3.0-beta.20 - 2017-04-13

### Added

- Add plugin support for HTTP repositories with .git extensions [#447]
- Run the global environment hook before checking out plugins [#445]

### Changed

- Use /bin/sh rather than /bin/bash when executing commands. This allows use in environments that don't have bash, such as Alpine Linux. (#448)

## 3.0-beta.19 - 2017-03-29

### Added

- `buildkite-agent start --disconnect-after-job` will run the agent, and automatically disconnect after running its first job. This has sometimes been referred to as "one shot" mode and is useful when you spin up an environment per-job and want the agent to automatically disconnect once it's finished its job
- `buildkite-agent start --disconnect-after-job-timeout` is the time in seconds the agent will wait for that first job to be assigned. The default value is 120 seconds (2 minutes). If a job isn't assigned to the agent after this time, it will automatically disconnect and the agent process will stop.

## 3.0-beta.18 - 2017-03-27

### Fixed

- Fixes a bug where log output would get sometimes get corrupted #441

## 2.5.1 - 2017-03-27

### Fixed

- Fixes a bug where log output would get sometimes get corrupted #441

## 3.0-beta.17 - 2017-03-23

### Added

- You can now specify a custom artifact upload destination with BUILDKITE_ARTIFACT_UPLOAD_DESTINATION #421
- git clean is now performed before and after the git checkout operation #418
- Update our version of lockfile which should fixes issues with running multiple agents on the same server #428
- Fix the start script for Debian wheezy #429
- The buildkite-agent binary is now built with Golang 1.8 #433
- buildkite-agent meta-data get now supports --default flag that allows you to return a default value instead of an error if the remote key doesn't exist #440

## [2.5] - 2017-03-23

### Added

- buildkite-agent meta-data get now supports --default flag that allows you to return a default value instead of an error if the remote key doesn't exist #440

## 2.4.1 - 2017-03-20

### Fixed

- 🐞 Fixed a bug where ^^^ +++ would be prefixed with a timestamp when ---timestamp-lines was enabled #438

## [2.4] - 2017-03-07

### Added

- Added a new option --timestamp-lines option to buildkite-agent start that will insert RFC3339 UTC timestamps at the beginning of each log line. The timestamps are not applied to header lines. #430

### Changed

- Go 1.8 [#433]
- Switch to govendor for dependency tracking [#432]
- Backport Google Cloud Platform meta-data to 2.3 stable agent [#431]

## 3.0-beta.16 - 2016-12-04

### Fixed

- "No command eval" mode now makes sure commands are inside the working directory 🔐
- Scripts which are already executable won't be chmoded 🔏

## 2.3.2 - 2016-11-28

### Fixed

- 🐝 Fixed an edge case that causes the agent to panic and exit if more lines are read a process after it's finished

## 2.3.1 - 2016-11-17

### Fixed

- More resilient init.d script (#406)
- Only lock if locks are used by the system
- More explicit su with --shell option

## 3.0-beta.15 - 2016-11-16

### Changed

- The agent now receives its "job status interval" from the Agent API (the number of seconds between checking if its current job has been remotely canceled)

## 3.0-beta.14 - 2016-11-11

### Fixed

- Fixed a race condition where the agent would pick up another job to run even though it had been told to gracefully stop (PR #403 by @grosskur)
- Fixed path to ssh-keygen for Windows (PR #401 by @bendrucker)

## [2.3] - 2016-11-10

### Fixed

- Fixed a race condition where the agent would pick up another job to run even though it had been told to gracefully stop (PR #403 by @grosskur)

## 3.0-beta.13 - 2016-10-21

### Added

- Refactored how environment variables are interpolated in the agent
- The buildkite-agent pipeline upload command now looks for .yaml files as well
- Support for the steps.json file has been removed

## 3.0-beta.12 - 2016-10-14

### Added

- Updated buildkite-agent bootstrap for Windows so that commands won't keep running if one of them fail (similar to Bash's set -e) behaviour #392 (thanks @essen)

## 3.0-beta.11 - 2016-10-04

### Added

- AWS EC2 meta-data tagging is now more resilient and will retry on failure (#381)
- Substring expansion works for variables in pipeline uploads, like \${BUILDKITE_COMMIT:0:7} will return the first seven characters of the commit SHA (#387)

## 3.0-beta.10 - 2016-09-21

### Added

- The buildkite-agent binary is now built with Golang 1.7 giving us support for macOS Sierra
- The agent now talks HTTP2 making calls to the Agent API that little bit faster
- The binary is a statically compiled (no longer requiring libc)
- meta-data-ec2 and meta-data-ec2-tags can now be configured using BUILDKITE_AGENT_META_DATA_EC2 and BUILDKITE_AGENT_META_DATA_EC2_TAGS environment variables

## [2.2] - 2016-09-21

### Added

- The buildkite-agent binary is now built with Golang 1.7 giving us support for macOS Sierra
- The agent now talks HTTP2 making calls to the Agent API that little bit faster
- The binary is a statically compiled (no longer requiring libc)
- meta-data-ec2 and meta-data-ec2-tags can now be configured using BUILDKITE_AGENT_META_DATA_EC2 and BUILDKITE_AGENT_META_DATA_EC2_TAGS environment variables

### Changed

- We've removed our dependency of libc for greater compatibly across \*nix systems which has had a few side effects:
  We've had to remove support for changing the process title when an agent starts running a job. This feature has only ever been available to users running 64-bit ubuntu, and required us to have a dependency on libc. We'd like to bring this feature back in the future in a way that doesn't have us relying on libc
- The agent will now use Golangs internal DNS resolver instead of the one on your system. This probably won't effect you in any real way, unless you've setup some funky DNS settings for agent.buildkite.com

## 3.0-beta.9 - 2016-08-18

### Added

- Allow fetching meta-data from Google Cloud metadata #369 (Thanks so much @grosskur)

## 2.1.17 - 2016-08-11

### Fixed

- Fix some compatibility with older Git versions 🕸

## 3.0-beta.8 - 2016-08-09

### Fixed

- Make bootstrap actually use the global command hook if it exists #365

## 3.0-beta.7 - 2016-08-05

### Added

- Support plugin array configs f989cde
- Include bootstrap in the help output 7524ffb

### Fixed

- Fixed a bug where we weren't stripping ANSI colours in build log headers 6611675
- Fix Content-Type for Google Cloud Storage API calls #361 (comment)

## 2.1.16 - 2016-08-04

### Fixed

- 🔍 SSH key scanning backwards compatibility with older openssh tools

## 2.1.15 - 2016-07-28

### Fixed

- 🔍 SSH key scanning fix after it got a little broken in 2.1.14, sorry!

## 2.1.14 - 2016-07-26

### Added

- 🔍 SSH key scanning should be more resilient, whether or not you hash your known hosts file
- 🏅 Commands executed by the Bootstrap script correctly preserve positional arguments and handle interpolation better
- 🌈 ANSI color sequences are a little more resilient
- ✨ Git clean and clone flags can now be supplied in the Agent configuration file or on the command line
- 📢 Docker Compose will now be a little more verbose when the Agent is in Debug mode
- 📑 $BUILDKITE_DOCKER_COMPOSE_FILE now accepts multiple files separated by a colon (:), like $PATH

## 3.0-beta.6 - 2016-06-24

### Fixed

- Fixes to the bootstrap when using relative paths #228
- Fixed hook paths on Windows #331
- Fixed default path of the pipeline.yml file on Windows #342
- Fixed issues surrounding long command definitions #334
- Fixed default bootstrap-command command for Windows #344

## 3.0-beta.5 - 2016-06-16

## [3.0-beta.3- 2016-06-01

### Added

- Added support for BUILDKITE_GIT_CLONE_FLAGS (#330) giving you the ability customise how the agent clones your repository onto your build machines. You can use this to customise the "depth" of your repository if you want faster clones BUILDKITE_GIT_CLONE_FLAGS="-v --depth 1". This option can also be configured in your buildkite-agent.cfg file using the git-clone-flags option
- BUILDKITE_GIT_CLEAN_FLAGS can now be configured in your buildkite-agent.cfg file using the git-clean-flags option (#330)
- Allow metadata value to be read from STDIN (#327). This allows you to set meta-data from files easier cat meta-data.txt | buildkite-agent meta-data set "foo"

### Fixed

- Fixed environment variable sanitisation #333

## 2.1.13 - 2016-05-30

### Added

- BUILDKITE_GIT_CLONE_FLAGS (#326) giving you the ability customise how the agent clones your repository onto your build machines. You can use this to customise the "depth" of your repository if you want faster clones `BUILDKITE_GIT_CLONE_FLAGS="-v --depth 1"`
- Allow metadata value to be read from STDIN (#327). This allows you to set meta-data from files easier `cat meta-data.txt | buildkite-agent meta-data set "foo"`

## 3.0-beta.2 - 2016-05-23

### Fixed

- Improved error logging when failing to capture the exit status for a job (#325)

## 2.1.12 - 2016-05-23

### Fixed

- Improved error logging when failing to capture the exit status for a job (#325)

## 2.1.11 - 2016-05-17

### Added

- A new --meta-data-ec2 command line flag and config option for populating agent meta-data from EC2 information (#320)
- Binaries are now published to download.buildkite.com (#318)

## 3.0-beta.1 - 2016-05-16

### Added

- New version number: v3.0-beta.1. There will be no 2.2 (the previous beta release)
- Outputs the build directory in the build log (#317)
- Don't output the env variable values that are set from hooks (#316)
- Sign packages with SHA512 (#308)
- A new --meta-data-ec2 command line flag and config option for populating agent meta-data from EC2 information (#314)
- Binaries are now published to download.buildkite.com (#318)

## 2.2-beta.4 - 2016-05-10

### Fixed

- Amazon Linux & CentOS 6 packages now start and shutdown the agent gracefully (#306) - thanks @jnewbigin
- Build headers now work even if ANSI escape codes are applied (#279)

## 2.1.10- 2016-05-09

### Fixed

- Amazon Linux & CentOS 6 packages now start and shutdown the agent gracefully (#290 #305) - thanks @jnewbigin

## 2.1.9 - 2016-05-06

### Added

- Docker Compose 1.7.x support, including docker network removal during cleanup (#300)
- Docker Compose builds now specify --pull, so base images will always attempted to be pulled (#300)
- Docker Compose command group is now expanded by default (#300)
- Docker Compose builds now only build the specified service’s image, not all images. If you want to build all set the environment variable BUILDKITE_DOCKER_COMPOSE_BUILD_ALL=true (#297)
- Step commands are now run with bash’s -o pipefail option, preventing silent failures (#301)

### Fixed

- BUILDKITE_DOCKER_COMPOSE_LEAVE_VOLUMES undefined errors in bootstrap.sh have been fixed (#283)
- Build headers now work even if ANSI escape codes are applied

## 2.2-beta.3 - 2016-03-18

### Addeed

- Git clean brokenness has been fixed in the Go-based bootstrap (#278)

## 2.1.8- 2016-03-18

### Added

- BUILDKITE_DOCKER_COMPOSE_LEAVE_VOLUMES (#274) which allows you to keep the docker-compose volumes once a job has been run

## 2.2-beta.2 - 2016-03-17

### Added

- Environment variable substitution in pipeline files (#246)
- Google Cloud Storage for artifacts (#207)
- BUILDKITE_DOCKER_COMPOSE_LEAVE_VOLUMES (#252) which allows you to keep the docker-compose volumes once a job has been run
- BUILDKITE_S3_ACCESS_URL (#261) allowing you set your own host for build artifact links. This means you can set up your own proxy/web host that sits in front of your private S3 artifact bucket, and click directly through to them from Buildkite.
- BUILDKITE_GIT_CLEAN_FLAGS (#270) allowing you to ensure all builds have completely clean checkouts using an environment hook with export BUILDKITE_GIT_CLEAN_FLAGS=-fqdx
- Various new ARM builds (#258) allowing you to run the agent on services such as Scaleway

### Fixed

- Increased many of the HTTP timeouts to ease the stampede on the agent endpoint (#259)
- Corrected bash escaping errors which could cause problems for installs to non-standard paths (#262)
- Made HTTPS the default for all artifact upload URLs (#265)
- Added Buildkite's bin dir to the end, not the start, of \$PATH (#267)
- Ensured that multiple commands separated by newlines fail as soon as a command fails (#272)

## 2.1.7- 2016-03-17

### Added

- Added support for BUILDKITE_S3_ACCESS_URL (#247) allowing you set your own host for build artifact links. This means you can set up your own proxy/web host that sits in front of your private S3 artifact bucket, and click directly through to them from Buildkite.
- Added support for BUILDKITE_GIT_CLEAN_FLAGS (#271) allowing you to ensure all builds have completely clean checkouts using an environment hook with export BUILDKITE_GIT_CLEAN_FLAGS=-fqdx
- Added support for various new ARM builds (#263) allowing you to run the agent on services such as Scaleway

### Fixed

- Updated to Golang 1.6 (26d37c5)
- Increased many of the HTTP timeouts to ease the stampede on the agent endpoint (#260)
- Corrected bash escaping errors which could cause problems for installs to non-standard paths (#266)
- Made HTTPS the default for all artifact upload URLs (#269)
- Added Buildkite's bin dir to the end, not the start, of \$PATH (#268)
- Ensured that multiple commands separated by newlines fail as soon as a command fails (#273)

## 2.1.6.1 - 2016-03-09

### Fixed

- The agent is now statically linked to glibc, which means support for Debian 7 and friends (#255)

## 2.1.6 - 2016-03-03

### Fixed

- git fetch --tags doesn't fetch branches in old git (#250)

## 2.1.5 2016-02-26

### Fixed

- Use TrimPrefix instead of TrimLeft (#203)
- Update launchd templates to use .buildkite-agent dir (#212)
- Link to docker agent in README (#225)
- Send desired signal instead of always SIGTERM (#215)
- Bootstrap script fetch logic tweaks (#243)
- Avoid upstart on Amazon Linux (#244)

## 2.2-beta.1 2015-10-20

### Changed

- Added some tests to the S3Downloader

## 2.1.4 - 2015-10-16

### Fixed

- yum.buildkite.com now shows all previous versions of the agent

## 2.1.3 - 2015-10-16

### Fixed

- Fixed problem with bootstrap.sh not resetting git checkouts correctly

## 2.1.2 - 2015-10-16

### Fixed

- Removed unused functions from the bootstrap.sh file that was causing garbage output in builds
- FreeBSD 386 machines are now supported

## 2.1.1 - 2015-10-15

### Fixed

- Fixed issue with starting the bootstrap.sh file on linux systems fork/exec error

## [2.1] - 2015-10-15

## 2.1-beta.3 - 2015-10-01

### Changed

- Added support for FreeBSD - see instructions here: https://gist.github.com/keithpitt/61acb5700f75b078f199
- Only fetch the required branch + commit when running a build
- Added support for a repository command hook
- Change the git origin when a repository URL changes
- Improved mime type coverage for artefacts
- Added support for pipeline.yml files, starting to deprecate steps.json
- Show the UUID in the log output when uploading artifacts
- Added graceful shutdown #176
- Fixed header time and artifact race conditions
- OS information is now correctly collected on Windows

## 2.1-beta.2 - 2015-08-04

### Fixed

- Optimised artifact state updating
- Dump artifact upload responses when --debug-http is used

## 2.1-beta.1 - 2015-07-30

### Fixed

- Debian packages now include the debian_version property 📦
- Artifacts are uploaded faster! We've optimised our Agent API payloads to have a smaller footprint meaning you can uploading more artifacts faster! 🚗💨
- You can now download artifacts from private S3 buckets using buildkite-artifact download ☁️
- The agent will now change its process title on linux/amd64 machines to report its current status: `buildkite-agent v2.1 (my-agent-name) [job a4f-a4fa4-af4a34f-af4]`

## 2.1-beta - 2015-07-3

## 2.0.4 - 2015-07-2

### Fixed

- Changed the format that --version returns buildkite-agent version 2.0.4, build 456 🔍

### Added

- Added post-artifact global and local hooks 🎣

## 2.0.3.761 - 2015-07-21

### Fixed

- Debian package for ARM processors
- Include the build number in the --version call

## 2.0.3 - 2015-07-21

## 2.0.1 - 2015-07-17

## [2.0] - 2015-07-14

### Added

- The binary name has changed from buildbox to buildkite-agent
- The default install location has changed from ~/.buildbox to ~/.buildkite-agent (although each installer may install in different locations)
- Agents can be configured with a config file
- Agents register themselves with a organization-wide token, you no longer need to create them via the web
- Agents now have hooks support and there should be no reason to customise the bootstrap.sh file
- There is built-in support for containerizing builds with Docker and Docker Compose
- Windows support
- There are installer packages available for most systems
- Agents now have meta-data
- Build steps select agents using key/value patterns rather than explicit agent selection
- Automatic ssh fingerprint verification
- Ability to specify commands such as rake and make instead of a path to a script
- Agent meta data can be imported from EC2 tags
- You can set a priority for the agent
- The agent now works better under flakey internet connections by retrying certain API calls
- A new command buildkite-agent artifact shasum that allows you to download the shasum of a previously uploaded artifact
- Various bug fixes and performance enhancements
- Support for storing build pipelines in repositories
