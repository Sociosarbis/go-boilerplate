package manager

import (
	"container/heap"

	hp "github.com/boilerplate/src/heap"
)

type TaskManager struct {
	taskIdToUserId   map[int]int
	taskIdToPriority map[int]int
	taskQueue        hp.Heap[[2]int]
}

func Constructor(tasks [][]int) TaskManager {
	n := len(tasks)
	taskIdToUserId := make(map[int]int, n)
	taskIdToPriority := make(map[int]int, n)
	taskQueue := hp.New[[2]int]([][2]int{}, func(a, b [2]int) bool {
		return a[1] > b[1] || (a[1] == b[1] && a[0] > b[0])
	})
	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		taskIdToUserId[taskId] = userId
		taskIdToPriority[taskId] = priority
		heap.Push(&taskQueue, [2]int{taskId, priority})
	}
	return TaskManager{
		taskIdToUserId,
		taskIdToPriority,
		taskQueue,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.taskIdToUserId[taskId] = userId
	this.taskIdToPriority[taskId] = priority
	heap.Push(&this.taskQueue, [2]int{taskId, priority})
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	this.taskIdToPriority[taskId] = newPriority
	heap.Push(&this.taskQueue, [2]int{taskId, newPriority})
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskIdToPriority, taskId)
	delete(this.taskIdToUserId, taskId)
}

func (this *TaskManager) ExecTop() int {
	if this.taskQueue.Len() == 0 {
		return -1
	}
	top := heap.Pop(&this.taskQueue).([2]int)
	priority, ok := this.taskIdToPriority[top[0]]
	for !ok || priority != top[1] {
		if this.taskQueue.Len() == 0 {
			return -1
		}
		top = heap.Pop(&this.taskQueue).([2]int)
		priority, ok = this.taskIdToPriority[top[0]]
	}
	userId := this.taskIdToUserId[top[0]]
	this.Rmv(top[0])
	return userId
}
