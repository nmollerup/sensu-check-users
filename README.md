[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/nmollerup/sensu-check-users)
![Go Test](https://github.com/nmollerup/sensu-check-users/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/nmollerup/sensu-check-users/workflows/goreleaser/badge.svg)

# Sensu logged in users check

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

The Sensu logged in users check is a [Sensu Check][1] that alerts on the
number of users currently logged into the system. It is a Go port of the
Ruby `check-users.rb` plugin from [sensu-checks-jppol-ruby][6], which shelled
out to `who | wc -l`. This check instead reads the system's login records
directly (`/var/run/utmp` and equivalent), with no dependency on external
commands.

## Usage examples

```
Check number of logged in users

Usage:
  check-users [flags]
  check-users [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --critical int   Number of logged in users that triggers a critical
  -h, --help           help for check-users
  -w, --warning int    Number of logged in users that triggers a warning

Use "check-users [command] --help" for more information about a command.
```

Both `--warning` and `--critical` are required, and `--critical` must be
greater than or equal to `--warning`.

```
check-users --warning 5 --critical 10
```

## Configuration

### Asset registration

[Sensu Assets][2] are the best way to make use of this plugin. If you're not
using an asset, please consider doing so! If you're using sensuctl 5.13 with
Sensu Backend 5.13 or later, you can use the following command to add the asset:

```
sensuctl asset add nmollerup/sensu-check-users
```

If you're using an earlier version of sensuctl, you can find the asset on the
[Bonsai Asset Index][3].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-users
  namespace: default
spec:
  command: >-
    sensu-check-users
    --warning 5
    --critical 10
  subscriptions:
  - system
  runtime_assets:
  - nmollerup/sensu-check-users
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
Asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable from this
source.

From the local path of the sensu-check-users repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][4].

[1]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[2]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[3]: https://bonsai.sensu.io/assets/nmollerup/sensu-check-users
[4]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[6]: https://git.rootdom.dk/KIT-Online/sensu-checks-jppol-ruby
