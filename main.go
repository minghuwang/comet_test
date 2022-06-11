package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// item 1
	scores := os.Args[1:]
	var s, w, e, n int32
	s = 5
	w = 5
	e = 5
	n = 5
	if len(scores) == 4 {
		south, err := strconv.Atoi(scores[0])
		if err != nil {
			fmt.Println("invalid argument")
			return
		}
		s = int32(south)
		west, err := strconv.Atoi(scores[1])
		if err != nil {
			fmt.Println("invalid argument")
			return
		}
		w = int32(west)
		east, err := strconv.Atoi(scores[2])
		if err != nil {
			fmt.Println("invalid argument")
			return
		}
		e = int32(east)
		north, err := strconv.Atoi(scores[3])
		if err != nil {
			fmt.Println("invalid argument")
			return
		}
		n = int32(north)

	}
	cm := ControlMethod(s, w, e, n)
	fmt.Println(cm)

}
