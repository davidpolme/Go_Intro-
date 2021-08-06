package main

import "fmt"

var estado bool

func main(){
	///if Statement
	
	estado = true
	//!reasigna el valor de la variable dentro del if
	if estado=false; estado == true {
		fmt.Println(estado)
	}else{
		fmt.Println("El estado es falso")
	}
	

	///switch Case Statement
	switch numero:=2; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}

}