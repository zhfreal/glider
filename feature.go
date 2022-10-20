package main

import (
	// comment out the services you don't need to make the compiled binary smaller.
	// _ "github.com/zhfreal/glider/service/xxx"

	// comment out the protocols you don't need to make the compiled binary smaller.
	_ "github.com/zhfreal/glider/proxy/http"
	_ "github.com/zhfreal/glider/proxy/kcp"
	_ "github.com/zhfreal/glider/proxy/mixed"
	_ "github.com/zhfreal/glider/proxy/obfs"
	_ "github.com/zhfreal/glider/proxy/pxyproto"
	_ "github.com/zhfreal/glider/proxy/reject"
	_ "github.com/zhfreal/glider/proxy/smux"
	_ "github.com/zhfreal/glider/proxy/socks4"
	_ "github.com/zhfreal/glider/proxy/socks5"
	_ "github.com/zhfreal/glider/proxy/ss"
	_ "github.com/zhfreal/glider/proxy/ssh"
	_ "github.com/zhfreal/glider/proxy/ssr"
	_ "github.com/zhfreal/glider/proxy/tcp"
	_ "github.com/zhfreal/glider/proxy/tls"
	_ "github.com/zhfreal/glider/proxy/trojan"
	_ "github.com/zhfreal/glider/proxy/udp"
	_ "github.com/zhfreal/glider/proxy/vless"
	_ "github.com/zhfreal/glider/proxy/vmess"
	_ "github.com/zhfreal/glider/proxy/ws"
)
