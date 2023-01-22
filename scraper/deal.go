package main

import (
	"time"
)

type AllDeals []Deal

type Deal struct {
	ID               int       `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Content          int       `json:"content"`
	PropCid          string    `json:"propCid"`
	Miner            string    `json:"miner"`
	DealID           int       `json:"dealId"`
	Failed           bool      `json:"failed"`
	Verified         bool      `json:"verified"`
	FailedAt         time.Time `json:"failedAt"`
	DtChan           string    `json:"dtChan"`
	TransferStarted  time.Time `json:"transferStarted"`
	TransferFinished time.Time `json:"transferFinished"`
	OnChainAt        time.Time `json:"onChainAt"`
	SealedAt         time.Time `json:"sealedAt"`
	ContentCid       string    `json:"contentCid"`
}
