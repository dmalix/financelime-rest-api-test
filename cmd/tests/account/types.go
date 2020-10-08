package account

import "time"

type signup struct {
	Email      string `json:"email"`
	InviteCode string `json:"inviteCode"`
	Language   string `json:"language"`
}

type resetPassword struct {
	Email string `json:"email"`
}

type getAccessToken struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Device   device `json:"device"`
}

type device struct {
	Platform string `json:"platform"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Language string `json:"language"`
	Timezone string `json:"timezone"`
}

type jwt struct {
	id              int64
	PublicSessionID string `json:"id"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
}

type session struct {
	PublicSessionID string    `json:"id"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Platform        string    `json:"platform"`
}

type refreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type revokeAccessToken struct {
	PublicSessionID string `json:"id"`
}
