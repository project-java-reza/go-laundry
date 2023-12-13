package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// menggunakan huruf besar karena digunakan diluar ketika dipanggil
// berkomunikasi antar layer menggunakan interface
type DbConnection interface {
	Conn() *sql.DB
}

// menggunakan huruf kecil karena hanya digunakan di file ini saja
type dbConnection struct {
	db  *sql.DB // data kosong
	cfg *Config
}

// pada saat d dipanggil maka otomatis mendapatkan db dan cfg yang ada di dbConnection
func (d *dbConnection) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.cfg.DbConfig.Host,
		d.cfg.DbConfig.Port,
		d.cfg.DbConfig.User,
		d.cfg.DbConfig.Password,
		d.cfg.DbConfig.Name,
	)

	db, err := sql.Open(d.cfg.DbConfig.Driver, dsn) // open
	if err != nil {
		return err
	}
	d.db = db // reassign dari data kosong db return pointer sql.DB setelah di open
	return nil
}

// function receiver nya dbConnection dan returnnya pointer sql.DB
func (d *dbConnection) Conn() *sql.DB {
	return d.db
}

// Constructor
// returnnya DbConnection dari interface, dan interface DbConnection cuma punya 1 method yang bisa diimplementasi yaitu Conn dan returnnya *sql.DB
func NewDbConnection(cfg *Config) (DbConnection, error) {
	// instance dbConnenction punya 2 field tapi disini hanya mengisi 1 field yaitu cfg
	// kenapa cuma 1 field ? karena db digunakan pada saat initDb otomatis dbConnection akan keisi pada saat pemanggil struct
	conn := &dbConnection{
		cfg: cfg,
	}
	// conn sekarang sudah punya db
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// clean architecture kalau butuh sebuah file hanya tinggal panggil Constructor dan codingan utamannya sudah ada di Constuctor
// sekarang kalau butuh DbConnection tinggal panggil Constructor
