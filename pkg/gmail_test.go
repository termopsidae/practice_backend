package pkg

import (
	"fmt"
	"testing"
)

func TestSendGmail(t *testing.T) {
	//err := SendGmailCode("guodayang999@gmail.com", "123456")
	err := SendGmailCode("pooluo.service@gmail.com", "kbnuygexyjsukjlc", "763917746@qq.com", "123456")
	fmt.Println("err:", err)
}
