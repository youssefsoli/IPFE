package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

var NUM_WORKERS = runtime.NumCPU()

func main() {
	client := http.Client{}

	// f, err := os.Open("ids.json")
	f, err := os.Open("ids.json")
	if err != nil {
		log.Fatalln("unable to read ids.json", err)
	}

	CIDs := []string{}
	if err := json.NewDecoder(f).Decode(&CIDs); err != nil {
		log.Fatalln("unable to decode ids.json", err)
	}

	typeConditions := Conditions{
		{
			// String must not contain strings in arr
			fn:  func(a, b string) bool { return !strings.Contains(a, b) },
			arr: []string{"octet-stream", "text/plain", "tar", "zip", "application/zstd"},
		},
		{
			// String must not match strings in arr
			fn:  func(a, b string) bool { return a != b },
			arr: []string{""},
		},
	}

	nameConditions := Conditions{
		{
			fn:  func(a, b string) bool { return a != b },
			arr: []string{"", "aggregate"},
		},
		{
			fn:  func(a, b string) bool { return !strings.HasPrefix(a, b) },
			arr: []string{"splitdir/"},
		},
	}

	writec := make(chan []byte)
	go writer(os.Stdout, writec)

	wg := sync.WaitGroup{}
	sem := make(chan struct{}, NUM_WORKERS) // token based semaphore
	for _, CID := range CIDs {
		CID := CID
		sem <- struct{}{} // block or insert a token
		wg.Add(1)
		go func() {
			defer func() {
				<-sem // take out token
				wg.Done()
			}()

			resp := ContentResp{}
			if err := fetchContent(&client, CID, &resp); err != nil {
				log.Printf("CID %s HEAD fail: %v\n", CID, err)
				return
			}

			for _, c := range resp {
				if !nameConditions.IsMet(c.Content.Name) || c.Content.Name == c.Content.Cid || c.Content.Size == 0 {
					continue
				}

				contentType, err := fetchContentTypeHEAD(&client, CID)
				if err != nil || !typeConditions.IsMet(contentType) {
					continue
				}

				body := fmt.Sprintf("---\nFile hash: %s\nFile name: %s\nUpload date: %s\nFile size: %d bytes\n", c.Content.Cid, c.Content.Name, c.Content.UpdatedAt.Local(), c.Content.Size)
				writec <- []byte(fmt.Sprintf("%sFile type: %s\n", body, contentType))
			}
		}()
	}

	wg.Wait()
	close(sem)
}

func writer(dst io.Writer, c chan []byte) {
	for data := range c {
		_, err := dst.Write(data)
		if err != nil {
			// TODO: send errors on a channel
		}
	}
}

func genericFetchJSON(client *http.Client, method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
