package main

import (
	"io"
// "os"
	"fmt"
	"flag"
	"os/exec"
// "os/signal"
	"github.com/jroimartin/gocui"
// "strings"
// "reflect"
	"time"
)

var (
	cmds = []string{"main", "b", "c", "d"}
//	views = []*gocui.View{}
	curViewName = cmds[0]
	output = map[string]string{}
	numb = 0
)

func nextView(g *gocui.Gui, v *gocui.View) error {

	//handle nil case. handle end of list case.
//	if curViewName == "b"{
//		numb = numb +1
//	}
//	if numb == 2 {
//		panic("being called when first called")
//	}
//	panic(v)
	if v.Name() == cmds[len(cmds)-1] {
//		panic("this is called")
		curViewName = cmds[0]

	} else {
		for i := range cmds {

			if v.Name() == cmds[i] {
				curViewName = cmds[i+1]

			}

		}
	}
//	b,_ := g.View("b")
//	panic(b.Name())

	return g.SetCurrentView(curViewName)
}

func previousView(g *gocui.Gui, v *gocui.View) error {

	//handle nil case. handle end of list case.
	if v.Name() == cmds[0] {
		curViewName = cmds[len(cmds)-1]
	} else {
		for i := range cmds {

			if v.Name() == cmds[i] {

				curViewName = cmds[i-1]

//				if temp, _ := g.View(curViewName); temp !=nil{
//					g.SetCurrentView(temp.Name())
////					panic(g.CurrentView().Name())
//				}

			}

		}
	}
//breaks here

//	return g.SetCurrentView(curViewName)
	return g.SetCurrentView(curViewName)
}

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, previousView); err != nil {
		return err
	}



	return nil
}
type Cmd struct {
	command exec.Cmd
	stdin   io.WriteCloser
	stderr  io.ReadCloser
	stdout  io.ReadCloser
}

// func parseCommands(commands []string) []Cmd {
//   cmds := make([]Cmd, len(commands))

//   for i, cmd := range commands {
//     fmt.Println(reflect.TypeOf(cmd))
//     cmdFull := strings.split(cmd," ")
//     cmdName := cmdFull[0]
//     cmdArgs := cmdFull[1:]

//     command := exec.Command(cmdName, cmdArgs)

//     stdin, err := command.StdinPipe()
//     if err != nil {
//       panic(err)
//     }
//     stdout, err := command.StdoutPipe()
//     if err != nil {
//       panic(err)
//     }
//     stderr, err := command.StderrPipe()
//     if err != nil {
//       panic(err)
//     }

//     cmds[i] = Cmd {
//       command: *command,
//       stdin: stdin,
//       stderr: stderr,
//       stdout: stdout,
//     }

//     command.Start()
//   }

//   return cmds
// }

func layout(g *gocui.Gui) error {
	//we know that curView is set
//	if curViewName == "b"{
//		panic(curViewName)
//	}


	// MUST create base views and initialize all of them here because
	// layout is called from key binding functions such as next and previous
	// which require other views (views other than the current view) to exist
	maxX, maxY := g.Size()
	for _, cmd := range cmds {
		if v, err := g.SetView(cmd, -1, -1, maxX, maxY); err != nil {
			if err != gocui.ErrorUnkView {
				return err
			}
			v.Autoscroll = true
			fmt.Fprintln(v, cmd)
			//todo: determine other view options


		}

	}

	g.SetCurrentView(curViewName)




	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}




//func createViews(commands []string, g *gocui.Gui) error {
//
//	maxX, maxY := g.Size()
//
//	for _, cmd := range cmds {
//
//		if v, err := g.SetView(cmd, -1, -1, maxX, maxY); err != nil {
//			if err != gocui.ErrorUnkView {
//				return err
//			}
//			v.Autoscroll=true
////			fmt.Fprintf(v, "%s", cmd)
//			//todo: determine other view options
//
//			views = append(views, v)
//
//			curView = v
//
//		}
//	}
//	//We know that all the views have been created
////	names := []string{}
////	for _,v := range views{
////		names = append(names, v.Name())
////	}
//
//	//know that generated views are inside g.views list.
////	my_view,_ :=g.View(names[3])
////	panic(my_view.Name())
//
//	return nil
//}


func updateViews(g *gocui.Gui) {
    for {

		time.Sleep(1000 * time.Millisecond)

		if sv,_ := g.View(curViewName); sv != nil {
            sv.Clear()

			fmt.Fprintln(sv, curViewName)

			curOutput := time.Now().Local().String()
			if previousVal, present := output[curViewName]; present {
				output[curViewName] = previousVal + "\n" + curOutput
			}else{
				output[curViewName] = curOutput
			}
            fmt.Fprintln(sv, output[curViewName])


			if err := g.Flush(); err != nil {
                return
            }

        }
    }
}


func main() {
	// sigchan := make(chan os.Signal, 1)
	// x := make(chan os.Signal, 1)
	// signal.Notify(x, os.Interrupt)
	// go func() {
	//   for _ = range sigchan {
	//     // <-x
	//     fmt.Println("Got interrupt")
	//   }
	// }()
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Given 0 commands to run.")
		// return
	}


	// cmds := parseCommands(flag.Args())
	// fmt.Println(cmds)

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		panic(err)
	}
	defer g.Close()


	g.BgColor = gocui.ColorBlack

//	createViews(cmds, g)

	//know that view is still inside g.views list even when we are out of function
//	my_view,_ := g.View("d")
//	panic(my_view.Name())

	g.SetLayout(layout)

	go updateViews(g)

	if err := initKeybindings(g); err != nil {
		panic(err)
	}

	err := g.MainLoop()
	if err != nil && err != gocui.Quit {
		panic(err)
	}

	// go io.Copy(stdin, os.Stdin)
	// go io.Copy(os.Stdout, stdout)
	// go io.Copy(os.Stderr, stderr)

	// for _, cmd := range cmds {
	//   if err := cmd.command.Wait(); err != nil {
	//     panic(err)
	//   }
}