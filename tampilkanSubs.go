package main
import "fmt"

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
        fmt.Println("-----------------------------------------")
    }
}