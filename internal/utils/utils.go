package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jfgrea27/loco-paster/internal/models"
)

func BuildEndpoint() string {
	var port int
	port_str := os.Getenv("LOCO_PASTER_API_PORT")
	if len(port_str) == 0 {
		port = 8000
	} else {
		p, err := strconv.Atoi(port_str)

		if err != nil {
			port = 8000
		} else {
			port = p
		}
	}
	return fmt.Sprintf("0.0.0.0:%v", port)
}

func FindPasteObjIndex(id int, pobjs []models.PasteObj) int {
	for i, o := range pobjs {
		if o.Id == id {
			return i
		}
	}
	return -1
}
