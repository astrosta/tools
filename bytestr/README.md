#性能测试结果
```shell
>go test -bench="." -benchmem
goos: windows
goarch: amd64
pkg: github.com/astrosta/tools/bytestr
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
Benchmark_NormalBytes2String-12         47796196                25.11 ns/op           48 B/op          1 allocs/op
Benchmark_Byte2String-12                1000000000               0.2413 ns/op          0 B/op          0 allocs/op
Benchmark_NormalString2Bytes-12         39800070                30.77 ns/op           48 B/op          1 allocs/op
Benchmark_String2Bytes-12               1000000000               0.2467 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/astrosta/tools/bytestr       3.098s
```