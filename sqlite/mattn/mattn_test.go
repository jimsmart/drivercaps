package mattn_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/sqlite"
)

const (
	PkgURL  = "github.com/mattn/go-sqlite3"
	PkgName = PkgURL
	DrvName = "sqlite3"
	ConnStr = ":memory:"
)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(sqlite.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
