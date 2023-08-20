package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_mhs string             `bson:"nama_mhs,omitempty" json:"nama_mhs,omitempty"`
	NPM      string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Jurusan  string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
}

type OrangTua struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_ortu    string             `bson:"nama_ortu,omitempty" json:"nama_ortu,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
}

type Matakuliah struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_matkul    string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	SKS            string             `bson:"sks,omitempty" json:"sks,omitempty"`
	Dosen_pengampu string             `bson:"dosen_pengampu,omitempty" json:"dosen_pengampu,omitempty"`
	Email          string             `bson:"email,omitempty" json:"email,omitempty"`
}

type Absensi struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_mk Matakuliah         `bson:"nama_mk,omitempty" json:"nama_mk,omitempty"`
	Tanggal string             `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Checkin string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
}

type Nilai struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NPM_ms       Mahasiswa          `bson:"npm_ms,omitempty" json:"npm_ms,omitempty"`
	Presensi     Absensi            `bson:"presensi,omitempty" json:"presensi,omitempty"`
	Nilai_akhir  string             `bson:"nilai_akhir,omitempty" json:"nilai_akhir,omitempty"`
	Grade        string             `bson:"grade,omitempty" json:"grade,omitempty"`
	Tahun_ajaran string             `bson:"tahun_ajaran,omitempty" json:"tahun_ajaran,omitempty"`
}

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fullname        string             `bson:"fullname,omitempty" json:"fullname,omitempty"`
	Username        string             `bson:"username,omitempty" json:"username,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
}
