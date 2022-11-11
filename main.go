package main

import (
	"fmt"
)

func main() {
	res := FindIps()
	fmt.Println("Map", res)
	keys := FindKeys()
	fmt.Println("Hostname", keys)

	Gay()
}
