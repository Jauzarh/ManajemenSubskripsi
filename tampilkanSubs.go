package main
import "fmt"

func showSubscriptions(subs []Subscription) {
    n := len(subs)
    
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if subs[j].Harga > subs[maxIdx].Harga {
                maxIdx = j
            }
        }
        subs[i], subs[maxIdx] = subs[maxIdx], subs[i]
    }
    
    fmt.Println("\nDaftar Subkripsi:")
   for i, sub := range subs {
        if !sub.Aktif {
        }
        fmt.Printf("%d. %s\n", i+1, sub.Nama)
        fmt.Printf("   - Harga: Rp%.0f\n", sub.Harga)
        fmt.Printf("   - Jatuh Tempo: tgl %s\n", sub.tanggalBiling)
        fmt.Println("   - Status: Akrif")
        fmt.Printf("   - Kategori: %s\n", sub.Kategori)
        fmt.Println("-----------------------------------------")
    }
}