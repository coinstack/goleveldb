package table

import (
	"testing"

	"bitbucket.org/cloudwallet/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
