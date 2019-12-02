#/usr/bin/env bash
_mysh_completions()
{
  COMPREPLY=($(compgen -W "$(cat $HOME/.local/share/mysh/completion)" -- "${COMP_WORDS[1]}"))
}

complete -F _mysh_completions mysh
