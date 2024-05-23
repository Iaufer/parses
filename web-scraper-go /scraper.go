package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Product struct {
	url, image, name, price string
}

func main() {
	c := colly.NewCollector()
	var products []Product

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := Product{}

		product.url = e.ChildAttr("a", "href")
		product.image = e.ChildAttr("img", "src")
		product.name = e.ChildText("h2")
		product.price = e.ChildText(".price")
		products = append(products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("prod.csv")

		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)

		headers := []string{
			"url",
			"image",
			"name",
			"price",
		}
		defer writer.Flush()

		writer.Write(headers)

		for _, product := range products{
			r := []string{
				product.url,
				product.image,
				product.name,
				product.price,
			}
			writer.Write(r)
		}



	})

	c.Visit("https://www.scrapingcourse.com/ecommerce/")
	fmt.Println("Hello, World!")
}
