package main
import "fmt"

func calculate(a,b int)[]int{
	sum:=a+b
	difference:=a-b
	product:=a*b
	quotient:=a/b
	results:=[]int{sum,difference,product,quotient}
	  return results
}
func main(){
	fmt.Println(calculate(20,10))
	fmt.Println(calculate(700,70))
}