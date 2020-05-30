package model

import "fmt"

type vino struct {
	id        int
	nombre    string
	resenha      string
	tipoDeUva  string
	tipo      string
	anho string
	origen string
	precio float32
	stock int
}

func alta(nombre, resenha, tipoDeUva, tipo, anho, origen  string; precio float32; stock int) {

	var vino vino
	vino.nombre = nombre
	vino.resenha = resenha
	vino.tipoDeUva = tipoDeUva
	vino.tipo = tipo
	vino.anho = anho
	vino.origen = origen
	vino.precio = precio
	vino.stock =stock

}

func baja(id int) {

}