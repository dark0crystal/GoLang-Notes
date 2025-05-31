package main
import(
	"fmt"
)

func ChangeValue(num *int){
	*num =777
}


func main(){
	//1. assign value to a variable 
	x := 7
	fmt.Println("x = ",x)
	// 2. get the reference to that variable memory block

	y := &x  // y = 0x14000104020
	fmt.Println("Reference to x is = ",y) //Reference to x is =  0x14000104020

	//3. Derefrencing using *
	*y = 10 
	fmt.Println("After Dereferencing, x is now = ",x) //After Dereferencing, x is now =  10


	ChangeValue(&x)
	fmt.Println("after x= ",x)
}