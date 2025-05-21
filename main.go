package main

import (
	"fmt"
	"os"
)

type Subscription struct {
	Name        string
	Price       float64
	BillingDate string
	IsActive    bool
	Category    string
}

var subscriptions []Subscription

func main() {
	subscriptions = []Subscription{
		{"Netflix", 149000, "15-01", true, "Hiburan"},
		{"Disney+", 79000, "10-03", true, "Hiburan"},
		{"Spotify", 49900, "05-02", true, "Musik"},
		{"iCloud", 29000, "20-02", true, "Penyimpanan"},
	}

	for {
		tampilanMenu()
		menuChoice()
	}
}

func tampilanMenu() {
	fmt.Println("\n=== MANAGEMENT SUBKRIPSI ===")
	fmt.Println("1. Lihat Subscription")
	fmt.Println("2. Tambah Subscription")
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
		showSubscriptions(subscriptions)
	case 2:
		addSubscription()
	case 3:
		searchSubscriptions()
	case 4:
		checkPaymentReminders()

	case 0:
		fmt.Println("Terima kasih!")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func subkripsi(sub Subscription) {
	status := "Non-Aktif"
	if sub.IsActive {
		status = "Aktif"
	}
	fmt.Printf("- %s (Rp%.0f, tgl %s, %s) [%s]\n",
		sub.Name, sub.Price, sub.BillingDate, sub.Category, status)
}