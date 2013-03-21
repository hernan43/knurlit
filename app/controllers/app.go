package controllers

import (
	"bytes"
	// "fmt"
	"github.com/robfig/revel"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Application struct {
	*revel.Controller
}

type Param struct {
	Name string
	Value string
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) KnurlIt(knurl string, method string, auth string, username string, password string, params []Param, headers []Param) revel.Result {
	c.RenderArgs["knurl"] = knurl
	c.RenderArgs["method"] = method
	c.RenderArgs["auth"] = auth
	c.RenderArgs["username"] = username
	c.RenderArgs["params"] = params
	c.RenderArgs["headers"] = headers
	
	if len(knurl) > 0 && len(method) > 0 {
		postBody := new(bytes.Buffer)		
		switch method {
			case "POST", "PUT":
				values := make(url.Values)
				for _, p := range params {
					if len(p.Name) > 0 {
						values.Set(p.Name, p.Value)				
					}
				}
				postBody = bytes.NewBufferString(values.Encode())
		}
		
		req, err := http.NewRequest(method, knurl, postBody)
		
		if auth == "yes" {
			req.SetBasicAuth(username, password)
		}
		
		for _, h := range headers {
			if len(h.Name) > 0 {
				req.Header.Add(h.Name, h.Value)
			}
		}

		resp, err := http.DefaultClient.Do(req) 
		if err != nil {
			// do something here
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		c.RenderArgs["response"] = string(body)
		
	}
	
	return c.Render()
}