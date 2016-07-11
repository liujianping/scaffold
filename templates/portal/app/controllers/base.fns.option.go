package controllers

import (
	"strings"

	"[[.project]]/app/models"
	"github.com/revel/revel"
)

var selections = map[string]interface{}{}

func GetSelection(code string) []models.SystemOption {
	if items, ok := selections[code]; ok {
		return items.([]models.SystemOption)
	}

	query := models.SystemOptionQuery{Code: code}
	sort := models.SystemOptionSort{2}
	page := models.SystemOptionPage{0, 0}
	_, items, _ := models.DefaultSystemOption.Search(db,
		query, sort, page)

	selections[code] = items
	return items
}

func GetOptionValue(code, option_code string) int64 {
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
	return "none"
}

func init() {
	revel.TemplateFuncs["selection"] = GetSelection
	revel.TemplateFuncs["option"] = GetOptionValue
	revel.TemplateFuncs["option_name"] = GetOptionName
}
