package indexer

import (
	"container/list"
	"fmt"
	"gobot/indexmap"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Download(url string) (html string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func Split(html string) (tokens *list.List) {
  reg, _ := regexp.Compile("[^A-Za-z0-9]+")

  sentences := regexp.MustCompile("<.*?>").Split(html, -1)

  tokens = list.New()
  for _, sentence := range sentences {
    tarray := Tokenize(reg.ReplaceAllString(sentence, " "))
    for _, t := range tarray {
      if len(t) > 0 {
        tokens.PushBack(t)
      }
    }
  }
  return tokens
}

func Tokenize(sentence string) (tokens []string) {
  tokens = strings.Split(sentence, " ")
  for i, token := range tokens {
    tokens[i] = strings.Trim(token, ".()!?,\t\r\n")
  }
  return tokens
}

func Indexing(url string) *list.List {
	var l = list.New()

	html, err := Download(url)
	if err != nil {
		return l
	}
	tokens := Split(html)
	for e := tokens.Front(); e != nil; e = e.Next() {
		s := fmt.Sprintf("%s %s %d", e.Value.(string), url, 0)
		l.PushBack(s)
	}
	return l
}

//func IndexingToMap(urls []string) *indexmap.IndexMap {
//	indexMap := indexmap.New()

//	for _, url := range urls {
//		indexingToMapSub(url, &indexMap)
//	}

//	return &indexMap
//}

func IndexingToMap(url string) *indexmap.IndexMap {
	index := indexmap.New()

	html, err := Download(url)
	if err != nil {
		return nil
	}

	tokens := Split(html)

	i := 1
	for e := tokens.Front(); e != nil; e = e.Next() {
		e_ := e.Value.(string)
		if len(e_) <= 0 {
			continue
		}
		//for i, e := range tokens {
		index.Add(e.Value.(string), url, i)
		i++
	}

	return index
}
