package violetlog

type LogRequest struct {
	ErrorLevel string `json:"error_level"`
	Message    string `json:"message,omitempty"`
	StackTrace string `json:"stack_trace"`
}

type Log struct {
	Id         int64  `json:"id"`
	AppId      int64  `json:"app_id"`
	ErrorLevel string `json:"error_level"`
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
	CreatedAt  int64  `json:"created_at"`
}
