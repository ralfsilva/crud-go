package rest_err

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

// lista dos campos incorretos na aplicação
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Recebe os campos definidos
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}
