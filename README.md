##  HTTP CHECK TOOL

Simple tool for checking http requests by response code or substring

Originally created as tool for [zabbix agent](http://www.zabbix.com) agent at Windows environment when you need to check 
some web services

### Usage

#### Check response by code

```
http-check <url> code <http-code>
```

return:  
  0 - if match ok  
  1 - if return code is different  
  
example:  
  
```
http-check http://localhost:7183/app code 200
```  

#### Check response content by substring
```
http-check <url> substring <substring>
```
  
return:  
  0 - if match ok  
  1 - substring not found  
    
example:  
  
```
http-check http://localhost:7183/app substring "42"
```  

#### Diagnostic mode  
```
http-check <url> diag <substring>
```  
  
return:  

  verbose response about what's going on  
    
example:  

```
http-check http://localhost:7183/app substring "42"
```

#### Usage with zabbix agent

edit `zabbix_agentd.conf`:

```
UserParameter=checkHealth,httptool.exe http://localhost:7183 code 200
```
  
  
### Build

Go: 1.9, not tested with older versions  

```
go build main.go
```

#### Support Windows XP

```
set GOARCH=386
set CGO_ENABLED=0
  
go build main.go
```