package actions

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/db"
)

func Click() (int, error) {
	const country = "GB"
	const ip = "127-0-0-1"
	clicks, err := db.Redis_client.Incr(context.Background(), "clicks").Result()
	if err != nil {
		return 0, err
	}
	logrus.Print("CLICKED")
	// var newUserMap map[string]models.User
	// newUser := &models.User{
	// 	Country:     country,
	// 	LastVisited: time.Now().Unix(),
	// }
	_, err = db.Redis_client.HIncrBy(context.Background(), "countries", country, 1).Result()
	if err != nil {
		return 0, err
	}
	logrus.Print("HAHSED")
	return int(clicks), nil
}
