package rtmbel

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAgendabel(topik string, hasil string, rencanaperbaikan string, penanggungjawab string, targetselesai string) (InsertedID interface{}) {
	var rtmdb Agendabel
	rtmdb.Topik = topik
	rtmdb.Hasil = hasil
	rtmdb.RencanaPerbaikan = rencanaperbaikan
	rtmdb.PenanggungJawab = penanggungjawab
	rtmdb.TargetSelesai = targetselesai

	return InsertOneDoc("DatabaseTugas3", "rtmdb", rtmdb)
}

func InsertPenjawab(nama string, divisi string) (InsertedID interface{}) {
	var jawab Penjawab
	jawab.Nama = nama
	jawab.Divisi = divisi

	return InsertOneDoc("DatabaseTugas3", "jawab", jawab)
}

func GetDataAgendabel(targetselesai string) (data []Agendabel) {
	user := MongoConnect("DatabaseTugas3").Collection("rtmdb")
	filter := bson.M{"targetselesai": targetselesai}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAgendabel :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataPenjawab(divisi string) (data []Penjawab) {
	user := MongoConnect("DatabaseTugas3").Collection("jawab")
	filter := bson.M{"divisi": divisi}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataPenjawab :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}