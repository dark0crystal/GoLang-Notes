package main

import(
	"fmt"
)


func main(){
	mybill := newBill("khalid's Bill")

	fmt.Println(mybill.format())
}


// output result:

// Bill Formated: 
// cake ...22.5 
// pie ...33.4 
// total ...$55.90