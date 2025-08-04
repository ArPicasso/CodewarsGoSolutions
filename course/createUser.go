package main

import (
	"fmt"
	"math/rand"
)
//Структура Юзера
type User struct{
	Name string
	Age int
}

//Рандомное создания Юзера 
func createUserName() string{
	names := []string{"alice", "kate", "steve", "max"}
	arrayIndex := rand.Intn(4)
	return names[arrayIndex]
}
func createUserAge() int{
	ages := []int{15,23,93,43}
	arrayIndex := rand.Intn(4)
	return ages[arrayIndex]
}




//Вызываем от юзера
func (user User) greeting(){
	fmt.Println("Hallo, ich bin", user.Name)
}

func (userus User) sayAge(){
	fmt.Println("Ich bin", userus.Age,"Jahre alt")
}

//Основная логика программы

func main(){
	user := User{
		Name: createUserName(),
		Age: createUserAge(),
	}
	user.greeting()
	user.sayAge()
}