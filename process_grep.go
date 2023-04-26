package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// wants revision to be able to use more than 1 capture group from a regular expression...
type StageResult string

type StageResults []StageResult

type MultipleStageResults map[int][]StageResults

func OpenFile(step grep_step, mPriorStageResult *StageResult) (*os.File, error) {
	var filename string
	patt := regexp.MustCompile("(.*)(<select_ref>)(.*)")
	results := patt.FindAllStringSubmatch(step.OutputFilenameTemplate, -1)
	if len(results) > 0 {
		result := results[0]
		filename = string(result[1]) + string(*mPriorStageResult) + string(result[3])
	} else {
		filename = step.GrepPattern
	}
	mfile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Unable to create file %s\n", filename)
		return nil, err
	}
	return mfile, nil
}

func OutputToFile(step grep_step, cmdoutput []byte, mPriorStageResult *StageResult) {
	fmt.Printf("Not implemented yet for grep\n")
	var endindex int
	var byte byte
	fmt.Printf("Output:\n%s\n", cmdoutput)
	startindex := 0
	mfile, err := OpenFile(step, mPriorStageResult)
	defer mfile.Close()
	if err != nil {
		fmt.Printf("Problem creating file\n")
	}
	for endindex, byte = range cmdoutput {
		if byte == '\n' {
			nBytesWritten, err := mfile.Write(cmdoutput[startindex : endindex+1])
			if err != nil {
				fmt.Printf("Problem writing file\n")
			}
			if nBytesWritten == 0 {
				fmt.Printf("Problem writing file\n")
			}
			startindex = endindex + 1
		}
		endindex++
	}
}

func GetResults(mGrepPattern string, step grep_step, mPriorStageResult *StageResult) (stageresults StageResults) {
	cmdoutput, err := exec.Command("grep", mGrepPattern, step.InputFile).Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}
	if step.Output == "Memory" {
		var endindex int
		var byte byte
		fmt.Printf("Output:\n%s\n", cmdoutput)
		startindex := 0
		for endindex, byte = range cmdoutput {
			if byte == '\n' {
				stageresults = append(stageresults, StageResult(cmdoutput[startindex:endindex]))
				startindex = endindex + 1
			}
			endindex++
		}
	} else if step.Output == "File" {
		OutputToFile(step, cmdoutput, mPriorStageResult)
	}
	return stageresults
}

func do_grep(step grep_step, mPriorStageResult *StageResult) *StageResults {
	var mstageresults StageResults
	var mGrepPattern string
	if len(step.Select_Ref) > 0 {
		// doesn't actually use the reference - which it should...
		// and it should loop...
		fmt.Printf("grep \"%s\" %s\n", string(*mPriorStageResult), step.InputFile)
		mGrepPattern = string(*mPriorStageResult)
		mstageresults = GetResults(mGrepPattern, step, mPriorStageResult)
	} else {
		fmt.Printf("grep \"%s\" %s\n", step.GrepPattern, step.InputFile)
		mGrepPattern = step.GrepPattern
		mstageresults = GetResults(mGrepPattern, step, mPriorStageResult)
	}
	return &mstageresults
}

func do_regexp(step grep_step, mPriorStageResult *StageResult) *StageResults {
	var stageresults StageResults
	//var mregexp string
	//mregexp = step.RegexpPattern
	patt, err := regexp.Compile(step.RegexpPattern)
	if err != nil {
		fmt.Printf("Problem compiling regular expression :%s\n", step.RegexpPattern)
	}
	results := patt.FindAllStringSubmatch(string(*mPriorStageResult), -1)
	if len(results) > 0 {
		result := results[0]
		if step.Output == "Memory" {
			stageresults = append(stageresults, StageResult(result[1]))
		}
	}
	return &stageresults
}

func ProcessGrep(mgrep_steps grep_steps) {
	var mStageResults StageResults
	var mBlankStageResult StageResult
	mStageResults = StageResults{}
	var mMultipleStageResults MultipleStageResults
	mMultipleStageResults = make(MultipleStageResults, 5)
	fmt.Printf("Welcome to Process Grep\n")
	for index, step := range mgrep_steps.Steps {
		if len(step.GrepPattern) > 0 {
			// do grep...
			// loop at this level if there is more than 1 prior stage result..
			if index == 0 {
				mNewStageResults := do_grep(step, &mBlankStageResult)
				mStageResults = *mNewStageResults
				mMultipleStageResults[step.Index] = append(mMultipleStageResults[step.Index], *mNewStageResults)
			} else {
				for _, mStageResults := range mMultipleStageResults[step.Index-1] {
					for _, mStageResult := range mStageResults {
						mNewStageResults := do_grep(step, &mStageResult)
						mStageResults = *mNewStageResults
						mMultipleStageResults[step.Index] = append(mMultipleStageResults[step.Index], *mNewStageResults)
					}
				}
			}
		} else if len(step.Select_Ref) > 0 {
			// use the output of the named step
			for _, mStageResult := range mStageResults {
				mNewStageResults := do_regexp(step, &mStageResult)
				mStageResults = *mNewStageResults
				mMultipleStageResults[step.Index] = append(mMultipleStageResults[step.Index], *mNewStageResults)
			}
		}

	}
	// fmt.Printf("Found %d sets of results", len(*mMultipleStageResults))
}
