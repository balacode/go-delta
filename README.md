## go-delta - A Go package and utility to generate and apply binary delta updates.

[![Go Report Card](https://goreportcard.com/badge/github.com/balacode/go-delta)](https://goreportcard.com/report/github.com/balacode/go-delta)
[![Build Status](https://travis-ci.org/balacode/go-delta.svg?branch=master)](https://travis-ci.org/balacode/go-delta)
[![Test Coverage](https://coveralls.io/repos/github/balacode/go-delta/badge.svg?branch=master&service=github)](https://coveralls.io/github/balacode/go-delta?branch=master)
[![Gitter chat](https://badges.gitter.im/balacode/go-delta.png)](https://gitter.im/go-delta/Lobby)
[![godoc](https://godoc.org/github.com/balacode/go-delta?status.svg)](https://godoc.org/github.com/balacode/go-delta)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Suggestions:

- Works best on text files, database dumps and any other files with lots of
  repeating patterns and few changes between updates.

- Generating deltas of compressed files is not recommended because a small
  change in the source data can lead to lots of changes in the compressed
  result, so generating a delta update may give you only minimal size
  reduction.

- Don't compress bytes returned by Delta.Bytes() because they are already
  compressed using ZLib compression.

- Every delta update adds about 156 bytes for the source and target hashes
  and various lengths, so it is not recommended for very miniscule updates.

## Demonstration:

```go
package main

import (
    "fmt"

    "github.com/balacode/go-delta"
)

func main() {
    fmt.Print("Binary delta update demo:\n\n")

    // The original data (20 bytes):
    source := []byte("quick brown fox, lazy dog, and five boxing wizards")
    fmt.Print("The original is:", "\n", string(source), "\n\n")

    // The updated data containing the original and new content (82 bytes):
    target := []byte(
        "The quick brown fox jumps over the lazy dog. " +
            "The five boxing wizards jump quickly.",
    )
    fmt.Print("The update is:", "\n", string(target), "\n\n")

    var dbytes []byte
    {
        // Use Make() to generate a compressed patch from source and target
        d := delta.Make(source, target)

        // Convert the delta to a slice of bytes (e.g. for writing to a file)
        dbytes = d.Bytes()
    }

    // Create a Delta from the byte slice
    d, err := delta.Load(dbytes)
    if err != nil {
        fmt.Println(err)
    }

    // Apply the patch to source to get the target
    // The size of the patch is much shorter than target.
    target2, err := d.Apply(source)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Print("Patched:", "\n", string(target2), "\n\n")
} //
```
