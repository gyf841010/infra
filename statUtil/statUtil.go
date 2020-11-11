package statUtil

import (
	"github.com/gyf841010/infra/redisUtil"
	. "github.com/gyf841010/infra/logging"
)

func PushStatData(jsonStr string) {
	if err := redisUtil.LpushString("statistics", jsonStr); err != nil {
		Log.Error("Failed to Send Stat Data while error occurs", WithError(err))
	}
}
