package dto

	type VerifyOTPRequest struct {
		Username string `json:"username" binding:"required"`
		Code string `json:"code" binding:"required"`
	}
