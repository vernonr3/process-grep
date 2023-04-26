package main

import (
	"testing"
)

func Test_DoGrep(t *testing.T) {
	var stageresults StageResults
	stageresults = StageResults{}
	//	var stageresult StageResult
	var step grep_step
	step = grep_step{
		Title:       "Step1",
		GrepPattern: "Application error",
		Output:      "Memory",
		InputFile:   "../../libp2p_issues/libp2p_tests/spawnrate2/clientexp7#1.txt",
	}
	do_grep(step, &stageresults)
}

func Test_DoRegexp(t *testing.T) {
	var stageresults StageResults
	var stageresult StageResult
	var mstep grep_step
	mstep = grep_step{
		Title:         "Step2",
		Select_Ref:    "Step1",
		RegexpPattern: ".* HandshakeDestConnID ([0-9abcdef]{8}).*",
		Output:        "Memory",
		//OutputFilenameTemplate: "<select_ref>.txt",
	}
	stageresult = "2023/04/23 16:48:32 client Closing connection HandshakeDestConnID 2459e732 DestConnID 2bf9818d5250cbb8cf SrcConnID <nil> with error: Application error 0x0 (local)"
	stageresults = append(stageresults, stageresult)
	do_regexp(mstep, &stageresults)

}

func Test_DoGrep2File(t *testing.T) {
	var stageresults StageResults
	var stageresult StageResult
	var step grep_step
	step = grep_step{
		Title:                  "Step3",
		Select_Ref:             "Step2",
		GrepPattern:            "<select_ref>",
		InputFile:              "../../libp2p_issues/libp2p_tests/spawnrate2/clientexp7#1.txt",
		Output:                 "File",
		OutputFilenameTemplate: "client_<select_ref>.txt",
	}
	stageresult = "3cbb9729"
	stageresults = append(stageresults, stageresult)
	do_grep(step, &stageresults)
}
