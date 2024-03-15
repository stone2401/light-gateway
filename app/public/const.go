package public

const (
	LoadTypeHttp = iota
	LoadTypeTcp
	LoadTypeGrpc

	HTTPRuleTypePrefixURL = 0
	HTTPRuleTypeDomainURL = 1
	EndMark               = "。"
)

var (
	LoadTypeSlice = []string{"HTTP", "TCP", "GRPC"}
)
