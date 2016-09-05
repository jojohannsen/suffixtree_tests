package suffixtree_tests

import (
	"testing"
	"github.com/jojohannsen/suffixtree"
	"fmt"
	"sort"
)

type int32arr []int32

func (a int32arr) Len() int           { return len(a) }
func (a int32arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int32arr) Less(i, j int) bool { return a[i] < a[j] }

func TestDepthVisitor(t *testing.T) {
	dvChan := make(chan []int32)
	dv := suffixtree.NewDepthVisitor(1, dvChan)
	dfs := suffixtree.NewDFS(dv)
	dataSource := suffixtree.NewStringDataSource("mississippi")
	ukkonen := suffixtree.NewUkkonen(dataSource)
	ukkonen.DrainDataSource()
	ukkonen.Finish()
	go func(dfs *suffixtree.DFS, node suffixtree.Node, dvChan chan []int32) {
		dfs.Traverse(node)
		close(dvChan)
	} (dfs, ukkonen.Tree().Root(), dvChan)
	resultsAsString := make(map[string]bool)
	expectedResults := []string{"1 4 7 10", "2 3 5 6", "8 9", "11", "0"}
	for _,s := range expectedResults {
		resultsAsString[s] = false
	}
	for suffixes := range dvChan {
		result := int32arr(suffixes)
		sort.Sort(result)
		s := fmt.Sprintf("%d", result[0])
		if len(result) > 1 {
			for _,n := range result[1:] {
				s = fmt.Sprintf("%s %d", s, n)
			}
		}
		val, ok := resultsAsString[s]
		if !ok {
			t.Errorf("Got unexpected result %s", s)
		}
		if val {
			t.Errorf("Got unexpected true value for %s", s)
		}
		resultsAsString[s] = true
	}
	for _,s := range expectedResults {
		if !resultsAsString[s] {
			t.Errorf("Did NOT find %s", s)
		}
	}
}
