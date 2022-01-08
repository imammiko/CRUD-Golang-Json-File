# CRUD-Golang-Json-File

CRUD golang sederhana dengan read dan write file json

datpat di jalankan di terminal dengan os.args

getAll 
create nama:<name> alamat:<address> pekerjaan:<job> alasan:<why>
read absen:<int>
update absen:<number> alamat:<new address>
delete absen:<absen> 

example:
$ go run main.go getAll 
$ go run main.go create nama:"Budi" alamat:"jakpus" pekerjaan:"programmer" alasan:"menambah kealihan"
$ go run main.go read absen:"1"
$ go run main.go update absen:"4" alamat:"Jaksel"
$ go run main.go delete absen:"4"
