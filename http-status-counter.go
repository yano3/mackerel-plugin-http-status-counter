package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

var graphdef = map[string]mp.Graphs{
	"http.body_bytes_sent": {
		Label: "HTTP Bytes Sent",
		Unit:  "bytes",
		Metrics: []mp.Metrics{
			{Name: "body_bytes_sent", Label: "Body Bytes Sent", Diff: true, Type: "unit64", Stacked: false},
		},
	},
	"http.latency": {
		Label: "HTTP Latency",
		Unit:  "float",
		Metrics: []mp.Metrics{
			{Name: "avg_request_time", Label: "HTTP Average Request Time", Diff: false, Type: "float64", Stacked: false},
			{Name: "avg_upstream_response_time", Label: "HTTP Average Upstream Response Time", Diff: false, Type: "float64", Stacked: false},
		},
	},
}

var graphdefStatus = mp.Graphs{
	Label: "HTTP Status Codes",
	Unit:  "integer",
	Metrics: []mp.Metrics{
		{Name: "100", Label: "100 Continue", Diff: true, Type: "unit64", Stacked: true},
		{Name: "101", Label: "101 Switching Protocols", Diff: true, Type: "uint64", Stacked: true},
		{Name: "102", Label: "102 Processing", Diff: true, Type: "uint64", Stacked: true},
		{Name: "200", Label: "200 OK", Diff: true, Type: "uint64", Stacked: true},
		{Name: "201", Label: "201 Created", Diff: true, Type: "uint64", Stacked: true},
		{Name: "202", Label: "202 Accepted", Diff: true, Type: "uint64", Stacked: true},
		{Name: "203", Label: "203 Non-Authoritative Information", Diff: true, Type: "uint64", Stacked: true},
		{Name: "204", Label: "204 No Content", Diff: true, Type: "uint64", Stacked: true},
		{Name: "205", Label: "205 Reset Content", Diff: true, Type: "uint64", Stacked: true},
		{Name: "206", Label: "206 Partial Content", Diff: true, Type: "uint64", Stacked: true},
		{Name: "207", Label: "207 Multi-Status", Diff: true, Type: "uint64", Stacked: true},
		{Name: "208", Label: "208 Already Reported", Diff: true, Type: "uint64", Stacked: true},
		{Name: "226", Label: "226 IM Used", Diff: true, Type: "uint64", Stacked: true},
		{Name: "300", Label: "300 Multiple Choices", Diff: true, Type: "uint64", Stacked: true},
		{Name: "301", Label: "301 Moved Permanently", Diff: true, Type: "uint64", Stacked: true},
		{Name: "302", Label: "302 Found", Diff: true, Type: "uint64", Stacked: true},
		{Name: "303", Label: "303 See Other", Diff: true, Type: "uint64", Stacked: true},
		{Name: "304", Label: "304 Not Modified", Diff: true, Type: "uint64", Stacked: true},
		{Name: "305", Label: "305 Use Proxy", Diff: true, Type: "uint64", Stacked: true},
		{Name: "306", Label: "306 Switch Proxy", Diff: true, Type: "uint64", Stacked: true},
		{Name: "307", Label: "307 Temporary Redirect", Diff: true, Type: "uint64", Stacked: true},
		{Name: "308", Label: "308 Permanent Redirect", Diff: true, Type: "uint64", Stacked: true},
		{Name: "400", Label: "403 Bad Request", Diff: true, Type: "uint64", Stacked: true},
		{Name: "401", Label: "401 Unauthorized", Diff: true, Type: "uint64", Stacked: true},
		{Name: "402", Label: "402 Payment Required", Diff: true, Type: "uint64", Stacked: true},
		{Name: "403", Label: "403 Forbidden", Diff: true, Type: "uint64", Stacked: true},
		{Name: "404", Label: "404 Not Found", Diff: true, Type: "uint64", Stacked: true},
		{Name: "405", Label: "405 Method Not Allowed", Diff: true, Type: "uint64", Stacked: true},
		{Name: "406", Label: "406 Not Acceptable", Diff: true, Type: "uint64", Stacked: true},
		{Name: "407", Label: "407 Proxy Authentication Required", Diff: true, Type: "uint64", Stacked: true},
		{Name: "408", Label: "408 Request Timeout", Diff: true, Type: "uint64", Stacked: true},
		{Name: "409", Label: "409 Conflict", Diff: true, Type: "uint64", Stacked: true},
		{Name: "410", Label: "410 Gone", Diff: true, Type: "uint64", Stacked: true},
		{Name: "411", Label: "411 Length Required", Diff: true, Type: "uint64", Stacked: true},
		{Name: "412", Label: "412 Precondition Failed", Diff: true, Type: "uint64", Stacked: true},
		{Name: "413", Label: "413 Payload Too Large", Diff: true, Type: "uint64", Stacked: true},
		{Name: "414", Label: "414 URI Too Long", Diff: true, Type: "uint64", Stacked: true},
		{Name: "415", Label: "415 Unsupported Media Type", Diff: true, Type: "uint64", Stacked: true},
		{Name: "416", Label: "416 Range Not Satisfiable", Diff: true, Type: "uint64", Stacked: true},
		{Name: "417", Label: "417 Expectation Failed", Diff: true, Type: "uint64", Stacked: true},
		{Name: "418", Label: "418 I'm a teapot", Diff: true, Type: "uint64", Stacked: true},
		{Name: "500", Label: "500 Internal Server Error", Diff: true, Type: "uint64", Stacked: true},
		{Name: "501", Label: "501 Not Implemented", Diff: true, Type: "uint64", Stacked: true},
		{Name: "502", Label: "502 Bad Gateway", Diff: true, Type: "uint64", Stacked: true},
		{Name: "503", Label: "503 Service Unavailable", Diff: true, Type: "uint64", Stacked: true},
		{Name: "504", Label: "504 Gateway Timeout", Diff: true, Type: "uint64", Stacked: true},
		{Name: "505", Label: "505 HTTP Version Not Supported", Diff: true, Type: "uint64", Stacked: true},
		{Name: "506", Label: "506 Variant Also Negotiates", Diff: true, Type: "uint64", Stacked: true},
		{Name: "507", Label: "507 Insufficient Storage", Diff: true, Type: "uint64", Stacked: true},
		{Name: "508", Label: "508 Loop Detected", Diff: true, Type: "uint64", Stacked: true},
		{Name: "510", Label: "510 Not Extended", Diff: true, Type: "uint64", Stacked: true},
		{Name: "511", Label: "511 Network Authentication Required", Diff: true, Type: "uint64", Stacked: true},
	},
}

var graphdefStatusGrouping = mp.Graphs{
	Label: "HTTP Status Codes",
	Unit:  "integer",
	Metrics: []mp.Metrics{
		{Name: "http_1xx", Label: "1xx Informational", Diff: true, Type: "unit64", Stacked: true},
		{Name: "http_2xx", Label: "2xx Success", Diff: true, Type: "unit64", Stacked: true},
		{Name: "http_3xx", Label: "3xx Redirection", Diff: true, Type: "unit64", Stacked: true},
		{Name: "http_4xx", Label: "4xx Client Error", Diff: true, Type: "unit64", Stacked: true},
		{Name: "http_5xx", Label: "5xx Server Error", Diff: true, Type: "unit64", Stacked: true},
	},
}

// HTTPStatusCounterPlugin mackerel plugin for http-status-counter
type HTTPStatusCounterPlugin struct {
	URI      string
	Grouping bool
}

// HTTPStatusCounterOutput http-status-counter metrics
type HTTPStatusCounterOutput struct {
	Status                      map[string]int
	BodyBytesSent               uint64  `json:"body_bytes_sent"`
	AverageRequestTime          float64 `json:"avg_request_time"`
	AverageUpstreamResponseTime float64 `json:"avg_upstream_response_time"`
}

// FetchMetrics interface for mackerelplugin
func (p HTTPStatusCounterPlugin) FetchMetrics() (map[string]interface{}, error) {
	resp, err := http.Get(p.URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var output HTTPStatusCounterOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	stat := make(map[string]interface{})

	if p.Grouping {
		stat = p.parseStatusGrouping(output)
	} else {
		stat = p.parseStatus(output)
	}

	stat["body_bytes_sent"] = output.BodyBytesSent
	stat["avg_request_time"] = output.AverageRequestTime
	stat["avg_upstream_response_time"] = output.AverageUpstreamResponseTime

	return stat, nil
}

func (p HTTPStatusCounterPlugin) parseStatusGrouping(stats HTTPStatusCounterOutput) map[string]interface{} {
	stat := make(map[string]interface{})

	http1xx := 0
	http2xx := 0
	http3xx := 0
	http4xx := 0
	http5xx := 0

	for code, count := range stats.Status {
		switch code[0:1] {
		case "1":
			http1xx += count
		case "2":
			http2xx += count
		case "3":
			http3xx += count
		case "4":
			http4xx += count
		case "5":
			http5xx += count
		}
	}

	stat["http_1xx"] = uint64(http1xx)
	stat["http_2xx"] = uint64(http2xx)
	stat["http_3xx"] = uint64(http3xx)
	stat["http_4xx"] = uint64(http4xx)
	stat["http_5xx"] = uint64(http5xx)

	return stat
}

func (p HTTPStatusCounterPlugin) parseStatus(stats HTTPStatusCounterOutput) map[string]interface{} {
	stat := make(map[string]interface{})

	for code, count := range stats.Status {
		stat[code] = uint64(count)
	}

	return stat
}

// GraphDefinition interface for mackerelplugin
func (p HTTPStatusCounterPlugin) GraphDefinition() map[string]mp.Graphs {
	if p.Grouping {
		graphdef["http.status"] = graphdefStatusGrouping
	} else {
		graphdef["http.status"] = graphdefStatus
	}
	return graphdef
}

func main() {
	optScheme := flag.String("scheme", "http", "Scheme")
	optHost := flag.String("host", "localhost", "Host")
	optPort := flag.String("port", "80", "Port")
	optPath := flag.String("path", "/status_count", "Path")
	optGrouping := flag.Bool("grouping", true, "Group by class")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	var httpStatusCounter HTTPStatusCounterPlugin
	httpStatusCounter.URI = fmt.Sprintf("%s://%s:%s%s", *optScheme, *optHost, *optPort, *optPath)
	httpStatusCounter.Grouping = *optGrouping

	helper := mp.NewMackerelPlugin(httpStatusCounter)
	helper.Tempfile = *optTempfile
	helper.Run()
}
