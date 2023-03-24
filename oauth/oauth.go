package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/*

package main
import (
	"fmt"
	"net/http"
	"github.com/coreos/go-oidc/v3/oidc"
//	"golang.org/x/net/context"
	"golang.org/x/oauth2" )

var (
	clientID = "a4db81022876dd2abef2"
	clientSecret = "12586c866e031c3130d49886d5c441abd878a884" )
func main() {

*/
//	ctx := context.Background()
//	provider, err := oidc.NewProvider(ctx, "https://qqcweixu.oneauth.cn/oauth/v1")
/*	if err != nil {
		fmt.Println(err)
	     return
	}
*/
/*
	config := oauth2.Config{
		ClientID: clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
			//provider.Endpoint(),
		RedirectURL: "http://127.0.0.1:5556/oneauth/callback",
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ready goto oneauth...")
		http.Redirect(w, r, config.AuthCodeURL(""), http.StatusFound) })
	http.HandleFunc("/oneauth/callback", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Oneauth done...")
		w.Write([]byte("hahahahhahha")) })
	fmt.Printf("listening on http://%s/", "0.0.0.0:5556")

	fmt.Println(http.ListenAndServe(":5556", nil)) }
*/

const clientID = "dingfv1zcj2uozqabeow"
const clientSecret = "otESK06u4P8EetuTqqKgGjU8nZYI0yHNtl1lfRFob38ePCFEEYJODAepESBWP1fj"

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// We will be using `httpClient` to make external HTTP requests later in our code
	httpClient := http.Client{}

	/*
		// Create a new redirect route route
		http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
			// First, we need to get the value of the `code` query param
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
				w.WriteHeader(http.StatusBadRequest)
			}
			code := r.FormValue("code")

			// Next, lets for the HTTP request to call the GitHub OAuth enpoint
			// to get our access token
			reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
			req, err := http.NewRequest(http.MethodPost, reqURL, nil)
			if err != nil {
				fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
				w.WriteHeader(http.StatusBadRequest)
			}
			// We set this header since we want the response
			// as JSON
			req.Header.Set("accept", "application/json")

			// Send out the HTTP request
			res, err := httpClient.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			defer res.Body.Close()

			// Parse the request body into the `OAuthAccessResponse` struct
			var t OAuthAccessResponse
			if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
				fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
				w.WriteHeader(http.StatusBadRequest)
			}

			// Finally, send a response to redirect the user to the "welcome" page
			// with the access token
			w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
			w.WriteHeader(http.StatusFound)
		})
	*/

	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		// First, we need to get the value of the `code` query param
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		r.FormValue("code")

		// Next, lets for the HTTP request to call the GitHub OAuth enpoint
		// to get our access token
		//reqURL := fmt.Sprintf("https://api.dingtalk.com/v1.0/oauth2/userAccessToken?client_id=%s&client_secret=%s&code=%sgrantType=authorization_code", clientID, clientSecret, code)
		reqURL := fmt.Sprintf("https://api.dingtalk.com/v1.0/oauth2/userAccessToken?client_id=%s&client_secret=%s", clientID, clientSecret)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		// We set this header since we want the response
		// as JSON
		req.Header.Set("accept", "application/json")

		// Send out the HTTP request
		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer res.Body.Close()

		// Parse the request body into the `OAuthAccessResponse` struct
		var t OAuthAccessResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		// Finally, send a response to redirect the user to the "welcome" page
		// with the access token
		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})
	http.ListenAndServe(":8080", nil)
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type DingOAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
