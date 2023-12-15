package SerialNumber

const (
	// XCheckPocNumberStrID Poc 抬头
	XCheckPocNumberStrID = XCheckNumberStrID + SeparatorStr + "Poc" // XCheck下的POC
)

const (
	// XCheckPocPayloadType Poc载荷类型抬头
	XCheckPocPayloadType = XCheckPocNumberStrID + SeparatorStr + "PayloadType" // XCheck下的POC载荷类型 =
)

const (
	// XSQLInjectionPocPayloadType SQL注入载荷类型
	XSQLInjectionPocPayloadType  = XCheckPocPayloadType + SeparatorStr + "SQLInjection"  // SQL注入载荷类型
	XReflectionXSSPocPayloadType = XCheckPocPayloadType + SeparatorStr + "ReflectionXSS" // 反射型XSS载荷类型
)
