package lang

import (
	"app/internal/help"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type messageMap map[string]string

var langData = map[string]messageMap{}

func loadMessages(lang string) (messageMap, error) {

	l, ok := langData[lang]
	if ok {
		return l, nil
	}

	data, err := os.ReadFile(help.Dir.Root + fmt.Sprintf("lang/%s.json", lang))
	if err != nil {
		log.Println("Failed to read lang file:", err)
		return nil, err
	}

	var messages messageMap
	err = json.Unmarshal(data, &messages)
	if err != nil {
		log.Println("Failed to parse lang data:", err)
		return nil, err
	}
	langData[lang] = messages
	return messages, nil
}

func Text(text, lang string) string {

	if lang == "" {
		lang = "en"
	}
	messages, err := loadMessages(lang)
	if err != nil {
		return text
	}
	msg, ok := messages[text]
	if !ok {
		return text
	}
	return msg
}
