package main
import "fmt"
func main(){
  fmt.Print("Enter your name -->")
  var name string
  fmt.Scan(&name)
  fmt.Print("Hello ,",name)
}
