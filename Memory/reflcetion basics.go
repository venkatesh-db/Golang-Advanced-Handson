
package main 

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}


func main(){

	var x int = 5;

	fmt.Println("type:", reflect.TypeOf(x))

	fmt.Println(" value", reflect.ValueOf(x))

	fmt.Println("type:", reflect.ValueOf(x).Kind())

	P := Person{Name: "Alice", Age: 30}

	fmt.Println("type:", reflect.TypeOf(P))
	
	fmt.Println("value:", reflect.ValueOf(P))


	for i:=0;i<reflect.TypeOf(P).NumField();i++{
		field := reflect.TypeOf(P).Field(i)
		value := reflect.ValueOf(P).Field(i)
		fmt.Printf("Field: %s, Value: %v\n", field.Name, value)
	}

	var x2 int=10 

	v:=reflect.ValueOf(&x2)
	val:=v.Elem()

	if val.CanSet(){
		val.SetInt(50)

		fmt.Println("Updated value of x2:", x2)
	}

}