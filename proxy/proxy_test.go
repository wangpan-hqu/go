package proxy

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func test1() {
	fmt.Println("Serve on :8080")
	http.Handle("/", &Pxy{})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func test2() {
	proxy := NewMultipleHostsReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:9091",
		},
		{
			Scheme: "http",
			Host:   "localhost:9092",
		},
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}

func test3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		director := func(req *http.Request) {
			req = r
			req.URL.Scheme = "http"
			req.URL.Host = r.Host
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func test4() {

	InitHttpClient()
	//http.Handle("/", )
	/*req, err := http.NewRequest("get", "http://localhost:9090/", nil)
	if err !=nil{
		os.Exit(0)
	}
	resp, err :=proxy.ProxyHttpClient.Do(req)
	if err !=nil{
		os.Exit(0)
	}
	defer resp.Body.Close()
	*/

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("sucess"))
	})
	http.ListenAndServe("localhost:9090", nil)
}
