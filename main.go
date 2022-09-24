package main

import (
	"bufio"
	"fmt"
	"log"
	Mods "main/Modules"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

var flags = pflag.FlagSet{SortFlags: false}

func init() {
	Mods.TextBanner.FNAME = "Modules/Text/banner.txt"
	Mods.TextBanner.Output_banner()
	flags.StringVarP(&Mods.O.FileFormat, "format", "f", "", "| Set the needed file format, all file formats must start like so *.")
	flags.BoolVarP(&Mods.O.Verbose, "verbose", "v", false, "| Set the verbosity, more logging and debug output")
	flags.StringVarP(&Mods.O.Output, "output", "o", "", " | Set the output file, this is required")
	flags.StringVarP(&Mods.O.Query, "query", "q", "", " | Set the google dorking query such as --query='RTFM filetype:pdf'")
	flags.IntVarP(&Mods.O.Resultspp, "results", "r", 10, "| Set the needed amount of results per page, default=10")
	flags.IntVarP(&Mods.O.Crawlpages, "pages", "p", 3, " | Set the needed amount of pages to crawl, default=3")
	flags.StringVarP(&Mods.O.ListPath, "listp", "t", "", "| Set the list filepath for a list of queries to run")
	flags.BoolVarP(&Mods.O.List, "list", "l", false, " | Required value and flag in order to define a path")
	flags.BoolVarP(&Mods.O.Help, "help", "h", false, " | Help menu")
	flags.Parse(os.Args[1:])
	if Mods.O.Help {
		fmt.Println(`
		
		Flag options:
			--format/-f  | This will set the file format, which will be the file extension you are dorking for, for example (*.pdf)
						  This FLAG MUST HAVE THE PREFIX *. BEFORE THE EXTENSION THIS IS REQUIRED
				
			--verbose/-v | This sets verbosity to true or false, do you want large output if so use it if not dont

			
			--output/-o  | This will set the output directory of all downloaded files this is REQUIRED

			
			--query/-q   | This will set the google dorking query, this is REQUIRED

			
			--results/-r | This will set the amount of results per page, for example for every one page you want to output 5 links you would use --results=5


			--pages/-p   | This will set the amount of pages you want to search, for example if you want 100 pages use --pages=100

			--listp/-t   | This is an option if you want to use a list of queries, this would be the file to the query list / dork list

			--list/-l    | This is an option that NEEDS to be set BEFORE you use --listp example (--list --listp="filepath to dorks")
		

			==========================================================================================================================

			Example usage;

			Normal        1: 			go run main.go --format=(*.pdf) --verbose --output="/home/user/Desktop/File" --query="BTFM filetype:pdf" 
			File of dorks 2:            go run main.go --format=(*.doc) --verbose --output="/home/user/Desktop/File" --list --listp="/home/fileofdorks.txt"
			Results       3:            go run main.go --format=(*.pdf) --verbose --output="/home/user/Desktop/File" --query="BTFM filetype:pdf" --results=900 --pages=1000
		`)
		os.Exit(0)
	}
	if Mods.O.Output == "" {
		fmt.Println("[!] OUTPUT FILE MUST BE SET! (-o/--output) EXAMPLE: --output /home/user/desktop")
		os.Exit(0)
	}
	if Mods.O.FileFormat == "" || !strings.Contains(Mods.O.FileFormat, "*.") {
		fmt.Println("ERR1     => [!] MUST CHOOSE A FILE FORMAT FROM THE LIST BELOW example: --format='*.pdf' ")
		for _, l := range Mods.Signatures {
			fmt.Println("\n\t | ", l.SuffixFile, " \n\t\t| ", l.FileFormat)
		}
		os.Exit(0)
	}
	if Mods.O.Query == "" {
		fmt.Println("[!] ERROR: Use the --query/-q option to set a dorkable query")
		os.Exit(0)
	}
	if Mods.O.List {
		if Mods.O.ListPath == "" {
			fmt.Println("[!] ERROR: In order to use a list you must provide a non empty value usage for the flag `--listp/-t`")
			os.Exit(0)
		} else {
			f, x := os.Open(Mods.O.ListPath)
			if x != nil {
				log.Fatal(x)
			} else {
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					Mods.Caller(Mods.O.Resultspp, Mods.O.Crawlpages, scanner.Text(), "google")
					Mods.Directories()
					filessaved := len(Mods.FP.Saves)
					fmt.Println("=================== SUMMARY =============== ")
					fmt.Println("Files saved to     -> ", Mods.O.Output)
					fmt.Println("Query googled      -> ", Mods.O.Query)
					fmt.Println("Results per page   -> ", Mods.O.Resultspp)
					fmt.Println("Setting verbose    -> ", Mods.O.Verbose)
					fmt.Println("results searched   -> ", Mods.O.Resultspp)
					fmt.Println("GOOD Requests made -> ", Mods.Counter_requests)
					fmt.Println("Files Saved        -> ", filessaved)
					os.Exit(0)
				}
			}
		}
	}
	if Mods.O.List && Mods.O.ListPath != "" {
		fmt.Println("[!] ERROR: In order to use the list of queries you MUST use the flag --listp/p")
	}
}

func main() {
	Mods.Caller(Mods.O.Resultspp, Mods.O.Crawlpages, Mods.O.Query, "google")
	Mods.Directories()
	filessaved := len(Mods.FP.Saves)
	fmt.Println("=================== SUMMARY =============== ")
	fmt.Println("Files saved to     -> ", Mods.O.Output)
	fmt.Println("Query googled      -> ", Mods.O.Query)
	fmt.Println("Results per page   -> ", Mods.O.Resultspp)
	fmt.Println("Setting verbose    -> ", Mods.O.Verbose)
	fmt.Println("results searched   -> ", Mods.O.Resultspp)
	fmt.Println("GOOD Requests made -> ", Mods.Counter_requests)
	fmt.Println("Files Saved        -> ", filessaved)
}
