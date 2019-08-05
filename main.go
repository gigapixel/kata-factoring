package main

import "fmt"

func main() {
	fmt.Println("Hello")
	primeFactor(20)
}

func primeFactor(num int) {
	fmt.Println("input ==> ", num)
	p := 2
	for ok := true; ok; ok = (num >= (p * p)) {
    if ((num % p) == 0) {
			fmt.Println("p = ", p)
			num = num / p
		} else {
			p = p + 1
		}
	}
	fmt.Println("num = ", num)
}
