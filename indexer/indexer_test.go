package indexer

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_Split(t *testing.T) {
	l := Split("<xxx>I am a boy.</xxx> \r\n haha\r\n\t-\r\n-")
	for e := l.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(string); ok {
			fmt.Printf(">%s\n", v)
		}
	}
}

func Test_Tokenize(t *testing.T) {
	words := Tokenize("I am a Boy!. \r\n haha\r\n\t-\r\n-")
	fmt.Printf("%q\n", words)

	if words[3] != "Boy" {
		t.Error("The third result. Expected : Boy, but " + words[3])
	}
}

func Test_Regexp(t *testing.T) {
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString("<xxx>I am a boy.</xxx> \r\n haha\r\n\t-\r\n-", -1)

	fmt.Printf("%q\n", words)
}

//func Test_Indexing(t *testing.T) {
//	url := "http://en.wikipedia.org/wiki/Peace"
//	l := Indexing(url)

//	i := 1
//	for w := l.Front(); w != nil; w = w.Next() {
//		w_ := w.Value.(string)
//		fmt.Println(w_, i)
//		i++
//	}
//}

func Test_IndexingToMap(t *testing.T) {
	url := "http://en.wikipedia.org/wiki/Peace"
	indexMap := IndexingToMap(url)

	html, err := Download(url)
	if err != nil {
		t.Error("Download fail !!!")
	}

	tokens := Split(html)
	i := 1
	for e := tokens.Front(); e != nil; e = e.Next() {
		w_ := e.Value.(string)
		if len(w_) <= 0 {
			continue
		}

		fmt.Printf("%s - %d\n", w_, len(w_))

		//i := 1
		//for w := l.Front(); w != nil; w = w.Next() {
		//	w_ := w.Value.(string)

		list, isExist := indexMap.QueryWordAndUrl(w_, url)

		if isExist == false {
			t.Error("Does not exist %s, %s", w_, url)
		}

		check := false
		for u := list.Front(); u != nil; u = u.Next() {
			//for _, u := range for w := l.Front(); w != nil; w = w.Next() { {
			if u.Value.(int) == i {
				break
			}
		}

		if check == false {
			t.Error("position is not exist %s, %s, %d", w_, url, i)
			break
		}

		i++
	}
}
