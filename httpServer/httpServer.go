package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	server := &http.Server{
		Addr:    ":443",
		Handler: yourHandler(),
	}

	println("Server is running on :443")
	err := server.ListenAndServeTLS("path/to/cert.pem", "path/to/key.pem")
	if err != nil {
		println("Error starting server:", err)
	}
}

func yourHandler() http.Handler {
	targetURL, _ := url.Parse("http://localhost:8080")
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 在这里可以添加自定义处理逻辑
		// 例如设置自定义头部、验证用户身份等

		// 执行反向代理
		proxy.ServeHTTP(w, r)
	})
}
