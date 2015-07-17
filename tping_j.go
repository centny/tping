package main

import (
	"fmt"
	"github.com/Centny/gwf/log"
	"github.com/Centny/gwf/tutil"
	"github.com/Centny/gwf/util"
	"time"
)

type Task struct {
	Name  string `m2s:"name",json:"name"`
	Type  string `m2s:"type",json:"type"`
	Host  string `m2s:"host",json:"host"`
	Cmds  string `m2s:"cmds",json:"cmds"`
	Delay int64  `m2s:"delay",json:"delay"`
	Times int    `m2s:"times",json:"times"`
}

func RunJ(in, e string) error {
	var tasks []Task
	err := util.J2Ss_f(in, &tasks)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.Times < 1 {
			task.Times = 1
		}
		switch task.Type {
		case TP_W:
			if len(task.Host) < 1 {
				log.W("run task(%v) by type(%v) err:host is empty ", task.Name, task.Type)
				break
			}
			delay, err := RunW(task.Host, time.Duration(task.Delay)*time.Millisecond, task.Times)
			if err != nil {
				log.E("run task(%v) by type(%v) err:%v", task.Name, task.Type, err.Error())
				break
			}
			var line string
			if task.Delay < 1 {
				line = fmt.Sprintf("%v/%v", delay, delay)
			} else {
				if task.Delay < delay {
					task.Delay = delay
				}
				line = fmt.Sprintf("%v/%v", task.Delay-delay, task.Delay)
			}
			err = tutil.Emma(e, task.Name, "1/1", "1/1", "1/1", line)
			if err != nil {
				return err
			}
		case TP_R:
			if len(task.Cmds) < 1 {
				log.W("run task(%v) by type(%v) err:cmds is empty ", task.Name, task.Type)
				break
			}
			delay, err := RunR(task.Cmds, time.Duration(task.Delay)*time.Millisecond, task.Times)
			if err != nil {
				log.E("run task(%v) by type(%v) err:%v", task.Name, task.Type, err.Error())
				break
			}
			var line string
			if task.Delay < 1 {
				line = fmt.Sprintf("%v/%v", delay, delay)
			} else {
				if task.Delay < delay {
					task.Delay = delay
				}
				line = fmt.Sprintf("%v/%v", task.Delay-delay, task.Delay)
			}
			err = tutil.Emma(e, task.Name, "1/1", "1/1", "1/1", line)
			if err != nil {
				return err
			}
		default:
			log.W("run task(%v) by type(%v) err:unknow type ", task.Name, task.Type)
		}
	}
	return nil
}