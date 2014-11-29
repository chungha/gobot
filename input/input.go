
package input

import (
  "fmt"
  "container/list"
  "bufio"
  "os"
)


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}


func Input(filePath string) *list.List {

  inputs := list.New();

  lines, err := readLines(filePath)
  if err != nil {
    fmt.Println("can not read file")
    return inputs
  }

  for _, line := range lines {
    // fmt.Println(i, line)
    inputs.PushBack(line)
  }

/* static
  inputs.PushBack("http://en.wikipedia.org/wiki/Go_(programming_language)");
  inputs.PushBack("http://en.wikipedia.org/wiki/Peace");
  inputs.PushBack("http://en.wikipedia.org/wiki/Apple_Inc.");
*/

/* for testing
  for e := inputs.Front(); e != nil; e = e.Next() {
       fmt.Println(e.Value)
  }
*/

  return inputs
}

