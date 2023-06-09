# Go Functional Library
[![ci](https://github.com/sghaida/fpv2/actions/workflows/ci.yaml/badge.svg)](https://github.com/sghaida/fpv2/actions/workflows/ci.yaml)
[![GitHub Clones](https://img.shields.io/badge/dynamic/json?color=success&label=Clone&query=count&url=https://gist.githubusercontent.com/sghaida/ddd9cdd153e34bb1f895e0a5b5128ebd/raw/clone.json&logo=github)](https://github.com/MShawon/github-clone-count-badge)
[![codecov](https://codecov.io/gh/sghaida/fpv2/branch/main/graph/badge.svg?token=T7LTPQKQIR)](https://codecov.io/gh/sghaida/fpv2)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2d9baa3db6cb4f9db65020013632dc1a)](https://app.codacy.com/gh/sghaida/fpv2/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## TL;DR
This is an Opinionated functional library that implements some aspects of **Functional paradigms**, which suppose to increase productivity and add to **GoLang** the missing beauty of functional programming which I have always longed for.

## Lib Primitives
The Library is still **WIP** while the entirety of the features is not yet finalized, below is the list of basic Primitives that is currently supported or would be supported in the future, for Detailed Support list please go through the code :D

The Library will Include `Mutable` and `Immutable` counterpart collections such as ` Array | Slice | Map | Set `

- [x] **_[Options](src/optional.go)_** `Some | None` along with all the `Monadic Operations`

- [x] **_[Either](src/either.go)_** `Left | Right` along with all `Monadic Operations`

- [x] **_[iter](src/iter)_** 
  
  - [x] **_[EmptyIter](src/iter/empty_iter.go)_** `Next | HasNext | Count | Size`
  
  - [x] **_[RangeIter](src/iter/range_iter.go)_** `Next | HasNext | Count | Size | FromSlice | ToSlice | Fold | FoldLeft | Map | Reduce | Filter | Foreach | Slice | Take | Drop | Contains |Clone`
  
  - [x] **_[SliceIter](src/iter/slice_iter.go)_**  `Next | HasNext | Count | Size | FromSlice | ToSlice | Fold | FoldLeft | Map | Reduce | Filter | Foreach | Slice | Take | Drop | Contains |Clone`
  
  - [x] **_[MapIter](src/iter/map_iter.go)_**  `Next | HasNext | Count | Size | FromMap | ToMap | Fold | FoldLeft | Map | Reduce | FilterByKey | FilterByValue | Foreach | ContainsKey | ContainsValue | GroupByValue | Clone`

- [ ] **_[Collections](src/collections)_**
  
  - [x] **_[Slice Ops](src/collections/list/slice_ops.go)_** `Size | Take | Map | Reduce | FoldLeft | Append | Prepend | Foreach | Flatten | Flatmap | Filter `
  
  - [x] **_[Map Ops](src/collections/dict/map_ops.go)_** `Size | Map | FoldLeft | Add | remove | Foreach | Contains | Keys | Values | KeysValues`

- [ ] **_[List](src/list/list.go)_** `Size | Head | Tail | Append | Prepend | Split | reverse | At | AppendList | PrependList | Concat | String | Find | Map | Filter | Reduce | FoldLeft | ToSlice | FromSlice | Find`

- [ ] **_Mutable | Immutable Set_** `Iter | Foreach | Union | Intersect | Add | Remove | Diff | Clone | ToSlice |  Map | FlatMap |  Flatten | Fold | Reduce | Zip | Filter | Add | Remove`

- [ ] **_Mutable | Immutable Array_** applies to `Slices | Arrays` with the following Operations `Iter | Foreach | Map | FlatMap |  Flatten | Fold | Reduce | Zip | ZipWithIndex | Filter | GroupBy | Head | Tail | AddAtIndex | Append | Prepend | ToMap | Clone`

- [ ] **_Mutable | Immutable Map_** applies to `Maps` with the following Operations ` Iter | Foreach | Map | FlatMap |  Flatten | Fold | Reduce | ToSlice | Clone`

