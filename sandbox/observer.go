// Copyright 2014 The gocui Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "log"
    "fmt"
    "github.com/jroimartin/gocui"
    "time"
)

func layout(g *gocui.Gui) error {
    maxX, _ := g.Size()

    if sv, err := g.SetView("status", -1, -1, maxX, 1); err != nil {
        if err != gocui.ErrorUnkView {
            return err
        }
        fmt.Fprintln(sv, time.Now().Local())
    }
//     if sv, err := g.SetView("sam", -1, -1, maxX, 1); err != nil {
//        if err != gocui.ErrorUnkView {
//            return err
//        }
//        fmt.Fprintln(sv, "hi Sam")
//    }

    return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.Quit
}

//time is updated constantly.
// we print to view constantly
func updateViews(g *gocui.Gui) {
    for {

//        time.Sleep(500 * time.Millisecond)
//        statusView,_ := g.View("status")
//        statusView.Clear()
//        fmt.Fprintln(statusView, time.Now().Local())
////        g.SetCurrentView(statusView.Name())
//        g.Flush()
////
////
////
////        g.SetCurrentView("sam")
////        g.Flush()
	time.Sleep(300 * time.Millisecond)
    if sv,_ := g.View("status"); sv != nil {
            sv.Clear()

            fmt.Fprintln(sv, time.Now().Local())

//            if err := g.Flush(); err != nil {
//                return
//            }
        }
    }
}
//we only flush to screen from time to time
func flushAction(g *gocui.Gui){
	for {
		time.Sleep(1000 * time.Millisecond)
		if err := g.Flush(); err != nil {
			return
		}
	}
}


func main() {
    var err error

    g := gocui.NewGui()
    if err := g.Init(); err != nil {
        log.Panicln(err)
    }
    defer g.Close()

    g.SetLayout(layout)
    if err := g.SetKeybinding("", gocui.KeyCtrlC, 0, quit); err != nil {
        log.Panicln(err)
    }

    go updateViews(g)
	go flushAction(g)

    err = g.MainLoop()
    if err != nil && err != gocui.Quit {
        log.Panicln(err)
    }
}