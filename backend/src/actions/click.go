package actions

import (
	"context"

	"github.com/sneakypanda64/portfolio/db"
)

func Click() (int, error) {
	const country = "GB"
	const ip = "127-0-0-1"
	err := db.Redis_client.Incr(context.Background(), "clicks").Err()
	if err != nil {
		return 0, err
	}
	_, err = db.Redis_client.HIncrBy(context.Background(), "countries", country, 1).Result()
	if err != nil {
		return 0, err
	}
	return int(2), nil
}
