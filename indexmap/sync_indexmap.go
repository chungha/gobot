package indexmap

import (
  "container/list"
  "fmt"
  "strings"
  "sync"
  "strconv"
)

type SyncIndexMap struct {
  sync.RWMutex
  indexMap_ map[string]map[string]*list.List
}

func NewSync() SyncIndexMap {
  return SyncIndexMap { indexMap_: make(map[string]map[string]*list.List) }
}


func (i SyncIndexMap) AddIndexStringList(l *list.List) {
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      data := strings.Split(v, " ")
      pos, _ := strconv.Atoi(data[2])
      i.Add(data[0], data[1], pos)
    }
  }
}

func (i *SyncIndexMap) Add(word string, url string, pos int) *SyncIndexMap {
  i.Lock()
  urlMap, ok := i.indexMap_[word]
  if !ok {
    urlMap = make(map[string]*list.List)
  }
  posList, ok := urlMap[url]
  if !ok {
    posList = list.New()
  }
  posList.PushBack(pos)
  urlMap[url] = posList
  i.indexMap_[word] = urlMap
  i.Unlock()
  return i
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
      if p, ok := e.Value.(int); ok {
       result += strconv.Itoa(p) + " "
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
