package main

//统计[t-3000,t]之间的元素个数
/**
把每一次t加入到队列中，然后从队列头开始判断，如果不符合，移除队列
*/
type RecentCounter struct {
	time []int
}

func Constructor() RecentCounter {
	return RecentCounter{time: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.time = append(this.time, t)
	for len(this.time) > 0 && this.time[0] < (t-3000) {
		this.time = this.time[1:]
	}

	return len(this.time)
}
