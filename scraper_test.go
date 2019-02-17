package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestgetLocation(t *testing.T){

		assert.Equal(t, "Senior Software Architecture Engineer at Casenet (Bedford, MA)", "Bedford")
		assert.Equal(t,"Senior Software Architecture Engineer at Casenet Here HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HRERE HERE HERE23412341234123423141234123421342134213421342134231423142134321 (12341234123412343214321432141234231412341234123421341234123412342134213421342134213423142314213421342134213412Bedford, MA)","12341234123412343214321432141234231412341234123421341234123412342134213421342134213423142314213421342134213412Bedford")
		assert.Equal(t,"S at C (B, MA)","B" )
		assert.Equal(t, "Senior Software Engineer at Microsoft (Boston, MA)", "Boston")
}

func TestgetTitle(t *testing.T) {

	assert.Equal(t,"Senior Software Architecture Engineer at Casenet (Bedford, MA)", "Senior Software Architecture Engineer at Casenet ")
	assert.Equal(t,"Senior Software Architecture Engineer at Casenet Here HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HRERE HERE HERE23412341234123423141234123421342134213421342134231423142134321 (12341234123412343214321432141234231412341234123421341234123412342134213421342134213423142314213421342134213412Bedford, MA)", "Senior Software Architecture Engineer at Casenet Here HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HRERE HERE HERE23412341234123423141234123421342134213421342134231423142134321 ")
	assert.Equal(t,"S at C (B, MA)", "S at C ")
}


func TestgetCompany(t *testing.T){

	assert.Equal(t,"Senior Software Architecture Engineer at Casenet (Bedford, MA)"," Casenet ")
	assert.Equal(t,"Senior Software Architecture Engineer at Casenet Here HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HRERE HERE HERE23412341234123423141234123421342134213421342134231423142134321 (12341234123412343214321432141234231412341234123421341234123412342134213421342134213423142314213421342134213412Bedford, MA)"," Casenet Here HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HERE HRERE HERE HERE23412341234123423141234123421342134213421342134231423142134321 ")
	assert.Equal(t,"S at C (B, MA)"," C ")

}
