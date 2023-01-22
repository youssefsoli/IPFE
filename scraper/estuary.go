package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func fetchUniqueCIDs(client *http.Client) []string {
	allMiners := AllMiners{}
	if err := fetchAllMiners(client, &allMiners); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("allMiners len: %d\n", len(allMiners))

	CIDs := make(map[string]struct{})

	allDeals := AllDeals{}
	dealNum := 0
	for minerNum, miner := range allMiners {
		if err := fetchDeals(client, miner.Addr, &allDeals); err != nil {
			log.Printf("error fetching miner deals (addr: %s): %v", miner.Addr, err)
			continue
		}

		log.Printf("minerNum: %d\n", minerNum)

		for _, deal := range allDeals {
			dealNum++
			log.Printf("dealNum: %d\n", dealNum)
			CIDs[deal.ContentCid] = struct{}{}
		}
	}

	CIDArr := make([]string, len(CIDs))
	i := 0
	for CID := range CIDs {
		CIDArr[i] = CID
		i++
	}
	return CIDArr
}

func fetchContent(client *http.Client, CID string, resp *ContentResp) error {
	res, err := genericFetchJSON(client, http.MethodGet, EndpointByCID(CID))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(resp)
}

func fetchContentTypeHEAD(client *http.Client, CID string) (string, error) {
	res, err := genericFetchJSON(client, http.MethodHead, EndpointContentData(CID))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return res.Header.Get("Content-Type"), nil
}

func fetchAllMiners(client *http.Client, allMiners *AllMiners) error {
	res, err := genericFetchJSON(client, http.MethodGet, EndpointAllMiners)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(allMiners)
}

// fetchDeals returns the passed miner's deals
func fetchDeals(client *http.Client, minerAddr string, allDeals *AllDeals) error {
	res, err := genericFetchJSON(client, http.MethodHead, EndpointMinerDeals(minerAddr))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(allDeals)
}
