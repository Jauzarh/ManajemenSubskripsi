package main

import "fmt"



func TotalPengeluaran() float64 {
    total := 0.0
    for _, sub := range subscriptions {
        if sub.IsActive {
            total += sub.Price
        }
    }
    return total
}

func TampilkanTotalPengeluaran() {
    total := TotalPengeluaran()
    fmt.Printf("\nTotal Pengeluaran Bulanan: Rp%.0f\n", total)
}
