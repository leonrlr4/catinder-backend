package dto

// RegisterInfo 用於用戶註冊請求的數據傳輸對象
type RegisterInfo struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Picture       string `json:"picture"`
	OAuthProvider string `json:"o_auth_provider"`
	JwtToken      string `json:"jwt_token"`
}

// UserInfo 用於用戶註冊請求的數據傳輸對象
type UserInfo struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Picture       string `json:"picture"`
	OAuthProvider string `json:"o_auth_provider"`
	JwtToken      string `json:"jwt_token"`
	CreatedAt     string `json:"created_at"`
}

type LocalLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReturnLoginInfo struct {
	ID            int    `json:"id"`
	UserName      string `json:"UserName"`
	Email         string `json:"Email"`
	Picture       string `json:"Picture"`
	OAuthProvider string `json:"OAuthProvider"`
	JwtToken      string `json:"JwtToken"`
	CreatedAt     string `json:"CreatedAt"`
}
