package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	splitter()
}

func splitter() {
	const (
		F1  = "merge1.txt"
		F2  = "merge2.txt"
		F3  = "merge3.txt"
		F4  = "merge4.txt"
		OUT = "merged-verbform-cards.txt"
	)

	files := []string{F1, F2, F3, F4}
	ss := make([][]string, len(files))

	for i := range files {
		a, err := os.Open(files[i])
		if err != nil {
			panic(err)
		}
		defer a.Close()
		sc := bufio.NewScanner(a)
		for sc.Scan() {
			ss[i] = append(ss[i], sc.Text())
		}
	}

	f, err := os.Create(OUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buffer := bufio.NewWriter(f)
	for i := 0; i < len(ss[0]); i++ {
		for j := 0; j < len(ss); j++ {
			if i < len(ss[j]) {
				x := ss[j][i] + "\n"
				_, e := buffer.WriteString(x)
				if e != nil {
					log.Fatal(e)
				}
			}
		}
	}
}
