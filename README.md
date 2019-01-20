## go-delta - Package to generate and apply binary delta updates between two byte arrays.

[![Go Report Card](https://goreportcard.com/badge/github.com/balacode/go-delta)](https://goreportcard.com/report/github.com/balacode/go-delta)
[![Build Status](https://travis-ci.org/balacode/go-delta.svg?branch=master)](https://travis-ci.org/balacode/go-delta)
[![Test Coverage](https://coveralls.io/repos/github/balacode/go-delta/badge.svg?branch=master&service=github)](https://coveralls.io/github/balacode/go-delta?branch=master)
[![Gitter chat](https://badges.gitter.im/balacode/go-delta.png)](https://gitter.im/go-delta/Lobby)
[![godoc](https://godoc.org/github.com/balacode/go-delta?status.svg)](https://godoc.org/github.com/balacode/go-delta)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Hello World:

```go
package main

import (
    "fmt"
    "github.com/balacode/go-delta"
)

func main() {
    fmt.Print("Binary delta update demo:\n\n")

    // The original data (20 bytes):
    var source = []byte("quick brown fox, lazy dog, and five boxing wizards")
    fmt.Print("The original is:", "\n", string(source), "\n\n")

    // The updated data containing the original and new content (82 bytes):
    var target = []byte(
        "The quick brown fox jumps over the lazy dog. " +
        "The five boxing wizards jump quickly.",
    )
    fmt.Print("The update is:", "\n", string(target), "\n\n")

    // Use MakeDiff to generate a compressed patch between source and target
    var dif = MakeDiff(source, target)

    // Apply the patch to source to get the target
    // The size of the patch is much shorter than target.
    var target2, err = dif.Apply(source)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Print("Patched:", "\n", string(target2), "\n\n")
} //                                                                        main
```
