package actions

import (
	"context"

	"github.com/sneakypanda64/portfolio/db"
)

func Click(country, ip string) (int, error) {
	// const country = "GB"
	// const ip = "127-0-0-1"
	clicks, err := db.Redis_client.Incr(context.Background(), "clicks").Result()
	if err != nil {
		return 0, err
	}
	_, err = db.Redis_client.HIncrBy(context.Background(), "countries", country, 1).Result()
	if err != nil {
		return 0, err
	}
	return int(clicks), nil
}
