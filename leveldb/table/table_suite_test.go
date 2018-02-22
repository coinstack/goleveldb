package table

import (
	"testing"

	"bitbucket.com/cloudwallet/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
