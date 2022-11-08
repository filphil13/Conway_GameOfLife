package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, err1 := strconv.Atoi(os.Args[1])
	seed, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil && err2 != nil {
		startGame(size, seed)
	} else {
		fmt.Println("error one or both of the args are not integers")
	}

}
