package main

import "play003/lab001/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
