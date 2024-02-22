package MorePossibilityApi

import (
	"context"
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
)

// RealTimeTrafficMirroring 实时流量镜像
type RealTimeTrafficMirroring interface {
	RealTimeTrafficMirroring(server BurpMorePossibilityApi.RealTimeTrafficMirroring_RealTimeTrafficMirroringServer) error
}

// IntruderPayloadProcessorServer 迭代载荷处理器
type IntruderPayloadProcessorServer interface {
	IntruderPayloadProcessor(context.Context, *BurpMorePossibilityApi.PayloadProcessorData) (*HttpStructureStandard.ByteData, error)
}

// IntruderPayloadGeneratorServer 迭代载荷生成器
type IntruderPayloadGeneratorServer interface {
	IntruderPayloadGeneratorProvider(context.Context, *BurpMorePossibilityApi.IntruderGeneratorData) (*BurpMorePossibilityApi.PayloadGeneratorResult, error)
}

// HttpReqEditBoxAssistServer http请求编辑框辅助
type HttpReqEditBoxAssistServer interface {
	ReqHttpEdit(context.Context, *BurpMorePossibilityApi.HttpEditBoxData) (*HttpStructureStandard.ByteData, error)
	IsReqHttpEditFor(context.Context, *BurpMorePossibilityApi.HttpEditBoxData) (*HttpStructureStandard.Boole, error)
}

// HttpResEditBoxAssistServer http响应编辑框辅助
type HttpResEditBoxAssistServer interface {
	ResHttpEdit(context.Context, *BurpMorePossibilityApi.HttpEditBoxData) (*HttpStructureStandard.ByteData, error)
	IsResHttpEditFor(context.Context, *BurpMorePossibilityApi.HttpEditBoxData) (*HttpStructureStandard.Boole, error)
}

// GetConTextMenuItemsServer 获取右键菜单
type GetConTextMenuItemsServer interface {
	GetConTextMenuItems(context.Context, *HttpStructureStandard.Str) (*BurpMorePossibilityApi.MenuInfo, error)
}

// ContextMenuItemsProviderServer 右键菜单执行
type ContextMenuItemsProviderServer interface {
	MenuItemsProvider(context.Context, *BurpMorePossibilityApi.ContextMenuItems) (*BurpMorePossibilityApi.MenuItemsReturn, error)
}

// ProxyRequestHandlerServer 代理请求处理器
type ProxyRequestHandlerServer interface {
	ProxyHandleRequestReceived(context.Context, *HttpStructureStandard.HttpReqGroup) (*BurpMorePossibilityApi.ProxyRequestAction, error)
}

// ProxyResponseHandlerServer 代理响应处理器
type ProxyResponseHandlerServer interface {
	ProxyHandleResponseReceived(context.Context, *HttpStructureStandard.HttpReqAndRes) (*BurpMorePossibilityApi.ProxyResponseAction, error)
}

// HttpFlowHandlerServer http流处理器
type HttpFlowHandlerServer interface {
	HttpHandleRequestReceived(context.Context, *BurpMorePossibilityApi.HttpFlowReqData) (*BurpMorePossibilityApi.HttpRequestAction, error)
	HttpHandleResponseReceived(context.Context, *BurpMorePossibilityApi.HttpFlowResData) (*BurpMorePossibilityApi.HttpResponseAction, error)
}
