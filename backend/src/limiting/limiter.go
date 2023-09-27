package limiting

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/db"
)

const script = `
local current
current = redis.call("INCR", KEYS[1])
if tonumber(current) == 1 then
	redis.call("EXPIRE", KEYS[1], 30)
end
return current
`

func RateLimit(key string, limit int64) (bool, error) {
	logrus.Infof("checking rate limit: %s, %d", key, limit)
	v, err := db.Redis_client.Eval(context.Background(), script, []string{key}).Result()
	if err != nil {
		return false, err
	}
	n, _ := v.(int64)
	return n > limit, nil
}
