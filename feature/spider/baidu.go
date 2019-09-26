
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "os"
)

// 爬取网页内容
func HttpGet(url string)(result string, err error) {
    resp, err1 := http.Get(url)
    if err1 != nil {
        // print err
        err = err1
        return
    }

    defer resp.Body.Close()

    // 读取网页 Body 内容
    buf := make([]byte, 1024 * 4)
    for {
        n, _ := resp.Body.Read(buf)
        if n == 0 { // 读取结束，或者出问题
            // print err
            break
        }

        result += string(buf[:n])
    }

    return
}

// 爬取一个网页
func SpidePage(i int, page chan<- int) {
    url := "http://tieba.baidu.com/f?kw=ios&ie=utf-8&pn=0&red_tag=y1558891520&pn=" + strconv.Itoa((i - 1) * 50)
    // print url

    // 2) 爬，将所有的网站的内容全部爬下来
    result, err := HttpGet(url)
    if err != nil {
    	fmt.Println("err")
    }

    // 把内容写入到文件
    fileName := strconv.Itoa((i - 1) * 50) + ".html"
    f, err1 := os.Create(fileName)
    if err1 != nil {
        // print err1
    	fmt.Println("err")
    }

    f.WriteString(result) // 写内容
    f.Close() // 关闭文件

    page <- i // 写内容
}


func DoWork(start, end int) {
    fmt.Printf("正在爬取 %d 到 %d 的页面\b", start, end)

    page := make(chan int)

    // 1) 明确目标（要知道你准备在哪个范围或者网站去搜索）
    // https://tieba.baidu.com/f?kw=绝地求生&ie=utf-8&pn=0 // 下一页加 50
    for i := start; i < end; i++ {
        go SpidePage(i, page)
    }

    for i := start; i < end; i++ {
        // 没爬完之前是阻塞的
        fmt.Printf("第%d个页面爬取完成\n", <-page)
    }
}

func main() {
    var start, end int
    fmt.Printf("请输入起始页(>= 1): ")
    fmt.Scan(&start)
    fmt.Printf("请输入终止页(>= 起始页): ")
    fmt.Scan(&end)

    DoWork(start, end)
}