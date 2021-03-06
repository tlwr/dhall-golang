# Changelog

## [Unreleased]

[Unreleased]: https://github.com/philandstuff/dhall-golang/compare/v1.0.0-rc.2...HEAD

## [1.0.0-rc.2] - 2020-02-16

[1.0.0-rc.2]: https://github.com/philandstuff/dhall-golang/compare/v1.0.0-rc.1...v1.0.0-rc.2

Another release candidate.  As promised, the `core` package is still
in flux and has undergone a huge refactor in this release.  Along with
that, the godoc has been vastly improved, and a new README has been
written to replace the previous scrappy development notes.

Also, this brings dhall-golang up to version 14.0.0 of the language
standard.

### Breaking changes

 - refactoring of the `core` package
   - moved `core.Term` and implementations to new package `term`
   - removed `core.AppValue` from public interface
   - renamed various types to remove `-Val` and `-Term` suffixes
 - (from Dhall 14.0.0): decimal Natural literals can no longer have
   leading 0 digits

### Changed

 - dhall-golang now supports [version 14.0.0][dhall-14.0.0] of the
   language.

[dhall-14.0.0]: https://github.com/dhall-lang/dhall-lang/releases/tag/v14.0.0

### Fixed

 - `dhall.Unmarshal()` now resolves imports and typechecks before
   evaluating

### Added

  - added `core.ListOf`, `core.OptionalOf`, `core.NoneOf` Value types
    to represent `List a`, `Optional a` and `None a` Values
    respectively

## [1.0.0-rc.1] - 2020-02-09

### Changed

 - Fixed a compile error in cbor.go :/

[1.0.0-rc.1]: https://github.com/philandstuff/dhall-golang/compare/v1.0.0-rc.0...v1.0.0-rc.1

## [1.0.0-rc.0] - 2020-02-09

First release candidate.  Note that some things are still in flux and
subject to change:

 - The `dhall` package is stable and will not have any breaking
   changes.  In particular, `dhall.Decode` and `dhall.Unmarshal` will
   not have any breaking changes before a v1.0.0 release.
 - The `parser` package is also stable and will not have any breaking
   changes.
 - The `core` package is still subject to change: in particular, names
   which are currently exported may be unexported before a v1.0.0
   release.

### Added

- Core Dhall functionality:
  - Parse Dhall source to Terms
  - Resolve Dhall imports
  - Use Dhall cache for imports
  - Typecheck Dhall Terms
  - Evaluate Dhall Terms to Values
  - Marshalling/unmarshalling to CBOR format
- Go bindings:
  - dhall.Decode to decode a Dhall Value into a Go variable
  - dhall.Unmarshal as a convenience all-in-one
    Dhall-source-to-Go-variable function

[1.0.0-rc.0]: https://github.com/philandstuff/dhall-golang/releases/tag/v1.0.0-rc.0
