package statUtil

import (
	"pz_backend/infra/redisUtil"
	. "pz_backend/infra/logging"
)

func PushStatData(jsonStr string) {
	if err := redisUtil.LpushString("statistics", jsonStr); err != nil {
		Log.Error("Failed to Send Stat Data while error occurs", WithError(err))
	}
}
