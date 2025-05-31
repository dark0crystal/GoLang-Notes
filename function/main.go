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
//labeled return values
func multi(x, y int) (num1 int , num2 int){
	num1 = x*y
	num2 = x*y+x
	return
}

func main(){
	sum := add(5 , 4);
	fmt.Println("Return one value: ",sum);
	sum2 , mult := add2(5,4);
	fmt.Printf("The sum is %d , and the multiply is: %d \n" ,sum2, mult);
	num1 , num2 := multi(4 , 5);
	fmt.Printf("First Num: %d , The second Num : %d", num1 , num2 );

}