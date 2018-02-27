package mattn_test

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-oci8"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/oracle"
)

const (
	PkgURL  = "github.com/mattn/go-oci8"
	PkgName = PkgURL
	DrvName = "oci8"
)

var ConnStr = fmt.Sprintf("%s/%s@%s:%s/%s", oracle.User, oracle.Pass, oracle.Host, oracle.Port, oracle.Schema)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(oracle.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
