package main

import "fmt"

func main() {
	primeFactor(1000)
}

func primeFactor(num int) {
	fmt.Println("input ==> ", num)
	p := 2
	fmt.Print("Prime ==> ")
	for ok := true; ok; ok = (num >= (p * p)) {
    if ((num % p) == 0) {
			fmt.Print(p)
			fmt.Print(" * ")
			num = num / p
		} else {
			p = p + 1
		}
	}
	fmt.Println(num)
}
