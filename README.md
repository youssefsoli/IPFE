<img src="https://user-images.githubusercontent.com/61663933/214231075-ede4477a-a982-49e0-9f65-fc9ca43ea61b.png" width="400" />

# IPFE: InterPlanetary File Explorer

### Explore the universe of distributed data in 3D – vizualize the connections between them.

---

<a href="https://youtu.be/rYsgF5YAVTI" target="_blank"><img src="https://user-images.githubusercontent.com/61663933/214232047-d5e2fe75-59b2-408c-a686-d50abf11194e.png" height="30" /></a> <a href="https://devpost.com/software/changeme" target="_blank"><img src="https://user-images.githubusercontent.com/61663933/214232534-7ac805d9-580d-4028-bdfd-088c894855e9.png" height="30" /></a>


### [Try it out!](https://ipfe.elguindi.xyz/)

![Planetary Explosion](https://user-images.githubusercontent.com/61663933/214230942-7e0a9879-daa5-4d13-b46f-f16fd0fdbc41.gif)

## What does it do?

This project visualizes the vast quantities of data stored on the InterPlanetary File System using an intuitive 3D graph. Connected nodes are nearby and have similar content; all nodes are colour-coded based on their file type, with larger nodes representing larger files. Hovering over a node gives more information about the file and clicking on the node downloads the file.

Nodes can be dragged with the cursor and the view of the graph can be zoomed in or out with the scroll wheel.

![Project Overview](https://user-images.githubusercontent.com/61663933/214231333-81bc0276-cdcb-4b1e-8527-35ebf1d4166c.png)

## How did we do it?

We built this application using 3 technologies: Golang, Python, and Three.js (JavaScript). Behind the scenes we used the Estuary to interface with files in the IPFS and Co:here's Embed platform in order to quantify the similarity of files.

Our pipeline consisted of:

1. Fetching the headers of 69,512 files on an IPFS with Estuary (only 1,998 files are visualized);
2. Embedding the texts into vectors with Co:here APIs;
3. Performing a reduction in vector space dimension with principal component analysis;
4. Classifying texts based on k-nearest neighbors;
5. Visualizing the resulting neighbors as a 3D graph.

---

Entered as a project for UofTHacks X 2023. For more information and details, see the [DevPost here](https://devpost.com/software/changeme).
