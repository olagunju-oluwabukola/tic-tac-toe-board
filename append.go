package main

import "fmt"

func appendSlice (){

	var s [] int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 3)
	printSlice(s)

	s = append(s, 2,4,5,6)
	 printSlice(s)

	 s = append(s, 3,4,5,6,7,)
	 printSlice(s)
}

func printSlice(s [] int){
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// range method

var pow = [] int {2,2,3,4,6,7,8,9}

func rangeMethod(){
	for i, v:=range pow{
	fmt.Printf("2**%d = %d\n", i, v)
	}
}