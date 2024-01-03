# Ethereum-Info
- A repository of using GraphQL to query Ethereum block information.
## Run Services
1. Generate images `make gen-images`
2. Run services `make service-up`
3. Stop services `make service-down`

## GraphQL
- [Schema](https://github.com/0x726f6f6b6965/ethereum-info/blob/main/api/graph-service/schema/schema.graphqls)
- Endpoint: "http://localhost:8866/"
- Query
  1. `block(num: Uint64!): Block!`
  2. `latestBlocks(num: Uint64!): Blocks!`
  3. `transaction(txHash: String!): Transaction!`
- Type
  - Block
    | Field        | Type          |
    | ------------ | ------------- |
    | blockNum     | BigInt        |
    | blockHash    | String        |
    | blockTime    | BigInt        |
    | parentHash   | String        |
    | transactions | Array[String] |
    | stable       | Boolean       |
  - Transaction
    | Field  | Type       |
    | ------ | ---------- |
    | txHash | String     |
    | from   | String     |
    | to     | String     |
    | nonce  | String     |
    | data   | String     |
    | value  | String     |
    | logs   | Array[Log] |
  - Log
    | Field  | Type   |
    | ------ | ------ |
    | txHash | String |
    | index  | Uint32 |
    | data   | String |
  - Blocks
    | Field  | Type         |
    | ------ | ------------ |
    | blocks | Array[Block] |
