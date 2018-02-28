package drivercaps

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
		s := fmt.Sprintf("t_%d %s", i, ci.DDL)
		types = append(types, s)
		cmd = append(cmd, "\t"+s)
		i++
	}
	for _, ci := range test.Columns {
		if ci.OnePerTable {
			continue
		}
		s := fmt.Sprintf("t_%d %s NOT NULL", i, ci.DDL)
		types = append(types, s)
		cmd = append(cmd, "\t"+s)
		i++
	}
	cmd = append(cmd, "\tPRIMARY KEY (t_0)")
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
		".Name",
		".DBTypeName",
		".Nullable",
		".DecimalSize",
		".Length",
		".ScanType",
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
				precStr = "MaxInt64"
			} else {
				precStr = strconv.FormatInt(precision, 10)
			}
			// TODO Factor this out.
			scaleStr := ""
			if scale == math.MaxInt64 {
				scaleStr = "MaxInt64"
			} else {
				scaleStr = strconv.FormatInt(scale, 10)
			}
			decsizeStr = fmt.Sprintf("(%s,%s)", precStr, scaleStr)
		}

		length, lengthok := ci[i].Length()
		lengthStr := "-"
		if lengthok {
			if length == math.MaxInt64 {
				lengthStr = "MaxInt64"
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
