package main
import "fmt"
func main(){
  half:= func (x int) (int,bool){
    var re bool
    if x%2==0{
      re=true
    }
    rew:=x/2
    return rew,re
  }
  fmt.Print(half(2))
}
