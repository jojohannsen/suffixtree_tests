package suffixtree_tests

import (
	"fmt"
	"github.com/jojohannsen/suffixtree"
	"testing"
)

//Users/jojo/genomes/apis_mellifera/stripped/ame_ref_Amel_4.5_chrLG16.fa

func TestBigFile(t *testing.T) {

	dataSource := suffixtree.NewFileDataSource("/Users/jojo/genomes/apis_mellifera/stripped/ame_ref_Amel_4.5_chrLG16.fa")
	ukkonen := suffixtree.NewUkkonen(dataSource)

	counter := 0
	for ukkonen.Extend() {
		counter++
		if (counter % 100000) == 0 {
			fmt.Print(".")
		}
	}
}
