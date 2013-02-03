// +build !foo,!bar

package main

import (
    "fmt"
)

func hello() {
    fmt.Println("I'm not FOOing or BARing!")
}
