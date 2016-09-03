package suffixtree_tests

import (
	"testing"
	"github.com/jojohannsen/suffixtree"
)

func TestStringTree(t *testing.T) {
	testString := "mississippi"
	runes := []rune(testString)
	s := suffixtree.NewStringDataSource(testString)
	incomingChannel := s.STKeys()
	for _, r := range runes {
		test := <-incomingChannel
		if test != suffixtree.STKey(r) {
			t.Error("channel did not provide the expected value")
		}
	}
}
