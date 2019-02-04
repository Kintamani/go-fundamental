package main

import (
	"basic/package/config"
	namaModel "basic/package/model" //namaModel = alias dari model
	"fmt"
)

func main() {
	fmt.Println(config.GetConnectionPosgres())

	newUser := namaModel.Users{ID: 1, Email: "Sandiaga@mail.com", Password: "secret", Name: "Sandiaga Ublag"}

	fmt.Println(newUser.ID)
}
