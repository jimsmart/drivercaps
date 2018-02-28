package drivercaps

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

func renderASCII(results [][]string, test *DriverTest, tim time.Time) {
	writeASCII(os.Stdout, results, test, tim)
}

func writeASCIIFile(results [][]string, test *DriverTest, tim time.Time) {
	filename := localFilename(asciiFilename)
	// log.Println("filename", filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufw := bufio.NewWriter(f)
	writeASCII(bufw, results, test, tim)
	bufw.Flush()
}

func writeASCII(w io.Writer, results [][]string, test *DriverTest, tim time.Time) {

	header := []string{
		"DDL Definition",
		".Name",
		".DBTypeName",
		".Nullable",
		".DecimalSize",
		".Length",
		".ScanType",
	}

	table := tablewriter.NewWriter(w)
	// table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	// table.SetHeader(results[0])
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	// table.SetColumnAlignment([]int{
	// 	tablewriter.ALIGN_LEFT,
	// 	tablewriter.ALIGN_LEFT,
	// 	tablewriter.ALIGN_LEFT,
	// 	tablewriter.ALIGN_LEFT,
	// 	tablewriter.ALIGN_CENTER,
	// 	tablewriter.ALIGN_RIGHT,
	// 	tablewriter.ALIGN_LEFT,
	// })
	// table.SetCaption(true, title)

	// Markdown format.
	// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	// table.SetCenterSeparator("|")

	results = results[1:len(results)]

	for i := range results {
		table.Append(results[i])
	}
	table.Render()
}
