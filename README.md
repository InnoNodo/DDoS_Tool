<img width="1120" alt="logo" src="https://github.com/user-attachments/assets/cb2a9585-4063-4421-aae8-ec517ad6e779" />

# DDoS Tool

**WARNING: THIS SOFTWARE CAN BE USED ONLY IN EDUCATIONAL PURPOSES** \

DDoS (distributed denial-of-service) tool is a soft that was writen in Go language



## Installation

* Download Go (if you don't have it):

* Clone repository
```bash
git clone https://github.com/InnoNodo/DDoS_Tool
```
* Install dependencies
```go
go mod tidy
```
* Add your proxies in file http.txt in root directory

```go
go run main.go
```

## How to add proxies?

* If you don't have proxies, you can use my ProxyParser

  ```link
  https://github.com/InnoNodo/ProxyParser
  ```

* If you have proxies with authorization, format should be:

  ```bash
  http://username:password@ip_adress:port
  ```

* If you have proxies without authorization, format should be:

  ```bash
  http://ip_adress:port
  ```
