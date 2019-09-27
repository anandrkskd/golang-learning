package main
import "fmt"
func varadic(sf ...float64) float64{
  var total float64
  for _,v:=range sf{
    total+=v
  }
  return total / float64(len(sf))
}
func main(){
  n:=varadic(20,12,13,14,15,16,17,18,19,11)
  fmt.Println(n)
}
