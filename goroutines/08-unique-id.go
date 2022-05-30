package main

import (
	"sync/atomic"
	"time"
)

func main() {
	for i := 0; i < 10000000; i++ {
		println(GetId())
		//time.Sleep(time.Millisecond * 2)
	}
}

//最大请求数量
const MAX_ID_NUM = 10000

//最大并发数量
const MAX_CHANEL_NUM = 10

var Ids = make(chan uint64, MAX_ID_NUM)
var info = make(chan []uint64, MAX_CHANEL_NUM)

var uniqueIdReqCount int64
var uniqueIdReqLockSeconds int64

func InitIds() {
	atomic.AddInt64(&uniqueIdReqCount, 1)
	defer func() {
		atomic.AddInt64(&uniqueIdReqCount, -1)
	}()

	println("协程启动中")

	var data []uint64
	var i uint64
	for i = 0; i < 5000; i++ {
		data = append(data, i)
	}

	time.Sleep(2 * time.Second)

	info <- data

	res := <-info
	for _, id := range res {
		Ids <- id
	}

	println("协程启动结束")
}

func GetId() uint64 {
	if len(Ids) < 2000 {
		if len(info) < MAX_CHANEL_NUM {
			println("接收channel小于5")
			if atomic.LoadInt64(&uniqueIdReqCount) < 10 {
				println("请求数小于10，启动协程")
				go InitIds()
			} else {
				println("请求数大于10，暂停请求")
				//设定1个锁定最大时间
				currentTimestamp := time.Now().Unix()
				ok := atomic.CompareAndSwapInt64(&uniqueIdReqLockSeconds, 0, currentTimestamp)
				if !ok {
					println("已经设定过，检查是否要释放")
					//锁定10秒
					if atomic.LoadInt64(&uniqueIdReqLockSeconds) < (currentTimestamp - 10) {
						print(atomic.LoadInt64(&uniqueIdReqLockSeconds))
						println("需要释放")
						atomic.StoreInt64(&uniqueIdReqCount, 0)
						atomic.StoreInt64(&uniqueIdReqLockSeconds, 0)
						go InitIds()
					}
				} else {
					println("设定成功")
				}
			}
		} else {
			println("接收channel大于5")
		}
	}

	return <-Ids
}
