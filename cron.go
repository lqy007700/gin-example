package main

import (
	"github.com/robfig/cron"
	"go-gin-example/models"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 3)
	log.Println("1")

	for {
		select {
		case <-t1.C:
			log.Println("2")
			t1.Reset(time.Second * 10)
		}
	}
}
