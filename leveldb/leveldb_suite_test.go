package leveldb

import (
	"testing"

	"bitbucket.org/cloudwallet/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
