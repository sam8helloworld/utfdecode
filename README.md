# utfdecode

[![go report card](https://goreportcard.com/badge/github.com/sam8helloworld/utfdecode)](https://goreportcard.com/report/github.com/sam8helloworld/utfdecode) 
[![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](https://gocover.io/github.com/sam8helloworld/utfdecode)
[![godocs](https://godoc.org/github.com/sam8helloworld/utfdecode?status.svg)](https://godoc.org/github.com/sam8helloworld/utfdecode) 

This is escaped unicode string decoder.

Use like this([Go Playground](https://go.dev/play/p/2qrSdCCpPd7)):

```go
s, _ := utfdecode.Decode(`\uD83D\uDE04あ\uD83D\uDE07い\uD83D\uDC7Aう`)
fmt.Println(s) // 😄あ😇い👺う
```

## Licence

MIT
