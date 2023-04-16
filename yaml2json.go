package main

import (
	"encoding/json"
  "fmt"
  "github.com/clevandowski/yamltool"
  "gopkg.in/yaml.v3"
  "io"
  "log"
  "os"
  "strings"
)

func yaml2json(document string) (string, error) {
  document = strings.TrimSpace(document)
  if len(document) < 2 {
    return "", nil
  }
  var m any
  err := yaml.Unmarshal([]byte(document), &m)
  if err != nil {
    return "", err
  }

  output, err := json.Marshal(&m)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("%v", string(output)), nil
}

func main() {
  stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	rawDocuments := string(stdin)

  documents := yamltool.SplitDocuments(rawDocuments)
  for _, document := range documents {
    output, err := yaml2json(document)
    if err != nil {
      log.Fatalf("error: %v", err)
    }
    fmt.Printf("%v\n", string(output))
  }
}
