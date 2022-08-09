package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := zero.New()
	engine.OnFullMatch("hello").Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("hello world!"))
	})
}

func main() {
	zero.RunAndBlock(zero.Config{
		NickName:      []string{"肥肥"},
		CommandPrefix: "/",
		SuperUsers:    []int64{2140579318}, //超级管理员

		Driver: []zero.Driver{
			driver.NewWebSocketClient("ws://127.0.0.1:6700/", ""),
		},
	}, nil)
}
