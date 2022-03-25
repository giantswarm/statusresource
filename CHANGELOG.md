# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Update apiextensions to v6.

## [4.0.1] - 2022-01-25

### Fixed

- Correct module major version from v3 to v4.

## [4.0.0] - 2022-01-25

### Changed

- Update `operatorkit` and `k8sclient` libraries both to `v7`.

## [3.1.0] - 2021-02-22

### Added

- Add additional error handling for timeouts and not-ready API servers.

## [3.0.0] - 2020-10-27

### Changed

- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.
- Prepare module v3.

## [2.0.0] - 2020-08-11

### Changed

- Update Kubernetes dependencies to v1.18.5.

## [0.4.0] 2020-06-18

### Changed

- Update `k8sclient` dependency to `3.1.0`.
- Update `tenantcluster` dependency to `2.0.0`.


## [0.3.0] 2020-04-20

### Changed

- Update `apiextensions` dependency to `v0.3.0`.
  - Notable change: Move from `DeepCopyTime()` to k8s built-in `metav1.Time`.


## [0.2.1] 2020-04-17

### Changed

- Update `errors` dependency.



## [0.2.0] 2020-03-25

### Changed

- Switch from dep to Go modules.
- Use architect orb.



## [0.1.0] 2020-03-25

### Added

- First release.



[Unreleased]: https://github.com/giantswarm/statusresource/compare/v4.0.1...HEAD
[4.0.1]: https://github.com/giantswarm/statusresource/compare/v4.0.0...v4.0.1
[4.0.0]: https://github.com/giantswarm/statusresource/compare/v3.1.0...v4.0.0
[3.1.0]: https://github.com/giantswarm/statusresource/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/giantswarm/statusresource/compare/v2.0.0...v3.0.0
[2.0.0]: https://github.com/giantswarm/statusresource/compare/v0.4.0...v2.0.0
[0.4.0]: https://github.com/giantswarm/statusresource/compare/v0.3.0...v0.4.0

[0.3.0]: https://github.com/giantswarm/statusresource/compare/v0.2.1...v0.3.0

[0.2.1]: https://github.com/giantswarm/statusresource/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/giantswarm/statusresource/compare/v0.1.0...v0.2.0

[0.1.0]: https://github.com/giantswarm/statusresource/releases/tag/v0.1.0
