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
		{"Netflix", 149000, "15-01", true, "Entertainment"},
		{"Disney+", 79000, "10-03", true, "Entertainment"},
		{"Spotify", 49900, "05-02", true, "Music"},
		{"iCloud", 29000, "20-02", true, "Cloud Storage"},
	}

	for {
		tampilanMenu()
		processMenuChoice()
	}
}

func tampilanMenu() {
	fmt.Println("\n=== APLIKASI KEUANGAN SEDERHANA ===")
	fmt.Println("1. Lihat Subscription")
	fmt.Println("2. Tambah Subscription")
	fmt.Println("3. Cari Layanan Langganan")
	fmt.Println("4. Periksa Pengingat Pembayaran")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func processMenuChoice() {
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

func showSubscriptions(subs []Subscription) {
    n := len(subs)
    
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if subs[j].Price > subs[maxIdx].Price {
                maxIdx = j
            }
        }
        subs[i], subs[maxIdx] = subs[maxIdx], subs[i]
    }
    
    fmt.Println("\nDaftar Subkripsi:")
   for i, sub := range subs {
        status := "Aktif"
        if !sub.IsActive {
            status = "Non-Aktif"
        }
        fmt.Printf("%d. %s\n", i+1, sub.Name)
        fmt.Printf("   - Harga: Rp%.0f\n", sub.Price)
        fmt.Printf("   - Jatuh Tempo: tgl %s\n", sub.BillingDate)
        fmt.Printf("   - Status: %s\n", status)
        fmt.Printf("   - Kategori: %s\n", sub.Category)
        fmt.Println("---------------------------------------------")
    }
}


func addSubscription() {
    var sub Subscription

    fmt.Print("Nama Subscription: ")
    fmt.Scanln(&sub.Name)

    fmt.Print("Harga: ")
    fmt.Scanln(&sub.Price)

    fmt.Print("Tanggal Pembayaran (DD-MM): ")
    var day, month int
    fmt.Scanf("%d-%d", &day, &month)

    month += 1
    if month > 12 {
        month = 1
    }

    sub.BillingDate = fmt.Sprintf("%02d-%02d", day, month)

    fmt.Print("Kategori: ")
    fmt.Scanln(&sub.Category)

    fmt.Print("Aktif (y/n): ")
    var active string
    fmt.Scanln(&active)
    sub.IsActive = active == "y" || active == "Y"

    subscriptions = append(subscriptions, sub)

    fmt.Println("Subscription berhasil ditambahkan!")
    fmt.Println("Jatuh tempo 1 bulan kemudian:", sub.BillingDate)
}

func searchSubscriptions() {
	fmt.Println("\nPilihan Pencarian:")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Kategori")
	fmt.Print("Pilih metode pencarian: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		searchSubscriptionByName()
	case 2:
		searchByCategory()
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func sequentialSearch(subscriptions []Subscription, name string) int {
	foundIndex := -1
	i := 0

	for foundIndex == -1 && i < len(subscriptions) {
		if subscriptions[i].Name == name {
			foundIndex = i
		}
		i++
	}

	return foundIndex
}

func searchSubscriptionByName() {
	fmt.Print("Masukkan nama subscription yang ingin dicari: ")
	var searchName string
	fmt.Scanln(&searchName)

	index := sequentialSearch(subscriptions, searchName)

	if index != -1 {
		fmt.Println("\nSubscription ditemukan pada indeks", index)
		jumlahSubkripsi(subscriptions[index])
	} else {
		fmt.Println("Subscription tidak ditemukan")
	}
}

func searchByCategory() {
	fmt.Print("Masukkan kategori yang dicari: ")
	var searchCategory string
	fmt.Scanln(&searchCategory)

	low := 0
	high := len(subscriptions) - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2

		if subscriptions[mid].Category == searchCategory {
			ketemu = true
			fmt.Println("\nHasil Pencarian:")
			jumlahSubkripsi(subscriptions[mid])

			left := mid - 1
			for left >= 0 && subscriptions[left].Category == searchCategory {
				jumlahSubkripsi(subscriptions[left])
				left--
			}

			right := mid + 1
			for right < len(subscriptions) && subscriptions[right].Category == searchCategory {
				jumlahSubkripsi(subscriptions[right])
				right++
			}
			break
		} else if subscriptions[mid].Category < searchCategory {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ditemukan subscription dengan kategori tersebut")
	}
}

func jumlahSubkripsi(sub Subscription) {
	status := "Non-Aktif"
	if sub.IsActive {
		status = "Aktif"
	}
	fmt.Printf("- %s (Rp%.0f, tgl %s, %s) [%s]\n",
		sub.Name, sub.Price, sub.BillingDate, sub.Category, status)
}

func checkPaymentReminders() {
	fmt.Print("Masukkan tanggal hari ini: ")
	var harisekarang, bulansekarang int
	fmt.Scanf("%02d-%02d", &harisekarang, &bulansekarang)

	fmt.Println("\nPengingat Pembayaran:")
	fmt.Printf("Tanggal: %02d-%02d\n", harisekarang, bulansekarang)

	hasReminders := false

	for _, sub := range subscriptions {
		if !sub.IsActive {
			continue
		}

		var haritagih, bulan int
		fmt.Sscanf(sub.BillingDate, "%02d-%02d", &haritagih, &bulan)

		if bulansekarang == bulan {
			daysUntilDue := haritagih - harisekarang
			switch {
			case daysUntilDue == 0:
				fmt.Printf("HARI INI! Bayar %s - Rp%.0f\n", sub.Name, sub.Price)
				hasReminders = true
			case daysUntilDue > 0 && daysUntilDue <= 3:
				fmt.Printf("%d hari lagi: %s - Rp%.0f (tgl %s)\n", 
					daysUntilDue, sub.Name, sub.Price, sub.BillingDate)
				hasReminders = true
			case daysUntilDue < 0:
				fmt.Printf("Terlambat %d hari: %s - Rp%.0f (seharusnya tgl %s)\n", 
					-daysUntilDue, sub.Name, sub.Price, sub.BillingDate)
				hasReminders = true
			}
		}
	}

	if !hasReminders {
		fmt.Println("Tidak ada tagihan yang perlu dibayar dalam 7 hari ke depan")
	}
}