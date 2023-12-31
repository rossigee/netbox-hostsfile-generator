package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

type IPRecord struct {
	Address     string `json:"address"`
	DNSName     string `json:"dns_name"`
	Description string `json:"description"`
}

type NetboxResponse struct {
	Results []IPRecord `json:"results"`
	Next    string     `json:"next"`
}

const (
	defaultPath = "/api/ipam/ip-addresses/"
	outputFile  = "hosts.txt"
)

func main() {
	netboxURL := os.Getenv("NETBOX_URL")
	if netboxURL == "" {
		log.Fatal("NETBOX_URL environment variable not set")
	}

	if !strings.HasSuffix(netboxURL, "/") {
		netboxURL += "/"
	}

	netboxToken := os.Getenv("NETBOX_TOKEN")
	if netboxToken == "" {
		log.Fatal("NETBOX_TOKEN environment variable not set")
	}

	ipRecords, err := fetchIPRecords(netboxURL, netboxToken)
	if err != nil {
		log.Fatalf("Failed to fetch IP records: %v", err)
	}

	hostsContent := generateHostsContent(ipRecords)

	if err := ioutil.WriteFile(outputFile, []byte(hostsContent), 0644); err != nil {
		log.Fatalf("Failed to write hosts file: %v", err)
	}

	fmt.Printf("Hosts file created: %s\n", outputFile)
}

func fetchIPRecords(netboxURL, netboxToken string) (map[string]IPRecord, error) {
	client := resty.New()
	url := netboxURL + defaultPath

	ipRecords := make(map[string]IPRecord)

	for {
		response, err := client.R().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", fmt.Sprintf("Token %s", netboxToken)).
			Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch IP records: %v", err)
		}

		var netboxResponse NetboxResponse
		if err := json.Unmarshal(response.Body(), &netboxResponse); err != nil {
			return nil, fmt.Errorf("failed to parse Netbox response: %v", err)
		}

		// Process IP records in the current page
		for _, record := range netboxResponse.Results {
			ipRecords[record.Address] = record
		}

		// Check if there are more pages
		if netboxResponse.Next == "" {
			break
		}

		// Set the next URL for the next page
		url = netboxResponse.Next
	}

	return ipRecords, nil
}

func generateHostsContent(ipRecords map[string]IPRecord) string {
	hostsContent := ""

	for ip, record := range ipRecords {
		if record.DNSName == "" {
			continue // Skip if DNS name is empty
		}

		ipWithoutSubnet := strings.Split(ip, "/")[0]
		comment := ""
		if record.Description != "" {
			comment = fmt.Sprintf("# %s", record.Description)
		}
		hostsContent += fmt.Sprintf("%s\t%s\t%s\n", ipWithoutSubnet, record.DNSName, comment)
	}

	return hostsContent
}
