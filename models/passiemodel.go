package models

import (
	"database/sql"
	"fmt"

	"github.com/reynaldineo/CRUD-Golang-Native/config"
	"github.com/reynaldineo/CRUD-Golang-Native/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel{
	conn, err := config.DBCOnnection()

	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) Create(passien entities.Pasien) bool{
	result, err := p.conn.Exec("INSERT INTO pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) VALUES (?, ?, ?, ?, ?, ?, ?)", passien.NamaLengkap, passien.NIK, passien.JenisKelamin, passien.TempatLahir, passien.TanggalLahir, passien.Alamat, passien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}