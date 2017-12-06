package main

import (
	"github.com/golang/mock/gomock"
	"play003/lab001/spider"
	"testing"
)

func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Errorf("Get wrong version %s", goVer)
	}
}
