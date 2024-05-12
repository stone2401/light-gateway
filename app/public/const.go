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

	StatusUp   = 1
	StatusDown = 0
)

var (
	LoadTypeSlice = []string{"HTTP", "TCP", "GRPC"}
)
