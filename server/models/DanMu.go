package models

type DanMu struct {
	GlobalModel
	UserId   string
	VideoId  string
	Time     float64
	Text     string
	Color    int
	Position int
	Type     uint
}

//"id": 弹幕id
//"userId": 发弹幕的userid
//"author": 作者,
//"time": 3.300442, //弹幕位置时间
//"text": "你好",//文本
//"color": 16777215,//颜色
//"type": 0// 位置
//"createTime" 创建弹幕时间
