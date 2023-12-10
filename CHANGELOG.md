# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Added a boolify filter that evaluates strings, numbers and bools as boolean (i.e. 'on', 'false', 'off', '0')
- Added an option to the parse commands to set the template directory to resolve includes

### Modified
- Resolve relative paths in the parse:file command against the template directory is specified
- Resolve a relative source dir in the parse:dir command against the template directory is specified
- If no template dir is specified in parse:file, the directory of the file is used as the templates dir
- If no template dir is specified in parse:dir, the directory of the source dir is used as the templates dir

## [1.3] - 2023-12-07

### Fixed
- Fixed a bug where .ini.twig files resulted in a .in file extension when parsing dirs.

## [1.2] - 2023-12-01

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
