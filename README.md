# Gunigram
A Go package to create Unigram(n=1 ngram) for text mining preprocessing, or any other purpose :D

Gunigram is a go package to simplifies creating a csv file of unigram representation of a given set of sentences

## Using
### go get
You need `go` installed and `GOPATH` in your `PATH` , then run :
```shell
$ go get github.com/kahalKevin/Gunigram
```

## Usage (Library)
```go
package main

import (
    "flag"
    "runtime"
    gunigram "github.com/kahalKevin/Gunigram"
)

var(
    file = flag.String("file", "bahan.txt", "The file which consist line of sentences")
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
    flag.Parse()
    gunigram.Unigram(*file)
}
```

## Output
Output will be saved ad ngram.csv
```console
aku,kamu,kita,dia,mereka,kata
2,1,0,0,0,0
0,1,1,1,0,0
1,0,0,1,1,0
0,0,1,0,1,1
1,0,1,1,0,0

```
