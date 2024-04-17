package msg

const (
	StatusSuccess        = "success"
	StatusError          = "error"
	StatusInvalidRequest = "invalid_request"
	StatusUnauthorized   = "unauthorized"
	StatusForbidden      = "forbidden"
	StatusNotFound       = "not_found"
	StatusConflict       = "conflict"
	StatusServerError    = "server_error"
	StatusMaintenance    = "maintenance"
	StatusRateLimit      = "rate_limit_exceeded"
	StatusRedirect       = "redirect"
	StatusCustomCode     = "custom_code"
)

type DataWrap struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Refer   interface{} `json:"refer,omitempty"`
}
