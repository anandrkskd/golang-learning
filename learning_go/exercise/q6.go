package main
import "fmt"
func main(){
  for i:=1;i<100;i++{
    if i%3==0 && i%5==0{
      fmt.Print("FizzBuzz\n")
    } else if i%3==0{
      fmt.Print("Fizz\n")
    } else if i%5==0{
      fmt.Print("Buzz\n")
    } else {
      fmt.Print(i,"\n")
    }
  }
}
