package main

import "fmt"

type usuario struct {
	id        int
	nombre    string
	dni       string
	telefono  string
	mail      string
	direccion string
}

func main() {
	alta("joaquin", "34534534", "34534543", "adsasdasd", "sdaasdasd")
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
