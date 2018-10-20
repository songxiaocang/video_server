package main

import "log"

type connLimiter struct {
	concurrentConn int
	bucket chan int
}

func newConnLimiter(cc int) *connLimiter{
	return &connLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),
	}
}

func(m *connLimiter) getConn() bool{
	//m.bucket
	//log.Print(len(m.bucket))
	//len(m.bucket)
	if len(m.bucket) > m.concurrentConn {
		log.Printf("over by max connection count:%v",m.concurrentConn)
		return false
	}
	m.bucket <- 1
	return true
}

func(m *connLimiter) releaseConn() {
	i := <-m.bucket
	log.Printf("current conn count:%v",i)
}