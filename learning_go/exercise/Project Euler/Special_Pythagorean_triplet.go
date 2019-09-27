package main
import "fmt"
func Pythagoreantriplet() (a,b,c int){
  for a=1;a<500;a++{
    for b=a+1;b<1000;b++{
      c = 1000 - a - b
      if b > c {
        break
      }
      if (a*a)+(b*b) == (c*c) {
        return
      }
    }
  }
  a = 0
  b = 0
  c = 0
  return
}
func main(){
  a,b,c:=Pythagoreantriplet()
  fmt.Println(a,b,c)
  fmt.Println(a*b*c)
}
