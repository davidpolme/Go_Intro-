package main

import "fmt"

func main(){
	fmt.Println(uno(5))
	fmt.Println(salida(3))

	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(Calculo(5,46))
	fmt.Println(Calculo(1))
	fmt.Println(Calculo(147,25,30,56))
}

func uno(numero int) int {
	return numero * 2
}

func salida(numero int) (z int){
	z = numero * 2
	return 
}

func dos(numero int) (int, bool) {
	if numero == 1 {
	return 5 , false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total:=0

	for item, num:= range numero{
		fmt.Printf("\n Item %d\n", item)
		total += num
	}
	return total
}