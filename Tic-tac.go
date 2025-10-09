package main

import (
	"fmt"
	"strings"
)
func board() {
	//board
    board := [] [] string{
		{"-","-","-" },
	 {"-","-","-" },
	 {"-","-","-" },
	}
//players result
board [0] [0] = "x"
board [2] [2] = "0"
board [1] [2] = "x"
board [1] [0] = "0"
board [0] [2] = "x"

for i := 0; i < len(board); i++{
	fmt.Printf("%s\n", strings.Join(board[i], " "))
}

// fmt.Println(board)
}

func main (){
board()
appendSlice()
rangeMethod()
mapTest()
fmt.Println(mp)
}