package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type grep_step struct {
	Title                  string `json:"title"`
	Index                  int    `json:"index"`
	Select_Ref             string `json:"select_ref"`
	GrepPattern            string `json:"grep_select_pattern"`
	RegexpPattern          string `json:"regexp_pattern"`
	Output                 string `step:"output"`
	InputFile              string `json:"input_file"`
	OutputFilenameTemplate string `json:"output_filename_template"`
}

type grep_steps struct {
	Steps []grep_step `json:"steps"`
}

func NewConfig(filename string) *grep_steps {
	var steps grep_steps
	var bytes []byte
	mfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(mfile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		bytes = append(bytes, scanner.Bytes()...)
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	mfile.Close()
	err = json.Unmarshal(bytes, &steps)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	return &steps
}
