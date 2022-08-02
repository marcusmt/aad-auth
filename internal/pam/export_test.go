package pam

const (
	DebugWelcome = debugWelcome
)

// WithPamLoggerFunc will call the given func instead of pam_syslog for tests.
func WithPamLoggerFunc(f func(pamh Handle, priority int, format string, a ...any)) Option {
	return func(o *options) {
		o.logWithPam = f
	}
}
