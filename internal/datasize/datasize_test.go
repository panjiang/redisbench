package datasize

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByteSize(t *testing.T) {
	require.Equal(t, "1023B", ByteSize(1023).String())
	require.Equal(t, "1K", ByteSize(1024).String())
	require.Equal(t, "1K", ByteSize(1025).String())
	require.Equal(t, "1M", ByteSize(1024*1024).String())
}
