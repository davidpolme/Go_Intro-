package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	numero1, numero2, numero3, texto := 5, 24, 12, "este es un texto"
	
	var moneda float64
	numero1 = int(moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(texto)
}
