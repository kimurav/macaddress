package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

const API_ADDRESS = "https://api.macaddress.io/v1"

func main() {
	macAddr := flag.String("addr", "", "the mac address to look up")
	secretKey := flag.String("key", "", "the secret api key")
	verbose := flag.Bool("v", false, "print out all data from the response")
	flag.Parse()
	client, err := NewMacAddressClient(API_ADDRESS, *secretKey)
	if err != nil {
		log.Fatal(err)
	}
	data, err := client.GetMacAddressDetails(*macAddr)
	if err != nil {
		log.Fatal(err)
	}
	if !*verbose {
		bytes, err := json.Marshal(map[string]string{"companyName": data.VendorDetails.CompanyName})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", bytes)
	} else {
		bytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", bytes)
	}
}
