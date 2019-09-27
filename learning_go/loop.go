package main
import "fmt"
func main(){
  for i:=0;i<10;i++{
    for j:=1;j<=i;j++{
      fmt.Printf("%d",j)
    }
    fmt.Printf("\n")
  }
}
