##  HTTP CHECK TOOL

Simple tool for checking http requests by response code or substring

Originally created as tool for zabbix agent when you need to check some web services

### Usage

#### Check response by code

```http-check <url> code <http-code>```

return:
  0 - if match ok
  1 - if return code is different

#### Check response content by substring
```http-check <url> substring <substring>```
return:
  0 - if match ok
  1 - substring not found
  
example
  
### Build

Go: 1.9, not tested with older versions

```go build main.go```

#### Support Windows XP

```
set GOARCH=386
set CGO_ENABLED=0
  
go build main.go
```