import requests as r
import json


if __name__ == "__main__":
    headers = {'Accept': 'application/json'}

    miners_response = r.request("GET", "https://api.estuary.tech/public/miners", headers=headers)
    # addresses = [miner.get("addr") for miner in miners_response.json()]
    miners = miners_response.json()

    content_ids = set()

    for i, miner in enumerate(miners):
        miner_deals_response = r.request("GET", f"https://api.estuary.tech/public/miners/deals/{miner['addr']}",
                                         headers=headers)
        deals = miner_deals_response.json()
        print(f'on miner #{i}\n{len(content_ids)} files so far')
        if deals:
            for deal in deals:
                content_ids.add(deal.get('contentCid'))

    with open('ids.json', 'w') as f:
        json.dump(list(content_ids), f)
