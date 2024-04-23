package request

// ปั้น struct update สำหรับรับ request
type CustomerUpdateRequest struct {
	Id   int    `validate:"required min=1,max=10" json:"id"`
	Name string `validate:"required min=1,max=100" json:"name"`
	Age  int    `validate:"required min=1,max=10" json:"age"`
}
