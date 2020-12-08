package authorization

import "time"

type signup struct {
	Email      string `json:"email"`
	InviteCode string `json:"inviteCode"`
	Language   string `json:"language"`
}

type resetPassword struct {
	Email string `json:"email"`
}

type requestAccessToken struct {
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
	PublicSessionID string `json:"sessionID"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
}

type session struct {
	SessionID string    `json:"sessionID"`
	UserData  string    `json:"userData"`
	UpdatedAt time.Time `json:"updatedAt"`
	Platform  string    `json:"platform"`
}
