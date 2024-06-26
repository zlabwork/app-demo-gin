## Run & Best Practices
```shell
./run.sh
```

```shell
# Use `export` instead of godotenv
# source .env
export `cat .env` # tested in mac and unix
export `cat .env | grep -v "#"` && go run main.go
```

---

## Request
```go
// 获取 Path Parameters
name := c.Param("name")

// 获取参数 GET
name := c.Query("name")

// 获取参数 POST PUT PATCH
name := c.PostForm("name")

// 获取 raw json
argv := struct {
    Name string
}{}
c.ShouldBind(&argv)

// 获取 body 任意请求类型
var r io.Reader = c.Request.Body
b, _ := io.ReadAll(r)
fmt.Println(string(b))
```


## Cookie
```go
// https://gin-gonic.com/docs/examples/cookie/

// 读取 cookie
cookie, err := c.Cookie("gin_cookie")

// 写入 cookie
c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
```

---

## Defined
```golang
// 分隔符
const (
    nul uint8 = 0x00 // 空字符
    lf  uint8 = 0x0A // 换行
    cr  uint8 = 0x0D // 回车键
    fs  uint8 = 0x1C // 文件分隔符
    gs  uint8 = 0x1D // 组分隔符
    rs  uint8 = 0x1E // 记录分隔符
    us  uint8 = 0x1F // 单元分隔符
)
```


## Rand Read
```golang
b := make([]byte, 32)
rand.Read(b)
```


## Copy
```golang
// bytes
bs[0] = 0x1F
copy(bs, bs1)
bytes.Split(bs, []byte{0x1F})
```


## Uint32/16/8 to bytes
```golang
// uint64 to []byte
// e.g. 1
var num uint64 = 258
b := make([]byte, 8)
binary.BigEndian.PutUint64(b, num)

// e.g. 2
var num uint16 = 2
bf := bytes.NewBuffer(nil)
binary.Write(bf, binary.BigEndian, num)
b := bf.Bytes() // [0 2]
```


## Bytes to uint32/16/8
```golang
bs := []byte{0x00, 0x00, 0x01, 0x02}
num := binary.BigEndian.Uint32(bs)

bs := []byte{0x01, 0x02}
num := binary.BigEndian.Uint16(bs)

bs := []byte{0x01}
num := bs[0]
```


## Sync
```golang
// 守护进程
var w sync.WaitGroup
w.Add(2)
go func () {
    // do something
    w.Done()
}
go func () {
    // do something
    w.Done()
}
w.Wait()
```


## Map & Array
```golang
// map是无序的, 数组是有序的
var foo = [3]int{1, 2, 3}
var bar = map[string]int64{
    "a": 1,
    "b": 2,
    "c": 3,
}

for index, item := range foo {
    fmt.Println(index, item)
}

for index, item := range bar {
    fmt.Println(index, item)
}

// 输出：
// 0 1
// 1 2
// 2 3

// c 3
// a 1
// b 2

// 解决方案 - 引入其他排序数组
import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
```


```go
// 切片特性 - 引用传递
func main() {
	alice := []int{0, 1, 2, 3, 4}
	fmt.Println(alice) // [0 1 2 3 4]

	bob := nothingDid(alice)
	bob[1] = 9

	fmt.Println(alice) // [0 9 2 3 4]
}

func nothingDid(data []int) []int {
	return data
}
```


```go
// 数组特性 - 值传递
func main() {

	alice := [5]int{0, 1, 2, 3, 4}
	fmt.Println(alice) // [0 1 2 3 4]

	bob := nothingDid(alice)
	bob[1] = 9

	fmt.Println(alice) // [0 1 2 3 4]
}

func nothingDid(data [5]int) [5]int {
	return data
}
```


## Libraries
Visit [here](https://github.com/avelino/awesome-go) for more about the libraries.  

| 包名                                                                               | 简介                                   |
|----------------------------------------------------------------------------------|--------------------------------------|
| [github.com/labstack/echo](https://github.com/labstack/echo)                     | web 框架                               |
| [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)                     | web 框架                               |
| [github.com/joho/godotenv](https://github.com/joho/godotenv)                     | 配置, env 环境变量                         |
| [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3)                                     | 配置, yaml                             |
| [github.com/spf13/viper](https://github.com/spf13/viper)                         | 配置, 支持多种格式                           |
| [gorm.io/gorm](https://gorm.io)                                                  | MySQL ORM                            |
| [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)         | 数据库驱动, MySQL                         |
| [github.com/lib/pq](https://github.com/lib/pq)                                   | 数据库驱动, Postgres                      |
| [github.com/gocql/gocql](https://github.com/gocql/gocql)                         | 数据库驱动, Cassandra                     |
| [go.mongodb.org/mongo-driver](https://go.mongodb.org/mongo-driver)               | 数据库驱动, mongoDB                       |
| [github.com/go-redis/redis](https://github.com/go-redis/redis)                   | 数据库驱动, redis                         |
| [github.com/go-resty/resty/](https://github.com/go-resty/resty/)                 | http 请求客户端                           |
| [github.com/valyala/fastjson](https://github.com/valyala/fastjson)               | json 解析                              |
| [github.com/tidwall/gjson](https://github.com/tidwall/gjson)                     | json 解析, 注意: []强制转为字符串后为[]           |
| [github.com/buger/jsonparser](https://github.com/buger/jsonparser)               | json 解析                              |
| [github.com/CloudyKit/jet](https://github.com/CloudyKit/jet)                     | jet 模版引擎                             |
| [github.com/flosch/pongo2](https://github.com/flosch/pongo2)                     | pongo2 模版引擎                          |
| [github.com/Masterminds/sprig](https://github.com/Masterminds/sprig)             | sprig 模版函数                           |
| [github.com/Sirupsen/logrus](https://github.com/Sirupsen/logrus)                 | 日志                                   |
| [github.com/uber-go/zap](https://github.com/uber-go/zap)                         | 日志                                   |
| [sqids.org](https://sqids.org/)                                                  | sqids.org                            |
| [github.com/bwmarrin/snowflake](https://github.com/bwmarrin/snowflake)           | ID生成                                 |
| [github.com/google/uuid](https://github.com/google/uuid)                         | UUID生成                               |
| [github.com/panjf2000/ants](https://github.com/panjf2000/ants)                   | ants 线程池                             |
| [github.com/silenceper/pool](https://github.com/silenceper/pool)                 | 线程池                                  |
| [github.com/apache/thrift](https://github.com/apache/thrift)                     | Thrift                               |
| [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5)             | Json web token                       |
| [github.com/panjf2000/gnet](https://github.com/panjf2000/gnet)                   | 基于事件驱动的网络框架                          |
| [github.com/cloudwego/netpoll](https://github.com/cloudwego/netpoll)             | NIO(Non-blocking I/O) 网络库，专注于 RPC 场景 |
| [github.com/lesismal/nbio](https://github.com/lesismal/nbio)                     | 异步网络框架                               |
| [github.com/ThreeDotsLabs/watermill](https://github.com/ThreeDotsLabs/watermill) | 发布订阅                                 |
| [github.com/asaskevich/EventBus](https://github.com/asaskevich/EventBus)         | EventBus                             |
| [github.com/sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)   | ChatGPT OPENAI                       |
| [github.com/RoaringBitmap/roaring](https://github.com/RoaringBitmap/roaring)     | RoaringBitmap                        |
| [github.com/kelindar/bitmap](https://github.com/kelindar/bitmap)                 | bitmap                               |
| [github.com/bits-and-blooms/bitset](https://github.com/bits-and-blooms/bitset)   | bitset                               |


## Technology
| 技术 | 简介 |
| --- | --- |
| [Zipkin](https://zipkin.io/) | 链路追踪 |
| [skywalking](https://skywalking.apache.org/) | 链路追踪 |
| [Jaeger](https://www.jaegertracing.io/)   | 链路追踪 |

| 特殊质数       | 备注                  |
|------------|---------------------|
| 1999999973 | 小于 2000000000 的最大质数 |
| 2147483647 | 小于 2^31 的最大质数       |
| 2147483629 | 小于 2^31 的大质数        |
| 4294967291 | 小于 2^32 的最大质数       |


## AI

| 名称                                                                             | 介绍       |
|--------------------------------------------------------------------------------|----------|
| [screenshot-to-code](https://github.com/abi/screenshot-to-code) | 根据截图生成网站 |

## Tools
https://platform.openai.com  
https://mholt.github.io/json-to-go  
https://www.jsonformatter.io  
https://www.processon.com  
https://www.figma.com  
https://swagger.io  


## Docs
https://bigbully.github.io/Dapper-translation  
https://www.zhihu.com/question/65502802  
https://github.com/avelino/awesome-go  
https://zhuanlan.zhihu.com/p/39326315  
https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1  
https://studygolang.com/articles/17467?fr=sidebar  
http://hbase.org.cn  
https://blog.boot.dev/golang/golang-project-structure  
https://t5k.org/lists/small/millions/  
http://www.prime-numbers.org/prime-number-2147480000-2147480000.htm  
http://www.prime-numbers.org/prime-number-4294960000-4294960000.htm  
