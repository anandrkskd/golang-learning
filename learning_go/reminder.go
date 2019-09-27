package main
import "fmt"
var x = 144

func reminder() func() int {
  return func() int {
      return x%13
    }
}

func main(){
  rem := reminder()
  fmt.Println(rem())
}
