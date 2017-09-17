1. 单独写操作测试结果：

   ```shell
   # 200000 次循环生成 name 版本结果
   yangs-Air:bbolt yang$ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bbolt
   BenchmarkUserServiceProvider_Create-4          2	 687744917 ns/op
   PASS
   ok  github.com/fengyfei/gopher/bbolt	2.241s
   # 1000000 次循环生成 name 版本结果
   yangs-Air:bbolt yang$ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bbolt
   BenchmarkUserServiceProvider_Create-4          1	2301492615 ns/op
   PASS
   ok  github.com/fengyfei/gopher/bbolt	2.337s
   ```

2. 单独读操作测试结果：

   ```shell
   # 1000000 次循环生成 name 再读取的结果
   yangs-Air:bbolt yang$ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bbolt
   BenchmarkUserServiceProvider_Get-4          1	2740592493 ns/op
   PASS
   ok  github.com/fengyfei/gopher/bbolt	2.757s
   ```

3. 写和读顺序执行测试结果：

   ```shell
   # 第一次运行，1000000 次循环写和读，数据库大小 80.7 M
   yangs-Air:bbolt yang$ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bbolt
   BenchmarkUserServiceProvider_Create-4          1	424240834381 ns/op
   BenchmarkUserServiceProvider_Get-4             1	3079024556 ns/op
   PASS
   ok  github.com/fengyfei/gopher/bbolt	427.403s

   # 第二次运行，各 1000000 次循环写和读，数据库变为 144.7 M
   yangs-Air:bbolt yang$ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bbolt
   BenchmarkUserServiceProvider_Create-4          1	2479825500 ns/op
   BenchmarkUserServiceProvider_Get-4             1	2970980138 ns/op
   PASS
   ok  github.com/fengyfei/gopher/bbolt	5.499s
   ```
