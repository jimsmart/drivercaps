package denisenkom_test

import (
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/mssql"
)

const (
	PkgURL  = "github.com/denisenkom/go-mssqldb"
	PkgName = PkgURL
	DrvName = "mssql"
)

var ConnStr = fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s", mssql.User, mssql.Pass, mssql.Host, mssql.Port)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(mssql.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
