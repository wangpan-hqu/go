package proxy

import (
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

//正向代理
type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)

	transport := http.DefaultTransport

	// step 1  代理接收到客户端的请求，复制了原来的请求对象，并根据数据配置新请求的各种参数（添加上 X-Forward-For 头部等）
	outReq := new(http.Request)
	*outReq = *req // this only does shallow copies of maps

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	// step 2 把新请求发送到服务器端，并接收到服务器端返回的响应
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// step 3 代理服务器对响应做一些处理，然后返回给客户端
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}

	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

var DefaultHttpClient *http.Client
var ProxyHttpClient *http.Client

func InitHttpClient() {
	// not use proxy
	DefaultHttpClient = http.DefaultClient

	// use proxy
	ProxyHttpClient = getProxyHttpClient()
}

func isAddressOpen(address string) bool {
	timeout := time.Millisecond * 100
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// cannot connect to address, proxy is not active
		return false
	}

	if conn != nil {
		defer conn.Close()
		fmt.Printf("Socks5 proxy enabled: %s\n", address)
		return true
	}

	return false
}

func getProxyHttpClient() *http.Client {
	httpProxy := "127.0.0.1:8080"
	if httpProxy == "" {
		return &http.Client{}
	}

	if !isAddressOpen(httpProxy) {
		return &http.Client{}
	}

	// https://stackoverflow.com/questions/33585587/creating-a-go-socks5-client
	dialer, err := proxy.SOCKS5("tcp", httpProxy, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{Dial: dialer.Dial}
	return &http.Client{
		Transport: tr,
	}
}

func GetHttpClient(url string) *http.Client {
	if strings.Contains(url, "githubusercontent.com") {
		return ProxyHttpClient
	} else {
		return DefaultHttpClient
	}
}
