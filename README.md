# 🪐 IPFE: InterPlanetary File Explorer

A 3D explorer of NLP-aggregated interconnected planetary files. Navigate the universe of distributed data by exploring connections between them.

### What does it do?

This project visualizes the vast quantities of data stored on the InterPlanetary File System using an intuitive 3D graph. Connected nodes are nearby and have similar content; all nodes are colour-coded based on their file type, with larger nodes representing larger files. Hovering over a node gives more information about the file and clicking on the node downloads the file.

Nodes can be dragged with the cursor and the view of the graph can be zoomed in or out with the scroll wheel.

### How did we do it?

We built this application using 3 technologies: Golang, Python, and Three.js (JavaScript). Behind the scenes we used the Estuary to interface with files in the IPFS and Co:here's Embed platform in order to quantify the similarity of files.

Our pipeline consisted of:

1. Fetching the headers of 69,512 files on an IPFS with Estuary (only 1,998 files are visualized);
2. Embedding the texts into vectors with Co:here APIs;
3. Performing a reduction in vector space dimension with principal component analysis;
4. Classifying texts based on k-nearest neighbors;
5. Visualizing the resulting neighbors as a 3D graph.

---

Entered as a project for UofTHacks X 2023. For more information and details, see the [DevPost here](https://devpost.com/software/changeme).
