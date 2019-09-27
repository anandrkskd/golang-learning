//(true&&false)||(false&&true)||!(false&&false)
package main
import "fmt"
func main(){
  var x bool
  x=(true&&false)||(false&&true)||!(false&&false)
  fmt.Print(x)
}
