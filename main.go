package main

import (
	"archive/tar"
	"bufio"
	"compress/bzip2"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		panic("Example:\n" + os.Args[0] + " eng cmn > eng-cmn.tsv")
	}

	lang1 := os.Args[1]
	lang2 := os.Args[2]

	langs := make(map[string]string)
	sents := make(map[string]string)

	f, err := os.Open("sentences.tar.bz2")
	if err != nil {
		panic("Cannot open sentences.tar.bz2; download from http://downloads.tatoeba.org/exports/sentences.tar.bz2")
	}
	defer f.Close()
	bzipReader := bzip2.NewReader(f)
	tarReader := tar.NewReader(bzipReader)
	tarReader.Next()

	scanner := bufio.NewScanner(tarReader)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "\t")
		lang := split[1]

		if lang == lang1 || lang == lang2 {
			id := split[0]
			sent := split[2]

			langs[id] = lang
			sents[id] = sent
		}

	}

	f1, err := os.Open("links.tar.bz2")
	if err != nil {
		panic("Cannot open links.tar.bz2; download from http://downloads.tatoeba.org/exports/links.tar.bz2")
	}
	defer f1.Close()
	bzipReader1 := bzip2.NewReader(f1)
	tarReader1 := tar.NewReader(bzipReader1)
	tarReader1.Next()

	scanner = bufio.NewScanner(tarReader1)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "\t")
		id1 := split[0]
		id2 := split[1]

		if langs[id1] != "" && langs[id2] != "" && langs[id1] != langs[id2] {
			if langs[id1] == lang1 {
				fmt.Printf("%s\t%s\n", sents[id1], sents[id2])
			} else {
				fmt.Printf("%s\t%s\n", sents[id2], sents[id1])
			}
		}

	}
}
