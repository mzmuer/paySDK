package wechat_pay

const (
	DomainApi   = "api.mch.weixin.qq.com"
	DomainApi2  = "api2.mch.weixin.qq.com"
	DomainApiHK = "apihk.mch.weixin.qq.com"
	DomainApiUS = "apius.mch.weixin.qq.com"

	Fail       = "FAIL"
	Success    = "SUCCESS"
	HMACSHA256 = "HMAC-SHA256"
	MD5        = "MD5"

	FieldSign     = "sign"
	FieldSignType = "sign_type"

	MicropayUrlSuffix         = "/pay/micropay"
	UnifiedorderUrlSuffix     = "/pay/unifiedorder"
	OrderqueryUrlSuffix       = "/pay/orderquery"
	ReverseUrlSuffix          = "/secapi/pay/reverse"
	CloseorderUrlSuffix       = "/pay/closeorder"
	RefundUrlSuffix           = "/secapi/pay/refund"
	RefundqueryUrlSuffix      = "/pay/refundquery"
	DownloadbillUrlSuffix     = "/pay/downloadbill"
	ReportUrlSuffix           = "/payitil/report"
	ShorturlUrlSuffix         = "/tools/shorturl"
	AuthcodetoopenidUrlSuffix = "/tools/authcodetoopenid"

	// sandbox
	SandboxMicropayUrlSuffix         = "/sandboxnew/pay/micropay"
	SandboxUnifiedorderUrlSuffix     = "/sandboxnew/pay/unifiedorder"
	SandboxOrderqueryUrlSuffix       = "/sandboxnew/pay/orderquery"
	SandboxReverseUrlSuffix          = "/sandboxnew/secapi/pay/reverse"
	SandboxCloseorderUrlSuffix       = "/sandboxnew/pay/closeorder"
	SandboxRefundUrlSuffix           = "/sandboxnew/secapi/pay/refund"
	SandboxRefundqueryUrlSuffix      = "/sandboxnew/pay/refundquery"
	SandboxDownloadbillUrlSuffix     = "/sandboxnew/pay/downloadbill"
	SandboxReportUrlSuffix           = "/sandboxnew/payitil/report"
	SandboxShorturlUrlSuffix         = "/sandboxnew/tools/shorturl"
	SandboxAuthcodetoopenidUrlSuffix = "/sandboxnew/tools/authcodetoopenid"
)
