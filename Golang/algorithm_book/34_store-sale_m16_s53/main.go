package main

import (
	"fmt"
)

type Store struct {
	Name        string
	WeeklySales float64
	SellerShare float64
}

func main() {
	var numStores int
	fmt.Print("Enter the number of chain stores: ")
	fmt.Scanln(&numStores)

	stores := make([]Store, numStores)
	maxSales := 0.0

	for i := 0; i < numStores; i++ {
		fmt.Printf("Enter the name of store %d: ", i+1)
		fmt.Scanln(&stores[i].Name)

		fmt.Printf("Enter the weekly sales for store %d: ", i+1)
		fmt.Scanln(&stores[i].WeeklySales)

		if stores[i].WeeklySales > maxSales {
			maxSales = stores[i].WeeklySales
		}

		if stores[i].WeeklySales > 2000 {
			stores[i].SellerShare = stores[i].WeeklySales * 0.05
		} else if stores[i].WeeklySales > 1000 {
			stores[i].SellerShare = stores[i].WeeklySales * 0.03
		}
	}

	fmt.Println("\nSeller's Share:")
	for _, store := range stores {
		fmt.Printf("Store: %s\nSeller's Share: $%.2f\n\n", store.Name, store.SellerShare)
	}

	fmt.Printf("Store with the highest sales: %s\n", getStoreWithHighestSales(stores, maxSales))
}

func getStoreWithHighestSales(stores []Store, maxSales float64) string {
	for _, store := range stores {
		if store.WeeklySales == maxSales {
			return store.Name
		}
	}
	return ""
}
