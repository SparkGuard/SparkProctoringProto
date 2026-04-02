// Package proto provides shared constants and helper types used across SparkProctoring modules.
//
// Data structures (EventPayload, SessionInfo, ChunkMeta, etc.) are defined in .proto files
// and generated into gen/proto/. This package only holds REST-specific constants,
// string mappings, and utilities that don't fit into protobuf definitions.
package proto

// Block data type bytes — used in SparkProctoringStorage block headers.
// These MUST match the DataType enum values in common.proto.
const (
	TypeUnknown    uint8 = 0x00
	TypeVideo      uint8 = 0x01 // H.264 MPEG-TS video chunk
	TypeEvent      uint8 = 0x02 // JSON event log
	TypeScreenshot uint8 = 0x03 // PNG/JPEG screenshot
	TypeMeta       uint8 = 0x04 // Session metadata
	TypeNetwork    uint8 = 0x05 // Network events
)

// Event type strings — used in JSON event payloads (REST transport).
// These MUST match the EventType enum semantics in common.proto.
const (
	EventMouse        = "mouse"
	EventKeyboard     = "keyboard"
	EventDNSQuery     = "dns_query"
	EventNetworkStats = "network_stats"
	EventAppSwitch    = "app_switch"
	EventClipboard    = "clipboard"
	EventDeviceChange = "device_change"
)

// Session status strings — used in REST API JSON responses.
const (
	StatusPending   = "pending"
	StatusActive    = "active"
	StatusPaused    = "paused"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

// Stream type strings — used in REST API and S3 path construction.
const (
	StreamScreen = "screen"
	StreamWebcam = "webcam"
)

// User role strings — used in JWT claims and REST API.
const (
	RoleStudent = "student"
	RoleTeacher = "teacher"
	RoleAdmin   = "admin"
)
