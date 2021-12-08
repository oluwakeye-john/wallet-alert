package jobs

import "github.com/jasonlvhit/gocron"

func SetupCronjobs() {
	gocron.Every(1).Day().At("10:00").Do(UnsubscribeStaleAccounts)
}
