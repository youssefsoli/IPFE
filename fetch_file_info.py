import requests as r
import json


if __name__ == "__main__":
    headers = {'Accept': 'application/json'}

    with open('ids.json', 'r') as f:
        content_ids = json.load(f)

    for content_id in content_ids[10:30]:
        try:
            content_response = r.request("HEAD", f"https://api.estuary.tech/get/{content_id}", headers=headers)
            print('------')
            print(f'https://api.estuary.tech/get/{content_id}')
            print(content_response.headers)

        except:
            pass
