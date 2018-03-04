package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

func main() {
	NS1_APIKEY := os.Getenv("NS1_APIKEY")
	if NS1_APIKEY == "" {
		fmt.Println("NS1_APIKEY environment variable is not set, giving up")
		os.Exit(1)
	}
	NS1_ZONE := os.Getenv("NS1_ZONE")
	if NS1_ZONE == "" {
		fmt.Println("NS1_ZONE environment variable is not set, giving up")
		os.Exit(1)
	}
	NS1_RECORD := os.Getenv("NS1_RECORD")
	if NS1_RECORD == "" {
		fmt.Println("NS1_RECORD environment variable is not set, giving up")
		os.Exit(1)
	}
	NS1_TYPE := os.Getenv("NS1_TYPE")
	if NS1_TYPE == "" {
		fmt.Println("NS1_TYPE environment variable is not set, giving up")
		os.Exit(1)
	}
	DEFAULT_IPV4 := os.Getenv("DEFAULT_IPV4")
	if DEFAULT_IPV4 == "" {
		fmt.Println("DEFAULT_IPV4 environment variable is not set, giving up")
		os.Exit(1)
	}
	httpClient := &http.Client{Timeout: time.Second * 10}
	client := api.NewClient(httpClient, api.SetAPIKey(NS1_APIKEY))
	zones, _, err := client.Zones.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Println(z.Zone)
	}
	record, _, err := client.Records.Get(NS1_ZONE, NS1_RECORD, NS1_TYPE)
	if err != nil {
		log.Fatal(err)
	}
	if len(record.Answers) > 1 {
		fmt.Println("Answer length is greater than a single record")
	} else if len(record.Answers) == 0 {
		updateRecord(*client, NS1_ZONE, NS1_RECORD, NS1_TYPE, DEFAULT_IPV4)
	} else {
		fmt.Println("Current dns record value: ", record.Answers[0].String())
		if record.Answers[0].String() != DEFAULT_IPV4 {
			updateRecord(*client, NS1_ZONE, NS1_RECORD, NS1_TYPE, DEFAULT_IPV4)
		}
	}
}

func updateRecord(client api.Client, zone string, domain string, dnsType string, answer string) {
	fmt.Println("Updating DNS")
	aRecord := dns.NewRecord(zone, domain, dnsType)
	aRecord.AddAnswer(dns.NewAv4Answer(answer))
	_, err := client.Records.Update(aRecord)
	if err != nil {
		log.Fatal(err)
	}

	//recordService, _ , err := client.Records.Update(aRecord)
}
