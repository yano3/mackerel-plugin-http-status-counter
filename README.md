mackerel-plugin-http-status-counter
===

Custom metrics plugin for [mackerel-agent](https://github.com/mackerelio/mackerel-agent) to get the number of http requests for each status code.

## Synopsis

```
mackerel-plugin-http-status-counter [-scheme=<http|https>] [-host=<host>] [-port=<port>] [-grouping=<true|false>] [-tempfile=<tempfile>]
```

## Requirements

- [http-status-counter](https://github.com/yano3/http-status-counter)
