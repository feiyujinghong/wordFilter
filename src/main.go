package main

import (
	"net/http"
	"net/url"
	"trie"
	"github.com/huichen/sego"
	"strings"
	"encoding/json"
	"fmt"
	"regexp"
)

type route struct{

}

var segmenter sego.Segmenter

func (r *route) ServeHTTP(w http.ResponseWriter, re *http.Request){
	text := ""
	paramName := "q"
	params,_ := url.ParseQuery(re.URL.RawQuery)

	if q,ret := params[paramName];ret{
		text = q[0]
	}

	wordsArr := KeyWords(text)

	keyWords := FindKeyWord(wordsArr)

	m := make(map[string]interface{})

	m["keywords"] = keyWords
	m["wordsArr"] = wordsArr

	ret := OutByJson(m)

	w.Write([]byte(ret))
}

//收集敏感词
func FindKeyWord(wordsArr []string) []string {
	var res []string

	for _,word := range wordsArr{

		reg  := regexp.MustCompile(`[\/v|\/nr]`)
		word =  reg.ReplaceAllString(word, "")
		if(true == trie.InitTrie().Search(word)){
			res = append(res,word)
		}
	}

	return res
}

/**
 * 获取分词表
 */
func KeyWords(word string) []string {
	segments := segmenter.Segment([]byte(word))

	words := (sego.SegmentsToString(segments,false))

	wordsArr := strings.Split(words," ");

	return wordsArr
}



func OutByJson(m map[string]interface{}) string {


	ret, err := json.Marshal(m)

	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return ""
	}

	return string(ret)
}

func main(){

	segmenter.LoadDictionary("dict/dictionary.txt")

	serverAdd := ":9090"
	trie.InitTrie()

	http.ListenAndServe(serverAdd,&route{})
}
