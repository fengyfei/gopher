1. SetOne 函数单独测试结果：

   ```shell
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bigcache
   BenchmarkCacheServiceProvider_SetOne-4   	 2000000	       676 ns/op
   PASS
   ok  	github.com/fengyfei/gopher/bigcache	2.101s
   ```

2. SetMany 函数单独测试结果，此函数里面循环了 1000 万次：

   ```shell
   2017/09/16 17:11:26 Allocated new queue in 160.36µs; Capacity: 585000
   2017/09/16 17:11:26 Allocated new queue in 315.279µs; Capacity: 1170000
   2017/09/16 17:11:26 Allocated new queue in 687.11µs; Capacity: 2340000
   2017/09/16 17:11:26 Allocated new queue in 1.857736ms; Capacity: 4680000
   2017/09/16 17:11:26 Allocated new queue in 3.243251ms; Capacity: 8388608
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bigcache
   BenchmarkCacheServiceProvider_SetMany-4   	       1	2505799703 ns/op
   PASS
   ok  	github.com/fengyfei/gopher/bigcache	2.554s
   ```

3. SetOne 和 GetOne 顺序执行测试结果：

   ```shell
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bigcache
   BenchmarkCacheServiceProvider_SetOne-4   	 2000000	       660 ns/op
   BenchmarkCacheServiceProvider_GetOne-4   	10000000	       134 ns/op
   PASS
   ok  	github.com/fengyfei/gopher/bigcache	3.544s
   ```

4. SetMany 函数单独测试结果：

   ```shell
   2017/09/16 17:40:04 Allocated new queue in 316.718µs; Capacity: 585000
   2017/09/16 17:40:04 Allocated new queue in 601.455µs; Capacity: 1170000
   2017/09/16 17:40:04 Allocated new queue in 1.153164ms; Capacity: 2340000
   2017/09/16 17:40:04 Allocated new queue in 2.319411ms; Capacity: 4680000
   2017/09/16 17:40:04 Allocated new queue in 6.971709ms; Capacity: 8388608
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bigcache
   BenchmarkCacheServiceProvider_SetMany-4   	       1	2543070007 ns/op
   PASS
   ok  	github.com/fengyfei/gopher/bigcache	2.580s
   ```

5. SetOne 和 GetMany 顺序执行测试结果：

   ```shell
   goos: darwin
   goarch: amd64
   pkg: github.com/fengyfei/gopher/bigcache
   BenchmarkCacheServiceProvider_SetOne-4   	 2000000	       689 ns/op
   BenchmarkCacheServiceProvider_GetAll-4   	       1	1168408083 ns/op
   PASS
   ok  	github.com/fengyfei/gopher/bigcache	3.290s
   ```

   ​
