import json
import numpy as np
from sklearn.neighbors import NearestNeighbors
import math
import seaborn as sns


if __name__ == "__main__":
    # data = json.load(open('graph_data.json'))
    #
    # file_types = set()
    # for node in data['nodes']:
    #     type_entry = node['name'].split(', ')[-1]
    #     type = type_entry.split(': ')[1]
    #     file_types.add(type)
    #
    # print(file_types)

    print(sns.color_palette('muted'))
