#/usr/bin/env fish

set -l hosts_subcmds show

# get
complete -c mysh -f -n '__fish_use_subcommand' -a get -d 'add repository and download hosts from it'
complete -c mysh -f -n '__fish_seen_subcommand_from get'

# help
complete -c mysh -f -n '__fish_use_subcommand' -a help -d 'print help message and exit'
complete -c mysh -f -n '__fish_seen_subcommand_from help'

# hosts
complete -c mysh -f -n '__fish_use_subcommand' -a hosts -d 'display all hosts'
complete -c mysh -f -n '__fish_seen_subcommand_from hosts'

# namespaces
complete -c mysh -f -n '__fish_use_subcommand' -a namespaces -d 'display all namespaces'
complete -c mysh -f -n '__fish_seen_subcommand_from namespaces'

# remotes
complete -c mysh -f -n '__fish_use_subcommand' -a remotes -d 'display all added remote repositories'
complete -c mysh -f -n '__fish_seen_subcommand_from remotes'

# show
complete -c mysh -f -n '__fish_use_subcommand' -a show -d 'display host information'
complete -c mysh -f -n '__fish_seen_subcommand_from show'

# update
complete -c mysh -f -n '__fish_use_subcommand' -a update -d 'refresh hosts from added remote repositories'
complete -c mysh -f -n '__fish_seen_subcommand_from update'

# version
complete -c mysh -f -n '__fish_use_subcommand' -a version -d 'print Mysh version'
complete -c mysh -f -n '__fish_seen_subcommand_from version'

complete -c mysh  -n "__fish_seen_subcommand_from $hosts_subcmds" -a (cat $HOME/.local/share/mysh/completion)
complete -c mysh -f -n '__fish_use_subcommand' -a (cat $HOME/.local/share/mysh/completion)
