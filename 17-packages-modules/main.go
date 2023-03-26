package main

import (
	"fmt"

	util "example.com/project/util"
)

func main() {
	fmt.Println("Hello", util.GetName())
	intSl := []int{1, 2, 3, 4}
	strSl := util.IntSliceToStrSlice(intSl)
	fmt.Printf("String slice - %v - %T", strSl, strSl)
}
