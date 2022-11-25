package scrapper

import (
	"strconv"
	"strings"
	"webs/structs"
)

func FormatUrls(product string, url string, index int64, finalIndex int64, structChan chan structs.Product) {

	defer close(structChan)

	var linksChan []chan []structs.Product
	var pageUrl string

	product = strings.Replace(product, " ", "-", -1)
	url = strings.Replace(url, "<<PRODUCT>>", product, 1)

	productIndex := 0
	for i := 1; i <= int(50*int64(finalIndex)+1); i += 50 {
		linksChan = append(linksChan, make(chan []structs.Product))
		formattedIndex := strconv.Itoa(i)
		pageUrl = strings.Replace(url, "<<INDEX>>", formattedIndex, 1)
		go ScrapUrls(pageUrl, linksChan[productIndex])
		productIndex++
	}
	for _, v := range linksChan {
		productValue := <-v
		for i, _ := range productValue {
			structChan <- productValue[i]
		}
	}
}
