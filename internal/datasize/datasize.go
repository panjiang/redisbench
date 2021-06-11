package datasize

import "fmt"

type ByteSize uint64

func (b ByteSize) String() string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%d%c",
		int64(float64(b)/float64(div)), "KMGTPE"[exp])
}
