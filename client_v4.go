package main

import(
	"fmt"
	"net"
	"encoding/gob"
)
type Persona struct{
	Nombre string
	Email []string
}
func client(persona Persona)  {
	c,err:=net.Dial("tcp",":9999")
	if err != nil {
		fmt.Print(err)
		return
	}	
	
	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Print(err)
		return
	}
	c.Close()
}
func main()  {
	persona := Persona{
		Nombre : "leon",
		Email: []string{
			"prueba@gmail.com",
			"prueba2@gmail.com",
		},
	}
	go client(persona)
	var input string
	fmt.Scanln(&input)
}