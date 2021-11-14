# gotor
A lightweight and simplistic Tor library for golang

```bash
go get github.com/ripmeep/gotor
```

```go
import "github.com/ripmeep/gotor"
```

# Usage
```go
t := tor.TorConnection{"127.0.0.1", 9050, 9051} // TorConnection{tor host, SOCKS5 port, Control Port}
tor_con, err := t.Connect("example.org", 80)

// Now you can use tor_con as a normal net.Dial socket!
/* do stuff.... */

t.Refresh("your tor password") 
/* 
  Replace with your control port password
  t.Refresh() will return (bool, error)
  If bool = false, the fresh failed and the error will be not be nil
  
  ok, err = t.Refresh("your tor password")
*/

// Now you have a new IP! //

tor_con, err = t.Connect("8.8.8.8", 53) // It can also work with IPs!
```
