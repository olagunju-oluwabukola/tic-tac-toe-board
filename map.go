package main

import "fmt"

type vert struct {
	lat, long float64
}

var mp = map [string]vert{
	"Bell labs" : vert{
	40.6666, -60.7777,
},
"google" : vert{
	56.77777, 90.6666,
},
}


func mapTest(){
	a:= make(map[string]int)
	a["answer"] = 9
	fmt.Println("the value:", a["answer"])

	a["answer3"] = 90
	fmt.Println("the value is:", a["answer"])
	a["answer2"] = 900
	fmt.Println("the value is:", a["answer"])
	v, ok := a["answer2"]
	fmt.Println("The value is:", v, "present?", ok)
	//check if the key "answer2" has a vlue and it dosent have, so it returns false
// map iteration
for key, value := range a{
	fmt.Println("key is ", key, "value is", value)
}


	b:= make(map[string]int)
	b["key2"] = 1000
	fmt.Print(b)

	check, result := b["key2"]
	fmt.Println("the value", check, "is present is", result )
}