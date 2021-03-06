package indexmap

import (
  "container/list"
  "fmt"
  "strings"
  "strconv"
)

type IndexMap struct {
  IndexMapData map[string]map[string]*list.List
}

func New() *IndexMap {
  return &IndexMap { make(map[string]map[string]*list.List) }
}

func (i *IndexMap) AddIndexStringList(l *list.List) {
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      data := strings.Split(v, " ")
      pos, _ := strconv.Atoi(data[2])
      i.Add(data[0], data[1], pos)
    }
  }
}

func (i *IndexMap) Add(word string, url string, pos int) *IndexMap {
  urlMap, ok := i.IndexMapData[word]
  if !ok {
    urlMap = make(map[string]*list.List)
  }
  posList, ok := urlMap[url]
  if !ok {
    posList = list.New()
  }
  posList.PushBack(pos)
  urlMap[url] = posList
  i.IndexMapData[word] = urlMap
  return i
}

func (i *IndexMap) QueryWord(w string) map[string]*list.List {
  return i.IndexMapData[w]
}

func (i *IndexMap) QueryWordAndUrl(word string, url string) (*list.List, bool) {
  urlMap, ok := i.IndexMapData[word]
  if !ok {
    return nil, false
  }
  posList, ok := urlMap[url]
  if !ok {
    return nil, false
  }
  return posList, true
}

func (i *IndexMap) QueryPrint(w string) string {
  urlMap, ok := i.IndexMapData[w]
  if !ok {
    return "none"
  }
  var result = w + "\n"
  for k, v := range urlMap {
    result += " - URL : " + k + ", POS [ "
    for e := v.Front(); e != nil; e = e.Next() {
      if p, ok := e.Value.(int); ok {
       result += strconv.Itoa(p) + " "
      }
    }
    result += "]\n"
  }
  return result
}

func (i *IndexMap) Print() {
  for k := range i.IndexMapData {
    fmt.Printf(i.QueryPrint(k))
  }
}
