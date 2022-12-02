package main

import (
	"encoding/json"
	"errors"
	"fmt"

	// "net/http"

	"github.com/AndersonNascimentoAFSN/golang/src/searchWeb"
	"github.com/AndersonNascimentoAFSN/golang/src/user"

	// "github.com/AndersonNascimentoAFSN/golang/src/webServer"
	"github.com/AndersonNascimentoAFSN/golang/src/multiThread"
)

func isBiggerThan(number int) (int, error) {
	if number <= 10 {
		return 0, errors.New("Numero não é maior que 10")
	}
	return number, nil
}

func array() {
	var x [10]int
	x[0] = 1
	x[1] = 2

	y := [10]int{1, 2, 3}
	fmt.Println(y)
}

func slices() { // array infinito?
	slice := make([]int, 2)
	slice[0] = 1
	slice[1] = 2
	slice = append(slice, 3, 4, 5, 6)
	fmt.Println(slice)

	sliceString := []string{
		"Anderson",
		"Yanni",
	}
	sliceString = append(sliceString, "Thyago")
	fmt.Println(sliceString)
}

func mapExample() {
	m := make(map[string]int)
	m["Anderson"] = 31
	m["Yanni"] = 13

	fmt.Println(m)
}

// type CarName string
// type CarYar int

type Car struct {
	CarName string `json:"name"`
	CarYear int    `json:"-"`
}

func (c Car) drive() {
	fmt.Println(c.CarName, "Andou!")
}

func typesExample() {
	// var x CarName
	// x = "Fusion"
	// fmt.Println(x)
	car1 := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}

	result, _ := json.Marshal(car1)

	fmt.Println("result", string(result))

	// fmt.Println(car1.CarName)
	// car1.drive()

	j := []byte(`{"name":"BMW", "year":"2020"}`)
	var car2 Car
	json.Unmarshal(j, &car2)
	fmt.Println("car2 name:", car2.CarName)
}

type vehicle interface {
	start() string
}

func example(car vehicle) {

}

func interfaceExemplo() {

}

func main() {
	var name string
	name = "Anderson"
	age := 31
	userName, userAge, _ := user.ShowUser(name, age)
	fmt.Println(userName, userAge)

	number := 12
	isBiggerThan10, isNotBiggerThan10 := isBiggerThan(number)
	if isNotBiggerThan10 != nil {
		panic("Error")
	}
	fmt.Println("O número", isBiggerThan10, "é maior que 10.")

	// slices()
	// mapExample()
	typesExample()
	searchWeb.GetUrl("https://go.dev/")
	// webServer.Server()
	multiThread.MultiThread()
}
