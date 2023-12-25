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
see the Usage section below on how point it to the data.

Alternatively, you may clone the repo and run it with Go:
```shell
git clone https://github.com/compressbronze/zs
cd zs
go run .
```

# Usage
```
Zendesk Search (zs)

It searches Zendesk.

Usage:
  zs [flags]

Flags:
      --config string       config file (default is $HOME/.config/zs/config.yaml)
  -d, --data-dir string     data directory (default "./data")
      --debug-file string   debug log file
  -h, --help                help for zs
  -v, --version             version for zs
```
This launches a terminal user interface (tui) that explains the available features.

The files `organaization.json`, `tickets.json`, and `users.json` MUST be present in that data directory.

# Demo
<img width="1200" src="./demo/demo.gif" />

# Design

## Data
The data has been abstracted behind the interface `stores.Stores`.

Initially, I considered creating a store for each document type (organization, ticket, user), so you will see stores for each type.
However, it turned out the code for each document type store is almost identical, the main difference being that organizations and users use integer keys, while tickets has UUIDs.

So I've combined the implementation of each document type store in the `implementations` package, which provides 2 implementations of `stores.Store`.
The test are all driven through this interface as well, so there are no tests for each document type store.

The first implementation, `HashStore`, uses a hash map indexed by the `_id` field to store the data.
This provides efficient querying of that field, but all other fields require scanning through all the data in the selected document type store.

The second implementation, `InvertedStore`, augments this with an inverted index of `(term, field)` pairs. That, is, for each top level field in a document, the value is tokenised, and a hash map of `(token, field)` to a list of document `_id`s is maintained.
Thus, given a document type, a field and a word, all the documents that match are returned in constant time, provided the data had been preprocessed to build the index.

There is a difference between the results returned by each implementation because the `HashStore` only supports matching the entire value of a field, while the `InvertedStore` matches any word.
I've assumed that typically, users will be either be searching fields that have short, relatively unique values, like a `name`, or have long blob of text that they only want to search one word in, like a `description`. Thus, querying a single word to get all documents that contain that word is appropriate.
We could extend this to support entering multiple words into the query and combining the results relatively easily, but I have not done so at this stage.

Because it will fail some tests designed for the `InvertedStore`, the `HashStore` has been deprecated and its tests have been skipped.

Because each document type requires its own store, some duplication is required to add a new document type store.
Perhaps generics could have been used to avoid such duplication, but it did not seem straightforward to implement.
This was hampered by limitations with Go generics, such as the inability have type parameters in methods.
If extensibility of this type is a frequent need, code generation is a potential avenue to pursue.

## UI
I decided to use [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) as a terminal user interface (TUI) framework.
I have never used it before, but I chose it because it uses a similar architecture to a framework I have used before, [Redux](https://redux.js.org/).
Both of these are based on the [Elm Architecture](https://guide.elm-lang.org/architecture/).

This enabled the creation of an interactive UI that allows selecting the document type, field and query term quite smoothly.
However, this framework is relatively new, and I am not very experienced with it, so there are likely to be many areas of improvement.
In particular, there may be some visual glitches.

I also did not take much effort to make error messages look good and provide actionable feed back.
I generally tended to bubble up the error value from the data layer.
Some kind of error presentation layer that uses Go's error wrapping functionality would be a good further direction to explore.

Another area to improve is the presentation of results. I could have spent some time on a better way to present each document and its related documents.
At the moment it is difficult to distinguish visually between documents that matched the search and those that were returned by the store just because they were related to a document that matched.
