package user

import (
	"errors"
	"fmt"
	"reflect"
)

func ShowUser(name string, age int) (string, int, error) {
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", age)
	if reflect.TypeOf(name).Kind() != reflect.String {
		return "", 0, errors.New("Digite um nome")
	}
	fmt.Println(reflect.TypeOf(name))
	return name, age, nil
}
