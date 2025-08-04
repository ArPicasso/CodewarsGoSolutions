package main

import (
	"fmt"
	"time"
)

type Person struct{
	Name string
	Rating float32
	ratingLocation *float32
}

func addRatind(user *Person){
	if user.Rating + 3.0 <= 10.0{
		fmt.Println("Обновляю рейтинг User'a с никнеймом:", user.Name)
		time.Sleep(2 * time.Second)
		newRating := user.Rating + 3.0
		user.ratingLocation = &newRating
		user.Rating = newRating
		fmt.Println("Рейтинг User'a с никнеймом:", user.Name,"обновлен.")
	} else{
		fmt.Println("User:",user.Name,"не сможет получить столько рейтинга, его рейтинг уже:", user.Rating)
	}
}
