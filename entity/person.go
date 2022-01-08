package entity

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
