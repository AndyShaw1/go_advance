package redis

// 问题1 ：使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
// 解答：相关命令参考文档：https://help.aliyun.com/document_detail/38689.html
/*
Usage: redis-benchmark [-h] [-p] [-c] [-n[-k]
 -h     Server hostname (default 127.0.0.1)
 -p     Server port (default 6379)
 -s     Server socket (overrides host and port)
 -c     Number of parallel connections (default 50)
 -n     Total number of requests (default 10000)
 -d      Data size of SET/GET value in bytes (default 2)
 -k      1=keep alive 0=reconnect (default 1)
 -r      Use random keys for SET/GET/INCR, random values for SADD
  Using this option the benchmark will get/set keys
  in the form mykey_rand:000000012456 instead of constant
  keys, the argument determines the max
  number of values for the random number. For instance
  if set to 10 only rand:000000000000 - rand:000000000009
  range will be allowed.
 -P       Pipelinerequests. Default 1 (no pipeline).
 -q       Quiet. Just show query/sec values
 --csv     Output in CSV format
 -l       Loop. Run the tests forever
 -t       Only run the comma-separated list of tests. The test
  names are the same as the ones produced as output.
 -I       Idle mode. Just open N idle connections and wait.
 */


// 在mac item2中，使用redis-benchmark命令
//结果如下：
/*
andyshaw@MacBook-Pro  ~  redis-benchmark -h 127.0.0.1 -p 31001 -d 10 -t get,set
====== SET ======
100000 requests completed in 56.97 seconds
50 parallel clients
10 bytes payload
keep alive: 1
1755.19 requests per second

====== GET ======
100000 requests completed in 57.00 seconds
50 parallel clients
10 bytes payload
keep alive: 1
1754.26 requests per second

//同样20
====== SET ======
100000 requests completed in 57.14 seconds
====== GET ======
100000 requests completed in 57.62 seconds

//50
====== SET ======
  100000 requests completed in 56.80 seconds
====== GET ======
  100000 requests completed in 57.07 seconds

//100
====== SET ======
  100000 requests completed in 61.71 seconds
====== GET ======
  100000 requests completed in 57.17 seconds

//200
====== SET ======
  100000 requests completed in 64.91 seconds
====== GET ======
  100000 requests completed in 72.64 seconds

//1k
====== SET ======
  100000 requests completed in 72.84 seconds
====== GET ======
  100000 requests completed in 60.54 seconds

//5k
====== SET ======
  100000 requests completed in 123.06 seconds
====== GET ======
  100000 requests completed in 58.07 seconds
 */


import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var client redis.UniversalClient
var ctx context.Context

const (
	ip   string = "127.0.0.1"
	port uint16 = 6380
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})

	ctx = context.Background()
}

func Compare() {
	write(10000, "len10_10k", generateValue(10))
	write(50000, "len10_50k", generateValue(10))
	write(500000, "len10_500k", generateValue(10))

	write(10000, "len1000_10k", generateValue(1000))
	write(50000, "len1000_50k", generateValue(1000))
	write(500000, "len1000_500k", generateValue(1000))

	write(10000, "len5000_10k", generateValue(5000))
	write(50000, "len5000_50k", generateValue(5000))
	write(500000, "len5000_500k", generateValue(5000))

	analysis()
}

func write(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := client.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}
