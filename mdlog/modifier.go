package mdlog

import "context"

// Mod is a Logger middleware
type Mod struct {
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

func (m *Mod) Fatal(ctx context.Context, err error, md map[string]any) {
	m.fatal(ctx, err, md, m.logger.Fatal)
}

func (m *Mod) Error(ctx context.Context, err error, md map[string]any) {
	m.error(ctx, err, md, m.logger.Error)
}

func (m *Mod) Warn(ctx context.Context, msg string, md map[string]any) {
	m.warn(ctx, msg, md, m.logger.Warn)
}

func (m *Mod) Info(ctx context.Context, msg string, md map[string]any) {
	m.info(ctx, msg, md, m.logger.Info)
}

func (m *Mod) Debug(ctx context.Context, msg string, md map[string]any) {
	m.debug(ctx, msg, md, m.logger.Debug)
}

// WithMods adds entry middleware
func WithMods(l Logger, em ErrMod, mm MsgMod) Logger {
	if em == nil {
		em = func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		}
	}

	if mm == nil {
		mm = func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		}
	}

	return &Mod{
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
	return &Mod{
		logger: l,
		fatal:  f,
		error:  f,
		warn: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		info: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		debug: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
	}
}

// WithMsgMod adds warn+info+debug entry middleware
func WithMsgMod(l Logger, f MsgMod) Logger {
	return &Mod{
		logger: l,
		fatal: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		error: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		warn:  f,
		info:  f,
		debug: f,
	}
}

// WithFatalMod adds fatal entry middleware
func WithFatalMod(l Logger, f ErrMod) Logger {
	return &Mod{
		logger: l,
		fatal:  f,
		error: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		warn: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		info: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		debug: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
	}
}

// WithErrorMod adds error entry middleware
func WithErrorMod(l Logger, f ErrMod) Logger {
	return &Mod{
		logger: l,
		fatal: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		error: f,
		warn: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		info: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		debug: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
	}
}

// WithWarnMod adds warn entry middleware
func WithWarnMod(l Logger, f MsgMod) Logger {
	return &Mod{
		logger: l,
		fatal: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		error: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		warn: f,
		info: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		debug: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
	}
}

// WithInfoMod adds info entry middleware
func WithInfoMod(l Logger, f MsgMod) Logger {
	return &Mod{
		logger: l,
		fatal: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		error: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		warn: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		info: f,
		debug: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
	}
}

// WithDebugMod adds debug entry middleware
func WithDebugMod(l Logger, f MsgMod) Logger {
	return &Mod{
		logger: l,
		fatal: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		error: func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
			f(ctx, err, md)
		},
		warn: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		info: func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
			f(ctx, msg, md)
		},
		debug: f,
	}
}
