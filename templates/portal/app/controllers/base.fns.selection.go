package controllers

import (
	"strings"

	"github.com/revel/revel"
)

var selections = map[string]interface{}{}

func GetSelection(code string) []models.Option {
	if items, ok := selections[code]; ok {
		return items.([]models.Option)
	}

	query := models.OptionQuery{Code: code}
	sort := models.OptionSort{1}
	page := models.OptionPage{0, 0}
	_, items, _ := models.DefaultOption.Search(db,
		query, sort, page)

	selections[code] = items
	return items
}

func GetOption(code, option_code string) int64 {
	opts := GetSelection(code)
	for _, opt := range opts {
		if strings.ToLower(opt.OptionCode) == strings.ToLower(option_code) {
			return opt.OptionValue
		}
	}
	return 0
}

func GetOptionName(code string, option_value int64) string {
	opts := GetSelection(code)
	for _, opt := range opts {
		if opt.OptionValue == option_value {
			return opt.OptionName
		}
	}
	return "未知"
}

func init() {
	revel.TemplateFuncs["selection"] = GetSelection
	revel.TemplateFuncs["option"] = GetOption
	revel.TemplateFuncs["option_name"] = GetOptionName
}
