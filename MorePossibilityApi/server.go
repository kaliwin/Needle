package MorePossibilityApi

import (
	"context"
	"fmt"
	"github.com/kaliwin/Needle/MorePossibilityApi/grpc/BurpMorePossibilityApi"
	"github.com/kaliwin/Needle/PublicStandard/newsletter/StandardHttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

// ApiTest api测试
type ApiTest struct {
	BurpMorePossibilityApi.UnimplementedHttpFlowHandlerServer
}

func (n ApiTest) HttpHandleRequestReceived(ctx context.Context, data *BurpMorePossibilityApi.HttpFlowReqData) (*BurpMorePossibilityApi.HttpRequestAction, error) {
	//TODO implement me
	return &BurpMorePossibilityApi.HttpRequestAction{Continue: true}, nil
}

func (n ApiTest) HttpHandleResponseReceived(ctx context.Context, data *BurpMorePossibilityApi.HttpFlowResData) (*BurpMorePossibilityApi.HttpResponseAction, error) {

	fmt.Println(data.GetHttpReqAndRes().GetReq().GetUrl())
	group, err := StandardHttp.BuildStandardHttpGroup(data.GetHttpReqAndRes(), nil)

	if err != nil {
		log.Println(err)

	} else {
		fmt.Println(group.Req.GetUrl())
		fmt.Println(group.Req.GetHead())
	}

	return &BurpMorePossibilityApi.HttpResponseAction{Continue: true}, nil
}

// NewGrpcServer 创建一个新的服务
func NewGrpcServer(address string, opt ...grpc.ServerOption) (BurpGrpcServer, error) {
	server := BurpGrpcServer{}
	server.serverStatus = ""
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return server, err
	}
	server.netList = listen
	server.Server = grpc.NewServer(opt...)
	return server, nil
}

// BurpGrpcServer burp的grpc服务
type BurpGrpcServer struct {
	Server       *grpc.Server
	netList      net.Listener
	serverStatus string
}

// RegisterRealTimeTrafficMirroring 注册实时流量镜像
func (s BurpGrpcServer) RegisterRealTimeTrafficMirroring(r RealTimeTrafficMirroring) {
	BurpMorePossibilityApi.RegisterRealTimeTrafficMirroringServer(s.Server, realTimeTrafficMirroring{realTimeTrafficMirroring: r})
}

// RegisterIntruderPayloadProcessorServer 注册负载处理器
func (s BurpGrpcServer) RegisterIntruderPayloadProcessorServer(i IntruderPayloadProcessorServer) {
	BurpMorePossibilityApi.RegisterIntruderPayloadProcessorServerServer(s.Server, intruderPayloadProcessor{intruderPayloadProcessorServer: i})
}

// RegisterIntruderPayloadGeneratorServer 注册负载生成器
func (s BurpGrpcServer) RegisterIntruderPayloadGeneratorServer(i IntruderPayloadGeneratorServer) {
	BurpMorePossibilityApi.RegisterIntruderPayloadGeneratorServerServer(s.Server, intruderPayloadGenerator{intruderPayloadGeneratorServer: i})
}

// RegisterHttpReqEditBoxAssistServer 注册http请求编辑框辅助
func (s BurpGrpcServer) RegisterHttpReqEditBoxAssistServer(h HttpReqEditBoxAssistServer) {
	BurpMorePossibilityApi.RegisterHttpReqEditBoxAssistServer(s.Server, httpReqEditBoxAssist{httpReqEditBoxAssistServer: h})
}

// RegisterHttpResEditBoxAssistServer 注册http响应编辑框辅助
func (s BurpGrpcServer) RegisterHttpResEditBoxAssistServer(h HttpResEditBoxAssistServer) {
	BurpMorePossibilityApi.RegisterHttpResEditBoxAssistServer(s.Server, httpResEditBoxAssist{httpResEditBoxAssistServer: h})
}

// RegisterGetConTextMenuItemsServer 注册右键菜单
func (s BurpGrpcServer) RegisterGetConTextMenuItemsServer(g GetConTextMenuItemsServer) {
	BurpMorePossibilityApi.RegisterGetConTextMenuItemsServerServer(s.Server, getConTextMenuItems{getConTextMenuItemsServer: g})
}

// RegisterContextMenuItemsProviderServer 注册右键菜单执行
func (s BurpGrpcServer) RegisterContextMenuItemsProviderServer(c ContextMenuItemsProviderServer) {
	BurpMorePossibilityApi.RegisterContextMenuItemsProviderServer(s.Server, contextMenuItemsProvider{contextMenuItemsProviderServer: c})
}

// RegisterProxyRequestHandlerServer 注册代理请求处理器
func (s BurpGrpcServer) RegisterProxyRequestHandlerServer(p ProxyRequestHandlerServer) {
	BurpMorePossibilityApi.RegisterProxyRequestHandlerServer(s.Server, proxyRequestHandler{proxyRequestHandlerServer: p})
}

// RegisterProxyResponseHandlerServer 注册代理响应处理器
func (s BurpGrpcServer) RegisterProxyResponseHandlerServer(p ProxyResponseHandlerServer) {
	BurpMorePossibilityApi.RegisterProxyResponseHandlerServer(s.Server, proxyResponseHandler{proxyResponseHandlerServer: p})
}

// RegisterHttpFlowHandlerServer 注册http流处理器
func (s BurpGrpcServer) RegisterHttpFlowHandlerServer(h HttpFlowHandlerServer) {
	BurpMorePossibilityApi.RegisterHttpFlowHandlerServer(s.Server, HttpFlowHandler{httpFlowHandlerServer: h})
}

//  ------------------------------------------------  //

// Start 启动服务 成功会陷入堵塞
func (s BurpGrpcServer) Start() error {
	return s.Server.Serve(s.netList)
}

// GetStatus 获取服务状态
func (s BurpGrpcServer) GetStatus() error {
	if s.serverStatus != "" {
		return fmt.Errorf(s.serverStatus)
	}
	return nil
}

// Stop 停止服务
func (s BurpGrpcServer) Stop() {
	s.Server.Stop()
}

// realTimeTrafficMirroring 实时流量镜像
type realTimeTrafficMirroring struct {
	BurpMorePossibilityApi.UnimplementedRealTimeTrafficMirroringServer
	realTimeTrafficMirroring RealTimeTrafficMirroring
}

// RealTimeTrafficMirroring 实时流量镜像
func (r realTimeTrafficMirroring) RealTimeTrafficMirroring(server BurpMorePossibilityApi.RealTimeTrafficMirroring_RealTimeTrafficMirroringServer) error {
	if r.realTimeTrafficMirroring != nil {
		return r.realTimeTrafficMirroring.RealTimeTrafficMirroring(server)
	}
	return status.Errorf(codes.Unimplemented, "method RealTimeTrafficMirroring not implemented")
}

// intruderPayloadProcessor 迭代载荷处理器
type intruderPayloadProcessor struct {
	BurpMorePossibilityApi.UnimplementedIntruderPayloadProcessorServerServer
	intruderPayloadProcessorServer IntruderPayloadProcessorServer
}

// IntruderPayloadProcessor 迭代载荷处理器
func (i intruderPayloadProcessor) IntruderPayloadProcessor(c context.Context, p *BurpMorePossibilityApi.PayloadProcessorData) (*BurpMorePossibilityApi.ByteData, error) {
	if i.intruderPayloadProcessorServer != nil {
		return i.intruderPayloadProcessorServer.IntruderPayloadProcessor(c, p)

	}
	return nil, status.Errorf(codes.Unimplemented, "method IntruderPayloadProcessor not implemented")
}

// intruderPayloadGenerator 迭代载荷生成器
type intruderPayloadGenerator struct {
	BurpMorePossibilityApi.UnimplementedIntruderPayloadGeneratorServerServer
	intruderPayloadGeneratorServer IntruderPayloadGeneratorServer
}

// IntruderPayloadGeneratorProvider 迭代载荷生成器
func (i intruderPayloadGenerator) IntruderPayloadGeneratorProvider(c context.Context, ig *BurpMorePossibilityApi.IntruderGeneratorData) (*BurpMorePossibilityApi.PayloadGeneratorResult, error) {
	if i.intruderPayloadGeneratorServer != nil {
		return i.intruderPayloadGeneratorServer.IntruderPayloadGeneratorProvider(c, ig)
	}
	return nil, status.Errorf(codes.Unimplemented, "method IntruderPayloadGeneratorProvider not implemented")
}

// httpReqEditBoxAssist http请求编辑框辅助
type httpReqEditBoxAssist struct {
	BurpMorePossibilityApi.UnimplementedHttpReqEditBoxAssistServer
	httpReqEditBoxAssistServer HttpReqEditBoxAssistServer
}

func (h httpReqEditBoxAssist) ReqHttpEdit(c context.Context, r *BurpMorePossibilityApi.HttpEditBoxData) (*BurpMorePossibilityApi.ByteData, error) {
	if h.httpReqEditBoxAssistServer != nil {
		return h.httpReqEditBoxAssistServer.ReqHttpEdit(c, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method ReqHttpEdit not implemented")
}
func (h httpReqEditBoxAssist) IsReqHttpEditFor(c context.Context, p *BurpMorePossibilityApi.HttpEditBoxData) (*BurpMorePossibilityApi.Boole, error) {
	if h.httpReqEditBoxAssistServer != nil {
		return h.httpReqEditBoxAssistServer.IsReqHttpEditFor(c, p)
	}
	return nil, status.Errorf(codes.Unimplemented, "method ReqHttpEdit not implemented")
}

// httpResEditBoxAssist http响应编辑框辅助
type httpResEditBoxAssist struct {
	BurpMorePossibilityApi.UnimplementedHttpResEditBoxAssistServer
	httpResEditBoxAssistServer HttpResEditBoxAssistServer
}

func (h httpResEditBoxAssist) ResHttpEdit(c context.Context, hb *BurpMorePossibilityApi.HttpEditBoxData) (*BurpMorePossibilityApi.ByteData, error) {
	if h.httpResEditBoxAssistServer != nil {
		return h.httpResEditBoxAssistServer.ResHttpEdit(c, hb)
	}

	return nil, status.Errorf(codes.Unimplemented, "method ResHttpEdit not implemented")
}
func (h httpResEditBoxAssist) IsResHttpEditFor(c context.Context, hb *BurpMorePossibilityApi.HttpEditBoxData) (*BurpMorePossibilityApi.Boole, error) {
	if h.httpResEditBoxAssistServer != nil {
		return h.httpResEditBoxAssistServer.IsResHttpEditFor(c, hb)
	}
	return nil, status.Errorf(codes.Unimplemented, "method IsResHttpEditFor not implemented")
}

// getConTextMenuItems 获取右键菜单
type getConTextMenuItems struct {
	BurpMorePossibilityApi.UnimplementedGetConTextMenuItemsServerServer
	getConTextMenuItemsServer GetConTextMenuItemsServer
}

// GetConTextMenuItems 获取右键菜单
func (g getConTextMenuItems) GetConTextMenuItems(c context.Context, s *BurpMorePossibilityApi.Str) (*BurpMorePossibilityApi.MenuInfo, error) {
	if g.getConTextMenuItemsServer != nil {
		return g.getConTextMenuItemsServer.GetConTextMenuItems(c, s)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetConTextMenuItems not implemented")
}

// contextMenuItemsProvider 右键菜单执行
type contextMenuItemsProvider struct {
	BurpMorePossibilityApi.UnimplementedContextMenuItemsProviderServer
	contextMenuItemsProviderServer ContextMenuItemsProviderServer
}

// MenuItemsProvider 右键菜单执行
func (c contextMenuItemsProvider) MenuItemsProvider(co context.Context, cm *BurpMorePossibilityApi.ContextMenuItems) (*BurpMorePossibilityApi.MenuItemsReturn, error) {
	if c.contextMenuItemsProviderServer != nil {
		return c.contextMenuItemsProviderServer.MenuItemsProvider(co, cm)
	}
	return nil, status.Errorf(codes.Unimplemented, "method MenuItemsProvider not implemented")
}

// proxyRequestHandler 代理请求处理器
type proxyRequestHandler struct {
	BurpMorePossibilityApi.UnimplementedProxyRequestHandlerServer
	proxyRequestHandlerServer ProxyRequestHandlerServer
}

// ProxyHandleRequestReceived 代理请求处理器
func (p proxyRequestHandler) ProxyHandleRequestReceived(c context.Context, h *BurpMorePossibilityApi.HttpReqGroup) (*BurpMorePossibilityApi.ProxyRequestAction, error) {
	if p.proxyRequestHandlerServer != nil {
		return p.proxyRequestHandlerServer.ProxyHandleRequestReceived(c, h)
	}

	return nil, status.Errorf(codes.Unimplemented, "method ProxyHandleRequestReceived not implemented")
}

// proxyResponseHandler 代理响应处理器
type proxyResponseHandler struct {
	BurpMorePossibilityApi.UnimplementedProxyResponseHandlerServer
	proxyResponseHandlerServer ProxyResponseHandlerServer
}

// ProxyHandleResponseReceived 代理响应处理器
func (p proxyResponseHandler) ProxyHandleResponseReceived(c context.Context, hr *BurpMorePossibilityApi.HttpReqAndRes) (*BurpMorePossibilityApi.ProxyResponseAction, error) {
	if p.proxyResponseHandlerServer != nil {
		return p.proxyResponseHandlerServer.ProxyHandleResponseReceived(c, hr)
	}
	return nil, status.Errorf(codes.Unimplemented, "method ProxyHandleResponseReceived not implemented")
}

// HttpFlowHandler http流处理器
type HttpFlowHandler struct {
	BurpMorePossibilityApi.UnimplementedHttpFlowHandlerServer
	httpFlowHandlerServer HttpFlowHandlerServer
}

// HttpHandleRequestReceived http请求处理器
func (h HttpFlowHandler) HttpHandleRequestReceived(c context.Context, req *BurpMorePossibilityApi.HttpFlowReqData) (*BurpMorePossibilityApi.HttpRequestAction, error) {
	if h.httpFlowHandlerServer != nil {
		return h.httpFlowHandlerServer.HttpHandleRequestReceived(c, req)
	}
	return nil, status.Errorf(codes.Unimplemented, "method HttpHandleRequestReceived not implemented")
}

// HttpHandleResponseReceived http响应处理器
func (h HttpFlowHandler) HttpHandleResponseReceived(c context.Context, res *BurpMorePossibilityApi.HttpFlowResData) (*BurpMorePossibilityApi.HttpResponseAction, error) {
	if h.httpFlowHandlerServer != nil {
		return h.httpFlowHandlerServer.HttpHandleResponseReceived(c, res)
	}
	return nil, status.Errorf(codes.Unimplemented, "method HttpHandleResponseReceived not implemented")
}

//
//
//
//
//
//
//
//
//
//
////
////
////
////
//////
////
////
////
//////
////
////
////
//////
////
////
////
//////
////
////
////
//////
////
////
////
////
