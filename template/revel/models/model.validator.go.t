package models

import (
	"regexp"

	"github.com/revel/revel"
)

var emailPattern = regexp.MustCompile("^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$")

func required() revel.Validator {
	return revel.Required{}
}

func min(size int) revel.Validator {
	return revel.MinSize{size}
}

func max(size int) revel.Validator {
	return revel.MaxSize{size}
}

func length(size int) revel.Validator {
	return revel.Length{size}
}

func email() revel.Validator {
	return revel.Match{emailPattern}
}

func match(exp string) revel.Validator {
	return revel.Match{regexp.MustCompile(exp)}
}
