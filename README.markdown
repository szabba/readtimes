# `readtimes`

`readtimes` reads out the current time aloud following a cron-like schedule.

Windows is the only supported OS.

## Build

```bash
$ go build .
```

This will succeed on unsupported platforms.
However the program will exit promptly reporting the lack of support.

> **NOTE:** On Windows you might want to add `-ldflags="-H windowsgui"` to the CLI.
>
> ```bash
> $ go build -ldflags="-H windowsgui" .
> ```

## Running

```bash
$ ./readtime
```

By default:

* time is read out aloud in a 24-hour format,

* time is read out at every full hour.


## Configuration

The following command will print out the supported CLI flags.

```bash
$ ./readtime -help

```

## Contributing

The project does not accept contributions to the source repository content.

## Copyright

Karol Marcjan &copy; 2024

Distributed under the European Public License.
See [`LICENSE`] for more details.

[`LICENSE`]: ./LICENSE