package model

type Wine struct {
	Id        int64
	Name      string
	Review    string
	GrapeType string
	Brand     string
	Year      int
	Origin    string
	Price     float32
	Stock     int
}

//Thing that it is not need
/*func add(name string, review string, grapeType string, brand string, year int, origin  string, price float32, stock int) {

	var vino vino
	vino.name = name
	vino.review = review
	vino.grapeType = grapeType
	vino.brand = brand
	vino.year = year
	vino.origin = origin
	vino.price = price
	vino.stock =stock

}*/
