package main

import "fmt"

func main()  {
	// paises:=make(map[string]string, 5)
	// fmt.Println(paises)
	
	// paises["Mexico"]="D.F."
	// fmt.Println(paises["Mexico"])
	// paises["Argentina"]="Buenos Aires"
	// fmt.Println(paises)

	campeonato:= map[string]int{
		"Barcelona": 4,
		"Real Madrid": 3,
		"Millonarios":1,
		"PSG":30	}

		fmt.Println(campeonato)

		campeonato["Barcelona"]=20
		fmt.Println(campeonato)
		
		delete(campeonato,"Real Madrid")
		fmt.Println(campeonato)


		for equipo, puntaje := range campeonato{
			fmt.Printf("%s tiene un puntaje de: %d \n", equipo, puntaje)
		}

		puntaje, existe := campeonato["Santa Fe"]
		fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
	