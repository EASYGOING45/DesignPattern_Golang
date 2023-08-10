# Go设计模式-单例模式

<img src="https://happygoing.oss-cn-beijing.aliyuncs.com/img/单例模式.png" alt="单例模式"  />

![image-20230810090455769](https://happygoing.oss-cn-beijing.aliyuncs.com/img/image-20230810090455769.png)

`go test -benchmem -bench="." -v` 命令用于运行当前包中的所有基准测试，并显示详细输出和内存分配。

解释输出结果的含义如下：

- `=== RUN TestGetLazyInstance` 和 `--- PASS: TestGetLazyInstance (0.00s)` 表示运行并通过了名为 `TestGetLazyInstance` 的测试函数。
- `=== RUN TestGetInstance` 和 `--- PASS: TestGetInstance (0.00s)` 表示运行并通过了名为 `TestGetInstance` 的测试函数。
- `goos: windows` 和 `goarch: amd64` 表示操作系统是 Windows，CPU 架构是 amd64。
- `pkg: DesignPattern_Golang/Singleton` 表示当前运行的测试包名为 `DesignPattern_Golang/Singleton`。
- `cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz` 表示当前运行的计算机 CPU 信息。
- `BenchmarkGetLazyInstanceParallel` 和 `BenchmarkGetLazyInstanceParallel-12` 表示运行并基准测试了名为 `BenchmarkGetLazyInstanceParallel` 的基准测试函数。
- `1000000000` 表示运行了 1000000000 次操作。
- `0.5663 ns/op` 表示每次操作的平均纳秒数。
- `0 B/op` 表示每次操作的平均内存分配字节数。
- `0 allocs/op` 表示每次操作的平均内存分配次数。
- `BenchmarkGetInstanceParallel` 和 `BenchmarkGetInstanceParallel-12` 表示运行并基准测试了名为 `BenchmarkGetInstanceParallel` 的基准测试函数。
- `PASS` 表示所有的测试和基准测试都通过了。
- `ok DesignPattern_Golang/Singleton 1.149s` 表示测试运行的总时间为 1.149 秒。

总结来说，输出结果表示所有的测试和基准测试都通过了，并显示了基准测试的运行时间、每次操作的平均时间、内存分配情况等重要指标。