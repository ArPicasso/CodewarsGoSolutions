package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i:=0; i<100; i++{
		n_re := fmt.Sprintf("%b",i)
		if i % 2 == 0{
			n_re = "10" + n_re 
		} else {
			n_re = "1" + n_re + "01"
		}
		R, _ := strconv.ParseInt(n_re,2,64)
		
		if i == 4{
			fmt.Println("Число:", i, "n_re:", n_re, "R:", R)
		}
		if R < 30{
			fmt.Println(i)
		}
	}
	
}
