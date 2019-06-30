package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhongmin/gtra"
)

func Process(src, dest string) {
	content := readAnkiFile(src) // word, en_sentence, paraphrase
	total := len(content)
	log.Println("number of words: ", total)
	var newContent [][]string
	for i, line := range content {
		enSen := removeHTMLTags(line[1])
		zhSen, err := googleTranslate(enSen)
		if err != nil {
			log.Println("word [", line[0], "]: ", err)
			log.Println(fmt.Sprintf("[%d/%d] word [%s]: %s", i, total, line[0], err))
			continue
		}
		phoneticSymbol, paraphrase := splitParaphrase(line[2])
		newLine := []string{line[0], line[1], zhSen, phoneticSymbol, paraphrase}
		newContent = append(newContent, newLine)
		log.Println(fmt.Sprintf("[%d/%d] word [%s] process done", i, total, line[0]))
		if i > 0 && i % 10 != 0 {
			time.Sleep(time.Second)
		}
	}
	writeAnkiFile(newContent, dest)
}

func readAnkiFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "\t")
		content = append(content, s)
	}
	return content
}

func removeHTMLTags(sentence string) string {
	s := strings.Replace(sentence, "<b><u>", "", -1)
	s = strings.Replace(s, "</u></b>", "", -1)
	return s
}

func writeAnkiFile(content [][]string, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("create file: ", err)
	}
	for _, line := range content {
		_, err := fmt.Fprintln(file, strings.Join(line, "\t"))
		if err != nil {
			log.Fatal("write file: ", err)
		}
	}
}

func googleTranslate(sentence string) (string, error) {
	t := gtra.NewTranslater()
	e, s := t.Translate(sentence)
	return s, e
}

func splitParaphrase(paraphrase string) (string, string) {
	if strings.Contains(paraphrase, "，英") {
		list := strings.Split(paraphrase, "，")
		return strings.Join(list[0:2], "，"), strings.Join(list[2:], "，")
	} else if strings.Contains(paraphrase, "美") || strings.Contains(paraphrase, "英") {
		list := strings.SplitN(paraphrase, "，", 2)
		return list[0], list[1]
	} else {
		return "", paraphrase
	}
}
