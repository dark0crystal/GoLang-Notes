package main
import(
	"fmt"
)

//return one value
func add(x, y int) int {
	 return x + y ; 
}
//return multiple values
func add2(x, y int) (int, int){
	return x+y , x*y;
}

func main(){
	sum := add(5 , 4);
	fmt.Println("Return one value: ",sum)
	sum2 , mult := add2(5,4);
	fmt.Printf("The sum is %d , and the multiply is: %d \n" ,sum2, mult)


}