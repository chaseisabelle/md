package mdlog

import "context"

// Modder is a Logger middleware
type Modder struct {
	logger Logger
	fatal  ErrMod
	error  ErrMod
	warn   MsgMod
	info   MsgMod
	debug  MsgMod
}

// ErrMod is a func to modify an error entry
type ErrMod func(context.Context, error, map[string]any, ErrFunc)

// MsgMod is a func to modify a message entry
type MsgMod func(context.Context, string, map[string]any, MsgFunc)

// Fatal mod a fatal error entry
func (m *Modder) Fatal(ctx context.Context, err error, md map[string]any) {
	m.fatal(ctx, err, md, m.logger.Fatal)
}

// Error mod an error entry
func (m *Modder) Error(ctx context.Context, err error, md map[string]any) {
	m.error(ctx, err, md, m.logger.Error)
}

// Warn mod a warning entry
func (m *Modder) Warn(ctx context.Context, msg string, md map[string]any) {
	m.warn(ctx, msg, md, m.logger.Warn)
}

// Info mod info entries
func (m *Modder) Info(ctx context.Context, msg string, md map[string]any) {
	m.info(ctx, msg, md, m.logger.Info)
}

// Debug mod debug entry
func (m *Modder) Debug(ctx context.Context, msg string, md map[string]any) {
	m.debug(ctx, msg, md, m.logger.Debug)
}

// WithMods adds entry middleware
func WithMods(l Logger, em ErrMod, mm MsgMod) Logger {
	if em == nil && mm == nil {
		return l
	}

	if em == nil {
		em = NopErrMod
	}

	if mm == nil {
		mm = NopMsgMod
	}

	return &Modder{
		logger: l,
		fatal:  em,
		error:  em,
		warn:   mm,
		info:   mm,
		debug:  mm,
	}
}

// WithErrMod adds fatal+error entry middleware
func WithErrMod(l Logger, f ErrMod) Logger {
	return &Modder{
		logger: l,
		fatal:  f,
		error:  f,
		warn:   NopMsgMod,
		info:   NopMsgMod,
		debug:  NopMsgMod,
	}
}

// WithMsgMod adds warn+info+debug entry middleware
func WithMsgMod(l Logger, f MsgMod) Logger {
	return &Modder{
		logger: l,
		fatal:  NopErrMod,
		error:  NopErrMod,
		warn:   f,
		info:   f,
		debug:  f,
	}
}

// WithFatalMod adds fatal entry middleware
func WithFatalMod(l Logger, f ErrMod) Logger {
	return &Modder{
		logger: l,
		fatal:  f,
		error:  NopErrMod,
		warn:   NopMsgMod,
		info:   NopMsgMod,
		debug:  NopMsgMod,
	}
}

// WithErrorMod adds error entry middleware
func WithErrorMod(l Logger, f ErrMod) Logger {
	return &Modder{
		logger: l,
		fatal:  NopErrMod,
		error:  f,
		warn:   NopMsgMod,
		info:   NopMsgMod,
		debug:  NopMsgMod,
	}
}

// WithWarnMod adds warn entry middleware
func WithWarnMod(l Logger, f MsgMod) Logger {
	return &Modder{
		logger: l,
		fatal:  NopErrMod,
		error:  NopErrMod,
		warn:   f,
		info:   NopMsgMod,
		debug:  NopMsgMod,
	}
}

// WithInfoMod adds info entry middleware
func WithInfoMod(l Logger, f MsgMod) Logger {
	return &Modder{
		logger: l,
		fatal:  NopErrMod,
		error:  NopErrMod,
		warn:   NopMsgMod,
		info:   f,
		debug:  NopMsgMod,
	}
}

// WithDebugMod adds debug entry middleware
func WithDebugMod(l Logger, f MsgMod) Logger {
	return &Modder{
		logger: l,
		fatal:  NopErrMod,
		error:  NopErrMod,
		warn:   NopMsgMod,
		info:   NopMsgMod,
		debug:  f,
	}
}

// NopErrMod is a no-op error mod func
func NopErrMod(ctx context.Context, err error, md map[string]any, f ErrFunc) {
	f(ctx, err, md)
}

// NopMsgMod is a no-op message mod func
func NopMsgMod(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
	f(ctx, msg, md)
}
