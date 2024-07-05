package base

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type BaseConfig struct {
	DebugModel   string `toml:"debug_model"`
	TimeLocation string `toml:"time_location"`
	WebURL       string `toml:"web_url"`
	WebName      string `toml:"web_name"`
}

type HTTPConfig struct {
	Addr           string   `toml:"addr"`
	ReadTimeout    int      `toml:"read_timeout"`
	WriteTimeout   int      `toml:"write_timeout"`
	MaxHeaderBytes int      `toml:"max_header_bytes"`
	AllowIP        []string `toml:"allow_ip"`
}

type LogFileWriterConfig struct {
	On              bool   `toml:"on"`
	LogPath         string `toml:"log_path"`
	RotateLogPath   string `toml:"rotate_log_path"`
	WfLogPath       string `toma:"wf_log_path"`
	RotateWfLogPath string `toml:"rotate_wf_log_path"`
}

type LogConsoleWriterConfig struct {
	On    bool `toml:"on"`
	Color bool `toml:"color"`
}

type LogConfig struct {
	LogLevel      string                 `toml:"log_level"`
	FileWriter    LogFileWriterConfig    `toml:"file_writer"`
	ConsoleWriter LogConsoleWriterConfig `toml:"console_writer"`
}

type Config struct {
	Base BaseConfig `toml:"base"`
	HTTP HTTPConfig `toml:"http"`
	Log  LogConfig  `toml:"log"`
}

func WriteToml(webName string, port int) {
	config := Config{
		Base: BaseConfig{
			DebugModel:   "debug",
			TimeLocation: "Asia/Shanghai",
			WebURL:       "http://127.0.0.1:8552/",
			WebName:      webName,
		},
		HTTP: HTTPConfig{
			Addr:           fmt.Sprintf(":%d", port),
			ReadTimeout:    10,
			WriteTimeout:   10,
			MaxHeaderBytes: 20,
			AllowIP:        []string{"127.0.0.1", "*"},
		},
		Log: LogConfig{
			LogLevel: "trace",
			FileWriter: LogFileWriterConfig{
				On:              false,
				LogPath:         "./log/golang_common.info.log",
				RotateLogPath:   "./log/golang_common.info.log",
				WfLogPath:       "./log/golang_common.wf.log",
				RotateWfLogPath: "./log/golang_common.wf.log",
			},
			ConsoleWriter: LogConsoleWriterConfig{
				On:    true,
				Color: true,
			},
		},
	}
	file, err := os.Create("conf/local/base_config.toml")
	if err != nil {
		fmt.Println("base配置文件创建失败")
		os.Exit(0)
	}
	defer file.Close()

	if _, err = file.WriteString("# This is base config\n"); err != nil {
		fmt.Println("base配置文件写入注释失败")
		os.Exit(0)
	}

	encoder := toml.NewEncoder(file)
	if err = encoder.Encode(config); err != nil {
		fmt.Println("base配置文件创建失败")
		os.Exit(0)
	}

	fmt.Println("成功创建和写入 conf/local/base_config.toml 文件")
}
