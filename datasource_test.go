package suffixtree_tests

import (
	"testing"

	"github.com/jojohannsen/suffixtree"
)

func TestStringDataSource(t *testing.T) {
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
	_, ok := <-incomingChannel
	if ok {
		t.Error("Unexpected value provided on incoming channel")
	}
}

func TestFileDataSource(t *testing.T) {
	fileContents := "mississippi"
	fileDataSource := suffixtree.NewFileDataSource("./test_data/mississippi.txt")
	runes := []rune(fileContents)
	incomingChannel := fileDataSource.STKeys()
	for _, r := range runes {
		test := <-incomingChannel
		if test != suffixtree.STKey(r) {
			t.Error("file channel did not provide the expected value")
		}
	}
	s := suffixtree.STKey('s')
	testS := fileDataSource.KeyAtOffset(2)
	if s != testS {
		t.Error("seek(2) failed, got %s, want %s", string(testS), string(s))
	}
	testS = fileDataSource.KeyAtOffset(3)
	if s != testS {
		t.Error("seek(2) failed, got %s, want %s", string(testS), string(s))
	}
	issi := fileDataSource.StringFrom(4, 7)
	if issi != "issi" {
		t.Errorf("Got %s, want %s", issi, "issi")
	}
}
