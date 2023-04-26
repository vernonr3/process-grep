package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DoGrep(t *testing.T) {
	var stageresult StageResult
	//	var stageresult StageResult
	var step grep_step
	step = grep_step{
		Title:       "Step1",
		GrepPattern: "Application error",
		Output:      "Memory",
		InputFile:   "clientexp7#1.txt",
	}
	do_grep(step, &stageresult)
}

func Test_DoRegexp(t *testing.T) {
	var stageresult StageResult
	var mstep grep_step
	mstep = grep_step{
		Title:         "Step2",
		Select_Ref:    "1",
		RegexpPattern: ".* HandshakeDestConnID ([0-9abcdef]{8}).*",
		Output:        "Memory",
	}
	stageresult = "2023/04/23 16:48:32 client Closing connection HandshakeDestConnID 2459e732 DestConnID 2bf9818d5250cbb8cf SrcConnID <nil> with error: Application error 0x0 (local)"
	results := do_regexp(mstep, &stageresult)
	assert.Equal(t, "2459e732", string((*results)[0]))

}

func Test_DoGrep2File(t *testing.T) {
	var stageresult StageResult
	var step grep_step
	step = grep_step{
		Title:                  "Step3",
		Select_Ref:             "2",
		GrepPattern:            "<select_ref>",
		InputFile:              "clientexp7#1.txt",
		Output:                 "File",
		OutputFilenameTemplate: "client_<select_ref>.txt",
	}
	stageresult = "3cbb9729"
	do_grep(step, &stageresult)
}

func Test_OpenFile(t *testing.T) {
	var mPriorStageResult StageResult
	var mstep grep_step
	mstep = grep_step{
		Title:                  "Step2",
		Select_Ref:             "1",
		RegexpPattern:          ".* HandshakeDestConnID ([0-9abcdef]{8}).*",
		Output:                 "Memory",
		OutputFilenameTemplate: "client_<select_ref>_suffix.txt",
	}
	mPriorStageResult = "1234"
	mfile, mfilename, err := OpenFile(mstep, &mPriorStageResult)
	assert.Equal(t, nil, err)
	assert.Equal(t, "client_1234_suffix.txt", mfilename)
	assert.NotNil(t, mfile)
	mfile.Close()
}

//GetResults(mGrepPattern string, step grep_step, mPriorStageResult *StageResult)
