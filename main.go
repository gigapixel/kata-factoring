package main

import "fmt"
import "strconv"
import "strings"

func main() {
	input := 11
	var ints []string
	index := 2
	for input > 0 {
		if(index > input) { break };
		if input % index == 0 {
			input = input / index
			ints = append(ints, strconv.Itoa(index));
		} else {
			index++
		}
	}
	fmt.Println("result => ", strings.Join(ints[:], " x "))
}
