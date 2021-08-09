package main

import "fmt"

//?var tabla [10]int
// *var tabla [10] int{5,41,88,61,81,15,4,6,8,41}
var matriz [5][7]int

func main() {
	//? tabla[0]=1	
	//? tabla[5]=15

	/*
	tabla := [10]int{5,41,88,61,81,15,4,6,8,41}
	for i:=0; i< len(tabla); i++{
		fmt.Println(tabla[i])
	}
	*/

	matriz[3][5]=1

	fmt.Println(matriz)
}
