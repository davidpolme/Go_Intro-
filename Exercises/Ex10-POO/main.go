package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct{
	us.Usuario
}

type david struct{
	us.Usuario
}


func main()  {
	d := new(pepe)
	d.AltaUsuario(1,"David",time.Now(), true)
	fmt.Println(d.Usuario)

	p := new(pepe)
	p.AltaUsuario(2,"Pepe",time.Now(), true)
	fmt.Println(p.Usuario)

}