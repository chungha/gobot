package indexmap

import (
  "container/list"
  "fmt"
  "strings"
  "sync"
)

type SyncIndexMap struct {
  sync.RWMutex
  indexMap_ map[string]map[string]*list.List
}

func NewSync() SyncIndexMap {
  return SyncIndexMap { indexMap_: make(map[string]map[string]*list.List) }
}


func (i SyncIndexMap) AddIndexStringList(l *list.List) {
  i.Lock()
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
  i.Unlock()
}

func (i SyncIndexMap) Query(w string) map[string]*list.List {
  i.RLock();
  result := i.indexMap_[w]
  i.RUnlock();
  return result
}

func (i SyncIndexMap) QueryPrint(w string) string {
  i.RLock();
  urlMap, ok := i.indexMap_[w]
  if !ok {
    i.RUnlock();
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
  i.RUnlock();
  return result
}

func (i SyncIndexMap) Print() {
  for k := range i.indexMap_ {
    fmt.Printf(i.QueryPrint(k))
  }
}
