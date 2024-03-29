package main

// perhatikan gambar unormal-example.png
// identifikasikan-lah semua attribut yang ada, lalu buatlah table seperti pada gambar dengan nama table "unormal"
// hint : ada sekitar 13 attribut yg bisa dijadikan kolom
// pada table "unormal", 1 baris database dapat berisi bermacam produk dipisahkan dengan tanda koma

import (
	"database/sql"
	// "fmt"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Unormal struct {
	NoBon      string `db:"no_bon"`
	NamaBarang string `db:"nama_barang"`
	Harga      string `db:"harga"`
	Jumlah     string `db:"jumlah"`
	Biaya      string `db:"biaya"`
	SubTotal   int   `db:"sub_total"`
	Discount   int  `db:"discount"`
	Total      int `db:"total"`
	Bayar      int `db:"bayar"`
	Kembalian  int `db:"kembalian"`
	Kasir      string `db:"kasir"`
	Tanggal    string `db:"tanggal"` 
	Waktu      string `db:"waktu"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama unormal dan insert data seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS unormal (
		no_bon VARCHAR(12),
		nama_barang VARCHAR(12),
		harga VARCHAR(10),
		jumlah VARCHAR(10),
		biaya VARCHAR(10),
		sub_total INT,
		discount INT,
		total INT,
		bayar INT,
		kembalian INT,
		kasir VARCHAR(10),
		tanggal VARCHAR(10),
		waktu VARCHAR(10));`
	// TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO
			unormal (no_bon, nama_barang, harga, jumlah, biaya, sub_total, discount, total, bayar, kembalian, kasir, tanggal, waktu)
			VALUES
			("00001"	, "Disket,Refil Tinta,CD Blank,Mouse"	, "4500,22500,1500,17500"	, "3,1,4,2"	, "13500,22500,6000,35000"	, 77000		, 0, 77000	, 100000	, 23000	, "Rosi"	, "04-05-2022"	, "12:00:00"),
			("00002"	, "Disket,Mouse,Flash Disk"				, "4500,17500,100000"		, "1,1,1"	, "4500,17500,100000"		, 122000	, 0, 122000	, 122000	, 0		, "Dewi"	, "04-05-2022"	, "12:00:00")
			`)
	// TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkDataExists:
// checkDataExists digunakan untuk melakukan pengecekan apakah data dengan
// nomor bon tertentu sudah ada di database atau belum
func checkDataExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT 1 FROM unormal WHERE no_bon = ?`
	// TODO: replace this

	row := db.QueryRow(sqlStmt, noBon)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// insert value hint
// ("00001"	, "Disket,Refil Tinta,CD Blank,Mouse"	, "4500,22500,1500,17500"	, "3,1,4,2"	, "13500,22500,6000,35000"	, 77000		, 0, 77000	, 100000	, 23000	, "Rosi"	, "04-05-2022"	, "12:00:00"),
// ("00002"	, "Disket,Mouse,Flash Disk"				, "4500,17500,100000"		, "1,1,1"	, "4500,17500,100000"		, 122000	, 0, 122000	, 122000	, 0		, "Dewi"	, "04-05-2022"	, "12:00:00"),
