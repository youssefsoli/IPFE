package main

import (
	"time"
)

type ContentResp []struct {
	Content struct {
		ID            int       `json:"id"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
		Cid           string    `json:"cid"`
		Name          string    `json:"name"`
		UserID        int       `json:"userId"`
		Description   string    `json:"description"`
		Size          int64     `json:"size"`
		Type          int       `json:"type"`
		Active        bool      `json:"active"`
		Offloaded     bool      `json:"offloaded"`
		Replication   int       `json:"replication"`
		AggregatedIn  int       `json:"aggregatedIn"`
		Aggregate     bool      `json:"aggregate"`
		Pinning       bool      `json:"pinning"`
		PinMeta       string    `json:"pinMeta"`
		Replace       bool      `json:"replace"`
		Origins       string    `json:"origins"`
		Failed        bool      `json:"failed"`
		Location      string    `json:"location"`
		DagSplit      bool      `json:"dagSplit"`
		SplitFrom     int       `json:"splitFrom"`
		PinningStatus string    `json:"pinningStatus"`
		DealStatus    string    `json:"dealStatus"`
	} `json:"content"`
	Deals []struct {
		ID                  int         `json:"ID"`
		CreatedAt           time.Time   `json:"CreatedAt"`
		UpdatedAt           time.Time   `json:"UpdatedAt"`
		DeletedAt           interface{} `json:"DeletedAt"`
		Content             int         `json:"content"`
		UserID              int         `json:"user_id"`
		PropCid             string      `json:"propCid"`
		DealUUID            string      `json:"dealUuid"`
		Miner               string      `json:"miner"`
		DealID              int         `json:"dealId"`
		Failed              bool        `json:"failed"`
		Verified            bool        `json:"verified"`
		Slashed             bool        `json:"slashed"`
		FailedAt            time.Time   `json:"failedAt"`
		DtChan              string      `json:"dtChan"`
		TransferStarted     time.Time   `json:"transferStarted"`
		TransferFinished    time.Time   `json:"transferFinished"`
		OnChainAt           time.Time   `json:"onChainAt"`
		SealedAt            time.Time   `json:"sealedAt"`
		DealProtocolVersion string      `json:"deal_protocol_version"`
		MinerVersion        string      `json:"miner_version"`
	} `json:"deals"`
}
