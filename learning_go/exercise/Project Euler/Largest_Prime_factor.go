/*The prime factors of 13195 are 5, 7, 13 and 29.
*/
package main
import "fmt"

var q uint =600851475143
var mid uint =600851475143/2

func isprime(n uint) uint{
  var k,i uint
  for i=2;i<n;i++{
    if n%i==0{
      k=1
    }
  }
  if k==0{
    return n
  }
  return 0
}

func main(){
  s := make([]uint, 100000)
  var r,i uint
  for i=2;i<mid;i++{
    r=isprime(i)
    if r!=0 && q%r==0{
      s = append(s, r)
      fmt.Print(r,"\n")
      }
  }
  max:=len(s)
  fmt.Print(s[max],"\n")
}
