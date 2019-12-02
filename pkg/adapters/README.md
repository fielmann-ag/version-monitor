# Version Adapters

Adapters are a good and simple way to extend the possibilities of fetching versions. There are already some simple
adapters, please have a look at the [example config](../../example.yaml) to get an idea how they work.

## Add a new Adapter

In order to add a new Adapter you need to ...

1. have a proper adapter folder structure. Currently all adapters are ordered by its technology and transport/method
(e.g. `http` is the technology and transport is `get`, so the adapter lives in [pkg/adapters/http/get/adapter.go`]()
but this is NOT a requirement. If you feel like there could be multiple adapters for the same technology
(like multiple transports, methods, commands, ...) then make sure to have a nested adapter pattern already applied.
2. register the [`AdapterConstructor`](types.go) for your [`AdapterType`](../monitor/types.go) in
[`adapters.go`](adapters.go). Each `AdapterType` has to be a unique string and should be defined in the same package as
the Adapter and its `AdapterConstructor`.
3. Have all required config parameters defined in the [config](../monitor/config.go). Remember: Config is parsed from
YAML, so all kinds of fancy nested types are supported.
4. Have a decent test coverage for your adapter: `go test ./... -cover`

### Adapter Example

For a pretty simple adapter have a look on the `ShellCommandAdapter`:

* Adapter: [./shell/command/adapter.go](./shell/command/adapter.go)
* AdapterConstructor: [./shell/command/constructor.go](./shell/command/constructor.go)
* AdapterType: [./shell/command/const.go](./shell/command/const.go)
* Register the adapter: [./adapters.go](./adapters.go) 
