package proto

// API version prefix. All REST endpoints are served under this path.
const APIPrefix = "/api/v1"

// --- REST endpoint paths (relative to APIPrefix) ---
// gRPC methods are defined in agent_service.proto and accessed via gRPC client stubs.
// These constants are ONLY for the REST API used by the web UI.

// Auth endpoints (REST — web UI login).
const (
	// POST — web UI login for teachers/admins.
	EndpointAuthLogin = "/auth/login"
)

// Session management endpoints (REST — web UI).
const (
	// GET — list sessions (query: ?status=active&limit=50&offset=0)
	EndpointSessions = "/sessions"

	// GET/DELETE — session by ID: /sessions/{id}
	EndpointSessionByID = "/sessions/{id}"

	// GET — HLS playlist: /sessions/{id}/playlist/{stream_type}.m3u8
	EndpointSessionPlaylist = "/sessions/{id}/playlist/{stream_type}.m3u8"

	// GET — telemetry events: /sessions/{id}/telemetry?from=...&to=...&type=...
	EndpointSessionTelemetry = "/sessions/{id}/telemetry"
)

// --- REST-only request/response types ---
// These are used by the web UI (JSON over HTTP).
// Agent communication uses gRPC types from gen/proto/.

// LoginRequest is sent by the web UI for teacher/admin authentication.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse is returned after successful web UI login.
type LoginResponse struct {
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"expires_at"`
}

// SessionListResponse wraps a paginated list of sessions for the web UI.
type SessionListResponse struct {
	Sessions []SessionSummary `json:"sessions"`
	Total    int              `json:"total"`
	Limit    int              `json:"limit"`
	Offset   int              `json:"offset"`
}

// SessionSummary is a lightweight session representation for list views.
type SessionSummary struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	Status        string `json:"status"`
	StartedAt     string `json:"started_at,omitempty"` // RFC 3339
	EndedAt       string `json:"ended_at,omitempty"`   // RFC 3339
	ChunksTotal   int    `json:"chunks_total"`
	EventsTotal   int    `json:"events_total"`
	BytesUploaded int64  `json:"bytes_uploaded"`
}

// SessionDetailResponse contains full session data for the teacher's player UI.
type SessionDetailResponse struct {
	Session    SessionSummary `json:"session"`
	Streams    []StreamInfo   `json:"streams"`
	EventCount EventCount     `json:"event_count"`
}

// StreamInfo describes an available video stream within a session.
type StreamInfo struct {
	StreamType  string `json:"stream_type"`
	PlaylistURL string `json:"playlist_url"` // presigned URL to m3u8
	ChunkCount  int    `json:"chunk_count"`
	DurationMs  int64  `json:"duration_ms"`
}

// EventCount provides a breakdown of telemetry events by type.
type EventCount struct {
	Keyboard     int `json:"keyboard"`
	Mouse        int `json:"mouse"`
	DNS          int `json:"dns"`
	NetworkStats int `json:"network_stats"`
	AppSwitch    int `json:"app_switch"`
	Clipboard    int `json:"clipboard"`
	Total        int `json:"total"`
}

// --- Error response ---

// ErrorResponse is the standard error format for all REST API endpoints.
type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

// Standard error codes.
const (
	CodeBadRequest      = "BAD_REQUEST"
	CodeUnauthorized    = "UNAUTHORIZED"
	CodeForbidden       = "FORBIDDEN"
	CodeNotFound        = "NOT_FOUND"
	CodeConflict        = "CONFLICT"
	CodeDuplicate       = "DUPLICATE"
	CodeInternalError   = "INTERNAL_ERROR"
	CodeSessionNotFound = "SESSION_NOT_FOUND"
	CodeSessionExpired  = "SESSION_EXPIRED"
	CodeInvalidToken    = "INVALID_TOKEN"
	CodeUploadFailed    = "UPLOAD_FAILED"
)
