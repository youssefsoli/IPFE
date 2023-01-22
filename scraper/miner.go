package main

type AllMiners []Miner

type Miner struct {
	Addr            string      `json:"addr"`
	Name            string      `json:"name"`
	Suspended       bool        `json:"suspended"`
	Version         string      `json:"version"`
	ChainInfo       interface{} `json:"chain_info"`
	SuspendedReason string      `json:"suspendedReason,omitempty"`
}
