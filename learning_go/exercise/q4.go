package main
import "fmt"
func main(){
  var a,b int
  fmt.Print("enter two numbers('one small and big numbers')")
  fmt.Scan(&a)
  fmt.Scan(&b)
  if a>b{
    fmt.Print("\nsol",a%b)
  } else{
    fmt.Print("\nsol",b%a)
  }
}
