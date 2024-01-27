package models

import (
	"database/sql"
	"fmt"
	"time"

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

func (p *PasienModel) FindAll() ([]entities.Pasien, error){

	rows, err := p.conn.Query("SELECT * FROM pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(
			&pasien.Id, 
			&pasien.NamaLengkap, 
			&pasien.NIK, 
			&pasien.JenisKelamin, 
			&pasien.TempatLahir, 
			&pasien.TanggalLahir, 
			&pasien.Alamat, 
			&pasien.NoHp)

		if pasien.JenisKelamin == "1" {
			pasien.JenisKelamin = "Laki-laki"
		} else {
			pasien.JenisKelamin = "Perempuan"
		}

		// 2006-01-02 => yyyy-mm-dd in Golang
		tgl_lahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)
		// 02 Januari 2006 => dd mmmm yyyy in Golang
		pasien.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataPasien = append(dataPasien, pasien)
	}		

	return dataPasien, nil
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

func (p *PasienModel) Find(id int64, pasien *entities.Pasien) error {

	return p.conn.QueryRow("SELECT * FROM pasien WHERE id = ?", id).Scan(
		&pasien.Id,
		&pasien.NamaLengkap, 
		&pasien.NIK, 
		&pasien.JenisKelamin, 
		&pasien.TempatLahir, 
		&pasien.TanggalLahir, 
		&pasien.Alamat, 
		&pasien.NoHp)
}

func (p *PasienModel) Update(pasien entities.Pasien) error {

	_, err := p.conn.Exec("UPDATE pasien SET nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? WHERE id = ?", pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp, pasien.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PasienModel) Delete(id int64) error {

	_, err := p.conn.Exec("DELETE FROM pasien WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}