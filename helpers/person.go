package helpers

import (
	"CRUD-Golang-Json-File/entity"
	"fmt"
)

func ArgInputToArray(dataInput []string) map[string]string {
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
func PersonDataToString(data ...entity.Person) {
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
