import json
import numpy as np
from sklearn.neighbors import NearestNeighbors
import math


if __name__ == "__main__":
    data = json.load(open('./3D-coordinates.json'))
    texts = data['texts'][:-1]
    embeddings = np.array(data['embeddings'][:-1])
    hashes = open('./hashes.txt').read().split()[2::3]

    nodes = [{
        "id": hashes[i],
        "name": texts[i],
        "coord": embeddings[i]
    } for i in range(len(hashes))]

    links = []

    neighbors = NearestNeighbors(n_neighbors=2, algorithm='ball_tree').fit(embeddings)
    distances, indices = neighbors.kneighbors(embeddings)

    for level, index in enumerate(indices):
        i, j = index
        links.append({
            "source": nodes[i]['id'],
            "target": nodes[j]['id']
        })

        nodes[i]['val'] = math.ceil(1/(distances[level][1]))
        # No longer need the coordinate
        del nodes[i]['coord']

    json_data = {
        "nodes": nodes,
        "links": links
    }

    with open("graph_data.json", "w") as f:
        json.dump(json_data, f)

