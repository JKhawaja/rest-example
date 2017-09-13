package logger

// Logger ...
type Logger interface {
	LogWithContext(interface{}, error) error
}
