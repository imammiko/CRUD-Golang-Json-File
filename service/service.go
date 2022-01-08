package service

import (
	"CRUD-Golang-Json-File/entity"
	"errors"
)

func GetPersons(datas *entity.Data) (*[]entity.Person, error) {
	return &datas.Data, nil
}

func FindAbsen(datas *entity.Data, index int) (*entity.Person, error) {

	for k, v := range datas.Data {
		if v.Absen == index {

			return &datas.Data[k], nil
		}
	}
	return nil, errors.New("absen not found")

}

func CreatPerson(data *entity.Data, person entity.Person) (*entity.Person, error) {
	data.Data = append(data.Data, person)
	return &person, nil
}

func RemoveIndex(s []entity.Person, index int) []entity.Person {
	return append(s[:index], s[index+1:]...)
}
func DeletePerson(datas *entity.Data, absen int) (*entity.Person, error) {
	for k, v := range datas.Data {
		if v.Absen == absen {
			personHadDelete := &datas.Data[k]
			datas.Data = RemoveIndex(datas.Data, k)
			return personHadDelete, nil
		}
	}
	return nil, errors.New("absen not found")
}
func UpdatePerson(datas *entity.Data, absen int, update entity.Person) (*entity.Person, error) {
	for k, v := range datas.Data {
		if v.Absen == absen {

			datas.Data[k] = update
			return &datas.Data[k], nil
		}
	}
	return nil, errors.New("absen not found")
}
