package main

import (
	"github.com/canvas/remote"
	"fmt"
)

func main() {

	c := remote.GetAirData("上海")
	fmt.Println(c)
}
