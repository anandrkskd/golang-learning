package main

import "fmt"

func newno(new int) int {
	return (new)
}
func main() {
	fmt.Println("hello =", newno(42))
	fmt.Println(hello("annad", "singh"))
}

func hello(lname, fname string) (string, string) {
	return fmt.Sprint(lname, fname), fmt.Sprint(fname, lname)
}
