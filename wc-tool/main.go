package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
	"unicode"
)

type fileStat struct {
	bytes int
	lines int
	words int
	chars int 
}

func getInfo (file *os.File) fileStat {
	var bytes, lines, words, chars int
	var inword bool
	reader := bufio.NewReader(file)

	for {
		r, sz, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if inword {
					words++
				}
				break
			} else {
				log.Fatal(err)
			}
		}
		if r == '\n' {
			lines++
		}
		bytes += int(sz)
		chars++
		if unicode.IsSpace(r) {
			if inword {
				words++
			}
			inword = false
		} else {
			inword = true
		}

	}
	return fileStat{bytes: bytes, lines: lines, words: words, chars: chars}

}

func main() {
	
	var getBytes, getLines, getWords, getChars bool 
	flag.BoolVar(&getBytes, "c", false, "boolean flag for bytes")
	flag.BoolVar(&getLines, "l", false, "boolean flag for lines")
	flag.BoolVar(&getWords, "w", false, "boolean flag for words")
	flag.BoolVar(&getChars, "m", false, "boolean flag for chars")

	flag.Parse()

	if !getBytes && !getLines && !getWords && !getChars {
		getBytes, getLines, getWords, getChars = true, true, true, true
	}

	filename := flag.CommandLine.Arg(0)

	stat, err := os.Stdin.Stat()

	if err != nil {
		log.Fatal(err)
	}

	var file *os.File

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		} 
		defer file.Close()
	}

	result := getInfo(file)

	var res []string
	if getLines {
		res = append(res, strconv.Itoa(result.lines))
	}
	if getWords {
		res = append(res, strconv.Itoa(result.words))
	}
	if getBytes {
		res = append(res, strconv.Itoa(result.bytes))
	}
	if getChars {
		res = append(res, strconv.Itoa(result.chars))
	}

	res = append(res, filename)
	fmt.Println(strings.Join(res, " "))
}