package main

import (
	"fmt"
)

func totalPengeluaran() {
	total := 0.0
	kategoriPengeluaran := make(map[string]float64)

	for _, sub := range subscriptions {
		if sub.IsActive {
			total += sub.Price
			kategoriPengeluaran[sub.Category] += sub.Price
		}
	}

	fmt.Printf("\n💰 Total Pengeluaran Bulanan: Rp%.0f\n", total)
	fmt.Println("\n📊 Rincian per Kategori:")
	for kategori, jumlah := range kategoriPengeluaran {
		fmt.Printf("- %s: Rp%.0f\n", kategori, jumlah)
	}
}
