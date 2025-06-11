package main
import(
	"fmt"
)

type bill struct { 
	name string
	items map[string] float64
	tip float64
}

func newBill(name string) bill {
	b:= bill{
		name: name,
		items: map[string] float64 {"pie":33.4 , "cake":22.5},
	}
	return b
}

// Reciver function 
func (b bill) format() string{
	fs :="Bill Formated: \n";
	var total float64 = 0

	// list the items
	for key ,value := range b.items{
		fs += fmt.Sprintf("%v ...%v \n", key ,value) 
		total +=value
	}

	// total
	fs +=fmt.Sprintf("%v ...$%0.2f" , "total" , total)

	return fs
} 
