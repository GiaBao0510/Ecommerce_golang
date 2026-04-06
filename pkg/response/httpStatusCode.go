package response

/*
	Đây là định nghĩa các mã lỗi và thông điệp code nội bộ giữa FrontEnd và Backend để biết,
thường thì sẽ thông bão lỗi chung chung trong client vì nếu thông báo lỗi chi tiết sẽ dễ bị hacker lợi dụng để tấn công 
vào hệ thống, nên sẽ có một mã lỗi chung chung để thông báo lỗi cho client, còn chi tiết lỗi sẽ được log lại trong server để developer 
có thể debug và fix lỗi.
*/
const (
	ErrorCodeSuccess = 2001			// Success
	ErrorCodeParamInvalid = 2003	// Email is invalid
	ErrorInvalidToken = 2004		// Unauthorized
)

//Message response: Giải thích chi tiết về mã lỗi, để developer có thể debug và fix lỗi dễ dàng hơn
var msg = map[int]string {
	ErrorCodeSuccess: "Success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorInvalidToken: "Token is invalid",
}