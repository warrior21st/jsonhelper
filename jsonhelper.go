package jsonhelper

import "os"
import "io/ioutil"
import "fmt"
import "encoding/json"
import "strings"
import "strconv"

func readFileBytes(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	return byteValue
}

func GetAppSetting(keys string) string {
	settingFilePath, _ := os.Getwd()
	settingFilePath += string(os.PathSeparator) + "appsettings.json"
	jsonByte := readFileBytes(settingFilePath)

	return ReadJsonValue(jsonByte, keys)
}

func ReadJsonValue(jsonByte []byte, keys string) string {
	var f interface{}
	err := json.Unmarshal(jsonByte, &f)
	if err != nil {
		fmt.Println(err)
	}
	val := ""
	keysArr := strings.Split(keys, ":")
	l := len(keysArr)
	for i := 0; i < l; i++ {
		m := f.(map[string]interface{})
		if i < l-1 {
			f = m[keysArr[i]]
		} else {
			switch t := m[keysArr[i]].(type) {
			default:
				fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
				break
			case bool:
				val = strconv.FormatBool(m[keysArr[i]].(bool))
				break
			case int:
				val = strconv.FormatInt(int64(m[keysArr[i]].(int)), 10)
				break
			case int64:
				val = strconv.FormatInt(m[keysArr[i]].(int64), 10)
				break
			case float32:
				val = strconv.FormatFloat(float64(m[keysArr[i]].(float32)), 'f', -1, 64)
				break
			case float64:
				val = strconv.FormatFloat(m[keysArr[i]].(float64), 'f', -1, 64)
				break
			case string:
				val = m[keysArr[i]].(string)
				break
			}
		}
	}

	return string(val)
}
