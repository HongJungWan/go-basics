package post

import "fmt"

// 우체국 전송 담당 구조체
type PostSender struct {
}

func (ps *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다.\n", parcel)
}
