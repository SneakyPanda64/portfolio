package actions

import (
	"context"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/sneakypanda64/portfolio/db"
)

func GetClicks() (int, error) {
	clicks, err := db.Redis_client.Get(context.Background(), "clicks").Result()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	logrus.Print("GOT CLICKS")
	i, err := strconv.Atoi(clicks)
	if err != nil {
		return 0, err
	}
	return i, nil
	// return i, nil
}
