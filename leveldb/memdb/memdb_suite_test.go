package memdb

import (
	"testing"

	"bitbucket.org/cloudwallet/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
