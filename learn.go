package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func variables_constants() {
	//2 ^ n gives us the values which can be presnted using the specified number of bits
	var number int16 = 32767 //this is the max number you can store in an integer of 16 bits 2 ^ 15 one bit is used for storing the sign
	number = number + 1      // this will overflow the number and cause it to loop back
	fmt.Println(number)
	// we have int(takes the default type of the number) int16 int32 int64
	var unsignedint uint16 = 65535 // this is the max value for unsigned int
	unsignedint = unsignedint + 1  // this will cause an overflow too
	fmt.Println(unsignedint)

	// floats can only be float32 or float64
	var decimal float32 = 12343.7
	fmt.Println(decimal)

	//operations can only be done on compatible types type cast very similar python
	a, b := 10, 11
	result := a / b
	fmt.Println(result)

	// booleans
	var boolean bool = true
	fmt.Println(boolean)

	// strings
	var hello string = "hello world"
	fmt.Println(hello)
	fmt.Println(len(hello))                    // this stores the umber of bytes used not the len of the string
	fmt.Println(utf8.RuneCountInString(hello)) // this returs the actual length of the string

	//rune is used to store single character asci values of strings
	var rune_example rune = 'a'
	fmt.Println(rune_example) // dumps the ascii value of the character

	//constants
	//constants just can be declared the need to be assigned and used
	const pi float32 = 3.14
	fmt.Println(pi)

	//cool example where the constant value is determined by the type it uses
	resultTest := 23.4 / 3
	fmt.Println(resultTest)

}

func printMe(printString string) {

	fmt.Println(printString)
}

func division(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("zero division error")
		return 0, 0, err
	}
	var div int = a / b
	var rem int = a % b
	return div, rem, err
}

func functionAndControlStrctures() {

	str := "i am an king"
	printMe(str)
	x, y := 20, 100
	var div, rem, err = division(x, y)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(div, rem)
	}
}

func arrays_maps_loops() {
	var newArray = []int{1, 2, 3, 4}
	fmt.Println(newArray)
	var slice_2 = make([]int, 3, 10)
	fmt.Println(slice_2)

	var string_ = []string{"a", "b", "c"}
	var new string
	var stringBuilder strings.Builder
	for i := range string_ {
		stringBuilder.WriteString(string_[i])
	}

	new = stringBuilder.String()
	fmt.Println(new)

}

func strings_and_runes() {
	var String string = "Résumé"
	fmt.Println(len(String))
	indexed := String[0]
	fmt.Println(indexed)

	// for i,v := range String{
	// 	fmt.Printf("the index is %v and the value is %v \n",i,v)
	// }

	var Strings_test = []rune("Résumé")
	for i, v := range Strings_test {
		fmt.Printf("%v and %v \n", i, v)
	}

}

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerName owner
}

type gasEngine1 struct {
	mpg     uint8
	gallons uint8
	owner
}
type electricEngine struct {
	kwph     uint8
	capacity uint8
	owner
}
type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwph * e.capacity
}

type engine interface {
	milesLeft() uint8
}

func canMakeit(e engine, miles uint8) {
	if e.milesLeft() > miles {
		fmt.Println("can reach destination")
	} else {
		fmt.Println("cannot reach destination")
	}

}

func structs_in_go() {
	var car1 gasEngine = gasEngine{10, 11, owner{"hero"}}
	fmt.Println(car1)
	var cars2 gasEngine1 = gasEngine1{10, 11, owner{"alex"}}
	fmt.Println(cars2)
	//anonymus structs cannot reuse these structs again
	var myEngine = struct {
		name string
		age  uint8
	}{"prajodh", 18}
	fmt.Println(myEngine)
	var result = car1.milesLeft()
	fmt.Println(result)
	var eEngine electricEngine = electricEngine{15, 11, owner{"king"}}
	var result2 = eEngine.milesLeft()
	fmt.Println(result2)
	canMakeit(eEngine, 42)
	canMakeit(car1, 42)

}

func main() {
	fmt.Println("helloWorld")
	variables_constants()
	functionAndControlStrctures()
	arrays_maps_loops()
	strings_and_runes()
	structs_in_go()
}
