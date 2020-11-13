package statUtil

import (
	. "pz-infra/logging"
	"pz-infra/redisUtil"
)

func PushStatData(jsonStr string) {
	if err := redisUtil.LpushString("statistics", jsonStr); err != nil {
		Log.Error("Failed to Send Stat Data while error occurs", WithError(err))
	}
}
