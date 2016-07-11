package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

var graphdef = map[string](mp.Graphs){
	"http.status": mp.Graphs{
		Label: "HTTP Status Codes",
		Unit:  "integer",
		Metrics: [](mp.Metrics){
			mp.Metrics{Name: "100", Label: "100 Continue", Diff: true, Stacked: true},
			mp.Metrics{Name: "101", Label: "101 Switching Protocols", Diff: true, Stacked: true},
			mp.Metrics{Name: "102", Label: "102 Processing", Diff: true, Stacked: true},
			mp.Metrics{Name: "200", Label: "200 OK", Diff: true, Stacked: true},
			mp.Metrics{Name: "201", Label: "201 Created", Diff: true, Stacked: true},
			mp.Metrics{Name: "202", Label: "202 Accepted", Diff: true, Stacked: true},
			mp.Metrics{Name: "203", Label: "203 Non-Authoritative Information", Diff: true, Stacked: true},
			mp.Metrics{Name: "204", Label: "204 No Content", Diff: true, Stacked: true},
			mp.Metrics{Name: "205", Label: "205 Reset Content", Diff: true, Stacked: true},
			mp.Metrics{Name: "206", Label: "206 Partial Content", Diff: true, Stacked: true},
			mp.Metrics{Name: "207", Label: "207 Multi-Status", Diff: true, Stacked: true},
			mp.Metrics{Name: "208", Label: "208 Already Reported", Diff: true, Stacked: true},
			mp.Metrics{Name: "226", Label: "226 IM Used", Diff: true, Stacked: true},
			mp.Metrics{Name: "300", Label: "300 Multiple Choices", Diff: true, Stacked: true},
			mp.Metrics{Name: "301", Label: "301 Moved Permanently", Diff: true, Stacked: true},
			mp.Metrics{Name: "302", Label: "302 Found", Diff: true, Stacked: true},
			mp.Metrics{Name: "303", Label: "303 See Other", Diff: true, Stacked: true},
			mp.Metrics{Name: "304", Label: "304 Not Modified", Diff: true, Stacked: true},
			mp.Metrics{Name: "305", Label: "305 Use Proxy", Diff: true, Stacked: true},
			mp.Metrics{Name: "306", Label: "306 Switch Proxy", Diff: true, Stacked: true},
			mp.Metrics{Name: "307", Label: "307 Temporary Redirect", Diff: true, Stacked: true},
			mp.Metrics{Name: "308", Label: "308 Permanent Redirect", Diff: true, Stacked: true},
			mp.Metrics{Name: "400", Label: "403 Bad Request", Diff: true, Stacked: true},
			mp.Metrics{Name: "401", Label: "401 Unauthorized", Diff: true, Stacked: true},
			mp.Metrics{Name: "402", Label: "402 Payment Required", Diff: true, Stacked: true},
			mp.Metrics{Name: "403", Label: "403 Forbidden", Diff: true, Stacked: true},
			mp.Metrics{Name: "404", Label: "404 Not Found", Diff: true, Stacked: true},
			mp.Metrics{Name: "405", Label: "405 Method Not Allowed", Diff: true, Stacked: true},
			mp.Metrics{Name: "406", Label: "406 Not Acceptable", Diff: true, Stacked: true},
			mp.Metrics{Name: "407", Label: "407 Proxy Authentication Required", Diff: true, Stacked: true},
			mp.Metrics{Name: "408", Label: "408 Request Timeout", Diff: true, Stacked: true},
			mp.Metrics{Name: "409", Label: "409 Conflict", Diff: true, Stacked: true},
			mp.Metrics{Name: "410", Label: "410 Gone", Diff: true, Stacked: true},
			mp.Metrics{Name: "411", Label: "411 Length Required", Diff: true, Stacked: true},
			mp.Metrics{Name: "412", Label: "412 Precondition Failed", Diff: true, Stacked: true},
			mp.Metrics{Name: "413", Label: "413 Payload Too Large", Diff: true, Stacked: true},
			mp.Metrics{Name: "414", Label: "414 URI Too Long", Diff: true, Stacked: true},
			mp.Metrics{Name: "415", Label: "415 Unsupported Media Type", Diff: true, Stacked: true},
			mp.Metrics{Name: "416", Label: "416 Range Not Satisfiable", Diff: true, Stacked: true},
			mp.Metrics{Name: "417", Label: "417 Expectation Failed", Diff: true, Stacked: true},
			mp.Metrics{Name: "418", Label: "418 I'm a teapot", Diff: true, Stacked: true},
			mp.Metrics{Name: "500", Label: "500 Internal Server Error", Diff: true, Stacked: true},
			mp.Metrics{Name: "501", Label: "501 Not Implemented", Diff: true, Stacked: true},
			mp.Metrics{Name: "502", Label: "502 Bad Gateway", Diff: true, Stacked: true},
			mp.Metrics{Name: "503", Label: "503 Service Unavailable", Diff: true, Stacked: true},
			mp.Metrics{Name: "504", Label: "504 Gateway Timeout", Diff: true, Stacked: true},
			mp.Metrics{Name: "505", Label: "505 HTTP Version Not Supported", Diff: true, Stacked: true},
			mp.Metrics{Name: "506", Label: "506 Variant Also Negotiates", Diff: true, Stacked: true},
			mp.Metrics{Name: "507", Label: "507 Insufficient Storage", Diff: true, Stacked: true},
			mp.Metrics{Name: "508", Label: "508 Loop Detected", Diff: true, Stacked: true},
			mp.Metrics{Name: "510", Label: "510 Not Extended", Diff: true, Stacked: true},
			mp.Metrics{Name: "511", Label: "511 Network Authentication Required", Diff: true, Stacked: true},
		},
	},
}

// HttpStatusCounterPlugin
type HttpStatusCounterPlugin struct {
	URI string
}

// FetchMetrics interface for mackerelplugin
func (p HttpStatusCounterPlugin) FetchMetrics() (map[string]interface{}, error) {
	resp, err := http.Get(p.URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	status_codes := strings.Split(strings.TrimSpace(string(body)), "\t")

	stat := make(map[string]interface{})
	for _, status_code := range status_codes {
		s := strings.Split(status_code, ":")
		code, count := s[0], s[1]
		stat[code] = count
	}

	return stat, nil
}

// GraphDefinition interface for mackerelplugin
func (p HttpStatusCounterPlugin) GraphDefinition() map[string](mp.Graphs) {
	return graphdef
}

func main() {
	optScheme := flag.String("scheme", "http", "Scheme")
	optHost := flag.String("host", "localhost", "Host")
	optPort := flag.String("port", "80", "Port")
	optPath := flag.String("path", "/status_count", "Path")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	var httpStatusCounter HttpStatusCounterPlugin
	httpStatusCounter.URI = fmt.Sprintf("%s://%s:%s%s", *optScheme, *optHost, *optPort, *optPath)

	helper := mp.NewMackerelPlugin(httpStatusCounter)
	if *optTempfile != "" {
		helper.Tempfile = *optTempfile
	} else {
		helper.Tempfile = fmt.Sprintf("/tmp/mackerel-plugin-http-status-counter")
	}
	helper.Run()
}
