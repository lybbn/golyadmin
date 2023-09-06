package crontab

import (
	"fmt"
	"testing"
	"time"
)

func testFunc() {
	time.Sleep(time.Second)
	fmt.Println("执行了AddTaskFunc定时任务")
}

func TestNewCronTasker(t *testing.T) {
	ct := NewCronTasker()
	_ct := ct.(*cronTasker)
	taskid, _ := ct.AddTaskFunc("task1", "@every 2s", testFunc)
	fmt.Printf("task1定时任务ID：%d \n", taskid)
	_, ok := _ct.taskList["task1"]
	if !ok {
		t.Error("no find task1")
	}
}
