package config

const (
	HttpError             = 10000 // 通用失败
	HttpSuccess           = 10001 // 通用成功
	HttpNotMethod         = 10002 // 没有方法
	HttpNotRoute          = 10003 // 没有路由
	HttpNeedLogin         = 10004 // 需要登录
	HttpNeedEmail         = 10005 // 需要邮箱
	HttpNeedPassword      = 10006 // 需要密码
	HttpNeedtype          = 10007 // 需要登录类型
	HttpNeedVertify       = 10008 // 需要第三方code
	HttpNeedCode          = 10009 // 需要code
	HttpEmailNotFound     = 10010 // 邮箱不存在
	HttpCodeValidate      = 10011 // 验证码不正确
	HttpPasswordValidate  = 10012 // 密码错误
	HttpTypeError         = 10013 // 类型错误
	HttpEmailError        = 10014 // 邮箱不正确
	HttpRegisterError     = 10015 // 注册失败
	HttpUserIsRegister    = 10016 // 用户已存在
	HttpRegisterSuccess   = 10017 // 用户创建成功
	HttpNeedToken         = 10018 // 需要token
	HttpUserIsNotRegister = 10019 // 用户不存在
	HttpNeedId            = 10020 // 请求id不存在
	HttpOrderNotFound     = 10021 // 订单不存在
)

const (
	HttpErrorMessage             = "请求失败"
	HttpSuccessMessage           = "" // 成功需要返回的字段
	HttpActionMessage            = "操作成功"
	HttpNotMethodMessage         = "请求方法不存在"
	HttpNotRouteMessage          = "请求路由不存在"
	HttpNeedLoginMessage         = "请先登录"
	HttpNeedEmailMessage         = "请填写邮箱"
	HttpNeedPasswordMessage      = "请填写密码"
	HttpNeedtypeMessage          = "请填写登录类型"
	HttpNeedVertifyMessage       = "第三方登录码有误"
	HttpNeedCodeMessage          = "请填写验证码"
	HttpEmailNotFoundMessage     = "邮箱不存在"
	HttpCodeValidateMessage      = "验证码不正确"
	HttpPasswordValidateMessage  = "密码错误"
	HttpTypeErrorMessage         = "类型错误"
	HttpEmailErrorMessage        = "邮箱格式不正确"
	HttpRegisterErrorMessage     = "注册失败"
	HttpUserIsRegisterMessage    = "用户已存在"
	HttpRegisterSuccessMessage   = "注册成功"
	HttpNeedTokenMessage         = "Token不存在"
	HttpUserIsNotRegisterMessage = "用户不存在"
	HttpNeedIdMessage            = "请求ID不存在"
	HttpOrderNotFoundMessage     = "订单不存在"
)

type Json struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

var (
	Toast = "toast"
	Modal = "modal"
)

var HttpCodeMesaage = map[int]Json{
	HttpError: {
		HttpError,
		HttpErrorMessage,
		Toast,
	},
	HttpSuccess: {
		HttpSuccess,
		HttpSuccessMessage,
		Toast,
	},
	HttpNotMethod: {
		HttpNotMethod,
		HttpNotMethodMessage,
		Toast,
	},
	HttpNotRoute: {
		HttpNotRoute,
		HttpNotRouteMessage,
		Toast,
	},
	HttpNeedLogin: {
		HttpNeedLogin,
		HttpNeedLoginMessage,
		Toast,
	},
	HttpNeedEmail: {
		HttpNeedEmail,
		HttpNeedEmailMessage,
		Toast,
	},
	HttpNeedPassword: {
		HttpNeedPassword,
		HttpNeedPasswordMessage,
		Toast,
	},
	HttpTypeError: {
		HttpTypeError,
		HttpTypeErrorMessage,
		Toast,
	},
	HttpEmailError: {
		HttpEmailError,
		HttpEmailErrorMessage,
		Toast,
	},
	HttpRegisterError: {
		HttpRegisterError,
		HttpRegisterErrorMessage,
		Toast,
	},
	HttpUserIsRegister: {
		HttpUserIsRegister,
		HttpUserIsRegisterMessage,
		Toast,
	},
	HttpNeedToken: {
		HttpNeedToken,
		HttpNeedTokenMessage,
		Toast,
	},
	HttpUserIsNotRegister: {
		HttpUserIsNotRegister,
		HttpUserIsNotRegisterMessage,
		Toast,
	},
	HttpNeedId: {
		HttpNeedId,
		HttpNeedIdMessage,
		Toast,
	},
	HttpOrderNotFound: {
		HttpOrderNotFound,
		HttpOrderNotFoundMessage,
		Toast,
	},
}
