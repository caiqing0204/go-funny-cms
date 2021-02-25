package help

import (
	//"bytes"
	//"math/big"
	//"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cast"
	"os"
	"time"
)

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}

// 获取随机字符串
func createRandomString(len int) string {
	//var container string
	//var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	//b := bytes.NewBufferString(str)
	//
	//length := b.Len()
	//bigInt := big.NewInt(int64(length))
	//
	//for i := 0; i < len; i++ {
	//	randomInt, _ := rand.Int(rand.Reader, bigInt)
	//}
	return ""
}

// 生成一个md5
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

// 获取偏移量
func GetOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

// 获取当前时间
func getCurrentTimestamp() string {
	return time.Now().Format("2006-01-02")
}

// 获取默认参数值
func GetDefaultParam(params ...interface{}) interface{} {
	if len(params) > 0 {
		return params[0]
	}
	return nil
}

// 递归
func TreeRecursion(data []interface{}, pid ...int) []interface{} {
	p_id := 0
	if len(pid) > 0 {
		p_id = pid[0]
	}

	var tree []interface{}
	for _, item := range data {
		column := item.(map[string]interface{})
		currentPid := cast.ToInt(column["pid"])
		if cast.ToInt(column["pid"]) == p_id {
			column["children"] = TreeRecursion(data, currentPid)
			tree = append(tree, column)
		}
	}

	return tree
}
