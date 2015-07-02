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
)

var (
  cmds = []string{"main", "b", "c", "d"}
  curCmd = "b"
)
func nextView(g *gocui.Gui, v *gocui.View) error {
  //handle nil case. handle end of list case.
  if v == nil || v.Name() == cmds[len(cmds)-1]{
    curCmd = cmds[0]
  } else {
    for i := range cmds{
      
      if v.Name() == cmds[i]{
        curCmd = cmds[i+1]    
      }
  
    }
  }
  fmt.Println(curCmd)
  //should never be called.
  return g.SetCurrentView(curCmd) 
}

func previousView(g *gocui.Gui, v *gocui.View) error {
  //handle nil case. handle end of list case.
  if v == nil || v.Name() == cmds[0]{
    curCmd = cmds[len(cmds)-1]
  } else {
    for i := range cmds{
    
      if v.Name() == cmds[i]{
        curCmd = cmds[i-1]    
      }
  
    }
  }
  
  //should never be called.
  return g.SetCurrentView(curCmd) 
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
  stdin io.WriteCloser
  stderr io.ReadCloser
  stdout io.ReadCloser
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
  // maxX, maxY := g.Size()
  
  // for i := range cmds {
  //   if v, err := g.SetView(cmds[i], 0, 0, maxX, maxY); err != nil {
  //     if err != gocui.ErrorUnkView {
  //       return err
  //     }
     
      
      
  //     v.Wrap = true
  //     v.Autoscroll = true
      
  //   }
  // }
  
  
  
  return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
  return gocui.Quit
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
  
  cmds = append(cmds,"alpha")
  
  // cmds := parseCommands(flag.Args())
  // fmt.Println(cmds)
  g := gocui.NewGui()
  if err := g.Init(); err != nil {
    panic(err)
  }
  defer g.Close()


  g.BgColor = gocui.ColorBlack
  maxX, maxY := g.Size()
  for i := range cmds {
    if v, err := g.SetView(cmds[i], 0, 0, maxX, maxY); err != nil {
      if err != gocui.ErrorUnkView {
        fmt.Println(err)
      }
     
      
      
      v.Wrap = true
      v.Autoscroll = true
      
    }
  }
  
  g.SetLayout(layout)
  


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