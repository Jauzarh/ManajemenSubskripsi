package main

import (
	"fmt"
	"os"
)

type Subkripsi struct {
	Nama        string
	Harga       float64
	jatuhTempo string
	Aktif    bool
	Kategori    string
}

var subscriptions []Subkripsi

func main() {
	subscriptions = []Subkripsi{
		{"iCloud", 29000, "20-02", true, "Penyimpanan"},
		{"Netflix", 149000, "15-01", true, "Hiburan"},
		{"Spotify", 49900, "05-02", true, "Musik"},
		{"Disney+", 79000, "10-03", true, "Hiburan"},
	}

	for {
		tampilanMenu()
		menuChoice()
	}
}

func tampilanMenu() {
	fmt.Println("\n=== AturIn ===")
	fmt.Println("1. Tambah Subkripsi")
	fmt.Println("2. Cari Subkripsi")
	fmt.Println("3. Lihat Subkripsi")
	fmt.Println("4. Periksa Pengingat")
	fmt.Println("5. Total Pengluaran")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func menuChoice() {
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addSubscription()
	case 2:
		carilangganan()
	case 3:
		menuTampilan(subscriptions)
	case 4:
		pengingatTenggat()
	case 5:
		tampilkantotal(subscriptions)
	case 0:
		fmt.Println("Terima kasih!")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func subkripsi(sub Subkripsi) {
	status := "Non-Aktif"
	if sub.Aktif {
		status = "Aktif"
	}
	fmt.Printf("- %s (Rp%.0f, tgl %s, %s) [%s]\n",
		sub.Nama, sub.Harga, sub.jatuhTempo, sub.Kategori, status)
}