package main

import "fmt"

func main()  {
	exponencia(2)
}
//Recursion es una funcion que se llama a si misma
func exponencia(num int)  {
	if num > 3000{
		return
	}
	fmt.Println(num)
	exponencia(num * 2)
}