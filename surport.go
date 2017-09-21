package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

//Curl ... HTTPS request
func (o *Aliyun) Curl(url string, method string, post string) (string, error) {
	client := http.Client{
		Timeout:   time.Duration(15 * time.Second), //15s超时
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("redirect")
		},
	}
	req, _ := http.NewRequest(method, url, strings.NewReader(post))
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}

//makeCommonURL ..
//build a common url and return signure randstr timestamp
func (o *Aliyun) makeCommonURL(_u string) (string, string, string) {
	com, randStr, t := o.makeCommon()
	_url := _u + com //生成对应的请求连接
	param := o.makeURLQuery(_url)
	s := o.hmacSha1(o.makeURLEncode("GET", param))
	return s, randStr, t
}

func (o *Aliyun) makeURLInternal(arr map[string]interface{}) (_url string) {
	com, _, _ := o.makeCommon()
	str := ""
	for k, v := range arr {
		if reflect.TypeOf(v).String() == "string" {
			str += k + "=" + v.(string) + "&"
		} else if reflect.TypeOf(v).String() == "[]string" {
			t, _ := json.Marshal(v.([]string))
			str += k + "=" + string(t) + "&"
		} else if reflect.TypeOf(v).String() == "int" {
			str += k + "=" + strconv.Itoa(v.(int)) + "&"
		}

	}
	_url = host + str[:len(str)-1] + com
	param := o.makeURLQuery(_url)
	s := o.hmacSha1(o.makeURLEncode("GET", param))
	_url += "&Signature=" + s
	return
}

//hmacSha1 ... signure HMAC-SHA1
func (o *Aliyun) hmacSha1(s string) string {
	key := []byte(o.Account.AccessSecret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(s))
	signure := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return o.percentEncode(signure)
}

//makeURLEncode ... StringToSign build a string which will be used to signup
func (o *Aliyun) makeURLEncode(m, str string) string {
	StringToSign := m + "&" + o.percentEncode("/") + "&" + o.percentEncode(str)
	return StringToSign
}

//percentEncode ...urlencode
func (o *Aliyun) percentEncode(str string) string {
	str1 := strings.Replace(url.QueryEscape(str), "+", "%20", -1)
	str2 := strings.Replace(str1, "*", "%2A", -1)
	str3 := strings.Replace(str2, "=", "%3D", -1)
	str4 := strings.Replace(str3, "/", "%2F", -1)
	return strings.Replace(str4, "%7E", "~", -1)
}

//build rand string to kepp the request uniqueness
func (o *Aliyun) makeRandStr(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//makeTime ...UTC Format 2006-01-02T15:04:05Z
func (o *Aliyun) makeTime() string {
	t := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	return t
}

//get the  query which is  part of url
func (o *Aliyun) makeURLQuery(_url string) string {
	query, _ := url.Parse(_url)
	param := query.Query()
	querys, arr := o.makeDictionarySort(param)
	str := ""
	for _, k := range querys {
		str += k + "=" + o.percentEncode(arr[k][0]) + "&"
	}
	return str[:len(str)-1]
}

func (o *Aliyun) makeCommon() (string, string, string) {
	t := o.makeTime()
	randStr := o.makeRandStr(20)
	com := fmt.Sprintf(CommonURLNoSign, o.Format, randStr, o.Account.AccessID, t)
	return com, randStr, t
}

//the query will be sort as DictionarySort which is  part of url
func (o *Aliyun) makeDictionarySort(arr map[string][]string) ([]string, map[string][]string) {
	keys := make([]string, len(arr))
	i := 0
	for k := range arr {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys, arr
}

//build a rand url strings
func (o *Aliyun) makeMapArgs(args []map[string]string) string {
	str := ""
	if len(args) > 0 {
		for _, v := range args {
			for kk, vv := range v {
				str += "&" + kk + "=" + vv + "&"
			}
		}
	}
	return str[:len(str)-1]
}

//Struct2Map ...
func (o *Aliyun) Struct2Map(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Tag.Get("json")] = objV.Field(i).Interface()
	}
	err = nil
	return
}

func (o *Aliyun) validator(arr map[string]interface{}) (data map[string]interface{}) {
	log.Println(11)
	data = make(map[string]interface{})
	for k, v := range arr {
		switch reflect.TypeOf(v).String() {
		case "string":
			if v != "" {
				data[k] = v
			}
		case "[]string":
			if len(v.([]string)) > 0 {
				data[k] = v
			}
		case "int":
			if v != 0 {
				data[k] = v
			}
		default:
			//
		}
	}
	return
}
