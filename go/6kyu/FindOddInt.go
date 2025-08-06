package main

import "fmt"

func FindOdd(numbers []int) int {
	counter := make(map[int]int)
	for _,number := range numbers{
		counter[number]++
	}
	
	for k,v := range counter{
		if v%2 != 0{
			return k
		}
	}
	return 0
}

func main(){
	fmt.Println(FindOdd([]int{1,2,2,3,3,3,4,3,3,3,2,2,1}))
}