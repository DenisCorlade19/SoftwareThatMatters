# SoftwareThatMatters
This document contains the code that was used during the [Research Project](https://github.com/TU-Delft-CSE/Research-Project) (2022 edition) of [TU Delft](https://github.com/TU-Delft-CSE).
## Instructions
This document will help one reproduce the results mentioned in Analyzing the Criticality of Maven Packages Through a Time-Dependent Dependency Graph

To set up go dependencies, run the following in the root directory (optional, since go should download deps automatically) :
```
go mod download
```

To run the graph algorithm, simply write in the terminal:
```
go run main.go start
```
This will open up a cli where various commands can be used.

The project requires a JSON file formatted the following way:
```
{"pkgs":[{
  "name": "react",
  "versions": {
    "1.00": {
      "timestamp": "06-05-2022T10:00:01",
      "dependencies": {
        "name": "^1.0.2"
      }
    }
  }
},
```

To process the packages metadata in this way, more instruction can be found on this [repository](https://github.com/DenisCorlade19/maven-package-metadata)

### License
The code's main license can be found in LICENSE.
It also re-uses some modified gonum code, for which the license can be found in GONUM_LICENSE