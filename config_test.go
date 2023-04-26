package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig(t *testing.T) {
	grep_steps := NewConfig("clientsteps_example.json")
	grep_step0 := grep_steps.Steps[0]
	grep_step1 := grep_steps.Steps[1]
	grep_step2 := grep_steps.Steps[2]
	assert.Equal(t, "step1", grep_step0.Title)
	assert.Equal(t, 1, grep_step0.Index)
	assert.Equal(t, "Application error", grep_step0.GrepPattern)
	assert.Equal(t, "Memory", grep_step0.Output)
	assert.Equal(t, "clientexp7#1.txt", grep_step0.InputFile)

	assert.Equal(t, "step2", grep_step1.Title)
	assert.Equal(t, "1", grep_step1.Select_Ref)
	assert.Equal(t, ".* HandshakeDestConnID ([0123456789abcdef]{8}).*", grep_step1.RegexpPattern)
	assert.Equal(t, "client7_<select_ref>_res.txt", grep_step2.OutputFilenameTemplate)
}
