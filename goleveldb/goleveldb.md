1. 单独写操作测试结果：

   ```shell
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_Create-4             300           7413778 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    2.742s
   ```

2. 单独顺序读操作测试结果：

   ```shell
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_Get-4                200           7633436 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    2.342s
   ```

3. 单独随机读操作测试结果：

   ```shell
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_GetRandom-4          100          10337479 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    1.052s
   ```

4. 写和读顺序执行测试结果：

   ```shell
   # 写入、顺序读取、随机读取顺序执行
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_Create-4             300           8146472 ns/op
   BenchmarkAccountServiceProvider_Get-4                200           7514547 ns/op
   BenchmarkAccountServiceProvider_GetRandom-4          200           7402362 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    7.468s
   ```

