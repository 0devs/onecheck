package plugin

type Result struct {
	Status string // OK, WARN, CRIT
	StatusMessage string
	Messages []string
	Errors []string
	Warnings []string
	Data interface{}
}

func (r *Result) SetStatus(status string) {
	r.Status = status
}

func (r *Result) SetStatusOk() {
	r.Status = "ok"
}

func (r *Result) SetStatusWarn() {
	r.Status = "warn"
}

func (r *Result) SetStatusCrit() {
	r.Status = "crit"
}

func (r *Result) SetStatusMessage(message string) {
	r.StatusMessage = message
}

func (r *Result) AddMessage(message string) {
	if r.Messages == nil {
		r.Messages = make([]string, 5)
	}

	r.Messages = append(r.Messages, message)
}

func (r *Result) AddError(message string) {
	if r.Errors == nil {
		r.Errors = make([]string, 5)
	}

	r.Errors = append(r.Errors, message)
}

func (r *Result) AddWarning(message string) {
	if r.Warnings == nil {
		r.Warnings = make([]string, 5)
	}

	r.Warnings = append(r.Warnings, message)
}

type Plugin interface {
	ValidateConfigFromMap(interface{}) error
	Run(interface{}) (error, Result)
}
