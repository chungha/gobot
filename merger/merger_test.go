package merger

import (
  "testing"
  "gobot/indexmap"
)

func TestMerger(t *testing.T) {
  var dst = indexmap.New();
  var src = indexmap.New();

  dst.Add("word1", "url1", 0);
  dst.Add("word2", "url1", 0);
  src.Add("word1", "url2", 0);
  src.Add("word2", "url1", 1);

  Merge(dst, src);

  posList, ok := dst.QueryWordAndUrl("word1", "url1")
  if !ok || posList.Len() != 1 {
    t.Error("Nope word1 and url1 or posList.Len() != 1")
    return
  }

  posList, ok = dst.QueryWordAndUrl("word1", "url2")
  if !ok || posList.Len() != 1 {
    t.Error("Nope word1 and url2 or posList.Len() != 1")
    return
  }

  posList, ok = dst.QueryWordAndUrl("word2", "url1")
  if !ok || posList.Len() != 2 {
    t.Error("Nope word2 and url1 or posList.Len() != 2")
    return
  }
}
