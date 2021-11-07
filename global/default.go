/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 10:08:32
 */
package global

import (
	"context"
	"flag"
	"goredismanager/comment"
	"goredismanager/model"
	"net"
	"path"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisServiceStorage map[string]RedisService
var HostName string
var Port string
var ConfigViper *viper.Viper

var UseClient GlobalClient

type GlobalClient struct {
	ConnectName string
	Db          int
	Client      *redis.Client
}

type RedisService struct {
	RedisService string
	Config       *redis.Options
	UseSsh       int
	SSHConfig    model.SSHConfig
	Client       *redis.Client
}

func init() {
	RedisServiceStorage = make(map[string]RedisService)

	//获取配置文件
	var configPath string
	flag.StringVar(&configPath, "c", "./config.yaml", "配置文件路径")
	flag.Parse()

	basePath := path.Base(configPath)
	fileInfo := strings.Split(basePath, ".")
	ConfigViper = viper.New()

	ConfigViper.AddConfigPath(path.Dir(configPath))
	ConfigViper.SetConfigName(fileInfo[0])
	ConfigViper.SetConfigType(fileInfo[1])
	err := ConfigViper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	HostName = ConfigViper.GetString("hostname")
	if HostName == "" {
		HostName = "127.0.0.1"
	}

	Port = ConfigViper.GetString("port")
	if Port == "" {
		Port = "8088"
	}

	connections := ConfigViper.Get("connections")

	if connections != nil {

		slice_conns := connections.([]interface{})

		for _, v := range slice_conns {
			vv := v.(map[interface{}]interface{})

			optionConfig := &redis.Options{
				Addr:     vv["host"].(string) + ":" + vv["port"].(string),
				Password: vv["password"].(string),
				DB:       0,
			}

			if vv["usessh"] == 1 {
				sshConfig := vv["sshconfig"].(map[interface{}]interface{})
				cli, err := comment.GetSSHClient(sshConfig["sshusername"].(string), sshConfig["sshpassword"].(string), sshConfig["sshhost"].(string)+":"+sshConfig["sshport"].(string))
				if nil != err {
					panic(err)
				}
				optionConfig.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
					return cli.Dial(network, addr)
				}
			}

			client := redis.NewClient(optionConfig)

			client.AddHook(comment.RedisLog{})

			_, err := client.Ping(context.Background()).Result()
			if err != nil {
				panic(vv["servicename"].(string) + "连接失败:" + err.Error())
			}

			RsSlice := RedisService{
				RedisService: vv["servicename"].(string),
				Config:       optionConfig,
				Client:       client,
			}

			RedisServiceStorage[vv["servicename"].(string)] = RsSlice

		}

		for name, conn := range RedisServiceStorage {
			//设置全局参数
			UseClient.ConnectName = name
			UseClient.Db = 0
			UseClient.Client = conn.Client
		}

	}

	// optionConfig := &redis.Options{
	// 	Addr:     "127.0.0.1:6380",
	// 	Password: "hengda",
	// 	DB:       0,
	// }

	// client := redis.NewClient(optionConfig)

	// RsSlice := RedisService{
	// 	RedisService: "docker",
	// 	Config:       optionConfig,
	// 	Client:       client,
	// }

	// RedisServiceStorage["docker"] = RsSlice

}
