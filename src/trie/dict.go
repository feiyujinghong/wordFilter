package trie

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

var tree *Trie

func readBlackDict(trie *Trie,path string){
	f,err := os.Open(path)

	if err != nil{
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}

	defer f.Close()

	fReader := bufio.NewReader(f)

	for{
		str,err := fReader.ReadString('\n')
		str = strings.Trim(str,"\n")
		trie.Add(str)

		if err == io.EOF{
			return
		}
	}
}

func InitTrie() *Trie{
	if tree == nil{
		tree = NewTrie()

		readBlackDict(tree,"dict/blackDict.txt")
	}

	return tree
}
