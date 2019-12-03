#/usr/bin/env fish

set -l hosts_subcmds show

# version
complete -c mysh -f -n '__fish_use_subcommand' -a version -d 'Display version and exit'
complete -c mysh -f -n '__fish_seen_subcommand_from version'

# remotes
complete -c mysh -f -n '__fish_use_subcommand' -a remotes -d 'Display all added remote repositories'
complete -c mysh -f -n '__fish_seen_subcommand_from remotes'

# update
complete -c mysh -f -n '__fish_use_subcommand' -a update -d 'Update hosts from all added remote repositories'
complete -c mysh -f -n '__fish_seen_subcommand_from update'

# namespaces
complete -c mysh -f -n '__fish_use_subcommand' -a namespaces -d 'Display all namespaces'
complete -c mysh -f -n '__fish_seen_subcommand_from namespaces'

# hosts
complete -c mysh -f -n '__fish_use_subcommand' -a hosts -d 'Display all hosts'
complete -c mysh -f -n '__fish_seen_subcommand_from hosts'

# show
complete -c mysh -f -n '__fish_use_subcommand' -a show -d 'Shows host information'
complete -c mysh -f -n '__fish_seen_subcommand_from show'


complete -c mysh  -n "__fish_seen_subcommand_from $hosts_subcmds" -a (cat $HOME/.local/share/mysh/completion)
complete -c mysh -f -n '__fish_use_subcommand' -a (cat $HOME/.local/share/mysh/completion)
