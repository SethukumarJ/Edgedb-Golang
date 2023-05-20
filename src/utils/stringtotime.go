package utils

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

const layout string = "2006-01-02"

func StringToTime(str string) (time.Time, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ResponseJSON(c gin.Context, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(data)
}
