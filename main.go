package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/AJMBrands/SoftwareThatMatters/ingest"
)

const limited_discovery_query string = "https://libraries.io/api/search?api_key=3dc75447d3681ffc2d17517265765d23&page=1&per_page=1&platforms=NPM"

const discovery_query string = "https://libraries.io/api/search?api_key=3dc75447d3681ffc2d17517265765d23&platforms=NPM&per_page=20"

const offline_file string = "data/100packages.json"
const versionPath string = "data/out/versions.csv"
const outPathTemplate string = "data/out/streamedout-%s.json"

var m1, m2 runtime.MemStats
var t1, t2 time.Time

//TODO: Make ingest process and file writing scalable
func main() {
	runtime.ReadMemStats(&m1) // Reading memory stats for debug purposes
	t1 = time.Now()
	//ingest.IngestFile(offline_file, outPath)
	// ingest.Ingest(limited_discovery_query, outPathTemplate, versionPath)

	amount := ingest.StreamParse("data/input.json", outPathTemplate)
	ingest.MergeJSON(outPathTemplate, amount)
	runtime.ReadMemStats(&m2)
	t2 = time.Now()
	fmt.Printf("Took %d ms", t2.UnixMilli()-t1.UnixMilli())
}
