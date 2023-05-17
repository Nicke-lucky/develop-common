package error

const (
	StatusSuccessfully              = 2000
	StatusErrorInternal             = 2200
	StatusErrorParameter            = 2201
	StatusErrorAddAppLanguages      = 2202
	StatusErrorDeleteAppLanguages   = 2203
	StatusErrorAddAppLanInfos       = 2204
	StatusErrorQueryAppLanInfos     = 2205
	StatusErrorUpdateApplanInfos    = 2206
	StatusErrorAddFileInfos         = 2207
	StatusErrorQueryFileSetInfos    = 2208
	StatusErrorUpdateFileInfos      = 2209
	StatusErrorAddGeneralInfos      = 2210
	StatusErrorQueryGeneralInfos    = 2211
	StatusErrorUpdateGeneralInfos   = 2212
	StatusErrorDeleteAppLanInfos    = 2213
	StatusErrorPublish              = 2214
	StatusErrorQueryAllSetLanguages = 2215
	StatusErrorQueryAppSetInfos     = 2216
	StatusErrorQueryWebSetInfos     = 2217
	StatusErrorUploadFile           = 2218
)

var statusText = map[int]string{
	StatusSuccessfully:              "Success",
	StatusErrorInternal:             "多语言配置服务内部错误",
	StatusErrorParameter:            "参数错误",
	StatusErrorAddAppLanguages:      "设置应用关联语种失败",
	StatusErrorQueryAllSetLanguages: "查询所有关联语种失败",
	StatusErrorDeleteAppLanguages:   "删除应用App关联语种失败",
	StatusErrorAddAppLanInfos:       "新增App多语言信息失败",
	StatusErrorQueryAppLanInfos:     "查询多语言配置信息失败",
	StatusErrorUpdateApplanInfos:    "修改多语言配置信息失败",
	StatusErrorAddFileInfos:         "新增图标、logo、文本信息的多语言配置信息失败",
	StatusErrorQueryFileSetInfos:    "查询图标、logo、文件信息的多语言配置信息失败",
	StatusErrorUpdateFileInfos:      "修改图标、logo、文本文件的多语言配置信息失败",
	StatusErrorAddGeneralInfos:      "新增App通用配置信息失败",
	StatusErrorQueryGeneralInfos:    "查询通用配置信息失败",
	StatusErrorUpdateGeneralInfos:   "修改通用配置信息失败",
	StatusErrorDeleteAppLanInfos:    "删除多语言配置信息失败",
	StatusErrorPublish:              "多语言配置发布生效失败",
	StatusErrorQueryAppSetInfos:     "app拉取多语言配置信息失败",
	StatusErrorQueryWebSetInfos:     "WEB拉取多语言配置信息失败",
	StatusErrorUploadFile:           "图标、logo、文本信息上传失败",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
