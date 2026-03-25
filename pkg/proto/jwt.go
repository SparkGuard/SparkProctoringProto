package proto

// JWT claim keys used in agent and server tokens.
const (
	ClaimSessionID = "session_id"
	ClaimUserID    = "user_id"
	ClaimRole      = "role"
)

// JWT expiration durations.
const (
	// AgentTokenTTLHours is how long an agent JWT stays valid.
	AgentTokenTTLHours = 4

	// WebTokenTTLHours is how long a web UI JWT stays valid.
	WebTokenTTLHours = 24
)

// JWTClaims represents the custom claims embedded in every token.
type JWTClaims struct {
	SessionID string `json:"session_id,omitempty"`
	UserID    string `json:"user_id"`
	Role      string `json:"role"` // one of RoleStudent, RoleTeacher, RoleAdmin
}
