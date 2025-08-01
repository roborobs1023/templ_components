package cfg

import (
	"fmt"
	"log"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type tcConfig struct {
	FileDirs struct {
		Assets     string `tc-config:"assets"`
		Css        string `tc-config:"css"`
		Javascript string `tc-config:"javascript"`
		Uploads    string `tc-config:"uploads"`
		Static     string `tc-config:"static"`
	} `tc-config:"file-dirs"`
	Theme struct {
		Colors struct {
			Primary   string `tc-config:"primary"`
			Secondary string `tc-config:"secondary"`
			Accent    string `tc-config:"accent"`
		} `tc-config:"colors"`
		FontFamily string `tc-config:"font-family"`
	} `tc-config:"theme"`
}

var Config *config.Config

var TCConfig tcConfig

func init() {
	Config = config.New("tc-config", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "tc-config"
	})

	Config.AddDriver(yaml.Driver)

	Config.Set("file-dirs.assets", "web/assets")
	Config.Set("file-dirs.css", "web/assets/css")
	Config.Set("file-dirs.javascript", "web/assets/js")
	Config.Set("file-dirs.staic", "web/assets")
	Config.Set("file-dirs.uploads", "web/uploads")

	err := Config.LoadFiles("./config.yaml")

	if err != nil {
		log.Printf("Error loading config file (expected if not exists): %v\n", err)
		log.Print("Using Default values.")

	}

	Config.MapOnExists("", &TCConfig)

	fmt.Println(TCConfig.FileDirs.Css)
}
