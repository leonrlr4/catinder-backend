package dto

// RegisterInfo 用於用戶註冊請求的數據傳輸對象
type RegisterInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
