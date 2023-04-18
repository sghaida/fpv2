# Go Functional Library
[![ci](https://github.com/sghaida/fpv2/actions/workflows/ci.yaml/badge.svg)](https://github.com/sghaida/fpv2/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/sghaida/fpv2/branch/main/graph/badge.svg?token=T7LTPQKQIR)](https://codecov.io/gh/sghaida/fpv2)


## TL;DR
This is an Opinionated functional library that implements some aspects of **Functional paradigms**, which suppose to increase productivity, and add to **GoLang** the missing beauty of functional programing which i always long for.

## Lib Primitives
The Library still **WIP** while the entirety of the features is not yet finalized, below is the list of basic Primitives that is currently supported or would be supported in the future, for Detailed Support list please go through the code :D

The Library will Include `Mutable` and `Immutable` counterpart collections such as ` Array | Slice | Map | Set `

- [x] **_[Options](src/optional.go)_** `Some | None` along with all the `Monadic Operations`

- [ ] **_Either_** `Left | Right` along with all `Monadic Operations`

- [ ] **_Mutable | Immutable Set_** `Iter | Foreach | Union | Intersect | Add | Remove | Diff | Clone | ToSlice |  Map | FlatMap |  Flatten | Fold | Reduce | Zip | Filter | Add | Remove`

- [ ] **_Mutable | Immutable Array_** applies to `Slices | Arrays` with the following Operations `Iter | Foreach | Map | FlatMap |  Flatten | Fold | Reduce | Zip | ZipWithIndex | Filter | GroupBy | Head | Tail | AddAtIndex | Append | Prepend | ToMap | Clone`

- [ ] **_Mutable | Immutable Map_** applies to `Maps` with the following Operations ` Iter | Foreach | Map | FlatMap |  Flatten | Fold | Reduce | ToSlice | Clone`

