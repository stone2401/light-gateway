package public

const (
	LoadTypeHttp = iota
	LoadTypeTcp
	LoadTypeGrpc

	HTTPRuleTypePrefixURL = 0
	HTTPRuleTypeDomainURL = 1
	EndMark               = "。"
	CAPTCHAKEY            = "captcha"
	CAPTCHALEN            = 5
)

var (
	LoadTypeSlice = []string{"HTTP", "TCP", "GRPC"}
)
