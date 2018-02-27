package gwenn_test

import (
	"testing"

	_ "github.com/gwenn/gosqlite"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/sqlite"
)

const (
	PkgURL  = "github.com/gwenn/gosqlite"
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
