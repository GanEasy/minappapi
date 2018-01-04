package minappapi

import (
	"fmt"
	"time"
)

// Que 可执行任务接口
type Que interface {
	Execute()
}

//NoticeDemo struct
type NoticeDemo struct {
	OpenID string
	Text   string
}

//Execute Notice.Run
func (p *NoticeDemo) Execute() {
	time.Sleep(time.Second)
	fmt.Printf("给 %s 发送内容 %s\n", p.OpenID, p.Text)
}

// Job represents the job to be run
type Job struct {
	Que Que
}

//JobQueue A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

//NewWorker new worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// job queue execute
				job.Que.Execute()
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//Dispatcher 工作池
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	MaxWorkers int
	WorkerPool chan chan Job
}

//NewDispatcher 新建工作池
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

//Run 工作池开始工作
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

// 工作池内部循环工作
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job

			}(job)
		}
	}
}

func init() {

	maxWorkers := 10
	maxQueue := 2000
	//初始化一个工作池,并指定它可以操作的 工人个数
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue) //指定任务的队列长度
	//并让它一直接运行
	dispatch.Run()
	// close(notice.JobQueue)
}
