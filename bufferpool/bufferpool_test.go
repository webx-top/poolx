package bufferpool_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webx-top/poolx/bufferpool"
)

func TestUsage(t *testing.T) {
	pool := bufferpool.New()

	buf := pool.Get()
	for i := 0; i < 10; i++ {
		if !assert.Equal(t, &bytes.Buffer{}, buf, `should be an empty buffer`) {
			return
		}
		pool.Release(buf)
	}
}
