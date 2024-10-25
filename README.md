# geoclue-tz

[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL] [![Release][Release-Image]][Release-URL]

Generate geoclue /etc/geolocation based on the current time zone.

```text
Usage:
  geoclue-tz [flags]

Flags:
  -c, --completion string   completion script for bash, zsh or fish
  -d, --dry-run             dry run debug mode
  -h, --help                help for geoclue-tz
  -l, --location string     enable custom location
      --version             display version
```

## Usage

Update `/etc/geolocation` based on the current time zone.

```zsh
geoclue-tz
```

Update `/etc/geolocation` based on your custom home location.

```zsh
geoclue-tz -l home
```

## Install

Make sure that `GOPATH` and `GOBIN` env vars are set. Then run:

```zsh
go install github.com/zquestz/geoclue-tz@latest
```

Arch Linux users can install from the AUR:

```zsh
yay -S geoclue-tz
```

## Configuration

To setup your own configuration just create `/etc/geoclue-tz.conf`. The configuration file is in UCL format. This makes it super easy to set the values for your custom locations, and restore them whenever you want.

For more information about UCL visit:
[https://github.com/vstakhov/libucl](https://github.com/vstakhov/libucl)

The following keys are supported:

* locations - array (custom locations)
* verbose - bool (verbose mode)
* dryRun - bool (dry run mode)

Here is a sample configuration, with a single custom location. The only required keys are `latitude`, `longitude`, and `name`.

```text
locations [
  {
    latitude = 19.520960
    longitude = 155.920517
    altitude = 0
    accuracy = 1000
    name = home
  }
]
```

## Shell Autocompletion

To set up autocompletion:

### Bash Linux

```zsh
geoclue-tz --completion bash > /etc/bash_completion.d/geoclue-tz
```

### Bash MacOS

```zsh
geoclue-tz --completion bash > /usr/local/etc/bash_completion.d/geoclue-tz
```

### Zsh

Generate a `_geoclue-tz` completion script and put it somewhere in your `$fpath`:

```zsh
geoclue-tz --completion zsh > /usr/local/share/zsh/site-functions/_geoclue-tz
```

### Fish

```zsh
geoclue-tz --completion fish > ~/.config/fish/completions/geoclue-tz.fish
```

## License

geoclue-tz is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-URL]: http://goreportcard.com/report/zquestz/geoclue-tz
[ReportCard-Image]: https://goreportcard.com/badge/github.com/zquestz/geoclue-tz
[Build-Status-URL]: https://app.travis-ci.com/github/zquestz/geoclue-tz
[Build-Status-Image]: https://app.travis-ci.com/zquestz/geoclue-tz.svg?branch=main
[Release-URL]: https://github.com/zquestz/geoclue-tz/releases/tag/v1.0.0
[Release-Image]: http://img.shields.io/badge/geoclue-tz-v1.0.0-1eb0fc.svg
