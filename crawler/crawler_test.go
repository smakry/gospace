package crawler

import (
	"testing"
)

func TestCrawler(t *testing.T) {
	DoCrawlerTask("http://forum.nderp.99.com/BugManage/Internal/bugOperationList.htm?taskId=100600")
}