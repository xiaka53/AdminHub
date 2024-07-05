package mysql

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
	"sync"
)

type BaseConfig struct {
	DriverName      string `toml:"driver_name"`
	DataSourceName  string `toml:"data_source_name"`
	MaxOpenConn     int    `toml:"max_open_conn"`
	MaxIdleConn     int    `toml:"max_idle_conn"`
	MaxConnLifeTime int    `toml:"max_conn_life_time"`
}

type ListConfig struct {
	Base BaseConfig `toml:"base"`
}

type Config struct {
	List ListConfig `toml:"list"`
}

func (sc *MysqlConf) writeToml(w *sync.WaitGroup) {
	defer w.Done()
	config := Config{
		List: ListConfig{
			Base: BaseConfig{
				DriverName:      "mysql",
				DataSourceName:  fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Asia%2FShanghai", sc.User, sc.Pass, sc.Url, sc.Port, sc.Database),
				MaxOpenConn:     20,
				MaxIdleConn:     10,
				MaxConnLifeTime: 10,
			},
		},
	}
	file, err := os.Create("conf/local/mysql_map.toml")
	if err != nil {
		fmt.Println("mysql配置文件创建失败")
		os.Exit(0)
	}
	defer func() {
		_ = file.Close()
	}()
	if _, err = file.WriteString("# this is mysql config\n"); err != nil {
		fmt.Println("mysql配置文件创建失败")
		os.Exit(0)
	}
	encoder := toml.NewEncoder(file)
	if err = encoder.Encode(config); err != nil {
		fmt.Println("mysql配置文件创建失败")
		os.Exit(0)
	}

	fmt.Println("mysql配置成功创建和写入 conf/local/mysql_map.toml 文件")
}
