package main

import "fmt"

type rect struct{
		length float32
		breadth float32
	}
type circle struct{
	radius float32
}

func (r rect) area() float32{
	return r.length * r.breadth
	
}

func (c circle) area() float32{
	const pi = 3.14
	return pi * c.radius * c.radius
}


func (r rect ) perimeter() float32{
	const con = 2
	return  con * r.length * 2 *r.breadth
}

func (c circle) perimeter() float32{
	const pi = 3.14
	const con = 2
	return con * pi * c.radius
}

type Shapes interface{
	area() float32
	perimeter() float32
}


func array_pointer(){
	var arr = [5]int{1,2,3,4,5}
	fmt.Printf("the the arr is at location %p", &arr)
	var result = square_arrays(&arr)
	fmt.Printf("the result is %v\n", result)
}

func square_arrays(arr1 *[5]int) [5]int {
	fmt.Printf("the adress for arr1 is %p", arr1)
    for i := range arr1{
		arr1[i] = arr1[i] * arr1[i]
	}
	return *arr1
}

func slice_experiment(){
	var slice = []int{1,2,3,4,5}
	var slice_copy = slice
	slice[2] = 10
	fmt.Printf("the memory address for slice is %p and the slice is %v", &slice, slice)
	fmt.Printf("the memory address for the copy slice is %p and the number is %v", &slice_copy, slice)
}

func pointer(){

	var p *int32 = new(int32)
	var i int32
	fmt.Printf("the pointer is %v and the vlue at pointer is %v\n", p, *p)
	*p = i
	fmt.Printf("the pointer value is %v\n", *p)
    i=1
	p = &i
	fmt.Printf("the pointer value is %v\n", *p)
}

func main(){
 var c1 = rect{10, 20}
 var r1 = circle{23}
 ci_result:=c1.area() 
 r1_result:= r1.area()
 ci_p_result := c1.perimeter()
 r1_p_result := r1.perimeter()
 fmt.Printf("area of circle is %v and area of rectangle is %v \n", ci_result, r1_result)
 fmt.Printf("perimeter of circle is %v and perimeter of rectngle is %v \n", ci_p_result, r1_p_result)
 pointer()
 array_pointer()
slice_experiment()

}
