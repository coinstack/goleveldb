package leveldb

import (
	"testing"

	"bitbucket.com/cloudwallet/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
