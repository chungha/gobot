package indexmap

import (
  "container/list"
  "fmt"
  "strings"
)

type IndexMap struct {
  indexMap_ map[string]map[string]*list.List
}

func New() *IndexMap {
  return &IndexMap { make(map[string]map[string]*list.List) }
}


func (i *IndexMap) AddIndexStringList(l *list.List) {
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      data := strings.Split(v, " ")
      urlMap, ok := i.indexMap_[data[0]]
      if !ok {
        urlMap = make(map[string]*list.List)
      }
      posList, ok := urlMap[data[1]]
      if !ok {
        posList = list.New()
      }
      posList.PushBack(data[2])
      urlMap[data[1]] = posList
      i.indexMap_[data[0]] = urlMap
    }
  }
}

func (i *IndexMap) Query(w string) map[string]*list.List {
  return i.indexMap_[w]
}

func (i *IndexMap) QueryPrint(w string) string {
  urlMap, ok := i.indexMap_[w]
  if !ok {
    return "none"
  }
  var result = w + "\n"
  for k, v := range urlMap {
    result += " - URL : " + k + ", POS [ "
    for e := v.Front(); e != nil; e = e.Next() {
      if p, ok := e.Value.(string); ok {
       result += p + " "
      }
    }
    result += "]\n"
  }
  return result
}

func (i *IndexMap) Print() {
  for k := range i.indexMap_ {
    fmt.Printf(i.QueryPrint(k))
  }
}
