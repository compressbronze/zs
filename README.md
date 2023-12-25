Zendesk Search
===

# Installation
Ensuring you have the [go toolchain v1.21+](https://go.dev/doc/install) installed, run
```shell
go install github.com/satrap-illustrations/zs@main
```
If your `PATH` includes `$GOPATH/bin` or `$GOBIN`, then you should be run the program in a terminal by executing
```shell
zs
```

Alternatively, you may checkout the source, `cd` into the checkout directory and run
```shell
go run .
```

# Usage
```
Zendesk Search (zs)

It searches Zendesk.

Usage:
  zs [flags]

Flags:
      --config string     config file (default is $HOME/config/zs/config.yaml)
  -d, --data-dir string   data directory (default "./data")
  -h, --help              help for zs
  -v, --verbose           verbose output
      --version           version for zs
```
This launches a terminal user interface (tui) that explains the available features.

If you run the program as a binary, you will most likely need to set the data directory that contains JSON files to search. The files `organaization.json`, `tickets.json`, and `users.json` MUST be present in that directory.
