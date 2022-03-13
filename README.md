# Go-Yandex-Translator
A simple Yandex Translate API v2 wrapper. Written in Go.
Usage:

```
package main

import (
  "fmt"
  "github.com/olekzonder/Go-Yandex-Translator/"
)

func main() {
  cli := translate.NewClient(API TOKEN HERE, Folder ID here)
  translation, err := cli.Translate([]string{"Hello World!"}, "en", "ru")
  fmt.Println(translation[0])
  //OUTPUT:
  //Привет, мир!
}
```
TO DO:
* Language Detection
