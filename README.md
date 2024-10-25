# geoclue-tz

[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL]

Write geoclue /etc/geolocation based on tz zone info.

```text
Usage:
  geoclue-tz [flags]

Flags:
  -d, --dry-run                  dry run mode
  -h, --help                     help for geoclue-tz
      --home                     enable home configuration
      --home-accuracy float32    home accuracy
      --home-altitude float32    home altitude
      --home-latitude float32    home latitude
      --home-longitude float32   home longitude
  -v, --verbose                  verbose mode
      --version                  display version
```

## Install

Make sure that `GOPATH` and `GOBIN` env vars are set. Then run:

```zsh
go install github.com/zquestz/geoclue-tz@latest
```

## Configuration

To setup your own home configuration just create `/etc/geoclue-tz.conf`. The configuration file is in UCL format. This makes it super easy to set the values for your home, and restore them whenever you want.

For more information about UCL visit:
[https://github.com/vstakhov/libucl](https://github.com/vstakhov/libucl)

The following keys are supported:

* home - bool (enable home configuration)
* homeLatitude - float32 (home latitude)
* homeLongitude - float32 (home longitude)
* homeAltitude - float32 (home altitude)
* homeAccuracy - float32 (home accuracy)
* verbose - bool (verbose mode)
* dryRun - bool (dry run mode)

## License

geoclue-tz is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-URL]: http://goreportcard.com/report/zquestz/geoclue-tz
[ReportCard-Image]: https://goreportcard.com/badge/github.com/zquestz/geoclue-tz
[Build-Status-URL]: https://app.travis-ci.com/github/zquestz/geoclue-tz
[Build-Status-Image]: https://app.travis-ci.com/zquestz/geoclue-tz.svg?branch=main
