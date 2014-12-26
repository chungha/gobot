package merger

import (
  "gobot/indexmap"
)

func Merge(dst *indexmap.IndexMap, src *indexmap.IndexMap) {
  for word, urls := range src.IndexMapData {
    for url, pos := range urls {
      for e := pos.Front(); e != nil; e = e.Next() {
        dst.Add(word, url, e.Value.(int))
      }
    }
    
  }
}
