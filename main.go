package main

import "fmt"

const con int = 10

func main() {
	m := make([]int, con)
	var num = 2
	for i := 0; i < con; i++ {
		m[i] = num * i
		num++
		fmt.Println(num)
		if m[i] > 20 {
			break
		}
	}
	fmt.Println(m)
}
