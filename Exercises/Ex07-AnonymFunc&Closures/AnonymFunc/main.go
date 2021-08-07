package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}
func main(){
	// fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5,7))

	// //Restamos
	// Calculo = func(num1 int, num2 int) int {
	// 	return num1-num2
	// }

	// fmt.Printf("Resto 10 - 7 = %d \n", Calculo(10,7))
	
	// //Dividimos
	// Calculo = func(num1 int, num2 int) int {
	// 	return num1 / num2
	// }
	
	// fmt.Printf("Divido 12 / 3 = %d \n", Calculo(12,3))

	Operaciones()
}

func Operaciones (){
	resultado:= func() int {
		var a int = 13
		var b int = 10
		return a + b
	} 

	fmt.Println(resultado())
}