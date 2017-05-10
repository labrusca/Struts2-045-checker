# Struts2-045-checker
#### 检查你的网站是否有编号为S2-045的漏洞

##### 关于S2-045漏洞的详细信息，点[这里](https://cwiki.apache.org/confluence/display/WW/S2-045)
___________________

#用法：
### 在linux下编译：
```
go build S2-045-checker.go
```

### 单个网站检测：
```
go run S2-045-checker.go http://xxx.xxx.xxx
```
例如：测试 http://www.github.com，你可以这么做：
```
go run S2-045-checker.go http://www.github.com
```

### 多个网站检测：
请把带http[s]://的一组网站存入TXT文件中，在确保换行符为“\n”而不是"\r\n"的情况下运行：
```
go run S2-045-checker.go -f urls.txt
```

如果目标网站有漏洞，输出信息：
```
[vul] http://xxx.xxx.xxx
```
如果没有：
```
[] http://xxx.xxx.xxx
```


### S2-045-checker的GUI版本

windows下执行build_with_ui.bat编译生成exe文件,点击执行，输入网站并点击launch!


# Lisence: GPL v3

