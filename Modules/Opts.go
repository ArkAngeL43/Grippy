package GRIP

type Opts struct {
	FileFormat     string // File format to look for EX MP4, GIF, PNG ETC
	Output         string // Output file of downloads
	Verbose        bool   // Output logging
	Delte_IF_Wrong bool   // Delete the file if the signature is not the same as a PDF
	Query          string // Dork query
	Resultspp      int    // Results per page
	Crawlpages     int    // how many pages of google to crawl
	List           bool   // Needs to be required for list
	ListPath       string // List of queries
	Help           bool
}
