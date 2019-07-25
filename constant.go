package wechatpay

const (
	// TODO: 动态容灾域名
	DomainApi   = "https://api.mch.weixin.qq.com"
	DomainApi2  = "api2.mch.weixin.qq.com"
	DomainApiHK = "apihk.mch.weixin.qq.com"
	DomainApiUS = "apius.mch.weixin.qq.com"

	Fail    = "FAIL"
	Success = "SUCCESS"

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
	TransfersUrlSuffix        = "/mmpaymkttransfers/promotion/transfers" // 企业付款到零钱
	GettransferinfoUrlSuffix  = "/mmpaymkttransfers/gettransferinfo"     // 查询企业付款

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
	SandboxTransfersUrlSuffix        = "/sandboxnew/mmpaymkttransfers/promotion/transfers"
	SandboxGettransferinfoUrlSuffix  = "/sandboxnew/mmpaymkttransfers/gettransferinfo"
)
