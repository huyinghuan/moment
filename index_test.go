package moment

import (
	"log"
	"testing"
	"time"
)

func Test_Today(t *testing.T) {
	now := time.Now()
	nowstr := now.Format("15:04:05")

	fromNow := New(now)
	fromUinx := NewFromUinx(now.Unix())
	if fromNow.Format("HH:mm:ss") != nowstr {
		t.Fail()
	} else {
		log.Println("pass 1")
	}

	if fromUinx.Format("HH:mm:ss") != nowstr {
		t.Fail()
	} else {
		log.Println("pass 2")
	}

	nowAddTime := now.AddDate(1, 0, 0)
	fromNowAddTime := fromNow.AddDate(1, 0, 0)
	fromUinxAddTime := fromUinx.AddDate(1, 0, 0)

	if fromNowAddTime.Format("HH:mm:ss") != nowAddTime.Format("15:04:05") {
		t.Fail()
	} else {
		log.Println("pass 3")
	}

	if fromUinxAddTime.Format("HH:mm:ss") != nowAddTime.Format("15:04:05") {
		t.Fail()
	} else {
		log.Println("pass 4")
	}

}
