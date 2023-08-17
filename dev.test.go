package NPM

import (
	"fmt"
	"testing"

	"github.com/AdeCandra12/BE_proyek1/model"
	"github.com/AdeCandra12/BE_proyek1/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertMahasiswa(t *testing.T) {
	nama_mhs := "ade candra"
	npm := "1214012"
	jurusan := "D4 teknik informatika"
	email := "adecandra12@gamil.com"
	insertedID, err := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama_mhs, npm, jurusan, email)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertOrangTua(t *testing.T) {
	nama_ortu := "Stephen kurniawan"
	phone_number := "081342565098"
	email := "steph.kur12@gmail.com"
	insertedID, err := module.InsertOrangTua(module.MongoConn, "orangtua", nama_ortu, phone_number, email)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertMatakuliah(t *testing.T) {
	nama_matkul := "proyek 1"
	sks := "4"
	dosen_pengampu := "rizki amiludin"
	email := "amiludin12@gmail.com"
	insertedID, err := module.InsertMatakuliah(module.MongoConn, "matakuliah", nama_matkul, sks, dosen_pengampu, email)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
func TestInsertAbsensi(t *testing.T) {
	nama_mk := model.Matakuliah{
		Nama_matkul:    "proyek 1",
		SKS:            "4",
		Dosen_pengampu: "rizki amilidin",
		Email:          "amiludin12@gmail.com",
	}
	tanggal := "12 februari 2023"
	checkin := "hadir"
	insertedID, err := module.InsertAbsensi(module.MongoConn, "absensi", nama_mk, tanggal, checkin)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestInsertNilai(t *testing.T) {
	nama_ms := model.Mahasiswa{
		Nama_mhs: "ade candra",
		NPM:      "1214012",
		Jurusan:  "D4 teknik informatika",
		Email:    "adecandra12@gmail.com",
	}
	presensi := model.Absensi{
		Nama_mk: model.Matakuliah{
			Nama_matkul:    "proyek 1",
			SKS:            "4",
			Dosen_pengampu: "rizki amilidin",
			Email:          "amiludin12@gmail.com",
		},
		Tanggal: "12 februari 2023",
		Checkin: "hadir",
	}
	nilai_akhir := "90"
	grade := "A"
	insertedID, err := module.InsertNilai(module.MongoConn, "monitoring", nama_ms, presensi, nilai_akhir, grade)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetMahasiswaFromID(t *testing.T) {
	id := "64a4e2e694cb7dd7f0d0f9f5"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	mahasiswa, err := module.GetMahasiswaFromID(objectID, module.MongoConn, "mahasiswa")
	if err != nil {
		t.Fatalf("error calling GetMahasiswaFromID: %v", err)
	}
	fmt.Println(mahasiswa)
}

func TestGetOrangTuaFromID(t *testing.T) {
	id := "64134e6931b8ceb1ad63d3ad"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	orangtua, err := module.GetOrangTuaFromID(objectID, module.MongoConn, "orangtua")
	if err != nil {
		t.Fatalf("error calling GetOrangTuaFromID: %v", err)
	}
	fmt.Println(orangtua)
}

func TestGetMatakuliahFromID(t *testing.T) {
	id := "64a4e2e994cb7dd7f0d0f9f7"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	matakuliah, err := module.GetMatakuliahFromID(objectID, module.MongoConn, "matakuliah")
	if err != nil {
		t.Fatalf("error calling GetMatakuliahFromID: %v", err)
	}
	fmt.Println(matakuliah)
}

func TestGetAbsensiFromID(t *testing.T) {
	id := "64a4e2ea94cb7dd7f0d0f9f8"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	absensi, err := module.GetAbsensiFromID(objectID, module.MongoConn, "absensi")
	if err != nil {
		t.Fatalf("error calling GetAbsensiFromID: %v", err)
	}
	fmt.Println(absensi)
}

func TestGetNilaiFromID(t *testing.T) {
	id := "64a4e2ea94cb7dd7f0d0f9f9"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	nilai, err := module.GetNilaiFromID(objectID, module.MongoConn, "nilai")
	if err != nil {
		t.Fatalf("error calling GetNilaiFromID: %v", err)
	}
	fmt.Println(nilai)
}

func TestGetAllMahasiswa(t *testing.T) {
	data := module.GetAllMahasiswa(module.MongoConn, "mahasiswa")
	fmt.Println(data)
}

func TestGetAllOrangTua(t *testing.T) {
	data := module.GetAllOrangTua(module.MongoConn, "orangtua")
	fmt.Println(data)
}

func TestGetAllMatakuliah(t *testing.T) {
	data := module.GetAllMatakuliah(module.MongoConn, "matakuliah")
	fmt.Println(data)
}

func TestGetAllAbsensi(t *testing.T) {
	data := module.GetAllAbsensi(module.MongoConn, "absensi")
	fmt.Println(data)
}

func TestGetAllNilai(t *testing.T) {
	data := module.GetAllNilai(module.MongoConn, "nilai")
	fmt.Println(data)
}

//Delete Data

func TestDeleteMahasiswaByID(t *testing.T) {
	id := "64a4e2e694cb7dd7f0d0f9f5" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMahasiswaByID(objectID, module.MongoConn, "mahasiswa")
	if err != nil {
		t.Fatalf("error calling DeleteMahasiswaByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetMahasiswaFromID
	_, err = module.GetMahasiswaFromID(objectID, module.MongoConn, "mahasiswa")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteOrangTuaByID(t *testing.T) {
	id := "64a532c90fbad6a10b05cb18" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteOrangTuaByID(objectID, module.MongoConn, "orangtua")
	if err != nil {
		t.Fatalf("error calling DeleteOrangTuaByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetOrangTuaFromID
	_, err = module.GetOrangTuaFromID(objectID, module.MongoConn, "orangtua")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteMatakuliahByID(t *testing.T) {
	id := "64a4e2e994cb7dd7f0d0f9f7" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMatakuliahByID(objectID, module.MongoConn, "matakuliah")
	if err != nil {
		t.Fatalf("error calling DeleteMatakuliahByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetMatakuliahFromID
	_, err = module.GetMatakuliahFromID(objectID, module.MongoConn, "matakuliah")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteAbsensiByID(t *testing.T) {
	id := "64a4e2ea94cb7dd7f0d0f9f8" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteAbsensiByID(objectID, module.MongoConn, "absensi")
	if err != nil {
		t.Fatalf("error calling DeleteAbsensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetAbsensiFromID
	_, err = module.GetAbsensiFromID(objectID, module.MongoConn, "absensi")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteNilaiByID(t *testing.T) {
	id := "64a635e4b5a0cc779e9a01ae" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteNilaiByID(objectID, module.MongoConn, "nilai")
	if err != nil {
		t.Fatalf("error calling DeleteNilaiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetNilaiFromID
	_, err = module.GetNilaiFromID(objectID, module.MongoConn, "nilai")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
