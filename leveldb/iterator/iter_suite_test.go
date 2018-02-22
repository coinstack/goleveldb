package iterator_test

import (
	"testing"

	"bitbucket.org/cloudwallet/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
