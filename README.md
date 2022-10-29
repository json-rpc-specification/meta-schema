# JSON RPC JSON Schemas

This repo contains the JSON Schema meta-schema that corresponds to the various JSON-RPC objects.

From these JSON Schema, types for various languages are generated.

meta schema is hosted here: https://meta.jsonrpc.net/

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
