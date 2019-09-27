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
    c+=n
  }
}

func main(){
  var j uint =2
  for ;j<=2000000;j++{
    var k,i uint
    for i=2;i<j;i++{
      if j%i==0{
        k=1
      }
    }
    if k==0{
      c+=j
      fmt.Println(j)
    }
  }
  fmt.Print("\n",c,"\n")
}
