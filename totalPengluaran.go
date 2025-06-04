package main

import "fmt"

func hitungTotal(sub []Subkripsi) (int, float64) {
    totalBiaya := 0.0
    jumlahSubs := 0

    for i := 0; i < len(sub); i++ {
        if sub[i].Aktif {
            totalBiaya += sub[i].Harga
            jumlahSubs++
        }
    }
    return jumlahSubs, totalBiaya
}

func tampilkantotal(subs []Subkripsi) {
    jumlah, total := hitungTotal(subs)
    
    fmt.Println("\n    Ringkasan Biaya    ")
    fmt.Printf("Jumlah Subskripsi Aktif: %d\n", jumlah)
    fmt.Printf("Total Pengeluaran Bulanan: Rp%.0f\n", total)
}
