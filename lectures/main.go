package main

import (
	"fmt"
	"reflect"
)

// we have a struct person...
type Person struct {
	Name string	 `required:"true"`
	Age  int	 `required:"true"`
	Level string `required:"false"`
}

func ValidatePersonData(s interface{}) error {
	t := reflect.TypeOf(s)
	fmt.Println("Numfield : ", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()

			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}

	}
	return nil
}

// create an instance of the Person struct
func main() {
	// init person struct instance p
	// p := Person{"Obie", 21}
	// t := reflect.TypeOf(p)
	// v := reflect.ValueOf(p)
	
	// // get the type of the person person object 
	// fmt.Println("Type : ",t)
	// fmt.Println("Value: ", v)
	// fmt.Println("Kind: ", t.Kind())

	newPerson := Person {
		Name : "",
		Age : 12,
		Level: "Beginner",
	}

	p := ValidatePersonData(newPerson)
	fmt.Println(p)
}