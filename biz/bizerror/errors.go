package bizerror

type BizError struct {
	ErrMsg   string     `json:"err_msg"`
	ErrCode  ErrorCode  `json:"err_code"`
	ErrLevel ErrorLevel `json:"err_level"`
}

func New(level ErrorLevel, code ErrorCode, msg string) *BizError {
	return &BizError{
		ErrMsg:   msg,
		ErrCode:  code,
		ErrLevel: level,
	}
}

func (e *BizError) Error() string {
	return e.ErrMsg
}
