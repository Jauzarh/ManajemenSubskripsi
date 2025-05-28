package main 
import "fmt"

func addSubscription() {
	var sub Subscription

	fmt.Print("Nama Subscription: ")
	fmt.Scanln(&sub.Nama)

	fmt.Print("Harga: ")
	fmt.Scanln(&sub.Harga)

	fmt.Print("Tanggal Pembayaran (DD-MM): ")
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

