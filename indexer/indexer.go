package parser

import (
  "fmt"
  "regexp"
  "io/ioutil"
  "net/http"
  "container/list"
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

func Split(html string) (words []string) {
  tokens := regexp.MustCompile("<.*?>").Split(html, -1)
  /*for _, token := range tokens {
    if (len(token) > 0) {
	  words := SplitWithWhiteSpace(token)
	}
  }*/
  return tokens
}

func Indexing(url string) *list.List {
  var l = list.New()

  html, err := Download(url)
  if (err != nil) {
    return l
  }
  words := Split(html)
  for index, word := range(words) {
    s := fmt.Sprintf("%s %s %d", word, url, index)
    l.PushBack(s)
  }
  return l
}
