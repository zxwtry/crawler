package main

import "github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/task"

func main() {
	// url := "https://music.163.com/#/song?id=341996"
	// errCode, errMsg, mp3Path, lrcPath := service.ServcieUrlMusicNete(url)
	// fmt.Printf("errCode:%d\n", errCode)
	// fmt.Printf("errMsg:%s\n", errMsg)
	// fmt.Printf("mp3Path:%s\n", mp3Path)
	// fmt.Printf("lrcPath:%s\n", lrcPath)

	task.TaskXinWenLianBoMp3()
	// task.TaskShenDuGuoJi()
}
