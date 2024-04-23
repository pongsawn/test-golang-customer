package response

// ปั้น struct  สำหรับ return response
type CustomerResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
