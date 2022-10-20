package smux

import "github.com/zhfreal/glider/proxy"

func init() {
	proxy.AddUsage("smux", `
Smux scheme:
  smux://host:port
`)
}
