package repository

import "context"

type ISMSRepository interface {
	// SendSMS gửi tin nhắn SMS đến số điện thoại được chỉ định với nội dung tin nhắn.
	SendSMS(ctx context.Context, phoneNumber string, message string) error
}
