// +build sample

package main

import (
	"fmt"

	"github.com/kyoh86/go-docbase"
)

func main() {
	fmt.Printf("A version of the Package %s is %s\n", "go-docbase", go-docbase.Version())
}
