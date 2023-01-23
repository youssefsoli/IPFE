import json
import numpy as np
from sklearn.neighbors import NearestNeighbors
import math


if __name__ == "__main__":
    data = json.load(open('./3D-coordinates-new.json'))
    texts = data['texts'][:-1]
    embeddings = data['embeddings'][:-1]
    hashes = open('./hashes-new.txt').read().split()[2::3]

    nodes = [{
        "id": i,
        "hash": hashes[i],
        "name": texts[i],
        "coord": embeddings[i]
    } for i in range(len(hashes))]

    file_types = {}
    max_group_id = 0

    deletion_queue = []
    for node in nodes:
        entry = node['name'].split(', ')
        if entry[0].startswith('File size'):
            deletion_queue.append(node)
        type_entry = entry[-1]
        type = type_entry.split(': ')[1]

        size_entry = entry[-2]
        size = size_entry.split(': ')[1]

        if type not in file_types:
            max_group_id += 1
            file_types[type] = max_group_id

        node['group'] = file_types[type]
        node['type'] = type
        node['val'] = math.log10(int(size.split(' ')[0]))
        node['size'] = int(size.split(' ')[0])

    for node in deletion_queue:
        nodes.remove(node)

    embeddings = np.array([node['coord'] for node in nodes])

    links = []

    neighbors = NearestNeighbors(n_neighbors=2, algorithm='ball_tree').fit(embeddings)
    distances, indices = neighbors.kneighbors(embeddings)

    for level, index in enumerate(indices):
        i = index[0]
        for j in index[1:]:
            links.append({
                "source": nodes[i]['id'],
                "target": nodes[j]['id']
            })

        # No longer need the coordinate
        del nodes[i]['coord']

    json_data = {
        "nodes": nodes,
        "links": links
    }

    with open("graph_data.json", "w") as f:
        json.dump(json_data, f)
