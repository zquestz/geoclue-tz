# geoclue-tz

[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL]

Generate geoclue /etc/geolocation based on tz zone info.

```text
Usage:
  geoclue-tz [flags]

Flags:
  -d, --dry-run           dry run debug mode
  -h, --help              help for geoclue-tz
  -l, --location string   enable custom location
      --version           display version
```

## Install

Make sure that `GOPATH` and `GOBIN` env vars are set. Then run:

```zsh
go install github.com/zquestz/geoclue-tz@latest
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

```
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

## License

geoclue-tz is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-URL]: http://goreportcard.com/report/zquestz/geoclue-tz
[ReportCard-Image]: https://goreportcard.com/badge/github.com/zquestz/geoclue-tz
[Build-Status-URL]: https://app.travis-ci.com/github/zquestz/geoclue-tz
[Build-Status-Image]: https://app.travis-ci.com/zquestz/geoclue-tz.svg?branch=main
