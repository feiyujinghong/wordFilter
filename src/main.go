package main

import (
	"net/http"
	"net/url"
	"trie"
)

type route struct{

}

func (r *route) ServeHTTP(w http.ResponseWriter, re *http.Request){
	text := ""
	paramName := "q"
	params,_ := url.ParseQuery(re.URL.RawQuery)

	if q,ret := params[paramName];ret{
		text = q[0]
	}

	res := trie.InitTrie().ReplaceWord(text)

	w.Write([]byte(res))

}

func main(){
	serverAdd := ":9090"
	trie.InitTrie()

	http.ListenAndServe(serverAdd,&route{})
}
