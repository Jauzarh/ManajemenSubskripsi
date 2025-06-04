package main

import "fmt"

func addSubscription() {
	var sub Subkripsi
	sub.Aktif = true

	fmt.Print("Nama Subscription: ")
	fmt.Scanln(&sub.Nama)

	for {
		fmt.Print("Harga: ")
		_, err := fmt.Scanln(&sub.Harga)
		if err != nil {
			fmt.Println("Error: input harus angka. Silakan coba lagi.")
			var salah string
			fmt.Scanln(&salah)
			continue
		}
		break
	}

	fmt.Print("Tanggal Pembayaran | Gunakan Format(DD-MM): ")
	var day, month int
	fmt.Scanf("%d-%d", &day, &month)
	fmt.Scanln()

	month += 1
	if month > 12 {
		month = 1
	}

	sub.tanggalBiling = fmt.Sprintf("%02d-%02d", day, month)

	fmt.Print("Kategori: ")
	fmt.Scanln(&sub.Kategori)

	subscriptions = append(subscriptions, sub)

	fmt.Println("Subscription berhasil ditambahkan!")
	fmt.Println("Jatuh tempo 1 bulan kemudian:", sub.tanggalBiling)
}
