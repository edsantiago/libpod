- [Release notes for](#release-notes-for)
- [Changelog since](#changelog-since)
  - [Changes by Kind](#changes-by-kind)
    - [Feature](#feature)
    - [Bug or Regression](#bug-or-regression)
    - [Uncategorized](#uncategorized)
  - [Dependencies](#dependencies)
    - [Added](#added)
    - [Changed](#changed)
    - [Removed](#removed)

# Release notes for 

[Documentation](https://docs.k8s.io/docs/home)
# Changelog since 

## Changes by Kind

### Feature

- Add new noquota option for volumes ([#14051](https://github.com/containers/podman/pull/14051), [@giuseppe](https://github.com/giuseppe))

### Bug or Regression

- Podman-remote no longer joins user NS ([#14010](https://github.com/containers/podman/pull/14010), [@vrothberg](https://github.com/vrothberg))

### Uncategorized

- Improved autocompletion for some commands ([#14053](https://github.com/containers/podman/pull/14053), [@Luap99](https://github.com/Luap99))
- Results from `podman search` are now truncated by default ([#14047](https://github.com/containers/podman/pull/14047), [@vrothberg](https://github.com/vrothberg))

## Dependencies

### Added
_Nothing has changed._

### Changed
- github.com/ProtonMail/go-crypto: [70ae35b → a948124](https://github.com/ProtonMail/go-crypto/compare/70ae35b...a948124)
- github.com/checkpoint-restore/checkpointctl: [54b4ebf → 33f4a66](https://github.com/checkpoint-restore/checkpointctl/compare/54b4ebf...33f4a66)
- github.com/container-orchestrated-devices/container-device-interface: [v0.3.2 → v0.4.0](https://github.com/container-orchestrated-devices/container-device-interface/compare/v0.3.2...v0.4.0)
- github.com/containernetworking/cni: [v1.0.1 → v1.1.0](https://github.com/containernetworking/cni/compare/v1.0.1...v1.1.0)
- github.com/containers/buildah: [5b8e791 → 8f2bb88](https://github.com/containers/buildah/compare/5b8e791...8f2bb88)
- github.com/containers/common: [49f1a40 → 4081e6b](https://github.com/containers/common/compare/49f1a40...4081e6b)
- github.com/containers/image/v5: [d1b6468 → be08568](https://github.com/containers/image/v5/compare/d1b6468...be08568)
- github.com/containers/storage: [eea4e0f → 8996869](https://github.com/containers/storage/compare/eea4e0f...8996869)
- github.com/fsnotify/fsnotify: [v1.5.2 → v1.5.4](https://github.com/fsnotify/fsnotify/compare/v1.5.2...v1.5.4)
- github.com/magefile/mage: [v1.13.0 → v1.12.1](https://github.com/magefile/mage/compare/v1.13.0...v1.12.1)
- github.com/onsi/gomega: [v1.16.0 → v1.19.0](https://github.com/onsi/gomega/compare/v1.16.0...v1.19.0)
- github.com/sylabs/sif/v2: [v2.4.2 → v2.6.0](https://github.com/sylabs/sif/v2/compare/v2.4.2...v2.6.0)
- golang.org/x/sys: a9b59b0 → 33da011

### Removed
_Nothing has changed._
