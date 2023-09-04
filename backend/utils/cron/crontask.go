package cron

import (
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type CronTasker interface {
	AddTaskFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error)
	AddTaskJob(taskName string, spec string, job interface{ Run() }, option ...cron.Option) (cron.EntryID, error)
	StartTask(taskName string)
	StopTask(taskName string)
	RemoveTask(taskName string, id int)
}

// cronTasker 定时任务
type cronTasker struct {
	taskList   map[string]*cron.Cron
	sync.Mutex //互斥锁
}

// AddTaskFunc 添加任务(通过函数的方法添加任务)、参数：spec任务时间(cron表达式)、cmd 任务函数、option 任务参数 如：cron.WithSeconds()则精确到秒级别，不添加，默认精确到分钟
func (c *cronTasker) AddTaskFunc(taskName string, spec string, cmd func(), option ...cron.Option) (cron.EntryID, error) {
	c.Lock()
	defer c.Unlock()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	option = append(option, cron.WithSeconds())
	option = append(option, cron.WithLocation(loc)) //指定任务时区
	if _, ok := c.taskList[taskName]; !ok {
		// 创建一个定时任务的实例，带上cron.WithSeconds()则精确到秒级别
		c.taskList[taskName] = cron.New(option...)
	}
	id, err := c.taskList[taskName].AddFunc(spec, cmd)
	c.taskList[taskName].Start()
	return id, err
}

// AddTaskJob 添加任务(通过接口的方法)、参数：spec任务时间(cron表达式)、job 、option 任务参数 如：cron.WithSeconds()则精确到秒级别，不添加，默认精确到分钟
func (c *cronTasker) AddTaskJob(taskName string, spec string, job interface{ Run() }, option ...cron.Option) (cron.EntryID, error) {
	c.Lock()
	defer c.Unlock()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	option = append(option, cron.WithSeconds())
	option = append(option, cron.WithLocation(loc)) //指定任务时区
	if _, ok := c.taskList[taskName]; !ok {
		// 创建一个定时任务的实例，带上cron.WithSeconds()则精确到秒级别
		c.taskList[taskName] = cron.New(option...)
	}
	id, err := c.taskList[taskName].AddJob(spec, job)
	c.taskList[taskName].Start()
	return id, err
}

// StartTask 开始指定任务
func (c *cronTasker) StartTask(taskName string) {
	c.Lock()
	defer c.Unlock()
	if value, ok := c.taskList[taskName]; ok {
		value.Start()
	}
}

// StopTask 停止指定任务
func (c *cronTasker) StopTask(taskName string) {
	c.Lock()
	defer c.Unlock()
	if value, ok := c.taskList[taskName]; ok {
		value.Stop()
	}
}

// RemoveTask 从taskName 删除指定任务
func (c *cronTasker) RemoveTask(taskName string, id int) {
	c.Lock()
	defer c.Unlock()
	if value, ok := c.taskList[taskName]; ok {
		value.Remove(cron.EntryID(id))
	}
}

func NewCronTasker() CronTasker {
	return &cronTasker{taskList: make(map[string]*cron.Cron)}
}
