# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][1] and this project adheres to [Semantic Versioning][2].

## [Unreleased]

### Added
- Initial release of `check-users` — a Go port of `check-users.rb` from [sensu-checks-jppol-ruby][3]
- Counts currently logged in users via the system login records (no dependency on the `who` binary)
- Configurable, required warning and critical thresholds

[1]: https://keepachangelog.com/en/1.0.0/
[2]: https://semver.org/spec/v2.0.0.html
[3]: https://git.rootdom.dk/KIT-Online/sensu-checks-jppol-ruby
