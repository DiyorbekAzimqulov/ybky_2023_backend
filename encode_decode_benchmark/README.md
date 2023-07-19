**Person Struct**

The `Person` struct represents a person's information with the following fields:

- `Name`: A string containing the person's name.
- `Age`: An integer representing the person's age.

**Encoding and Decoding Methods**

This package includes encoding and decoding methods for the `Person` struct using different formats:

1. JSON Encoding/Decoding: Utilizing the `encoding/json` package for JSON serialization and deserialization.

2. BSON Encoding/Decoding: Using the `go.mongodb.org/mongo-driver/bson` package for BSON serialization and deserialization.

3. Protocol Buffers Encoding/Decoding: Employing Google's Protocol Buffers with Go's `google.golang.org/protobuf/proto` package.

4. GOB (Golang Object) Encoding/Decoding: Utilizing Go's built-in `encoding/gob` package for encoding and decoding.

**Benchmark Functions**

The package contains benchmark functions for each encoding/decoding method to compare their performance when handling the `Person` struct.

**How to Use**

To run the benchmarks, use the standard `go test` command with the `-bench` flag followed by the package name:

```bash
go test -bench=.
```

## Literature Review

(Benchmarking Go Programs)[https://gobyexample.com/testing-and-benchmarking]

To learn more about GOB objects check this out: (GOB Objects)[https://go.dev/blog/gob]

Here is a good source about Protocol Buffers: (Protocol Buffers)[https://protobuf.dev/overview/]

To learn about difference between JSON and ProtoBuf this Youtube video helps a lot: (JSON vs ProtoBuf)[https://www.youtube.com/watch?v=uGYZn6xk-hA]

To learn more about BSON and why it is used check this out: (BSON vs JSON)[https://www.mongodb.com/json-and-bson]
