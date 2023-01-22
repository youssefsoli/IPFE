package main

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
)

type CIDDir []CIDDirEntry

type CIDDirEntry struct {
	filename string
	cid      string
}

func fetchCIDDirHTML(client *http.Client, CID string) (*http.Response, error) {
	return genericFetchJSON(client, http.MethodGet, EndpointsDirCID(CID))
}

func parseCIDHTML(CID string, r io.Reader) CIDDir {
	tokenizer := html.NewTokenizer(r)

	cidDir := CIDDir{}
	prevIsLi := false
	isLi := false
	isA := false

PARSE:
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			break PARSE

		case html.StartTagToken:
			t := tokenizer.Token()

			prevIsLi = isLi
			isLi = t.Data == "li"
			isA = t.Data == "a"

		case html.TextToken:
			t := tokenizer.Token()

			if prevIsLi && isA {
				cidDir = append(cidDir, CIDDirEntry{
					filename: t.Data,
					cid:      CID + "/" + t.Data,
				})
			}

			isLi = false
			isA = false
		}
	}

	return cidDir
}
