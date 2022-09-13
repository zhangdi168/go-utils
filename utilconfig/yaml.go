// Package utilconfig
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utilconfig

import (
	"fmt"
	"github.com/zhangdi168/go-utils/utilerror"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Env struct {
	Wechat Wechat `json:"wechat" yaml:"wechat"`
	Data   Data   `json:"data" yaml:"data"`
}
type Wechat struct {
	Appid      string `yaml:"appid" json:"appid,omitempty"`
	Secret     string `yaml:"secret" json:"secret,omitempty"`
	Userid     string `yaml:"userid" json:"userid,omitempty"`
	Templateid string `yaml:"template_id" json:"templateid,omitempty"`
	City       string `yaml:"city" json:"city,omitempty"`
	TaskTime   string `yaml:"task_time" json:"task_time,omitempty"` //定时任务时间 格式20:03
}

type Data struct {
	First    string `yaml:"first" json:"first,omitempty"`
	Remark   string `yaml:"remark" json:"remark,omitempty"`
	Keyword1 string `yaml:"keyword1" json:"keyword1,omitempty"`
	Keyword2 string `yaml:"keyword2" json:"keyword2,omitempty"`
	Keyword3 string `yaml:"keyword3" json:"keyword3,omitempty"`
	Keyword4 string `yaml:"keyword4" json:"keyword4,omitempty"`
	Keyword5 string `yaml:"keyword5" json:"keyword5,omitempty"`
	Keyword6 string `yaml:"keyword6" json:"keyword6,omitempty"`
	Keyword7 string `yaml:"keyword7" json:"keyword7,omitempty"`
	Keyword8 string `yaml:"keyword8" json:"keyword8,omitempty"`
}

var Conf *Env

const YamlPath_ string = "config.yaml"

func InitConfigs() {
	defer utilerror.GetErrorRecover()
	Conf = YamlRead(YamlPath_)
}

// YamlRead
// @Description: 读取Yaml文件到数据结构类型，需要读取完成后需要收到转换类型
// @param yamlPath
// @return env
func YamlRead(yamlPath string) (env *Env) {
	fileYamlBytes, errFile := os.ReadFile(yamlPath)
	if errFile != nil {
		fmt.Println("文件读取失败", errFile)
		return
	}
	if err := yaml.Unmarshal(fileYamlBytes, &env); err != nil {
		fmt.Println(err)
		return nil
	}

	return env
}

// YamlWrite YamlRead
// @Description: 将数据写入到YAMl文件
// @param yamlPath
// @return env
func YamlWrite(yamlPath string, newEnv *Env) {
	data, err := yaml.Marshal(newEnv)
	checkError(err)
	err = ioutil.WriteFile(yamlPath, data, 0777)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
