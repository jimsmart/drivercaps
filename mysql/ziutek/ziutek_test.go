package ziutek_test

import (
	"fmt"
	"testing"

	_ "github.com/ziutek/mymysql/godrv"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/mysql"
)

const (
	PkgURL  = "github.com/ziutek/mymysql"
	PkgName = "github.com/ziutek/mymysql/godrv"
	DrvName = "mymysql"
)

var ConnStr = fmt.Sprintf("tcp:%s:%s*%s/%s/%s", mysql.Host, mysql.Port, mysql.DB, mysql.User, mysql.Pass)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(mysql.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
