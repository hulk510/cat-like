package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	n := flag.Bool("n", false, "投資番号付与")
	flag.Parse()
	files := flag.Args()
	currendDir, err := os.Getwd()
	if err != nil {
		log.Fatal("ファイルのパスの読み込みが失敗してるよ")
	}
	for _, file := range files {
		f, err := os.Open(filepath.Join(currendDir, file))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		bufio.NewReader(f)
		scanner := bufio.NewScanner(f)

		for i := 1; scanner.Scan(); i++ {
			if *n {
				fmt.Printf("%d: ", i)
			}
			fmt.Println(scanner.Text())
		}

		if err = scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
