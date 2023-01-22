package main

var (
	EndpointAPI         = "https://api.estuary.tech/"
	EndpointPublic      = EndpointAPI + "public/"
	EndpointAllMiners   = EndpointPublic + "miners/"
	EndpointMinerDeals  = func(minerAddr string) string { return EndpointAllMiners + "deals/" + minerAddr }
	EndpointByCID       = func(CID string) string { return EndpointPublic + "by-cid/" + CID }
	EndpointContentData = func(CID string) string { return "https://" + CID + ".ipfs.dweb.link/" }
)
