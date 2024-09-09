package fedex

import "fmt"

// Fedex 전송 담당 구조체
type FedexSender struct {
}

func (fs *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex Sends %v, parcel\n", parcel)
}
