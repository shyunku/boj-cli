#!/bin/sh
set -e

# ── helpers ────────────────────────────────────────────────────────────────

append_if_missing() {
  file="$1"
  line="$2"
  grep -qF "$line" "$file" 2>/dev/null || echo "$line" >> "$file"
}

set_editor_env() {
  editor="$1"
  shell_rc="$HOME/.zshrc"
  [ -f "$HOME/.bashrc" ] && shell_rc="$HOME/.bashrc"

  append_if_missing "$shell_rc" "export EDITOR=$editor"
  echo "  → EDITOR=$editor set in $shell_rc"
}

# ── build & install boj ────────────────────────────────────────────────────

echo ""
echo "Building boj..."
go build -o boj .
mv boj "$HOME/go/bin/boj"
echo "Installed: $(which boj 2>/dev/null || echo $HOME/go/bin/boj)"

# ── editor setup ──────────────────────────────────────────────────────────

echo ""
echo "Choose an editor to use with boj:"
echo "  1) vim    (classic, syntax on by default)"
echo "  2) nvim   (modern vim, tree-sitter highlighting)"
echo "  3) hx     (helix — built-in highlighting, no config needed)"
echo "  4) micro  (VSCode-like, Ctrl+S / Ctrl+Q)"
echo "  5) nano   (simple, beginner-friendly)"
echo "  6) skip"
printf "Select (1-6): "
read choice

case "$choice" in
  1)
    echo "Installing vim..."
    pkg install -y vim
    append_if_missing "$HOME/.vimrc" "syntax on"
    append_if_missing "$HOME/.vimrc" "set number"
    append_if_missing "$HOME/.vimrc" "set tabstop=2 shiftwidth=2 expandtab"
    set_editor_env vim
    ;;
  2)
    echo "Installing neovim..."
    pkg install -y neovim
    mkdir -p "$HOME/.config/nvim"
    nvim_init="$HOME/.config/nvim/init.vim"
    append_if_missing "$nvim_init" "syntax on"
    append_if_missing "$nvim_init" "set number"
    append_if_missing "$nvim_init" "set tabstop=2 shiftwidth=2 expandtab"
    set_editor_env nvim
    ;;
  3)
    echo "Installing helix..."
    pkg install -y helix
    set_editor_env hx
    echo "  → syntax highlighting is built-in, no config needed"
    ;;
  4)
    echo "Installing micro..."
    pkg install -y micro
    set_editor_env micro
    ;;
  5)
    echo "Installing nano..."
    pkg install -y nano
    set_editor_env nano
    ;;
  6|*)
    echo "Skipping editor setup."
    ;;
esac

# ── done ──────────────────────────────────────────────────────────────────

echo ""
echo "Done! Run: boj --help"
echo "To apply EDITOR setting now: source ~/.zshrc"
