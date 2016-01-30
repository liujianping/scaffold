package models

[[range $k, $v := .options]]//![[$k]]
const ([[range $v]]
[[.Code | camel | lint]][[.OptionCode | camel | lint ]] int64 = [[.OptionValue]] //[[.OptionName]][[end]]
)

func [[$k | camel | lint]]ToCode(status int64) string {
    switch status {[[range $v]]
    case [[.Code | camel | lint]][[.OptionCode | camel | lint]]:
        return "[[.OptionCode]]"[[end]]
    }
    return "unset"
}

func [[$k | camel | lint]]ToString(status int64) string {
    switch status {[[range $v]]
    case [[.Code | camel | lint]][[.OptionCode | camel | lint]]:
        return [[.OptionName | quote]][[end]]
    }
    return "未知"
}
[[end]]
