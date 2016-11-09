mackerel-plugin-http-status-counter [![wercker status](https://app.wercker.com/status/51a3cb233487aab7399122d982e20a31/s/master "wercker status")](https://app.wercker.com/project/byKey/51a3cb233487aab7399122d982e20a31)
===

Custom metrics plugin for [mackerel-agent](https://github.com/mackerelio/mackerel-agent) to get the number of http requests for each status code.

## Synopsis

```
mackerel-plugin-http-status-counter [-scheme=<http|https>] [-host=<host>] [-port=<port>] [-path=<path>] [-grouping=<true|false>] [-tempfile=<tempfile>]
```

## Requirements

- [http-status-counter](https://github.com/yano3/http-status-counter)
