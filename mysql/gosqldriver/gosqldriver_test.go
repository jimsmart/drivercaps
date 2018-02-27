package gosqldriver_test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/mysql"
)

const (
	PkgURL  = "github.com/go-sql-driver/mysql"
	PkgName = PkgURL
	DrvName = "mysql"
)

var ConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysql.User, mysql.Pass, mysql.Host, mysql.Port, mysql.DB)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(mysql.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
