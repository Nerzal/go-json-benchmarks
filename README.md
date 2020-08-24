# go-json-benchmarks
Some json benchmarks


```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench . -v -count=1

goos: linux
goarch: amd64
pkg: github.com/Nerzal/go-json-benchmarks
Benchmark_Std_Canada-8         	      21	  56231227 ns/op	 3933007 B/op	  116727 allocs/op
Benchmark_json_iter_Canada-8   	      30	  38248309 ns/op	 5448086 B/op	  221219 allocs/op
Benchmark_ffjson_Canada-8      	      22	  50847856 ns/op	 3913833 B/op	  116473 allocs/op
Benchmark_pkg_json_Canada
    Benchmark_pkg_json_Canada: /home/tobias/go/src/github.com/Nerzal/go-json-benchmarks/benchmark_test.go:72: decodeValue: unhandled type: struct
--- FAIL: Benchmark_pkg_json_Canada
Benchmark_go_jay_Canada
    Benchmark_go_jay_Canada: /home/tobias/go/src/github.com/Nerzal/go-json-benchmarks/benchmark_test.go:86: Cannot unmarshal JSON to type '*bench.Canada'
--- FAIL: Benchmark_go_jay_Canada
FAIL
exit status 1
FAIL	github.com/Nerzal/go-json-benchmarks	5.866s
```