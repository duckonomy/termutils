package main

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
)

func main() {
  args := os.Args[1:]
  if len(args) == 0 {
    fmt.Println(getProjectRoot(""))
  } else {
    fmt.Println(getProjectRoot(args[0]))
  }
}

func getProjectRoot(path string) string {
  wd, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  if path != "" {
    wd = path
  }

  home, err := os.UserHomeDir()
  if err != nil {
    panic(err)
  }

  for {
    parent := filepath.Dir(wd)

    if wd == "/" {
      break 
    }

    if wd == home {
      break
    }

    var fileList = listDirectory(wd)

    for  _, file := range fileList {
      if isProjectType(file.Name()) {
        return wd;
      }
    }

    wd = parent
  }

  return wd
}

func listDirectory(path string) []os.FileInfo {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  files, err := f.Readdir(-1)
  f.Close()
  if err != nil {
    log.Fatal(err)
  }

  return files
}

func isProjectType(path string) bool {
  switch path {
  case ".git":
    return true
  case ".ccls":
    return true
  case "Cargo.toml":
    return true
  case "stack.yaml":
    return true
  case "DESCRIPTION":
    return true
  case "Eldev":
    return true
  case "Cask":
    return true
  case "shard.yml":
    return true
  case "Gemfile":
    return true
  case ".bloop":
    return true
  case "deps.edn":
    return true
  case "build.boot":
    return true
  case "project.clj":
    return true
  case "build.sc":
    return true
  case "build.sbt":
    return true
  case "application.properties":
    return true
  case "gradlew":
    return true
  case "build.gradle":
    return true
  case "pom.xml":
    return true
  case "poetry.lock":
    return true
  case "info.rkt":
    return true
  case "pubspec.yaml":
    return true
  case "dune-project":
    return true
  case "Pipfile":
    return true
  case "tox.ini":
    return true
  case "setup.py":
    return true
  case "requirements.txt":
    return true
  case "manage.py":
    return true
  case "angular.json":
    return true
  case "package.json":
    return true
  case "gulpfile.js":
    return true
  case "Gruntfile.js":
    return true
  case "mix.exs":
    return true
  case "rebar.config":
    return true
  case "composer.json":
    return true
  case "CMakeLists.txt":
    return true
  case "Makefile":
    return true
  case "debian/control":
    return true
  case "default.nix":
    return true
  case "meson.build":
    return true
  case "SConstruct":
    return true
  case "go.mod":
    return true
  default:
    return false
  }
}
