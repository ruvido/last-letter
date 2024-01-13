package letter

import (
	"log"
	// "github.com/spf13/viper"
	"bufio"
	"os"
	"github.com/BurntSushi/toml"
)

type Item struct {
	Date       string
	Collection string
	Markdown   string
}

func Schedule( filePath string ) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var items map[string]Item
	if _, err := toml.DecodeReader(bufio.NewReader(file), &items); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", items)
	for _,ii := range items {
		log.Println(ii.Date)
	}
}

