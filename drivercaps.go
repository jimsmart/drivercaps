package drivercaps

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jimsmart/schema"
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
	PKType    string
}

type ColumnDefn struct {
	DDL string // e.g. VARCHAR(255)
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
	defer func() {
		err = db.Close()
		if err != nil {
			log.Printf("db.Close error %v", err)
		}
	}()

	types := make([]string, len(test.Columns)*2)
	cis := make([]*sql.ColumnType, len(test.Columns)*2)
	i := 0

	testType := func(typeDDL string) {
		j := i
		i++
		tname := "drivercaps_" + randomString(8)
		// log.Println(tname)
		create := fmt.Sprintf("CREATE TABLE %s (\n\tpk %s PRIMARY KEY,\n", tname, test.PKType)
		s := fmt.Sprintf("t_%d %s", j, typeDDL)
		types[j] = s
		create += fmt.Sprintf("\t%s\n)", s)
		err := exec(db, create)
		if err != nil {
			log.Printf("create error %v", err)
			return
		}

		// query caps
		ci, err := schema.Table(db, tname)
		if err != nil {
			log.Printf("metadata query error %v", err)
			return
		}
		cis[j] = ci[1]

		err = exec(db, "DROP TABLE "+tname)
		if err != nil {
			log.Printf("drop error %v", err)
		}
	}

	for _, t := range test.Columns {
		testType(t.DDL)
	}
	for _, t := range test.Columns {
		testType(t.DDL + " NOT NULL")
	}

	return assembleResults(types, cis), nil
}

func assembleResults(types []string, ci []*sql.ColumnType) [][]string {
	var results [][]string

	header := []string{
		"DDL Definition",
		".Name",
		".DBTypeName",
		".Nullable",
		".DecimalSize",
		".Length",
		".ScanType",
	}
	results = append(results, header)

	for i := range ci {

		nullable, nullableok := ci[i].Nullable()
		nullableStr := "-"
		if nullableok {
			nullableStr = strconv.FormatBool(nullable)
		}

		precision, scale, decsizeok := ci[i].DecimalSize()
		decsizeStr := "-"
		if decsizeok {
			decsizeStr = fmt.Sprintf("(%s,%s)", formatInt64(precision), formatInt64(scale))
		}

		length, lengthok := ci[i].Length()
		lengthStr := "-"
		if lengthok {
			lengthStr = formatInt64(length)
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

func formatInt64(n int64) string {
	if n == math.MaxInt64 {
		return "MaxInt64"
	} else {
		return strconv.FormatInt(n, 10)
	}
}

func localFilename(filename string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd + filename
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
