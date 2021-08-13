package main

import "fmt"

/* 
* Middlewares
Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos tipos de variables
*/

var result int

func main()  {
	fmt.Println("Inicio")
	
	result = operacionesMidd(sumar)(2,3)
	fmt.Printf("Suma: %d \n",result)

	result = operacionesMidd(restar)(10,6)
	fmt.Printf("Resta: %d \n",result)

	result = operacionesMidd(multiplicar)(4,2)
	fmt.Printf("Multiplicacion: %d \n",result)
}

func operacionesMidd(f func(int, int) int ) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a,b)
	}
}


func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}