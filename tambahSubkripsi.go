package main
import "fmt"

func tambahSubkripsi() {
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
