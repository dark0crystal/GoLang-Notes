package main

import(
	"fmt"
)

// It's like a custom type that we want to use 
type Student struct{
	id int
	name string
	grade float64
}



func main(){
	var std Student = Student{id:111222 , name:"omar" , grade:87.4}

	fmt.Printf("Student Name is : %s \nStudent ID is: %d \nStudent grade is : %.2f", std.name, std.id, std.grade)

}