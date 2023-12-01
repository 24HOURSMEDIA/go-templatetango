# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2]

### Added
- Added support for include, embed, macros, and layouts.
- Add support for parsing files in an entire directory and outputting them to another directory.

### Fixed
- Fixed some small errors in the docs

## [1.1] - 2023-11-30

### Added

- Filters for use in json templates: json_value, json_casted_value, json_escape
- Filters for url: rawurlencode (path escaping
- json_decode filter allows passing json strings in environment variables and returning them as maps or arrays

## [1.0] - 2023-11-24

### Added

- Initial release.
