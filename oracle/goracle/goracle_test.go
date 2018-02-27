package goracle_test

import (
	"fmt"
	"testing"

	_ "gopkg.in/goracle.v2"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/oracle"
)

const (
	PkgURL  = "github.com/go-goracle/goracle"
	PkgName = "gopkg.in/goracle.v2"
	DrvName = "goracle"
)

var ConnStr = fmt.Sprintf("%s/%s@%s:%s/%s", oracle.User, oracle.Pass, oracle.Host, oracle.Port, oracle.Schema)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(oracle.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
