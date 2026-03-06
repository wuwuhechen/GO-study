package main

import "fmt"

func main() {

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} //单个条件

	fmt.Println("-----part 1 end-----")

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	} //经典的for循环

	fmt.Println("-----part 2 end-----")

	for {
		fmt.Println("loop")
		break
	} //无限循环 = while(true)

	fmt.Println("-----part 3 end-----")

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
