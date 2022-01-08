package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Person struct {
	Absen     int    `json:"absen"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan"`
}
type Data struct {
	Data []Person `json:"data"`
}

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

	DataLoad := Data{}

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

func menuArgs(datas *Data, menu string, dataInput ...string) {
	switch menu {
	case "getAll":
		dataPersons, err := getPersons(datas)
		if err != nil {
			fmt.Println(err)
		} else {
			personDataToString(*dataPersons...)
		}
	case "create":
		handlerCreate(dataInput, menu, datas)
	case "read":
		handlerFindAbsen(dataInput, datas)
	case "update":
		handlerUpdate(dataInput, datas)
	case "delete":
		handlerDelete(dataInput, datas)
	}

}

func handlerDelete(dataInput []string, datas *Data) {
	updateInput := argInputToArray(dataInput)
	i, _ := strconv.Atoi(updateInput["absen"])
	_, err := deletePerson(datas, i)
	if err != nil {
		fmt.Println(err)
	}
	file, _ := json.MarshalIndent(datas, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}

func handlerUpdate(dataInput []string, datas *Data) {
	updateInput := argInputToArray(dataInput)
	i, _ := strconv.Atoi(updateInput["absen"])
	findedAbsen, err := findAbsen(datas, i)
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
	updatePerson(datas, i, *findedAbsen)
	file, _ := json.MarshalIndent(datas, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
	fmt.Println(datas)

}

func handlerFindAbsen(dataInput []string, datas *Data) {

	absen := argInputToArray(dataInput)
	i, _ := strconv.Atoi(absen["absen"])
	findedAbsen, err := findAbsen(datas, i)
	if err != nil {
		fmt.Println(err)
	}
	personDataToString(*findedAbsen)

}

func argInputToArray(dataInput []string) map[string]string {
	var saveLoad = make(map[string]string)
	for _, value := range dataInput {
		var keyLoad, valueLoad string
		flag := false
		for _, v := range value {
			if string(v) == ":" {
				flag = true
				continue
			}
			if flag == false {
				keyLoad += string(v)
			} else {
				valueLoad += string(v)
			}
		}
		saveLoad[string(keyLoad)] = string(valueLoad)

	}
	return saveLoad
}
func handlerCreate(dataInput []string, menu string, datas *Data) {

	var personCreate = Person{}
	saveLoad := argInputToArray(dataInput)
	fmt.Println(saveLoad)
	if val, ok := saveLoad["nama"]; ok {
		//do something here
		personCreate.Nama = val
		// fmt.Println(val, ok)
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
	creatPerson(datas, personCreate)
	file, _ := json.MarshalIndent(datas, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}

func personDataToString(data ...Person) {
	fmt.Println("data Person :")
	for _, v := range data {
		fmt.Println("###########")
		fmt.Printf("Absen: %d \n", v.Absen)
		fmt.Printf("Nama: %s \n", v.Nama)
		fmt.Printf("Alamat: %s \n", v.Alamat)
		fmt.Printf("Pekerjaan: %s \n", v.Pekerjaan)
		fmt.Printf("Alasan: %s \n", v.Alasan)

	}
}

func getPersons(datas *Data) (*[]Person, error) {
	return &datas.Data, nil
}

func findAbsen(datas *Data, index int) (*Person, error) {

	for k, v := range datas.Data {
		if v.Absen == index {

			return &datas.Data[k], nil
		}
	}
	return nil, errors.New("absen not found")

}

func creatPerson(data *Data, person Person) (*Person, error) {
	data.Data = append(data.Data, person)
	return &person, nil
}

func removeIndex(s []Person, index int) []Person {
	return append(s[:index], s[index+1:]...)
}
func deletePerson(datas *Data, absen int) (*Person, error) {
	for k, v := range datas.Data {
		if v.Absen == absen {
			personHadDelete := &datas.Data[k]
			datas.Data = removeIndex(datas.Data, k)
			return personHadDelete, nil
		}
	}
	return nil, errors.New("absen not found")
}
func updatePerson(datas *Data, absen int, update Person) (*Person, error) {
	for k, v := range datas.Data {
		if v.Absen == absen {

			datas.Data[k] = update
			return &datas.Data[k], nil
		}
	}
	return nil, errors.New("absen not found")
}
