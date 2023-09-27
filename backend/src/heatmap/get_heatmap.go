package heatmap

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/biter777/countries"
	"github.com/sneakypanda64/portfolio/db"
	"github.com/sneakypanda64/portfolio/models"
)

func GetStats() ([]byte, error) {
	hm, err := GetHeatmap()
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(hm)
	if err != nil {
		return nil, err
	}
	resp := &models.SendData{
		Type:  "stats",
		Value: string(body),
	}
	fbody, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return fbody, nil
}

func GetHeatmap() (map[int64]string, error) {
	result, err := db.Redis_client.HGetAll(context.Background(), "countries").Result()
	if err != nil {
		return map[int64]string{}, err
	}
	total_count := 0
	for _, count := range result {
		i, err := strconv.Atoi(count)
		if err != nil {
			i = 0
		}
		total_count = total_count + i
	}
	hexas := make(map[int64]string)
	for country, count := range result {
		i, err := strconv.Atoi(count)
		if err != nil {
			i = 0
		}
		percent := (math.Ceil((float64(i)/float64(total_count))*100) / 100)
		a := percent
		if a >= 1.0 {
			a = 1.0
		}
		al := int(math.Floor(a*100) / 10)
		al = al * 10
		colours := []string{
			"#ffffff",
			"#fffdc4",
			"#fffb80",
			"#ffe14a",
			"#ffcc4a",
			"#ffbd4a",
			"#ffa44a",
			"#ff8c4a",
			"#ff6e4a",
			"#ff564a",
			"#ff564a",
		}
		m49 := int64(countries.ByName(country))
		hexas[m49] = fmt.Sprintf(colours[al/10])
	}
	return hexas, nil
}
