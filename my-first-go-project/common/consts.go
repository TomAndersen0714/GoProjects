package common

import "time"

const (
	AppTimeZoneStr      = "Asia/Shanghai"
	TimerFormat         = "2006-01-02 15:04:05.000"
	DefaultTimeFormat   = "2006-01-02 15:04:05"
	XingHuanCompany     = "星环"
	DefaultRemindSource = "智能质检"
	DefaultRemindTitle  = "会话异常"
	AdminAccount        = "产品开通"

	EmployeeWork = 1
	EmployeeQuit = 2
	InitPassword = "a123456"

	DepartmentDepth = 4

	SnickStatus_Invalid = 0
	SnickStatus_Valid   = 1

	GroupMaxDepth    = 5
	PreSaleDepName   = "售前组"
	AfterSaleDepName = "售后组"

	DefaultTplCompany   = "产品开通" // 用该公司的模板为其他公司开通
	DefaultTplNameRegex = "专家"   // 包含专家字样的模板需要复制
)

var AppTimeZone, _ = time.LoadLocation(AppTimeZoneStr)

const PasswordFormat = `^[\da-zA-Z]{6,18}$`

const LabelNameFormat = "^[\u4e00-\u9fa5_A-Za-z0-9]{1,15}$"

var AllowFileType = map[string]bool{
	".xlsx": true,
	".Xls":  true,
	".doc":  true,
	".png":  true,
	".docx": true,
	".ppt":  true,
	".pptx": true,
	".pdf":  true,
	".jpg":  true,
	".bmp":  true,
	".zip":  true,
}

var BoolText = map[bool]string{
	true:  "是",
	false: "否",
}

var SexMap = map[int]string{
	1: "男",
	2: "女",
}

var AlertLevelMap = map[AlertLevel]string{
	AlertLevel_Low:  "初级告警",
	AlertLevel_Mid:  "中级告警",
	AlertLevel_High: "高级告警",
}

type BoolInt int

const (
	BoolInt_All   BoolInt = 0
	BoolInt_False BoolInt = 1
	BoolInt_True  BoolInt = 2
)

type AlertLevel int

const (
	AlertLevel_No   = 0
	AlertLevel_Low  = 1
	AlertLevel_Mid  = 2
	AlertLevel_High = 3
)
