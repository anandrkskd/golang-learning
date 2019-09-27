package main
import "fmt"
var maxa int
func ispalindrome(n int) bool{
  var k,tmp int
  k=n
  for n>0{
    tmp,n=tmp*10+n%10,n/10
  }
  if tmp==k{
    return true
  }
  return false
}
func main(){
  for i:=100;i<=999;i++{
    for j:=100;j<999;j++{
      if product:=i*j; ispalindrome(product) && product>maxa{
        maxa=product
      }
    }
  }
  fmt.Print("max=",maxa)
}
