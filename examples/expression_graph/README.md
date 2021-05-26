# Expression Graph

Make sure to have [Graphviz](https://graphviz.org/download/) installed in your system.

Graphviz (the tool is called `dot`) generates SVG and PNG images from *dot* files.

In this example, the computational graph is serialized in this format in order to represent it graphically. Note,
however, that spaGO does not need Graphviz to work.

Run this from the current directory:

```console
go run main.go | dot -Tsvg > example.svg
```

Enjoy ;)