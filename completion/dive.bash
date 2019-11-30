#/usr/bin/env bash
_dive_completions()
{
    
  COMPREPLY=($(compgen -W "$(cat $HOME/.local/share/dive/completion)" -- "${COMP_WORDS[1]}"))
}

complete -F _dive_completions dive
