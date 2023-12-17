package snowflake

import (
	"testing"
	"time"
)

func TestSnowflake(t *testing.T) {
	var a, b, c int64
	a = ID()
	b = ID()
	c = ID()
	if a == b || a == c || b == c {
		t.Errorf("repeated snowflake")
		return
	}
	t.Log(a, b, c, time.Now().UnixMilli())
}
