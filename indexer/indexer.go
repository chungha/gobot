package indexer

import (
  "fmt"
  "strings"
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

func Split(html string) (tokens *list.List) {
  sentences := regexp.MustCompile("<.*?>").Split(html, -1)

  tokens = list.New()
  for _, sentence := range sentences {
    tarray := Tokenize(sentence)
    for _, t := range tarray {
      tokens.PushBack(t)
    }
  }
  return tokens
}

func Tokenize(sentence string) (words []string) {
  return strings.Split(sentence, " ")
}

func Indexing(url string) *list.List {
  var l = list.New()

  html, err := Download(url)
  if (err != nil) {
    return l
  }
  tokens := Split(html)
  for e := tokens.Front(); e != nil; e = e.Next() {
    s := fmt.Sprintf("%s %s %d", e.Value.(string), url, 0)
    l.PushBack(s)
  }
  return l
}
