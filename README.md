# JSON RPC JSON Schemas

This repo contains the JSON Schema meta-schema that corresponds to the various JSON-RPC objects.

From these JSON Schema, types for various languages are generated.

## Installing

### Typescript

`npm install --save @json-rpc-specification/meta-schema`

### Golang

`go get github.com/json-rpc-specification/meta-schema`


### Rust

`cargo install jsonrpctypes`

## Using

### Typescript
```typescript
import { JSONRPCRequest } from "@json-rpc-specification/meta-schema"
```

### Contributing

How to contribute, build and release are outlined in [CONTRIBUTING.md](CONTRIBUTING.md), [BUILDING.md](BUILDING.md) and [RELEASING.md](RELEASING.md) respectively. Commits in this repository follow the [CONVENTIONAL_COMMITS.md](CONVENTIONAL_COMMITS.md) specification.
