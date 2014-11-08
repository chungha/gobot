package reducer

import (
  "strings"
  "container/list"
)

var indexMap = make(map[string]map[string]*list.List)

func Reduce(l *list.List) {
  for e := l.Front(); e != nil; e = e.Next() {
    if v, ok := e.Value.(string); ok {
      data := strings.Split(v, " ")
      urlMap, ok := indexMap[data[0]]
      if !ok {
        urlMap = make(map[string]*list.List)
      }
      posList, ok := urlMap[data[1]]
      if !ok {
        posList = list.New()
      }
      posList.PushBack(data[2])
      urlMap[data[1]] = posList
      indexMap[data[0]] = urlMap
    }
  }
}

func Query(w string) string {
  urlMap, ok := indexMap[w]
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
