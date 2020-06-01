package model

type vino struct {
	id        int
	name      string
	review    string
	grapeType string
	brand     string
	year      string
	origin    string
	price     float32
	stock     int
}

func addWine(name string, review string, grapeType string, brand string, year string, origin string, price float32, stock int) {

	var vino vino
	vino.name = name
	vino.review = review
	vino.grapeType = grapeType
	vino.brand = brand
	vino.year = year
	vino.origin = origin
	vino.price = price
	vino.stock = stock

}

func deleteWine(id int) {

}
