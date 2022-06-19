# SoftwareThatMatters

## Instructions
This document will help one reproduce the results mentioned in Analyzing the Criticality of NPM Packages Through a Time-Dependent Dependency Graph

To set up go dependencies, run the following in the root directory (optional, since go should download deps automatically) :
```
go mod download
```

To run the graph algorithm, simply write in the terminal:
```
go run main.go start
```
This will open up a cli where various commands can be used.

### License
The code's main license can be found in LICENSE.
It also re-uses some modified gonum code, for which the license can be found in GONUM_LICENSE