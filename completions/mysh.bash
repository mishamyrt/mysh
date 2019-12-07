#/usr/bin/env bash

complete -F _go go

_get_mysh_hosts() {
    hosts="$(cat $HOME/.local/share/mysh/completion 2> /dev/null)"
    if [ -z "$hosts" ]
    then
        echo ""
    else
        echo $hosts
    fi
}

_mysh_completions()
{
  cur="${COMP_WORDS[COMP_CWORD]}"
  case "${COMP_WORDS[COMP_CWORD-1]}" in
    "mysh")
        cmds="get update help remotes namespaces hosts show version"
        hosts="$(_get_mysh_hosts)"
        COMPREPLY=($(compgen -W "${cmds} ${hosts}" -- ${cur}))
    ;;
    "show")
        COMPREPLY=($(compgen -W "$(_get_mysh_hosts)" -- ${cur}))
    ;;
  esac
  return 0
}

complete -F _mysh_completions mysh
