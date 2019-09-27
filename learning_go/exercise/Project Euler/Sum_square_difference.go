package main
import "fmt"
func main(){
  sq:=0
  susq:=0
  for i:=1;i<=100;i++{
    sq,susq=sq+i*i,susq+i
  }
  susq*=susq
  sub:=susq-sq
  fmt.Print("\nsub=",sub)
}
