package main

import (
	"log"
	"os"
	"webs/scrapper"
	"webs/structs"

	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	product := "adaptador usb c"
	var index int8 = 1
	var finalIndex int8 = 1
	url := "https://lista.mercadolivre.com.br/<<PRODUCT>>_Desde_<<INDEX>>_NoIndex_True"

	productChan := make(chan structs.Product)
	go scrapper.FormatUrls(product, url, index, finalIndex, productChan)
	products := []structs.Product{}
	for prod := range productChan {
		products = append(products, prod)
	}
	csv, err := gocsv.MarshalString(products)

	if err != nil {
		log.Fatal(err.Error())
	}
	file, _ := os.Create("products.csv")

	file.Write([]byte(csv))
}
