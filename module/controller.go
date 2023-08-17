package module

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"errors"

	"github.com/AdeCandra12/BE_proyek1/model"
	"github.com/aiteung/atdb"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "proyek1_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMahasiswa(db *mongo.Database, col string, nama_mhs string, npm string, jurusan string, email string) (insertedID primitive.ObjectID, err error) {
	mahasiswa := bson.M{
		"nama_mhs": nama_mhs,
		"npm":      npm,
		"jurusan":  jurusan,
		"email":    email,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), mahasiswa)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertOrangTua(db *mongo.Database, col string, nama_ortu string, phone_number string, email string) (insertedID primitive.ObjectID, err error) {
	orangtua := bson.M{
		"nama_ortu":    nama_ortu,
		"phone_number": phone_number,
		"email":        email,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), orangtua)
	if err != nil {
		fmt.Printf("InsertOrangTua: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertMatakuliah(db *mongo.Database, col string, nama_matkul string, sks string, dosen_pengampu string, email string) (insertedID primitive.ObjectID, err error) {
	matakuliah := bson.M{
		"nama_matkul":    nama_matkul,
		"sks":            sks,
		"dosen_pengampu": dosen_pengampu,
		"email":          email,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), matakuliah)
	if err != nil {
		fmt.Printf("InsertMatakuliah: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
func InsertAbsensi(db *mongo.Database, col string, nama_mk model.Matakuliah, tanggal string, checkin string) (insertedID primitive.ObjectID, err error) {
	absensi := bson.M{
		"nama_mk": nama_mk,
		"tanggal": tanggal,
		"checkin": checkin,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), absensi)
	if err != nil {
		fmt.Printf("InsertAbsensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertNilai(db *mongo.Database, col string, npm_ms model.Mahasiswa, presensi model.Absensi, nilai_akhir string, grade string) (insertedID primitive.ObjectID, err error) {
	nilai := bson.M{
		"npm_ms":      npm_ms,
		"presensi":    presensi,
		"nilai_akhir": nilai_akhir,
		"grade":       grade,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), nilai)
	if err != nil {
		fmt.Printf("InsertNilai: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// UPDATE DATA

func UpdateMahasiswa(db *mongo.Database, col string, id primitive.ObjectID, nama_mahasiswa string, npm int, jurusan string, email string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":    nama_mahasiswa,
			"npm":     npm,
			"jurusan": jurusan,
			"email":   email,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMahasiswa: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateOrangTua(db *mongo.Database, col string, id primitive.ObjectID, nama_orangtua string, phone_number string, email string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_ot":      nama_orangtua,
			"phone_number": phone_number,
			"email":        email,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateOrangTua: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateMatakuliah(db *mongo.Database, col string, id primitive.ObjectID, nama_matkul string, sks string, dosen_pengampu string, email string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_matkul":    nama_matkul,
			"sks":            sks,
			"dosen_pengampu": dosen_pengampu,
			"email":          email,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMatakuliah: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}
func UpdateAbsensi(db *mongo.Database, col string, id primitive.ObjectID, npm_mhs model.Mahasiswa, nama_mk model.Matakuliah, tanggal string, checkin string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"npm_mhs": npm_mhs,
			"nama_mk": nama_mk,
			"tanggal": tanggal,
			"checkin": checkin,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateAbsensi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func UpdateNilai(db *mongo.Database, col string, id primitive.ObjectID, npm_ms model.Mahasiswa, matkul model.Matakuliah, presensi model.Absensi, nilai_akhir string, grade string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"npm_ms":      npm_ms,
			"matkul":      matkul,
			"presensi":    presensi,
			"nilai_akhir": nilai_akhir,
			"grade":       grade,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateNilai: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

// DELETE DATA

func DeleteMahasiswaByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	mahasiswa := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := mahasiswa.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteOrangTuaByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	orangtua := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := orangtua.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteMatakuliahByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	matkul := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := matkul.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteAbsensiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	absen := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := absen.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func DeleteNilaiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	nilai := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := nilai.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// Get Data pakai ID

func GetMahasiswaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mhs model.Mahasiswa, errs error) {
	siswa := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := siswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mhs, fmt.Errorf("no data found for ID %s", _id)
		}
		return mhs, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mhs, nil
}

func GetOrangTuaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (ot model.OrangTua, errs error) {
	ortu := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := ortu.FindOne(context.TODO(), filter).Decode(&ot)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ot, fmt.Errorf("no data found for ID %s", _id)
		}
		return ot, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return ot, nil
}

func GetMatakuliahFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mk model.Matakuliah, errs error) {
	matkul := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := matkul.FindOne(context.TODO(), filter).Decode(&mk)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mk, fmt.Errorf("no data found for ID %s", _id)
		}
		return mk, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mk, nil
}

func GetAbsensiFromID(_id primitive.ObjectID, db *mongo.Database, col string) (abs model.Absensi, errs error) {
	absen := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := absen.FindOne(context.TODO(), filter).Decode(&abs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return abs, fmt.Errorf("no data found for ID %s", _id)
		}
		return abs, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return abs, nil
}

func GetNilaiFromID(_id primitive.ObjectID, db *mongo.Database, col string) (nl model.Nilai, errs error) {
	nilai := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := nilai.FindOne(context.TODO(), filter).Decode(&nl)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nl, fmt.Errorf("no data found for ID %s", _id)
		}
		return nl, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return nl, nil
}

// GET ALL FUNCTION

func GetAllNilai(db *mongo.Database, col string) (data []model.Nilai) {
	nilai := db.Collection(col)
	filter := bson.M{}
	cursor, err := nilai.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllNilai :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllMahasiswa(db *mongo.Database, col string) (data []model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{}
	cursor, err := mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMahasiswa :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllOrangTua(db *mongo.Database, col string) (data []model.OrangTua) {
	orangtua := db.Collection(col)
	filter := bson.M{}
	cursor, err := orangtua.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllOrangTua :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllMatakuliah(db *mongo.Database, col string) (data []model.Matakuliah) {
	matakuliah := db.Collection(col)
	filter := bson.M{}
	cursor, err := matakuliah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAlMatakuliah :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetAllAbsensi(db *mongo.Database, col string) (data []model.Absensi) {
	absensi := db.Collection(col)
	filter := bson.M{}
	cursor, err := absensi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllAbsensi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func InsertOneDoc2(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertUser(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		// fmt.Printf("InsertOneDoc: %v\n", err)
		return insertedID, fmt.Errorf("kesalahan server")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetUserFromUsername(username string, db *mongo.Database, col string) (result model.User, err error) {
	collection := db.Collection(col)
	filter := bson.M{"username": username}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, fmt.Errorf("username tidak ditemukan")
		}
		return result, fmt.Errorf("kesalahan server")
	}
	return result, nil
}

func SignUp(db *mongo.Database, col string, insertedDoc2 model.User) (insertedID primitive.ObjectID, err error) {
	if insertedDoc2.Fullname == "" || insertedDoc2.Username == "" || insertedDoc2.Password == "" || insertedDoc2.Confirmpassword == "" {
		return insertedID, fmt.Errorf("Data tidak boleh kosong")
	}
	if err = checkmail.ValidateFormat(insertedDoc2.Username); err != nil {
		return insertedID, fmt.Errorf("username tidak valid")
	}
	if !strings.Contains(insertedDoc2.Username, "username") {
		return insertedID, fmt.Errorf("username harus menggunakan domain username")
	}
	userExists, _ := GetUserFromUsername(insertedDoc2.Username, db, col)
	if insertedDoc2.Username == userExists.Username {
		return insertedID, fmt.Errorf("username sudah terdaftar")
	}
	if insertedDoc2.Confirmpassword != insertedDoc2.Password {
		return insertedID, fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc2.Password, " ") {
		return insertedID, fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc2.Password) < 8 {
		return insertedID, fmt.Errorf("password terlalu pendek")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc2.Password), nil, 1, 64*1024, 4, 32)
	insertedDoc2.Password = hex.EncodeToString(hashedPassword)
	insertedDoc2.Confirmpassword = ""
	return InsertUser(db, col, insertedDoc2)
}

func LogIn(db *mongo.Database, col string, insertedDoc2 model.User) (userName string, err error) {
	if insertedDoc2.Username == "" || insertedDoc2.Password == "" {
		return userName, fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc2.Username); err != nil {
		return userName, fmt.Errorf("username tidak valid")
	}
	existsDoc, err := GetUserFromUsername(insertedDoc2.Username, db, col)
	if err != nil {
		return
	}
	hash := argon2.IDKey([]byte(insertedDoc2.Password), nil, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return userName, fmt.Errorf("password salah")
	}
	return existsDoc.Fullname, nil
}
