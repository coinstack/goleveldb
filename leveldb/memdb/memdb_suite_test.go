package memdb

import (
	"testing"

	"bitbucket.com/cloudwallet/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
