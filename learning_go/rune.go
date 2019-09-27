package main
import "fmt"
func main(){
  i:=0
  for i<20000{
    fmt.Println(i,"-",string(i), "-",[]byte(string(i)))
    i++
  }
}
