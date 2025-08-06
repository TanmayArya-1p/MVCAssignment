package workers

import (
	"inorder/pkg/config"
	"inorder/pkg/models"
	"log"
	"time"
)

func StartCleanupWorker() (chan bool, error) {
	exitChan := make(chan bool)

	var ticker *time.Ticker = time.NewTicker(time.Duration(config.Config.InOrder.JTI_CLEANUP_INTERVAL) * time.Second)

	go cleanupWorker(exitChan, ticker)
	return exitChan, nil
}

func cleanupWorker(exitChan chan bool, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			err := models.DeleteExpiredJTIs()
			if err != nil {
				log.Fatalln("CLEANUP JOB ERROR:", err)
			}
		case <-exitChan:
			ticker.Stop()
			return
		}
	}
}
