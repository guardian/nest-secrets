package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
)

func TestAsKV(t *testing.T) {
	// Prefixed name
	name1 := "foo/name"
	value1 := "value"

	// Period, slash, hyphen replacement
	name2 := "name.baz/bar-bat"
	value2 := "value"

	// Value escaping single quote, to prevent injection attack
	name3 := "name"
	value3 := "value' \" ${hax}"

	param1 := ssm.Parameter{Name: &name1, Value: &value1}
	param2 := ssm.Parameter{Name: &name2, Value: &value2}
	param3 := ssm.Parameter{Name: &name3, Value: &value3}

	params := []*ssm.Parameter{
		&param1,
		&param2,
		&param3,
	}

	paramString := asKV(params, "foo")

	expectedString := 
		"name='value'\n" + 
		"name_baz_bar_bat='value'\n" +
		"name='value\\' \" ${hax}'\n"
		
	if paramString != expectedString {
		t.Error("Failed: " + paramString + " does not match " + expectedString)
	}
}