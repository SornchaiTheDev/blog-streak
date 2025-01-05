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

func setup() {
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
	if err != nil {
		log.Fatalf("Cannot parse the data in %s to Streak struct", fileName)
	}

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Cannot write to %s file", fileName)
	}
}

func (s *streakService) Get() (int, int) {
	file, err := os.Open(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		setup()
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

	days := int(hoursDiff/24) + 1

	return days, streak.Longest

}

func (s *streakService) Update() {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if errors.Is(err, fs.ErrNotExist) {
		setup()
		file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
		if err != nil {
			log.Fatalf("Cannot open %s file", fileName)
		}
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

	if time.Now().Sub(streak.LatestDate).Hours() > 24 {
		diffHours := streak.LatestDate.Sub(streak.StartedDate).Hours()
		diffDays := int(diffHours/24) + 1
		streak.Longest = diffDays

		streak.StartedDate = time.Now()
		streak.LatestDate = time.Now()
	} else {
		streak.LatestDate = time.Now()
		streak.Longest++
	}

	data, err := json.MarshalIndent(streak, "", "  ")
	if err != nil {
		log.Fatalf("Cannot stringify the Streak struct in %s", fileName)
	}

	_, err = file.WriteAt(data, 0)
	if err != nil {
		log.Fatalf("Cannot write to %s file", fileName)
	}
}
