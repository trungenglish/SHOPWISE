package handler

// Deprecated request/response types remain for API contract compatibility.
// Resume links are issued and consumed only by internal/resume_session.
type GenerateResumeTokenRequest struct {
	ExpiresInSeconds int `json:"expires_in_seconds"`
}

type GenerateResumeTokenResponse struct {
	Token string `json:"token"`
}
