package iqshell

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"path/filepath"
)

const (
	STATUS_OK = iota
	//process error
	STATUS_ERROR
	//local error
	STATUS_HALT
)

const (
	BLOCK_BITS = 22 // Indicate that the blocksize is 4M
	BLOCK_SIZE = 1 << BLOCK_BITS
)

var (
	PATH_ROOT  = []string{"path.root", "path.root_path"}
	PATH_ACCDB = []string{"path.accdb", "path.acc_db_path"}
	PATH_ACC   = []string{"path.acc", "path.acc_path"}
	HOST_UP    = []string{"hosts.up", "hosts.up_host"}
	HOST_RS    = []string{"hosts.rs", "hosts.rs_host"}
	HOST_RSF   = []string{"hosts.rsf", "hosts.rsf_host"}
	HOST_IO    = []string{"hosts.io", "hosts.io_host"}
	HOST_API   = []string{"hosts.api", "hosts.api_host"}
	ACCESS_KEY = []string{"access_key"}
	SECRET_KEY = []string{"secret_key"}
)

func UpHostBindPFlag(val *pflag.Flag) {
	for _, key := range HOST_UP {
		viper.BindPFlag(key, val)
	}
}

func AccessKey() string {
	return viper.GetString(ACCESS_KEY[0])
}

func SecretKey() string {
	return viper.GetString(SECRET_KEY[0])
}

func RootPath() string {
	path := viper.GetString(PATH_ROOT[0])
	if path != "" {
		return path
	}
	return viper.GetString(PATH_ROOT[1])
}

func SetRootPath(val string) {
	for _, key := range PATH_ROOT {
		viper.Set(key, val)
	}
}

func AccDBPath() string {
	path := viper.GetString(PATH_ACCDB[0])
	if path != "" {
		return path
	}
	return viper.GetString(PATH_ACCDB[1])
}

func SetAccDBPath(val string) {
	for _, key := range PATH_ACCDB {
		viper.Set(key, val)
	}
}

func SetDefaultAccDBPath(val string) {
	for _, key := range PATH_ACCDB {
		viper.SetDefault(key, val)
	}
}

func AccPath() string {
	path := viper.GetString(PATH_ACC[0])
	if path != "" {
		return path
	}
	return viper.GetString(PATH_ACC[1])
}

func SetAccPath(val string) {
	for _, key := range PATH_ACC {
		viper.Set(key, val)
	}
}

func SetDefaultAccPath(val string) {
	for _, key := range PATH_ACC {
		viper.SetDefault(key, val)
	}
}

func OldAccPath() string {
	acc_path := AccPath()
	if acc_path == "" {
		return ""
	}
	return filepath.Join(filepath.Dir(acc_path), "old_"+filepath.Base(acc_path))
}

func UpHost() string {
	host := viper.GetString(HOST_UP[0])
	if host != "" {
		return host
	}
	return viper.GetString(HOST_UP[1])
}

func SetUpHost(val string) {
	for _, key := range HOST_UP {
		viper.Set(key, val)
	}
}

func SetDefaultUpHost(val string) {
	for _, key := range HOST_UP {
		viper.SetDefault(key, val)
	}
}

func RsHost() string {
	host := viper.GetString(HOST_RS[0])
	if host != "" {
		return host
	}
	return viper.GetString(HOST_RS[1])
}

func SetRsHost(val string) {
	for _, key := range HOST_RS {
		viper.Set(key, val)
	}
}

func SetDefaultRsHost(val string) {
	for _, key := range HOST_RS {
		viper.SetDefault(key, val)
	}
}

func RsfHost() string {
	host := viper.GetString(HOST_RSF[0])
	if host != "" {
		return host
	}
	return viper.GetString(HOST_RSF[1])
}

func SetRsfHost(val string) {
	for _, key := range HOST_RSF {
		viper.Set(key, val)
	}
}

func SetDefaultRsfHost(val string) {

	for _, key := range HOST_RSF {
		viper.SetDefault(key, val)
	}
}

func IoHost() string {
	host := viper.GetString(HOST_IO[0])
	if host != "" {
		return host
	}
	return viper.GetString(HOST_IO[1])
}

func SetIoHost(val string) {
	for _, key := range HOST_IO {
		viper.Set(key, val)
	}
}

func SetDefaultIoHost(val string) {
	for _, key := range HOST_IO {
		viper.SetDefault(key, val)
	}
}

func ApiHost() string {
	host := viper.GetString(HOST_API[0])
	if host != "" {
		return host
	}
	return viper.GetString(HOST_API[1])
}

func SetApiHost(val string) {
	for _, key := range HOST_API {
		viper.Set(key, val)
	}
}

func SetDefaultApiHost(val string) {
	for _, key := range HOST_API {
		viper.SetDefault(key, val)
	}
}

type FetchItem struct {
	RemoteUrl string
	Bucket    string
	Key       string
}
