package main
import "fmt"

func pengingatTenggat() {
	fmt.Print("Masukkan tanggal hari ini: ")
	var harisekarang, bulansekarang int
	fmt.Scanf("%02d-%02d", &harisekarang, &bulansekarang)

	fmt.Println("\nPengingat Pembayaran:")
	fmt.Printf("Tanggal: %02d-%02d\n", harisekarang, bulansekarang)

	pengingat := false

	for _, sub := range subscriptions {
		if !sub.Aktif {
			continue
		}

		var haritagih, bulan int
		fmt.Sscanf(sub.tanggalBiling, "%02d-%02d", &haritagih, &bulan)

		if bulansekarang == bulan {
			jatuhtempo := haritagih - harisekarang
			switch {
			case jatuhtempo == 0:
				fmt.Printf("HARI INI! Bayar %s - Rp%.0f\n", sub.Nama, sub.Harga)
				pengingat = true
			case jatuhtempo > 0 && jatuhtempo <= 3:
				fmt.Printf("%d hari lagi: %s - Rp%.0f (tgl %s)\n", 
					jatuhtempo, sub.Nama, sub.Harga, sub.tanggalBiling)
				pengingat = true
			case jatuhtempo < 0:
				fmt.Printf("Terlambat %d hari: %s - Rp%.0f (seharusnya tgl %s)\n", 
					-jatuhtempo, sub.Nama, sub.Harga, sub.tanggalBiling)
				pengingat = true
			}
		}
	}

	if !pengingat {
		fmt.Println("Tidak ada tagihan yang perlu dibayar dalam 7 hari ke depan")
	}
}