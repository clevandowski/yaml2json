package main

import (
  "testing"
)

func TestYaml2JsonEmpty(t *testing.T) {
  yaml := ``

  want := ``

  got, _ := yaml2json(yaml)

  if want != got {
    t.Fatalf("yaml2json(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestYaml2JsonBasicMap(t *testing.T) {
  yaml := `coincoin: 5
`

  want := `{"coincoin":5}`

  got, _ := yaml2json(yaml)

  if want != got {
    t.Fatalf("yaml2json(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestYaml2JsonBasicMap2(t *testing.T) {
  yaml := `---
coincoin: 5
`

  want := `{"coincoin":5}`

  got, _ := yaml2json(yaml)

  if want != got {
    t.Fatalf("yaml2json(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestYaml2JsonBasicArray(t *testing.T) {
  yaml := `- coincoin
- cuicui
- coucou
`

  want := `["coincoin","cuicui","coucou"]`

  got, _ := yaml2json(yaml)

  if want != got {
    t.Fatalf("yaml2json(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestYaml2JsonBasicArray2(t *testing.T) {
  yaml := `---
- coincoin
- cuicui
- coucou
`

  want := `["coincoin","cuicui","coucou"]`

  got, _ := yaml2json(yaml)

  if want != got {
    t.Fatalf("yaml2json(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestYaml2JsonAdvanced(t *testing.T) {
  yaml := `a: Easy!
b:
  c: 2
  d: true
  my.super/key.test:
  - une\vilaine` + "`string`" + `
  - 4
  - another.key.too:
      x: 1
      y: 5
      z: 6.666666666666666666666666
  e:
  - true
  - false
  f: |
    coincoin
      cuicui\npouetpouet
    je\mets\des\antislash
  g: 3.1415
  h: 3.141592654
  i: "true"
`
  want := `{"a":"Easy!","b":{"c":2,"d":true,"e":[true,false],"f":"coincoin\n  cuicui\\npouetpouet\nje\\mets\\des\\antislash\n","g":3.1415,"h":3.141592654,"i":"true","my.super/key.test":["une\\vilaine` + "`string`" + `",4,{"another.key.too":{"x":1,"y":5,"z":6.666666666666667}}]}}`

  got, _ := yaml2json(yaml)

  if want != got {
    t.Fatalf("yaml2json(): want:\n%v\n, got:\n%v", want, got)
  }
}
