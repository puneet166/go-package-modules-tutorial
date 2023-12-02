// main.go

package main

import (
	"go-ethereum/package-tutorial/package_1"
	"go-ethereum/package-tutorial/package_2"
	"go-ethereum/package-tutorial/package_3"

	"github.com/ethereum/go-ethereum/core/types"

	"fmt"
)

func main() {
	fmt.Println("--------start------")

	package_1.Function_1()
	package_2.Function_2()
	package_3.Function_3()
	types.EncodeNonce(5) // this is how use go-ethereum code's function and packages
	fmt.Println("--------end------")
}
