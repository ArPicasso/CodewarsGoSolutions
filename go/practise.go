package main

import (
	"fmt"
)

func main() {
	secretInfo := "******"
	recievedSignal := &secretInfo

	fmt.Println(recievedSignal)
	fmt.Println("Recieved password:",workWithSignal(recievedSignal))
}

func workWithSignal(signal *string) string{
	return *signal
}
