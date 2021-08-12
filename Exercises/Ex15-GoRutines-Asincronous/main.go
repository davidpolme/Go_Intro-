package main


import (
	"fmt"
	"strings"
	"time"
)
func main()  {
	go miNombreLentooo("David Polania Mejia")

	fmt.Println("Estoy aqui")

	var x string
	fmt.Scanln(&x)
}

func miNombreLentooo(nombre string){
	letras := strings.Split(nombre,"")
	for _, letra := range letras {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(letra)
	}
}