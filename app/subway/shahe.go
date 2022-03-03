package subway

import (
	"encoding/json"
	"fmt"
	"mercury/app/utils"
	"time"
)

type Success struct {
	AppointmentId   string
	Balance         int
	StationEntrance string
}

type Subway struct {
	LineName           string `json:"lineName" binding:"required"`
	SnapshotWeekOffset int    `json:"snapshotWeekOffset" binding:"-"`
	StationName        string `json:"stationName" binding:"required"`
	EnterDate          string `json:"enterDate" binding:"required"`
	SnapshotTimeSlot   string `json:"snapshotTimeSlot" binding:"required"`
	TimeSlot           string `json:"timeSlot" binding:"required"`
}

const layout = "2006-01-02 15:04:05"

func getEnterDate() string {
	now := time.Now()
	tt, _ := time.ParseDuration("24h")
	dd := now.Add(tt)
	return dd.Format("20060102")
}

// 时间对比
func isStart() (bool, error) {
	time1 := time.Now().Format("2006-01-02 ") + "11:59:59"
	time2 := time.Now().Format("2006-01-02 ") + "13:00:00"
	time3 := time.Now().Format("2006-01-02 ") + "19:59:59"
	now := time.Now().Format(layout)
	fmt.Println(now)
	t1, err1 := time.Parse(layout, time1)
	t2, err2 := time.Parse(layout, time2)
	t3, err3 := time.Parse(layout, time3)
	nw, err4 := time.Parse(layout, now)
	if err1 != nil && err2 != nil && err3 != nil && err4 != nil {
		return false, nil
	}
	// 上午
	if nw.Before(t2) {
		return t1.Before(nw), nil
	} else {
		return t3.Before(nw), nil
	}
	// 下午
}

// 获取请求的参数
func getStringParams(timeSlot string) (string, error) {
	enterDate := getEnterDate()
	params := &Subway{
		LineName:           "昌平线",
		EnterDate:          enterDate,
		SnapshotTimeSlot:   "0630-0930",
		SnapshotWeekOffset: 0,
		StationName:        "沙河站",
		TimeSlot:           timeSlot,
	}

	json, err := json.Marshal(params)
	fmt.Println(params)
	if err != nil {
		fmt.Println("err---<", err)
		return "", err
	}
	// bytes.NewBuffer 是一个缓冲byte类型的缓冲器，这个缓冲器里面存放的都是byte
	return string(json), err
}

// 获取请求头
func getRequestHeader() map[string]string {
	authorization, _ := utils.ReadFile()
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": authorization,
		"Origin":        "https://webui.mybti.cn",
	}
}

func request(timeSlot string) {
	params, _ := getStringParams(timeSlot)
	var header = getRequestHeader()
	body, err := utils.HTTPPost("https://webapi.mybti.cn/Appointment/CreateAppointment", params, header)

	if err != nil {
		fmt.Println(err)
	}
	var data Success // TopTracks
	err = json.Unmarshal(body, &data)
	fmt.Println("success---->", data.Balance)
}

func myChannel(id int, timeSlot string) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		var count = 0
		for range ticker.C {
			isBegin, _ := isStart()
			if isBegin {
				count++
				request(timeSlot)
				if count == 2 {
					ticker.Stop()
				}
			}
		}
	}()
}

func isLogin() {
	var header = getRequestHeader()
	body, err := utils.HTTPGet("https://webapi.mybti.cn/AppointmentRecord/GetCurrentAppointment", header)
	if err != nil {
		fmt.Println("登录状态失效---->>", err)
		//res, err := utils.HTTPPost("https://webapi.mybti.cn/User/SignUp", )
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
}
func Start() {
	isLogin()
	myChannel(1, "0740-0750")
	//myChannel(2, "0630-0640")
	//myChannel(3, "0750-0800")
}
