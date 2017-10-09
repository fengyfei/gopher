1. 单独写操作测试结果：

   ```shell
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_Create-4          100000             12619 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    1.430s
   ```

2. 单独顺序读操作测试结果：

   ```shell
     300000             17967 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    5.513s
   ```

3. 写和读顺序执行测试结果：

   ```shell
   # 写入、顺序读取顺序执行
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_Create-4          200000             12058 ns/op
   BenchmarkAccountServiceProvider_Get-4             200000              5638 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    4.558s

   # 写入、顺序读取、随机读取顺序执行
   gopher/goleveldb git/master*  
   ❯ go test -bench=.
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/goleveldb
   BenchmarkAccountServiceProvider_Create-4          100000             12035 ns/op
   BenchmarkAccountServiceProvider_Get-4             200000              6127 ns/op
   BenchmarkAccountServiceProvider_GetRandom-4       200000              6725 ns/op
   PASS
   ok      github.com/fengyfei/gopher/goleveldb    4.046s
   ```

