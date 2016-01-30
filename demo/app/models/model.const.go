package models

//!user_accounts_sex
const (
	UserAccountsSexMale   int64 = 1 //男
	UserAccountsSexFemale int64 = 2 //女
)

func UserAccountsSexToCode(status int64) string {
	switch status {
	case UserAccountsSexMale:
		return "male"
	case UserAccountsSexFemale:
		return "female"
	}
	return "unset"
}

func UserAccountsSexToString(status int64) string {
	switch status {
	case UserAccountsSexMale:
		return "男"
	case UserAccountsSexFemale:
		return "女"
	}
	return "未知"
}

//!user_accounts_status
const (
	UserAccountsStatusUnactived int64 = 1 //<span class="label label-info">待激活</span>
	UserAccountsStatusNormal    int64 = 2 //<span class="label label-success">正常</span>
	UserAccountsStatusForbidden int64 = 3 //<span class="label label-danger">禁用</span>
)

func UserAccountsStatusToCode(status int64) string {
	switch status {
	case UserAccountsStatusUnactived:
		return "unactived"
	case UserAccountsStatusNormal:
		return "normal"
	case UserAccountsStatusForbidden:
		return "forbidden"
	}
	return "unset"
}

func UserAccountsStatusToString(status int64) string {
	switch status {
	case UserAccountsStatusUnactived:
		return "<span class=\"label label-info\">待激活</span>"
	case UserAccountsStatusNormal:
		return "<span class=\"label label-success\">正常</span>"
	case UserAccountsStatusForbidden:
		return "<span class=\"label label-danger\">禁用</span>"
	}
	return "未知"
}

//!user_posts_status
const (
	UserPostsStatusDraft  int64 = 1 //<span class="label label-primary">待发布</span>
	UserPostsStatusPublic int64 = 2 //<span class="label label-success">已发布</span>
	UserPostsStatusTrash  int64 = 3 //<span class="label label-default">已作废</span>
)

func UserPostsStatusToCode(status int64) string {
	switch status {
	case UserPostsStatusDraft:
		return "draft"
	case UserPostsStatusPublic:
		return "public"
	case UserPostsStatusTrash:
		return "trash"
	}
	return "unset"
}

func UserPostsStatusToString(status int64) string {
	switch status {
	case UserPostsStatusDraft:
		return "<span class=\"label label-primary\">待发布</span>"
	case UserPostsStatusPublic:
		return "<span class=\"label label-success\">已发布</span>"
	case UserPostsStatusTrash:
		return "<span class=\"label label-default\">已作废</span>"
	}
	return "未知"
}
