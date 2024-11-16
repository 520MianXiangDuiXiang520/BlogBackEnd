package initialization

import (
	"JuneBlog/patch/logger"
	"JuneBlog/patch/logger/logfile"
	"log/slog"
	"path"
	"time"
)

func InitLogger(logPath string) error {
	stdOut, err := logfile.NewLogFile(path.Join(logPath, "blog_std.log"),
		logfile.WithSplit(256*logfile.KB))
	if err != nil {
		return err
	}
	errOut, err := logfile.NewLogFile(path.Join(logPath, "blog_err.log"),
		logfile.WithSplit(256*logfile.KB))
	if err != nil {
		return err
	}

	logger.SetDefault(logger.OutputCtrl{
		DebugOut: stdOut, InfoOut: stdOut,
		WarnOut: stdOut, ErrorOut: errOut,
	},
		logger.WithLevel(slog.LevelDebug),
		logger.WithSource(),
		logger.WithTime(),
		logger.WithTimeFormat(time.DateTime))
	return nil
}
