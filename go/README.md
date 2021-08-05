# Termutils

## Launch
- It should be unix-style in design
- Generic launcher extender using fuzzy selectors
- It should be like 
  - `find_file | launch --menu='fzf' --run='nvim' --system='tmux' --pattern='recurse --base='$PATH' | fuzzy'`
  - `grep_file | launch --menu='fzf' --run='nvim'`
- It should chain commands but how?
- Inspired by Clap
- Config file (Maybe)
- Supports various terminal launchers
 - Skim
 - Fzf
 - Fzy
 - Selecta
 - Pick

## Vim Plugin for launch
- Supports various vim functionality
 - Buffers
 - Windows
 - Lsp
 - Coc
 - Treesitter
 - Lines

## Git stuff for launch

## Notmuch stuff for launch

## MPD stuff for launch

## Global File Buffer
- FIFO buffer manager for editors.
- Can refresh and kill.
- `set nobuflisted` for every vim file

## Clip Port
- Clipboard that supports various ports as registers
- Direct support for vim/nvim (via clip_port.vim/clip_port.lua) as a means for inter-editor instance sharing
- Also in sync with unnamed or global clipboard

## Prj
- Project root finder inspired by projectile

## TermDap (like vimspector)
- Generic Dap Client utilizing IPC for various debug operations
- Creates debug instances and allows you to connect to them with various window operations
- Works with any `$EDITOR` that supports jumping and IPC (separate editor plugins for this)
 - Things like Nano might have to restart
 - Vim
 - Emacs

## Tmux History
- Keeps track of ordering and provide "modern" Ctrl-tab behavior that makes sense
