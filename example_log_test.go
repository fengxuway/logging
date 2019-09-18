package logging_test

import (
	"fmt"
	"time"

	// "github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"

	"git.intra.weibo.com/adx/logging"
)

func Example_DefaultLogger() {
	logger := logging.DefaultLogger()
	logger.SetLevel(log.DEBUG)
	logger.Debug("debug")
	logger.Info("Info")
	logger.Warn("warn")
	logger.Error("error")
}

func Example_CustomLogging() {
	logger := logging.NewLoggerWithConfig(&logging.LogConfig{
		Level: "debug",
		File: &logging.LogFile{
			FileName:           "adx.log",
			LogPath:            "/tmp/test-adx-logging",
			FileNameDateFormat: "20060102.150405",
			FileNameDateAlign:  true,
			RotationDuration:   logging.Duration{5 * time.Second},
			RotationCount:      3,
		},
	})

	for i := 0; i < 20; i++ {
		logger.Info("message...")
		time.Sleep(500 * time.Millisecond) // 0.5s
	}
}

// func Example_CustomLoggingFromToml() {
//     yamlConfigStr := `
// title = "demo"

// [logging]
// level = "info"
//   [file]
//   filename = "demolog"
//   path = "/tmp/demologging"
//   file_name_date_format = "20060102.150405"
//   file_name_date_align = true
//   rotation_duration = "5s" # 10m, 24h, 1d, ...
//   rotation_count = 3
//     `

//     type Config struct {
//         logging.LogConfig `toml:"logging" json:"logging"`
//     }

//     var config Config
//     if _, err := toml.Decode(yamlConfigStr, &config); err != nil {
//         fmt.Printf("load config err: %s\n", err.Error())
//         return
//     }
//     fmt.Printf("config : %#v\n", config)

//     logger := logging.NewLoggerWithConfig(&config.LogConfig)

//     for i := 0; i < 20; i++ {
//         logger.Info("message...")
//         time.Sleep(500 * time.Millisecond) // 0.5s
//     }
// }
