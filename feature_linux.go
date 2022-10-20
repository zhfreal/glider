package main

import (
	// comment out the services you don't need to make the compiled binary smaller.
	_ "github.com/zhfreal/glider/service/dhcpd"

	// comment out the protocols you don't need to make the compiled binary smaller.
	_ "github.com/zhfreal/glider/proxy/redir"
	_ "github.com/zhfreal/glider/proxy/tproxy"
	_ "github.com/zhfreal/glider/proxy/unix"
	_ "github.com/zhfreal/glider/proxy/vsock"
)
