package services

import (
	"blogstreak/models"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"time"
)

type streakService struct{}

const fileName = "./streak.json"

func NewStreakService() *streakService {
	return &streakService{}
}

func (s *streakService) setup() {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create %s file", fileName)
	}
	defer file.Close()

	streak := models.Streak{
		StartedDate: time.Now(),
		LatestDate:  time.Now(),
	}

	data, err := json.MarshalIndent(streak, "", " ")

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Cannot write to %s file", fileName)
	}
}

func (s *streakService) Get() int {
	file, err := os.Open(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		s.setup()
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("An error occur when reading %s", fileName)
	}

	var streak models.Streak
	err = json.Unmarshal(content, &streak)
	if err != nil {
		log.Fatalf("Cannot parse the data in %s to Streak struct", fileName)
	}

	hoursDiff := streak.LatestDate.Sub(streak.StartedDate).Hours()

	days := int(hoursDiff / 24)

	return days

}
