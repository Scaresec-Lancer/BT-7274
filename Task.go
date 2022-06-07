package main

import (
	"fmt"
	"sort"
)

var (
	//设置两个切片，一个是未完成，一个是已完成
	Undone []Task
	Done   []Task
)

type Task struct {
	//任务的结构体，一个元素是时间，一个元素是内容
	Time    int
	Content string
}

//创建新任务的函数
func createTask(Time int, Content string) {

	var newTask Task

	//如果两个输入都不为空，创建任务
	if Time != 0 && Content != "" {
		newTask.Time = Time
		newTask.Content = Content
	}

	//赋值完成，将任务添加到未完成中
	Undone = append(Undone, newTask)
	fmt.Println("[+] 新任务创建成功！")

}

//删除任务的函数
func delTask(Time int, Content string) {
	for index, value := range Undone {
		if Time == value.Time && Content == value.Content {

			//将指定元素前后切断，并重新拼接
			Undone = Undone[:index+copy(Undone[index:], Undone[index+1:])]
			fmt.Println("[-] 任务删除成功！")
		}
	}
}

//修改任务的函数
func editTask(Time int, Content string) {
	var (
		time    int
		content string
	)
	for _, value := range Undone {
		if Time == value.Time && Content == value.Content {

			fmt.Println("当前时间为", Time, "，请输入要修改为的时间：")
			fmt.Scanln(&time)
			if time != 0 {
				value.Time = time
			}

			fmt.Println("当前任务内容为", Content, "，请输入要修改为的内容：")
			fmt.Scanln(&content)
			if content != "" {
				value.Content = content
			}
		}
		fmt.Println("[+] 任务变更完成！")
	}
}

//对任务列表根据时间进行排序
func sortTask() {
	if len(Undone) > 1 {
		sort.Slice(Undone, func(i, j int) bool {
			return Undone[i].Time > Undone[j].Time
		})
	}
}

//展示任务列表
func showTask() {
	fmt.Printf("%-10s %-20s", "时间", "内容")
}
