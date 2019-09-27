package main
import "fmt"

func even_fibo()  {
  a,b,tmp:=1,2,0
  for b<4000000{
    a,b=b,a+b
    if b%2==0{
      tmp+=b
    }
  }
  fmt.Print("\nsum =",tmp)
}

func main(){
  even_fibo()
}
