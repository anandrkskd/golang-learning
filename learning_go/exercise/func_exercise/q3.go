package main
import "fmt"
func funct(sf ...float64) float64{
  var max float64
  for _, v := range sf {
		max += v
  }
  return max
}
func main()  {
  n:=funct(12,14,333,161,191,456,12,1,1,23,12,1234,61,100,1987,6133,1,221,1)
  fmt.Print(n)
}
