package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	// "avenssi"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func GetCurrentTimestampSec() int {
	// ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	ts := int(time.Now().Unix())
	return ts
}

func SendDeleteVideoRequest(id string) {
	log.Println("id:", id)
	// addr := config.GetLbAddr() + ":10001"
	addr := "127.0.0.1:10001" // scheduler的地址
	url := "http://" + addr + "/video-delete-record/" + id
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Sending deleting video request error: %s", err)
	}
}
