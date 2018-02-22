package iterator_test

import (
	"testing"

	"bitbucket.com/cloudwallet/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
