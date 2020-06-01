package model

type user struct {
	id     int
	name   string
	dni    string
	phone  string
	mail   string
	adress string
}

func addUser(name, dni, phone, mail, adress string) {

	var user user
	user.name = name
	user.dni = dni
	user.phone = phone
	user.mail = mail
	user.adress = adress

}

func deleteUser(id int) {

}
