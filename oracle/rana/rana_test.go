package rana_test

import (
	"fmt"
	"testing"

	_ "gopkg.in/rana/ora.v4"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/oracle"
)

const (
	PkgURL  = "github.com/rana/ora"
	PkgName = "gopkg.in/rana/ora.v4"
	DrvName = "ora"
)

var ConnStr = fmt.Sprintf("%s/%s@%s:%s/%s", oracle.User, oracle.Pass, oracle.Host, oracle.Port, oracle.Schema)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(oracle.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
