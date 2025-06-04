package main

import (
	"fmt"
	"os"
)

type Subkripsi struct {
	Nama        string
	Harga       float64
	tanggalBiling string
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
	fmt.Println("1. Lihat Subkripsi")
	fmt.Println("2. Tambah Subkripsi")
	fmt.Println("3. Cari Layanan Langganan")
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
		menuTampilan(subscriptions)
	case 2:
		addSubscription()
	case 3:
		carilangganan()
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
		sub.Nama, sub.Harga, sub.tanggalBiling, sub.Kategori, status)
}