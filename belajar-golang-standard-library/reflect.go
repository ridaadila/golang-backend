package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"20"`
	Address string `required:"true" max:"20"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name: ", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type: ", valueField.Type)
		fmt.Println("Tag required: ", valueField.Tag.Get("required"))
		fmt.Println("Tag max: ", valueField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()

			result = data != ""

			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	// readField(Sample{"Rida"})
	// readField(Person{"Budi", "Tes"})
	fmt.Println(IsValid(Person{"Rida", "Ok"}))
}
