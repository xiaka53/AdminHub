package redis

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
	"sync"
)

type BaseConfig struct {
	ProxyList    []string `toml:"proxy_list"`
	MaxIdle      int      `toml:"max_idle"`
	MaxActive    int      `toml:"max_active"`
	ConnTimeout  int      `toml:"conn_timeout"`
	ReadTimeout  int      `toml:"read_timeout"`
	WriteTimeout int      `toml:"write_timeout"`
}

type ListConfig struct {
	Base BaseConfig `toml:"base"`
}

type Config struct {
	List ListConfig `toml:"list"`
}

func (rc *RedisConf) writeToml(w *sync.WaitGroup) {
	defer w.Done()
	config := Config{
		List: ListConfig{
			Base: BaseConfig{
				ProxyList:    []string{fmt.Sprintf("%s:%d", rc.Url, rc.Port), rc.Pass},
				MaxIdle:      2000,
				MaxActive:    2000,
				ConnTimeout:  50,
				ReadTimeout:  100,
				WriteTimeout: 100,
			},
		},
	}
	if err := os.MkdirAll(filepath.Dir("conf/local/redis_config.toml"), os.ModePerm); err != nil {
		fmt.Println("创建配置文件夹失败,err:", err)
		os.Exit(0)
	}
	file, err := os.Create("conf/local/redis_config.toml")
	if err != nil {
		fmt.Println("redis配置文件创建失败,err:", err)
		os.Exit(0)
	}
	defer file.Close()
	if _, err = file.WriteString("# this is redis config file\n"); err != nil {
		fmt.Println("redis写入注释失败")
		os.Exit(0)
	}
	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		fmt.Println("redis配置文件创建失败,err:", err)
		os.Exit(0)
	}

	fmt.Println("成功创建和写入 conf/local/redis_config.toml 文件")
}
