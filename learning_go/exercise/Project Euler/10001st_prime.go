package main
import "fmt"
var c uint
func isprime(n uint) {
  var k,i uint
  for i=2;i<n;i++{
    if n%i==0{
      k=1
    }
  }
  if k==0{
    c++
  }
}

func main(){
  var i uint
  for c<=10001{
    isprime(i)
    i++
  }
  fmt.Print("\n",i,"\n")
}
