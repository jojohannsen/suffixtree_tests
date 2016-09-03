package suffixtree_tests

import (
	"fmt"
	"testing"

	"github.com/jojohannsen/suffixtree"
)

func TestLocationString(t *testing.T) {
	location := suffixtree.NewLocation(suffixtree.NewRootNode())
	want := "Location on Node 1"
	s := fmt.Sprintf("%s", location)
	if s != want {
		t.Errorf("NewLocation() got %s, want %s", s, want)
	}
	location.Edge = suffixtree.NewEdge(10, 20)
	s = fmt.Sprintf("%s", location)
	want = "Location on Node 1"
	if s != want {
		t.Errorf("After setting edge, got %s, want %s", s, want)
	}
	location.OffsetFromTop = 9999
	location.OnNode = false
	want = "Location on Edge, OffsetFromTop=9999, node is 1"
	s = fmt.Sprintf("%s", location)
	if s != want {
		t.Errorf("After setting offsetFromTop, got %s, want %s", s, want)
	}
	location.Base = suffixtree.NewRootNode()
	s = fmt.Sprintf("%s", location)
	want = "Location on Edge, OffsetFromTop=9999, node is 2"
	if s != want {
		t.Errorf("After setting node, got %s, want %s", s, want)
	}
	location.OnNode = false
	s = fmt.Sprintf("%s", location)
	want = "Location on Edge, OffsetFromTop=9999, node is 2"
	if s != want {
		t.Errorf("After setting onInternalNode, got %s, want %s", s, want)
	}
}
