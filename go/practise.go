package main

import (
	"fmt"
)

func main() {
	fmt.Println("UnderstandRating = 1",UnderstandRatind(1))
	fmt.Println("UnderstandRating = 5",UnderstandRatind(5))
	secretInfo := "******"
	recievedSignal := &secretInfo

	fmt.Println(recievedSignal)
	fmt.Println("Recieved password:",workWithSignal(recievedSignal))
}

func workWithSignal(signal *string) string{
	return *signal
}

func UnderstandRatind(rating float32) float32{
	switch rating{
	case 1:
		return rating + 3.2
	default:
		return rating
	}
}
