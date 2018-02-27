package drivercaps

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/djherbis/times"
	"github.com/jimsmart/schema"
	"github.com/olekukonko/tablewriter"
	// . "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

const (
	csvFilename      = "/report.csv"
	asciiFilename    = "/report.txt"
	markdownFilename = "/README.md"
)

type DriverTest struct {
	Database  string
	PkgURL    string
	PkgName   string
	DrvName   string
	ConnStr   string
	Columns   []*ColumnDefn
	CreateOpt string
}

type ColumnDefn struct {
	DDL         string // e.g. VARCHAR(255)
	OnePerTable bool
	// *sql.ColumnType return values.
	// Name                 string
	// DatabaseTypeName     string
	// DecimalSizeOk        bool
	// DecimalSizePrecision int64
	// DecimalSizeScale     int64
	// LengthOk             bool
	// Length               int64
	// NullableOk           bool
	// Nullable             bool
	// // ScanType reflect.Type
}

func ReportForDriver(test *DriverTest) error {

	results, err := exerciseDriver(test)
	if err == nil && len(results) > 1 {
		writeCSVFile(results, test)
		// renderCSV(results, test)
	}

	// if err != nil {
	// 	return err
	// }

	// TODO Write new CSV on successful results, but always regenerate the other views from existing csv

	results, tim, err2 := readCSVFile()
	if os.IsNotExist(err2) {
		return nil
	}
	if err2 != nil {
		return err2
	}
	if len(results) < 2 {
		return nil
	}

	renderASCII(results, test, tim)
	writeASCIIFile(results, test, tim)
	// renderMarkdown(results, test, tim)
	writeMarkdownFile(results, test, tim)

	return nil
}

func exerciseDriver(test *DriverTest) ([][]string, error) {
	if len(test.Columns) == 0 {
		return nil, nil
	}
	title := fmt.Sprintf("sql.Driver %s %q", test.PkgName, test.DrvName)
	log.Printf(title)
	db, err := sql.Open(test.DrvName, test.ConnStr)
	if err != nil {
		return nil, err
	}

	// tname := "drivercaps_test_table_" + randomString(8)
	tname := "drivercaps_" + randomString(8)
	// log.Println(tname)

	defer func() {
		drop := fmt.Sprintf("DROP TABLE %s", tname)
		exec(db, drop)
		err = db.Close()
		if err != nil {
			log.Printf("db.Close error %v", err)
		}
	}()

	// Build a CREATE TABLE command, using given list of fields.
	create := fmt.Sprintf("CREATE TABLE %s (\n", tname)
	i := 0
	var types []string
	var cmd []string
	// TODO(js) Consider making some columns uppercase, to detect whether driver honours case.
	for _, ci := range test.Columns {
		s := fmt.Sprintf("column_%d %s", i, ci.DDL)
		types = append(types, s)
		cmd = append(cmd, "\t"+s)
		i++
	}
	for _, ci := range test.Columns {
		if ci.OnePerTable {
			continue
		}
		s := fmt.Sprintf("column_%d %s NOT NULL", i, ci.DDL)
		types = append(types, s)
		cmd = append(cmd, "\t"+s)
		i++
	}
	cmd = append(cmd, "\tPRIMARY KEY (column_0)")
	create += strings.Join(cmd, ",\n")
	create += fmt.Sprintf("\n) %s", test.CreateOpt)
	err = exec(db, create)
	if err != nil {
		return nil, err
	}

	ci, err := schema.Table(db, tname)
	if err != nil {
		return nil, err
	}

	return assembleResults(types, ci), nil
}

func assembleResults(types []string, ci []*sql.ColumnType) [][]string {
	var results [][]string

	// header := []string{
	// 	"DDL Definition",
	// 	"Name",
	// 	"DB Type Name",
	// 	"Nullable",
	// 	"Decimal Size",
	// 	"Length",
	// 	"Scan Type",
	// }
	header := []string{
		"DDL Definition",
		"ct.Name",
		"ct.DBTypeName",
		"ct.Nullable",
		"ct.DecimalSize",
		"ct.Length",
		"ct.ScanType",
	}
	results = append(results, header)

	for i := range types {

		nullable, nullableok := ci[i].Nullable()
		nullableStr := "-"
		if nullableok {
			nullableStr = strconv.FormatBool(nullable)
		}

		precision, scale, decsizeok := ci[i].DecimalSize()
		decsizeStr := "-"
		if decsizeok {
			// TODO Factor this out.
			precStr := ""
			if precision == math.MaxInt64 {
				precStr = "math.MaxInt64"
			} else {
				precStr = strconv.FormatInt(precision, 10)
			}
			// TODO Factor this out.
			scaleStr := ""
			if scale == math.MaxInt64 {
				scaleStr = "math.MaxInt64"
			} else {
				scaleStr = strconv.FormatInt(scale, 10)
			}
			decsizeStr = fmt.Sprintf("(%s,%s)", precStr, scaleStr)
		}

		length, lengthok := ci[i].Length()
		lengthStr := "-"
		if lengthok {
			if length == math.MaxInt64 {
				lengthStr = "math.MaxInt64"
			} else {
				lengthStr = strconv.FormatInt(length, 10)
			}
		}

		scantype := ci[i].ScanType()
		scantypeStr := "nil"
		if scantype != nil {
			scantypeStr = scantype.String()
		}

		row := []string{
			// strconv.Quote(types[i]),
			types[i],
			ci[i].Name(),
			ci[i].DatabaseTypeName(),
			nullableStr,
			decsizeStr,
			lengthStr,
			scantypeStr,
		}
		results = append(results, row)
	}

	return results
}

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
		"ct.Name",
		"ct.DBTypeName",
		"ct.Nullable",
		"ct.DecimalSize",
		"ct.Length",
		"ct.ScanType",
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

func renderCSV(results [][]string, test *DriverTest) {
	writeCSV(os.Stdout, results, test)
}

func writeCSVFile(results [][]string, test *DriverTest) {
	filename := localFilename(csvFilename)
	// log.Println("filename", filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufw := bufio.NewWriter(f)
	writeCSV(bufw, results, test)
	bufw.Flush()
}

func writeCSV(w io.Writer, results [][]string, test *DriverTest) {
	csvw := csv.NewWriter(w)
	csvw.WriteAll(results) // calls Flush internally
	if err := csvw.Error(); err != nil {
		log.Fatalf("csv.WriteAll error %v", err)
	}
}

func readCSVFile() ([][]string, time.Time, error) {
	filename := localFilename(csvFilename)
	// log.Println("filename", filename)
	// return nil, nil

	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, time.Time{}, err
	}
	defer f.Close()

	bufr := bufio.NewReader(f)
	csvr := csv.NewReader(bufr)
	results, err := csvr.ReadAll()
	if err != nil {
		return nil, time.Time{}, err
	}

	t, err := times.Stat(filename)
	if err != nil {
		return nil, time.Time{}, err
	}
	return results, t.ModTime(), nil
}

func localFilename(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd + filename
}

func renderMarkdown(results [][]string, test *DriverTest, tim time.Time) {
	writeMarkdown(os.Stdout, results, test, tim)
}

func writeMarkdownFile(results [][]string, test *DriverTest, tim time.Time) {

	// TODO(js) This is a bad code smell. We override the default headings from csv,
	// both here and when dealing with ascii :/
	header := []string{
		"DDL Definition",
		"ct.Name",
		"ct.DBTypeName",
		"ct.Nullable",
		"ct.DecimalSize",
		"ct.Length",
		"ct.ScanType",
	}

	var res [][]string
	res = append(res, header)
	res = append(res, results[1:len(results)]...)
	results = res

	filename := localFilename(markdownFilename)
	// log.Println("filename", filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufw := bufio.NewWriter(f)
	writeMarkdown(bufw, results, test, tim)
	bufw.Flush()
}

// func markdownFilename() string {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return wd + "/README.md"
// }

func writeMarkdown(w io.Writer, results [][]string, test *DriverTest, tim time.Time) {

	// TODO(js) Externalise this (markdown) template.
	const tpl = `
# Driver &#91;&#93;&#42;sql.ColumnType Capability Report

- Package "{{.Package}}" ({{.Driver}})
- {{.Database}}

<table>
	<thead>
		<tr>
			{{range .Headings}}<th>{{ . }}</th>{{end}}
		</tr>
	</thead>
	<tbody>{{range $row := .Rows}}
		<tr>{{range $row}}
			{{if and (ne . "-") (ne . "")}}<td nowrap><code>{{ . }}</code></td>{{end}}{{if eq . "-"}}<td>-</td>{{end}}{{if eq . ""}}<td/>{{end}}{{end}}
		</tr>{{end}}
	</tbody>
</table>

Report for [{{.Package}}](https://{{.URL}}) ({{.Driver}})<br/>
Test timestamp {{.Timestamp}}<br/>
Generated by [drivercaps](https://github.com/jimsmart/drivercaps)

`

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatalf("template.Parse error %v", err)
	}

	data := struct {
		Database  string
		URL       string
		Package   string
		Driver    string
		Headings  []string
		Rows      [][]string
		Timestamp string
	}{
		Database:  test.Database,
		URL:       test.PkgURL,
		Package:   test.PkgName,
		Driver:    test.DrvName,
		Headings:  results[0],
		Rows:      results[1:len(results)],
		Timestamp: tim.Format(time.RFC3339),
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatalf("template.Execute error %v", err)
	}
}

func exec(db *sql.DB, ddl string) error {
	// log.Printf("executing %s", ddl)
	_, err := db.Exec(ddl)
	if err != nil {
		log.Printf("db.Exec error %v for %s", err, ddl)
		return err
	}
	// log.Println("executed ok")
	return nil
}

// const unmistakableChars = "23456789ABCDEFGHJKLMNPQRSTWXYZabcdefghijkmnopqrstuvwxyz"
const unmistakableChars = "23456789abcdefghijkmnopqrstuvwxyz"

func randomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = unmistakableChars[r.Int63()%int64(len(unmistakableChars))]
	}
	return string(b)
}
