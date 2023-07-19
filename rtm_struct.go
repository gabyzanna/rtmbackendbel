package rtmbel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agendabel struct {
	ID         			primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Topik				string             `bson:"topik" json:"topik"`
	Hasil 				string             `bson:"hasil" json:"hasil"`
	RencanaPerbaikan	string             `bson:"rencanaperbaikan" json:"rencanaperbaikan"`
	PenanggungJawab		string             `bson:"penanggungjawab" json:"penanggungjawab"`
	TargetSelesai		string             `bson:"targetselesai" json:"targetselesai"`
}

type Penjawab struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nama		string             `bson:"nama" json:"nama"`
	Divisi 		string             `bson:"divisi" json:"divisi"`
}