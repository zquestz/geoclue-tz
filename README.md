# geoclue-tz

[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL]

Write geoclue /etc/geolocation based on tz zone info.

```text
Usage:
  geoclue-tz [flags]

Flags:
      --default-accuracy float32    default accuracy
      --default-altitude float32    default altitude
      --default-latitude float32    default latitude
      --default-longitude float32   default longitude
  -d, --dry-run                     dry run mode
  -h, --help                        help for geoclue-tz
  -v, --verbose                     verbose mode
      --version                     display version
```

## Install

Make sure that `GOPATH` and `GOBIN` env vars are set. Then run:

```zsh
go install github.com/zquestz/geoclue-tz@latest
```

## Configuration

To setup your own default configuration just create `/etc/geoclue-tz.conf`. The configuration file is in UCL format.

For more information about UCL visit:
[https://github.com/vstakhov/libucl](https://github.com/vstakhov/libucl)

The following keys are supported:

* defaultLatitude (default latitude if match isn't found)
* defaultLongitude (default longitude if match isn't found)
* defaultAltitude (default altitude if match isn't found)
* defaultAccuracy (default accuracy if match isn't found)
* verbose (verbose mode)
* dryRun (dry run mode)

## License

geoclue-tz is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-URL]: http://goreportcard.com/report/zquestz/geoclue-tz
[ReportCard-Image]: https://goreportcard.com/badge/github.com/zquestz/geoclue-tz
[Build-Status-URL]: https://app.travis-ci.com/github/zquestz/geoclue-tz
[Build-Status-Image]: https://app.travis-ci.com/zquestz/geoclue-tz.svg?branch=main
