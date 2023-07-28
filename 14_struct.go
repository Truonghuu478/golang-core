package main

import "fmt"

type TypeImage struct {
	name  string
	typeF string
}

type Product struct {
	prdName  string
	prdPrice float64
	prdImage TypeImage
}

func main() {
	var listProduct = []Product{
		{
			prdName:  "A",
			prdPrice: 100000,
			prdImage: TypeImage{
				name:  "adsadasdsa",
				typeF: "png",
			},
		},
		{
			prdName:  "A",
			prdPrice: 100000,
			prdImage: TypeImage{
				name:  "adsadasdsa",
				typeF: "png",
			},
		},
	}

	for _, product := range listProduct {

		fmt.Println("Product:", product.prdName)
	}
}
