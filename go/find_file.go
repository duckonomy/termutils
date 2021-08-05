package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  // "path/filepath"
)

func main() {
  // TODO Should parse cli args and config instead of const editor
  // Should fallback to env if not provided
  editor := os.Getenv("EDITOR")

  if editor == "" {
    log.Fatal("No editor selected!")
  } 

  execEditor(editor, "~/cools")
}

func fromTmux(cmd string) {
  if [ "$(tmux display-message -p -F "#{session_name}")" = "popup" ];then
    tmux detach-client
  else
    tmux popup -d '#{pane_current_path}' -xC -yC -w$width -h$height -E "eval $me_executable"

}

func launchTmuxSplit() {

}

func launchTmuxReplace() {

}

// func getFile(root string) {

// }

func execEditor(editor string, path string) {
  cmd := exec.Command(editor, path)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  err := cmd.Run()
  fmt.Println(err)
}

// func execFzf() {

// }

// func outFifo() {

// }

// func inFifo() {

// }

// func mkFifo() {

// }
