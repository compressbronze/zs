Zendesk Search
===

# Installation
Ensuring you have the [go toolchain](https://go.dev/doc/install) installed, run
```shell
go install github.com/satrap-illustrations/zs@main
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

If you run the program as a binary, you will most likely need to set the data directory that contains JSON files to search.
