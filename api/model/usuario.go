package model

import "fmt"

type usuario struct {
	id        int
	nombre    string
	dni       string
	telefono  string
	mail      string
	direccion string
}

func alta(nombre, dni, telefono, mail, direccion string) {

	var usuario usuario
	usuario.nombre = nombre
	usuario.dni = dni
	usuario.telefono = telefono
	usuario.mail = mail
	usuario.direccion = direccion

	fmt.Println(usuario)

}

func baja(id int) {

}
