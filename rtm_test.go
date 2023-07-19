package rtmbel

import (
	"fmt"
	"testing"
)


func TestInsertAgendabel(t *testing.T) {
	topik := "Pendaftaran"
	hasil := "Tidak sesuai target"
	rencanaperbaikan := "Membatasi waktu pendaftaran"
	penanggungjawab := "Kelompok A"
	targetselesai := "Secepatnya"
	hsl := InsertAgendabel(topik, hasil, rencanaperbaikan, penanggungjawab, targetselesai)
	fmt.Println(hsl)
}

func TestInsertPenjawab(t *testing.T) {
	nama := "Bella"
	divisi := "Manajemen"
	hsl := InsertPenjawab(nama, divisi)
	fmt.Println(hsl)
}

func TestGetDataPenjawab(t *testing.T) {
	divisi := "Manajemen"
	dt := GetDataPenjawab(divisi)
	fmt.Println(dt)
}

func TestGetDataAgendabel(t *testing.T) {
	targetselesai := "Secepatnya"
	dt := GetDataPenjawab(targetselesai)
	fmt.Println(dt)
}