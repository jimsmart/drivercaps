package drivercaps

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"

	"github.com/djherbis/times"
)

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
