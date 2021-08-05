package main

import (
	"bytes"
  "fmt"
  "log"
  "io/ioutil"
  "os"
  "os/exec"
  "strings"
)

const fzfCommandFile = "/tmp/fzf-cmd-file"
const fzfTargetFile = "/tmp/fzf-target-file"
const fzfTempFile = "/tmp/fzf-temp-file"

func main() {
  // TODO Should parse cli args and config instead of const editor
  // Should fallback to env if not provided
  // Also should indicate path or default is current dir
  editor := os.Getenv("EDITOR")

  if editor == "" {
    log.Fatal("No editor selected!")
  } 

  execFzf()
}

func execFzf() {
  // TODO Replace with dir supplied as arg (through prj)
  wd, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  fzfCommand := "fzf --layout=reverse --bind ctrl-d:page-down,ctrl-u:page-up,ctrl-k:kill-line > " + fzfCommandFile
  cmd := exec.Command("tmux", "popup", "-d", wd, "-xC", "-yC", "-E", fzfCommand)
    
  cmd.Stdout = os.Stdout
  cmd.Run()

  content, err := ioutil.ReadFile(fzfCommandFile)

  if err != nil {
    log.Fatal(err)
  }

  runCmd := string(content)

  execTmux(strings.TrimSpace(runCmd))
  // fmt.Println(strings.TrimSpace(runCmd))
}

func cleanUp() {
  e := os.Remove(fzfCommandFile)
  if e != nil {
    log.Fatal(e)
  }
}

func execTmux(path string) {
  launchCmd := exec.Command("tmux", "display-message", "-p", "'#{session_name}:#{window_index}.#{pane_index}'")

  var out bytes.Buffer
  launchCmd.Stdout = &out
  err := launchCmd.Run()
  if err != nil {
    log.Fatal(err)
  }

  cmd := exec.Command("tmux", "respawn-pane", "-k", "nvim " + path + "; zsh")
  cmd.Run()
}

func getFile(root string) {

}

func execEditor(editor string, path string) {
  cmd := exec.Command(editor, path)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  err := cmd.Run()
  fmt.Println(err)
}
