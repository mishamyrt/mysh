#/usr/bin/env bash

complete -F _go go

_get_mysh_hosts() {
    $(compgen -W "$(cat $HOME/.local/share/mysh/completion)" -- "${COMP_WORDS[1]}")
}

_mysh_completions()
{
  cur="${COMP_WORDS[COMP_CWORD]}"
  case "${COMP_WORDS[COMP_CWORD-1]}" in
    "mysh")
        cmds="get update help remotes namespaces hosts show version"
        hosts="$(cat $HOME/.local/share/mysh/completion)"
        COMPREPLY=($(compgen -W "${cmds} ${hosts}" -- ${cur}))
    ;;
    "show")
        COMPREPLY=($(compgen -W "$(cat $HOME/.local/share/mysh/completion)" -- ${cur}))
    ;;
  esac
  return 0
}

complete -F _mysh_completions mysh
