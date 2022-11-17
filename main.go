package main

import (
	"fmt"
	"os"

	"github.com/grailbio/base/tsv"
)

type row struct {
	Tconst         string `tsv:"tconst"`
	TitleType      string `tsv:"titleType"`
	PrimaryTitle   string `tsv:"primaryTitle"`
	OriginalTitle  string `tsv:"originalTitle"`
	IsAdult        byte   `tsv:"isAdult"`
	StartYear      uint16 `tsv:"startYear"`
	EndYear        string `tsv:"endYear"`
	RuntimeMinutes uint16 `tsv:"runtimeMinutes"`
	Genres         string `tsv:"genres"`
}

func ReadFilePlain() {
	file, err := os.Open("/static/data.tsv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := tsv.NewReader(file)
	r.HasHeaderRow = true
	r.UseHeaderNames = true
	for i := 0; i < 1000; i++ {
		var v row
		err = r.Read(&v)
		if err == nil {
			fmt.Printf("%+v\n", v)
		} else {
			fmt.Println(err)
		}
	}
}

func ReadFileGoRoutines() {
	file, err := os.Open("/static/data.tsv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := tsv.NewReader(file)
	r.HasHeaderRow = true
	r.UseHeaderNames = true
	for i := 0; i < 1000; i++ {
		go func() {
			var v row
			err = r.Read(&v)
			if err == nil {
				fmt.Printf("%+v\n", v)
			} else {
				fmt.Println(err)
			}
		}()
	}
}

func main() {
	ReadFilePlain()
}
