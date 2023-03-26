package main

import "fmt"

var pl = fmt.Println

func main() {
	/* var name type
	- name: begin with letter, may contains letters, digits
	- Capital name (ex: DoSomething()): exported, can be access outside of the package
	- Variable is a mutable data type, you can change its value but cannot change its type
	*/
	var vName string = "Hoang"
	var v1, v2 int = 1, 2
	var v3 = 3.14

	v4 := "Hello"
	v4 = "World"

	pl(vName, v1, v2, v3, v4)
}
