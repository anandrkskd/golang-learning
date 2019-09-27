package main
import "fmt"
const (
  b=iota
  c=iota
)
func main(){
  var KB int32 = 1 << (c * 10)
  fmt.Printf("%b\t", KB)
  fmt.Printf("%d\n", KB)
}
