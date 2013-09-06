package controllers

import (
	"bytes"
	"fmt"
	"github.com/robfig/revel"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Application struct {
	*revel.Controller
}

type Param struct {
	Name  string
	Value string
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) KnurlIt(knurl string, method string, auth string, username string, password string, params []Param, postBodyString string, headers []Param) revel.Result {
	c.RenderArgs["knurl"] = knurl
	c.RenderArgs["method"] = method
	c.RenderArgs["auth"] = auth
	c.RenderArgs["username"] = username
	c.RenderArgs["params"] = params
	c.RenderArgs["postBodyString"] = postBodyString
	c.RenderArgs["headers"] = headers

	if len(knurl) > 0 && len(method) > 0 {
		postBody := new(bytes.Buffer)
		switch method {
		case "POST", "PUT":
			if len(postBodyString) > 0 {
				postBody = bytes.NewBufferString(postBodyString)
			} else {
				values := make(url.Values)
				for _, p := range params {
					if len(p.Name) > 0 {
						values.Set(p.Name, p.Value)
					}
				}
				postBody = bytes.NewBufferString(values.Encode())
			}
		}

		req, err := http.NewRequest(method, knurl, postBody)
		switch method {
		case "POST", "PUT":
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}

		if auth == "yes" {
			req.SetBasicAuth(username, password)
		}

		for _, h := range headers {
			if len(h.Name) > 0 {
				req.Header.Add(h.Name, h.Value)
			}
		}

		var requestBuffer bytes.Buffer
		requestBuffer.WriteString(fmt.Sprintf("%s %s %s\n", req.Method, req.URL.Path, req.Proto))
		requestBuffer.WriteString(fmt.Sprintf("Host: %s\n", req.Host))
		requestBuffer.WriteString(fmt.Sprintf("Content-length: %d\n", req.ContentLength))
		for k, v := range req.Header {
			requestBuffer.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ",")))
		}
		requestBuffer.WriteString(fmt.Sprintf("\n%s\n", postBody))
		c.RenderArgs["request"] = requestBuffer.String()

		begin := time.Now()
		resp, err := http.DefaultClient.Do(req)
		elapsed := time.Since(begin)
		if err != nil {
			// do something here
			return c.Render()
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		var responseBuffer bytes.Buffer
		responseBuffer.WriteString(fmt.Sprintf("%s %s\n", resp.Proto, resp.Status))
		for k, v := range resp.Header {
			responseBuffer.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ",")))
		}
		responseBuffer.WriteString(string(body))
		c.RenderArgs["response"] = responseBuffer.String()
		c.RenderArgs["responseTime"] = elapsed.String()

	}

	return c.Render()
}
