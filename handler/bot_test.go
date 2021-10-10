package handler

import "testing"

func TestBotReply(t *testing.T) {
	askEn := "test"
	askZh := "测试"
	if TlBot(askEn) == "" || TlBot(askZh) == "" {
		t.Error("tlBot error")
	}
}
