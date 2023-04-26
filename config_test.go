package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig(t *testing.T) {
	grep_steps := NewConfig("clientsteps.json")
	grep_step0 := grep_steps.Steps[0]
	grep_step1 := grep_steps.Steps[1]
	assert.Equal(t, "step1", grep_step0.Title)
	assert.Equal(t, "Closing connection HandshakeDestConnID [0123456789abcdef]* .*SrcConnID <nil> with error: application error 0x0 (local)", grep_step0.GrepPattern)
	assert.Equal(t, "memory", grep_step0.Output)
	assert.Equal(t, "clientexp7#1.txt", grep_step0.InputFile)

	assert.Equal(t, "step2", grep_step1.Title)
	assert.Equal(t, "step1", grep_step1.Select_Ref)
	assert.Equal(t, "(.*) HandshakeDestConnID ([0123456789{8}])(.*)", grep_step1.RegexpPattern)
}
