package filesystem

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func FileExist(filename string) bool {
	s, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return false
	}
	return !s.IsDir()
}
