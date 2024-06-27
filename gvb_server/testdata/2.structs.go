package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"please enter the title" structs:"Title"`
	Href   string `json:"href" binding:"required,url" msg:"illegal jump link" structs:"Href"`
	Images string `json:"images" binding:"required,url" msg:"illegal address of the image" structs:"Images"`
	IsShow bool   `json:"isShow" binding:"required" msg:"please choose whether to show up" structs:"IsShow"`
}

func main() {
	u1 := AdvertRequest{
		Title:  "xxx",
		Href:   "xxx",
		Images: "xxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
