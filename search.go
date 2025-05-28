package main
import "fmt"

func searchSubscriptions() {
	fmt.Println("\nCARI :")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Kategori")
	fmt.Print("Pilih metode pencarian: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		searchSubscriptionByName()
	case 2:
		searchByCategory()
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func sequentialSearch(subscriptions []Subscription, name string) int {
	foundIndex := -1
	i := 0

	for foundIndex == -1 && i < len(subscriptions) {
		if subscriptions[i].Nama == name {
			foundIndex = i
		}
		i++
	}

	return foundIndex
}

func searchSubscriptionByName() {
	fmt.Print("Masukkan nama subscription yang ingin dicari: ")
	var searchName string
	fmt.Scanln(&searchName)

	index := sequentialSearch(subscriptions, searchName)

	if index != -1 {
		fmt.Println("\nSubscription ditemukan pada indeks", index)
		subkripsi(subscriptions[index])
	} else {
		fmt.Println("Subscription tidak ditemukan")
	}
}

func searchByCategory() {
	fmt.Print("Masukkan kategori yang dicari: ")
	var searchCategory string
	fmt.Scanln(&searchCategory)

	low := 0
	high := len(subscriptions) - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2

		if subscriptions[mid].Kategori == searchCategory {
			ketemu = true
			fmt.Println("\nHasil Pencarian:")
			subkripsi(subscriptions[mid])

			left := mid - 1
			for left >= 0 && subscriptions[left].Kategori == searchCategory {
				subkripsi(subscriptions[left])
				left--
			}

			right := mid + 1
			for right < len(subscriptions) && subscriptions[right].Kategori == searchCategory {
				subkripsi(subscriptions[right])
				right++
			}
			break
		} else if subscriptions[mid].Kategori < searchCategory {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ditemukan subscription dengan kategori tersebut")
	}
}

