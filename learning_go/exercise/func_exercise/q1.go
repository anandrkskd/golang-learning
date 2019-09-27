package main
import "fmt"
func half(x int) (int,bool) {
  var re bool
  if x%2==0{
    re=true
  }
  rew:=x/2
  return rew,re
}
func main(){
  fmt.Print(half(1))
}
