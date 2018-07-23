//usr/bin/env go run "$0" "$@"; exit "$?"
package main

import (
	"fmt"

	"github.com/mhf-air/sh"
)

func main() {
	fmt.Println("hello")

	fmt.Println(sh.Pwd())
}
