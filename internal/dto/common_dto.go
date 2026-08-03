package dto

// IDParam — Dung cho tat ca endpoint co :id tren URL
// VD: /status/:id, /roles/:id, /permissions/:id
//
// Tag giai thich:
//   uri:"id"       → Gin se lay gia tri tu URL parameter co ten "id"
//   binding:"required,gt=0" → Bat buoc phai co, va phai lon hon 0
//
// Khi dung: thay vi c.Param("id") + validation thu cong,
//           chi can c.ShouldBindUri(&param)
type ID_Param struct {
	ID int32 `uri:"id" binding:"required;gt=0"`
}

// UUIDParam — Dung cho endpoint co :uuid tren URL
// VD: /users/:uuid
//
// Tag giai thich:
//   uri:"uuid"            → Lay gia tri tu parameter "uuid" tren URL
//   binding:"required,uuid" → Bat buoc + phai dung dinh dang UUID
type UUID_Param struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

type Email_Param struct {
	Email string `uri:"email" binding:"required;email"`
}

type Phone_Param struct {
	Phone string `uri:"phone" binding:"required;regex=^(0?)(3[2-9]|5[6|8|9]|7[0|6-9]|8[0-6|8|9]|9[0-4|6-9])[0-9]{7}$"`
}

type PaginationQuery struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}
