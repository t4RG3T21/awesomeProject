package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name   string
	Price  float64
	Rating float64
}

type Comparator struct {
	Field string // Поле для сортировки: "name", "price", "rating"
}

type ProductSorter struct {
	Products   []Product
	Comparator Comparator
}

func (ps ProductSorter) Len() int {
	return len(ps.Products)
}

func (ps ProductSorter) Swap(i, j int) {
	ps.Products[i], ps.Products[j] = ps.Products[j], ps.Products[i]
}

func (ps ProductSorter) Less(i, j int) bool {
	p1 := ps.Products[i]
	p2 := ps.Products[j]

	switch ps.Comparator.Field {
	case "price":
		return p1.Price < p2.Price
	case "rating":
		return p1.Rating > p2.Rating
	default:
		return p1.Name < p2.Name
	}
}

func SortProducts(products []Product, field string) {
	sorter := ProductSorter{
		Products:   products,
		Comparator: Comparator{Field: field},
	}
	sort.Sort(sorter)
}

func main() {
	products := []Product{
		{"Laptop", 1200.0, 4.5},
		{"Smartphone", 800.0, 4.2},
		{"Tablet", 500.0, 4.7},
		{"Headphones", 150.0, 4.3},
	}

	SortProducts(products, "name")
	fmt.Println("По имени:")
	for _, p := range products {
		fmt.Printf("%-10s $%.2f (рейтинг: %.1f)\n", p.Name, p.Price, p.Rating)
	}

	SortProducts(products, "price")
	fmt.Println("\nПо цене (возрастание):")
	for _, p := range products {
		fmt.Printf("%-10s $%.2f\n", p.Name, p.Price)
	}

	SortProducts(products, "rating")
	fmt.Println("\nПо рейтингу (убывание):")
	for _, p := range products {
		fmt.Printf("%-10s ★%.1f\n", p.Name, p.Rating)
	}
}
