package main

import (
	"fmt"
	"github.com/yujie0121/canvas/remote"
)

func main() {

	c := remote.GetAirData("上海")
	fmt.Println(c)
}
