package scrapper

import (
	"fmt"
	"webs/structs"

	"github.com/gocolly/colly"
)

func ScrapUrls(url string, productChan chan []structs.Product) {
	defer close(productChan)

	var productChans []chan structs.Product
	products := []structs.Product{}
	collyCollector := colly.NewCollector()
	var index int = 0
	collyCollector.OnHTML(".ui-search-item__group__element.shops__items-group-details.ui-search-link", func(e *colly.HTMLElement) {
		productChans = append(productChans, make(chan structs.Product))
		go ExtractProductData(e.Attr("href"), productChans[index])
		index++
	})

	err := collyCollector.Visit(url)

	if err != nil {
		fmt.Printf("Cannot get URL: %s \n ", url)
	}

	for _, v := range productChans {
		products = append(products, <-v)
	}
	productChan <- products
}
