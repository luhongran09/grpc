package main

import (
	"fmt"
	"sync"
)

func main() {
	//创建一个工作池
	workerCount := 5
	bufferSize := 10
	wp := NewWorkerPool(workerCount, bufferSize)
	//启动工作池中的工作线程
	for i := 0; i < workerCount; i++ {
		go wp.worker(i)
	}

	//提交任务到工作池
	tasks := []Task{
		{URL: "https://example.com/page1"},
		{URL: "https://example.com/page2"},
		{URL: "https://example.com/page3"},
		// 添加更多任务...
	}
	for _, task := range tasks {
		wp.SubmitTask(task)
	}
	wp.Close()
}

// 定义任务类型
type Task struct {
	URL string
}

// 定义工作池
type WorkerPool struct {
	WorkerCount int
	Tasks       chan Task
	waitGroup   sync.WaitGroup
}

// 定义一个工作函数，用于处理任务
func (wp *WorkerPool) worker(id int) {
	defer wp.waitGroup.Done()
	for Task := range wp.Tasks {
		//执行爬取任务
		fmt.Print(Task)
	}
}

// 初始化工作池
func NewWorkerPool(workerCount, buffersize int) *WorkerPool {
	wp := &WorkerPool{
		WorkerCount: workerCount,
		Tasks:       make(chan Task, buffersize),
	}
	return wp
}

// 提交任务到工作池
func (wp *WorkerPool) SubmitTask(task Task) {
	wp.waitGroup.Add(1)
	wp.Tasks <- task
}

// 关闭任务通道
func (wp *WorkerPool) Close() {
	close(wp.Tasks)
	wp.waitGroup.Wait()
}
