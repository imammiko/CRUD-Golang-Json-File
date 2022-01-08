package main

import (
	"CRUD-Golang-Json-File/entity"
	"CRUD-Golang-Json-File/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println(
			"\nmasukan perintah sebagai berikut:\n\ngetAll \ncreate nama:<name> alamat:<address> pekerjaan:<job> alasan:<why>\nread absen:<int>\nupdate absen:<number> alamat:<new address>\ndelete absen:<absen> \n ")

		fmt.Printf("Error occured %T", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
func main() {
	defer catch()

	file, _ := ioutil.ReadFile("persons.json")

	DataLoad := entity.Data{}

	_ = json.Unmarshal([]byte(file), &DataLoad)

	menu := []string{"getAll", "create", "read", "update", "delete"}
	findMenu := false
	for _, v := range menu {
		if v == os.Args[1] {
			findMenu = true
		}
	}
	if findMenu == false {
		fmt.Println(
			"\nmasukan perintah sebagai berikut:\n\ngetAll \ncreate nama:<name> alamat:<address> pekerjaan:<job> alasan:<why>\nread absen:<int>\nupdate absen:<number> alamat:<new address>\ndelete absen:<absen> \n ")
	} else {
		loadDataArgs := []string{}
		for _, v := range os.Args[2:] {
			loadDataArgs = append(loadDataArgs, v)
		}
		menuArgs(&DataLoad, os.Args[1], loadDataArgs...)

	}

}

func menuArgs(datas *entity.Data, menu string, dataInput ...string) {
	switch menu {
	case "getAll":
		// handlers.HandlerGetAll(datas)
		handlers.HandlerGetAll(datas)
	case "create":
		handlers.HandlerCreate(dataInput, menu, datas)
	case "read":
		handlers.HandlerFindAbsen(dataInput, datas)
	case "update":
		handlers.HandlerUpdate(dataInput, datas)
	case "delete":
		handlers.HandlerDelete(dataInput, datas)
	}

}
