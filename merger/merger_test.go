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

}
