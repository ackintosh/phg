package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		src, err := ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
		}

		phpDelimiter := "<?php"
		pos := 0

		for {
			if strings.HasPrefix(string(src)[pos:], phpDelimiter) {
				pos = pos + len(phpDelimiter)
				l := new(Lexer)

				l.Init(strings.NewReader(string(src)[pos:]))
				yyParse(l)
				_, err = Evaluate(l.result)
				if err != nil {
					panic(err)
				}
			}

			break
		}
	}
}
