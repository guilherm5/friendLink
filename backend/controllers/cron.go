package controllers

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Minute().Do(func() {

		_, err := DB.Exec(`DELETE FROM story WHERE dt_exclusao <= NOW()`)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("deu certo")

	})

	s.Every(24).Day().Do(func() {

	})

	s.StartBlocking()
}
