from random import randint

buckets = {}

s = "File hash: {}\nFile name: {}\nFile size: {} bytes\nFile type: {}\n"

cur = {
        "File hash": "",
        "File name": "",
        "File size": "",
        "File type": ""
        }

f = open("new.txt", "r")
for line in f:
    if line.startswith("---"):
        continue
    
    line = line.split(":", 1)
    key = line[0]
    val = line[1][:-1]

    cur[key] = val

    if line[0] == "File type":
        val = val.split(";")[0]
        if val in buckets:
            buckets[val].append(cur.copy())
        else:
            buckets[val] = [cur.copy()]

f.close()

NUMFILES = 2000
keys = []
for key in buckets.keys():
    if len(buckets[key]) >= (NUMFILES // len(buckets)):
        keys.append(key)


postings = []

need = 2000 // len(keys)
for key in keys:
    bucket = buckets[key]
    bucketLen = len(bucket)
    start = 0
    if bucketLen - need - 1 > 0:
        start = randint(0, bucketLen - need - 1)

    for i in range(start, start + need, 1):
        postings.append(bucket[i].copy())

for p in postings:
    print("---")
    for key in p:
        print(f"{key}: {p[key]}")
