package handlers

import (
	"CRUD-Golang-Json-File/entity"
	"CRUD-Golang-Json-File/helpers"
	"CRUD-Golang-Json-File/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func HandlerGetAll(datas *entity.Data) {
	dataPersons, err := service.GetPersons(datas)
	if err != nil {
		fmt.Println(err)
	} else {
		helpers.PersonDataToString(*dataPersons...)
	}
}
func HandlerDelete(dataInput []string, datas *entity.Data) {
	updateInput := helpers.ArgInputToArray(dataInput)
	i, _ := strconv.Atoi(updateInput["absen"])
	_, err := service.DeletePerson(datas, i)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := json.MarshalIndent(datas, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}

func HandlerUpdate(dataInput []string, datas *entity.Data) {
	updateInput := helpers.ArgInputToArray(dataInput)
	i, _ := strconv.Atoi(updateInput["absen"])
	findedAbsen, err := service.FindAbsen(datas, i)
	if err != nil {
		fmt.Println(err)
	}
	if _, ok := updateInput["nama"]; ok {
		findedAbsen.Nama = updateInput["nama"]
	}
	if _, ok := updateInput["alamat"]; ok {
		findedAbsen.Alamat = updateInput["alamat"]
	}
	if _, ok := updateInput["pekerjaan"]; ok {
		findedAbsen.Pekerjaan = updateInput["pekerjaan"]
	}
	if _, ok := updateInput["alasan"]; ok {
		findedAbsen.Alasan = updateInput["alasan"]
	}
	// func updatePerson(datas *Data, absen int, update Person) (*Person, error)
	service.UpdatePerson(datas, i, *findedAbsen)
	file, _ := json.MarshalIndent(datas, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)

}

func HandlerFindAbsen(dataInput []string, datas *entity.Data) {

	absen := helpers.ArgInputToArray(dataInput)
	i, _ := strconv.Atoi(absen["absen"])
	findedAbsen, err := service.FindAbsen(datas, i)
	if err != nil {
		fmt.Println(err)
	}
	helpers.PersonDataToString(*findedAbsen)

}

func HandlerCreate(dataInput []string, menu string, datas *entity.Data) {

	var personCreate = entity.Person{}
	saveLoad := helpers.ArgInputToArray(dataInput)

	if val, ok := saveLoad["nama"]; ok {

		personCreate.Nama = val

	} else {

		personCreate.Nama = "null"

	}
	if val, ok := saveLoad["alamat"]; ok {
		personCreate.Alamat = val
	} else {

		personCreate.Alamat = "null"

	}
	if val, ok := saveLoad["pekerjaan"]; ok {
		personCreate.Pekerjaan = val
	} else {

		personCreate.Pekerjaan = "null"

	}
	if val, ok := saveLoad["alasan"]; ok {
		personCreate.Alasan = val
	} else {

		personCreate.Alasan = "null"

	}
	maxAbsen := 0
	for _, v := range datas.Data {
		if v.Absen > maxAbsen {
			maxAbsen = v.Absen
		}
	}
	personCreate.Absen = maxAbsen + 1
	service.CreatPerson(datas, personCreate)
	file, _ := json.MarshalIndent(datas, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}
