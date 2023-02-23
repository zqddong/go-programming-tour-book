#### pprof 性能剖析

- 通过浏览器访问

```go
import (
	_ "net/http/pprof"
	...
)

open URL http://127.0.0.1:6060/debug/pprof/

/debug/pprof/

Types of profiles available:
Count	Profile
3	allocs
0	block
0	cmdline
8	goroutine
3	heap
0	mutex
0	profile
11	threadcreate
0	trace
full goroutine stack dump

allocs：查看过去所有内存分配的样本，访问路径为 $HOST/debug/pprof/allocs。
block：查看导致阻塞同步的堆栈跟踪，访问路径为 $HOST/debug/pprof/block。
cmdline： 当前程序的命令行的完整调用路径。
goroutine：查看当前所有运行的 goroutines 堆栈跟踪，访问路径为 $HOST/debug/pprof/goroutine。
heap：查看活动对象的内存分配情况， 访问路径为 $HOST/debug/pprof/heap。
mutex：查看导致互斥锁的竞争持有者的堆栈跟踪，访问路径为 $HOST/debug/pprof/mutex。
profile： 默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件，访问路径为 $HOST/debug/pprof/profile。
threadcreate：查看创建新 OS 线程的堆栈跟踪，访问路径为 $HOST/debug/pprof/threadcreate。
```

- 终端交互式访问

```shell
$  go tool pprof http://127.0.0.1:6060/debug/pprof/profile\?seconds\=60

...

$ (pprof) top10

flat：函数自身的运行耗时。
flat%：函数自身在 CPU 运行耗时总比例。
sum%：函数自身累积使用 CPU 总比例。
cum：函数自身及其调用函数的运行总耗时。
cum%：函数自身及其调用函数的运行耗时总比例。
Name：函数名。


$ go tool pprof http://localhost:6060/debug/pprof/heap

# inuse_space：分析应用程序的常驻内存占用情况。
go tool pprof -inuse_space http://localhost:6060/debug/pprof/heap

# alloc_objects：分析应用程序的内存临时分配情况。
go tool pprof -alloc_objects http://localhost:6060/debug/pprof/heap
```


#### trace

```go
import (
    "os"
    "runtime/trace"
)

func main() {
    trace.Start(os.Stderr)
    defer trace.Stop()

    ch := make(chan string)
    go func() {
        ch <- "Go 语言编程之旅"
    }()

    <-ch
}

//$ go run main.go 2> trace.out
//$ go tool trace trace.out

// View trace：查看跟踪
// Goroutine analysis：Goroutine 分析
// Network blocking profile：网络阻塞概况
// Synchronization blocking profile：同步阻塞概况
// Syscall blocking profile：系统调用阻塞概况
// Scheduler latency profile：调度延迟概况
// User defined tasks：用户自定义任务
// User defined regions：用户自定义区域
// Minimum mutator utilization：最低 Mutator 利用率
```

### GODEBUG

```go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}

// $ GODEBUG=schedtrace=1000 go run main.go
// sched：每一行都代表调度器的调试信息，后面提示的毫秒数表示启动到现在的运行时间，输出的时间间隔受 schedtrace 的值影响。
// gomaxprocs：当前的 CPU 核心数（GOMAXPROCS 的当前值）。
// idleprocs：空闲的处理器数量，后面的数字表示当前的空闲数量。
// threads：OS 线程数量，后面的数字表示当前正在运行的线程数量。
// spinningthreads：自旋状态的 OS 线程数量。
// idlethreads：空闲的线程数量。
// runqueue：全局队列中中的 Goroutine 数量，而后面的 [0 0 1 1] 则分别代表这 4 个 P 的本地队列正在运行的 Goroutine 数量。

// $ GODEBUG=scheddetail=1,schedtrace=1000 go run main.go
```


#### gops

```shell
$ go install github.com/google/gops@latest
```
```go
import (
	...
	"github.com/google/gops/agent"
)

func main() {
	// 创建并监听 gops agent，gops 命令会通过连接 agent 来读取进程信息
	// 若需要远程访问，可配置 agent.Options{Addr: "0.0.0.0:6060"}，否则默认仅允许本地访问
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`Go 语言编程之旅 `))
	})
	_ := http.ListenAndServe(":6060", http.DefaultServeMux)
}
```


```shell
$ gops help

# 查看指定进程信息
$ gops <pid>

# 查看调用栈信息
$ gops stack <pid>

# 查看内存使用情况
$ gops memstats <pid>

# 查看运行时信息
$ gops stats <pid>

# 查看 trace 信息
$ gops trace <pid>

# 查看 profile 信息
$ gops pprof-cpu <pid>
```

#### 逃逸分析

```shell
# 怎么确定是否逃逸
# 1. 通过编译器提供的指令 -gcflags
$ go build -gcflags '-m -l' main.go

# 2. 通过反编译命令查看
$ go tool compile -S main.go

```
