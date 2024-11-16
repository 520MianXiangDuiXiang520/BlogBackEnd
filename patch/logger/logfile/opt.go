package logfile

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
)

type ctx struct {
	splitSize  int
	namePrefix string
}

func newDefaultCtx() *ctx {
	return &ctx{}
}

func (c *ctx) isSplit() bool { return c.splitSize > 0 }

type Opt func(c *ctx)

func WithSplit(size int) Opt {
	return func(c *ctx) {
		if size <= 0 {
			panic("LogFile: split size mast > 0")
		}
		c.splitSize = size
	}
}

func WithPrefix(prefix string) Opt {
	return func(c *ctx) {
		c.namePrefix = prefix
	}
}
