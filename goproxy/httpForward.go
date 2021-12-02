package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func httpForward(pattern string, port int) {
	http.HandleFunc(pattern, doGo)
	strPort := strconv.Itoa(port)
	fmt.Print("listenning on :", " ", pattern, " ", strPort, "\n")
	err := http.ListenAndServe(":"+strPort, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func doGo(w http.ResponseWriter, r *http.Request) {

	//r.host不带http,r.url是完整的url
	//fmt.Println(r.Host, " ", r.URL, "\n")

	fmt.Println("url: ", r.URL)

	//查看url各个信息
	// str := "hi ,it is working"
	// b := []byte(str)
	//w.Write(b)
	//fmt.Print(r.Host, " ", r.Method, " \nr.URL.String ", r.URL.String(), " r.URL.Host ", r.URL.Host, " r.URL.Fragment ", r.URL.Fragment, " r.URL.Hostname ", r.URL.Hostname(), " r.URL.RequestURI ", r.URL.RequestURI(), " r.URL.Scheme ", r.URL.Scheme)

	cli := &http.Client{}

	//不建议用readfull，对于body大小难以判断，容易出错
	// body := make([]byte, 2048000)
	// n, err := io.ReadFull(r.Body, body)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print("io.ReadFull(r.Body, body) ", err.Error())
		//return,没有数据也是可以的，不需要直接结束
	}
	fmt.Print("req count :", len(body), "\n")

	//fmt.Print(len(body))
	//reqUrl := r.Host + r.URL.String()

	reqUrl := r.URL.String()

	//url里带了协议类型，不需要用scheme
	// if r.URL.Scheme != "" {
	//     reqUrl = r.URL.Scheme + reqUrl
	// } else {
	//     reqUrl = "http://" + reqUrl
	// }

	req, err := http.NewRequest(r.Method, reqUrl, strings.NewReader(string(body)))
	if err != nil {
		fmt.Print("http.NewRequest ", err.Error())
		return
	}

	//用遍历header实现完整复制
	//contentType := r.Header.Get("Content-Type")
	//req.Header.Set("Content-Type", contentType)

	for k, v := range r.Header {
		req.Header.Set(k, v[0])
	}
	res, err := cli.Do(req)
	if err != nil {
		fmt.Print("cli.Do(req) ", err.Error())
		return
	}
	defer res.Body.Close()

	// n, err = io.ReadFull(res.Body, body)
	// if err != nil {
	//     fmt.Print("io.ReadFull(res.Body, body) ", err.Error())
	//     return
	// }
	//fmt.Print("count body bytes: ", n, "\n")

	for k, v := range res.Header {
		w.Header().Set(k, v[0])
	}
	io.Copy(w, res.Body)

	//这样复制对大小控制较差，不建议。用copy即可
	// io.WriteString(w, string(body[:n]))
	// fmt.Print(string(body))
}

func main() {
	httpForward("/", 7889)
}
