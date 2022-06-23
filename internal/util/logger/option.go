package logger

import "io"

type Option interface {
	apply(logger *Logger)
}

type option func(logger *Logger)

func (f option) apply(logger *Logger) {
	f(logger)
}

func AddAppendix(appendix Appendix) Option {
	return option(func(logger *Logger) {
		logger.appendices = append(logger.appendices, appendix)
	})
}

func Output(writer io.Writer) Option {
	return option(func(logger *Logger) {
		logger.output = writer
	})
}

func ErrorOutput(writer io.Writer) Option {
	return option(func(logger *Logger) {
		logger.errorOutput = writer
	})
}
