package main
import "fmt"
func check(n uint) bool{
  var i uint =2
  for ;i<=20;i++{
    if n%i!=0{
      return false
    }
  }
  return true
}
func main(){
  var no uint = 2520
  var i bool =true
  for i==true{
    if check(no){
      break
    }
    no++
  }
  fmt.Print("\n",no);
}
