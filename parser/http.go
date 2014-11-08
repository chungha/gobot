package parser

import (
    "regexp"
	"io/ioutil"
    "net/http"
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
  return tokens
}

func Indexing(url string) (results []string, err error) {
  html, err := Download(url)
  if (err != nil) {
    return nil, err
  }
  words := Split(html)
  
  return words, nil
}