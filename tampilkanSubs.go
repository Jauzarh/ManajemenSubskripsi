package main

import "fmt"

func menuTampilan(subs []Subkripsi) {
    for {
        fmt.Println("\n    Lihat Subkripsi    ")
        fmt.Println("1. Urutkan berdasarkan Harga (Termahal)")
        fmt.Println("2. Urutkan berdasarkan Kategori")
        fmt.Println("3. Kembali ke Menu Utama")
        
        var pilihan int
        fmt.Print("Pilih opsi: ")
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            sesuaiHarga(subs)
        case 2:
            kategori(subs)
        case 3:
            return
        default:
            fmt.Println("Pilihan tidak valid!")
        }
    }
}

func tampilkanTabel(subs []Subkripsi) {
    fmt.Println("---------------------------------------------------------------------------")
    fmt.Println("| No |      Nama      |   Harga   | Jatuh Tempo  |  Status  |  Kategori   |")
    fmt.Println("----------------------------------------------------------------------------")
    
    for i, sub := range subs {
        status := "Nonaktif"
        if sub.Aktif {
            status = "Aktif"
        }
        
        fmt.Printf("| %-2d | %-14s | Rp%-7.0f | %-12s | %-8s | %-11s |\n", 
            i+1,
            sub.Nama,
            sub.Harga,
            sub.tanggalBiling,
            status,
            sub.Kategori)
    }
    
    fmt.Println("----------------------------------------------------------------------------")
}

func sesuaiHarga(subk []Subkripsi) {
    urutan := make([]Subkripsi, len(subk))
    copy(urutan, subk)
    
    n := len(urutan)
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if urutan[j].Harga > urutan[maxIdx].Harga {
                maxIdx = j
            }
        }
        urutan[i], urutan[maxIdx] = urutan[maxIdx], urutan[i]
    }

    fmt.Println("\nDaftar Subkripsi (urut harga termahal):")
    tampilkanTabel(urutan)
}

func kategori(subk []Subkripsi) {
    urutan := make([]Subkripsi, len(subk))
    copy(urutan, subk)
    
    n := len(urutan)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if urutan[j].Kategori < urutan[minIdx].Kategori {
                minIdx = j
            }
        }
        urutan[i], urutan[minIdx] = urutan[minIdx], urutan[i]
    }

    fmt.Println("\nDaftar Subkripsi (urut kategori):")
    tampilkanTabel(urutan)
}