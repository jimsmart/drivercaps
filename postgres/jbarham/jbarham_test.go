package jbarham_test

import (
	"fmt"
	"testing"

	_ "github.com/jbarham/gopgsqldriver"

	"github.com/jimsmart/drivercaps"
	"github.com/jimsmart/drivercaps/postgres"
)

const (
	PkgURL  = "github.com/jbarham/gopgsqldriver"
	PkgName = PkgURL
	DrvName = "postgres"
)

var ConnStr = fmt.Sprintf("user=%s host=%s port=%s sslmode=disable", postgres.User, postgres.Host, postgres.Port)

func TestDriver(t *testing.T) {
	err := drivercaps.ReportForDriver(postgres.NewTest(PkgURL, PkgName, DrvName, ConnStr))
	if err != nil {
		t.Errorf("error %v", err)
	}
}
