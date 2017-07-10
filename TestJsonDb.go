package main

import (
	"encoding/json"
	"fmt"
	"os"

  "database/sql"
  _"github.com/lib/pq"
)

const (
  DB_USER     = "postgres"
  DB_PASSWORD = "123456"
  DB_NAME     = "DataMahasiswa"
)

type Mahasiswa struct {
  ID      int
  Nama    string
  Jurusan string
}

var mahasiswa []Mahasiswa

func main() {

  fmt.Println("Initialize DB")
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
      DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err)
  defer db.Close()

  fmt.Println("Read Data From DB")
  rows, err := db.Query("SELECT * FROM userinfo")
  checkErr(err)

  for rows.Next() {
      var uid int
      var nama string
      var jurusan string
      err = rows.Scan(&uid, &nama, &jurusan)
      checkErr(err)
      mahasiswa = append(mahasiswa, Mahasiswa{ID: uid, Nama: nama, Jurusan: jurusan})
  }
  fmt.Println("Read Data From DB Succes\n")

  fmt.Println("Bentuk Data ke JSON")
	b, err := json.Marshal(mahasiswa)
	if err != nil {
		fmt.Println("error:", err)
  }
  os.Stdout.Write(b)
  fmt.Println("\n")

  fmt.Println("Dari JSON bentuk ke Struct")
  var dataMahasiswa []Mahasiswa
	err = json.Unmarshal(b, &dataMahasiswa)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", dataMahasiswa)
}

func checkErr(err error) {
  if err != nil {
      panic(err)
  }
}
