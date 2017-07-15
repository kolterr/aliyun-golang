package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

//Curl ...
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
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return string(body), nil
}

//hmacSha1 ... signure
func (o *Aliyun) hmacSha1(s string) string {
	key := []byte(o.Account.AccessSecret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(s))
	signure := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return signure[:len(signure)-1]
}

//makeURLEncode ...
func (o *Aliyun) makeURLEncode(m, str string) string {
	StringToSign := m + "&" + o.percentEncode("/") + "&" + o.percentEncode(str)
	return StringToSign
}

//percentEncode ...
func (o *Aliyun) percentEncode(str string) string {
	str1 := strings.Replace(url.QueryEscape(str), "+", "%20", -1)
	str2 := strings.Replace(str1, "*", "%2A", -1)
	return strings.Replace(str2, "%7E", "~", -1)
}

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

//makeTime ...
func (o *Aliyun) makeTime() string {
	t := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	return t
}

func (o *Aliyun) makeURLQuery(_url string) string {
	query, _ := url.Parse(_url)
	param := query.Query()
	querys, arr := o.makeDictionarySort(param)
	str := ""
	for _, k := range querys {
		str += k + "=" + arr[k][0] + "&"
	}
	return str[:len(str)-1]
}

func (o *Aliyun) makeCommon() (string, string, string) {
	t := o.makeTime()
	randStr := o.makeRandStr(20)                                        //生成随机字符串
	com := fmt.Sprintf(CommonURLNoSign, randStr, o.Account.AccessID, t) //拿到没有签名的公共参数部分
	return com, randStr, t
}

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
