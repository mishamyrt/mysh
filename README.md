# Mysh [![Build status][actions-badge]][actions]  [![Go Report result][goreport-badge]][goreport] 

<img align="right" width="110" height="122"
     alt="Mouse, logo of Mysh"
     src="https://mishamyrt.github.io/mysh/logo@2x.png">

Mys(s)h — wrapper over SSH, which helps not to clog your head with unnecessary things. In Mysh, you can specify a remote repository with SSH hosts and connect to it by knowing only the name. If the host address changes, you don't have to edit the configuration manually, just update from the repository.

## Features

* Hosts repository
* SSH compatible syntax (partly)
* Local aliases

## Installation

### Ubuntu/Debian

```sh
wget https://github.com/mishamyrt/mysh/releases/download/v0.1.0beta/mysh_0.1-0beta-amd64.deb
sudo apt install ./mysh_0.1-0beta-amd64.deb
```

### macOS

```sh
brew tap mishamyrt/mysh
brew install mysh
```

## How to use

First of all, add the host repository to the mysh.

```sh
$ mysh get https://yourcompany.com/hosts/mysh.yaml
'yourcompany' config successfully added
```

Now use the `hosts` command to see a list of all available hosts.

```sh
$ mysh hosts
- yourcompany:mercury
- yourcompany:may
- yourcompany:deacon
```

Edit `~/.local/share/mysh/global.yaml` to define username for userless hosts.

```
namespaces:
  yourcompany:
    user: mishamyrt
```

As you can see, all the hosts are prefixed with the namespace. You can enter a hostname with or without a namespace. Usually you need a namespace to avoid collisions.

```sh
$ mysh mercury
freddie@mercury:~# 
```

There is `show` command, that prints information about the host.

```sh
$ mysh show mercury
Host: 10.10.9.5
User: freddie
```

To update all your remotes you can use `update` command.

```sh
$ mysh update
'yourcompany' is updated
'petproject' is updated
```

And now my favorite: you can put the `.mysh.yaml` file in the project folder, prescribe host aliases there.

```yaml
aliases:
    test: mercury
```

Now, being in the folder with this file, you can connect to the host using the alias.

```sh
$ mysh test
freddie@mercury:~# 
```

[actions-badge]:  https://github.com/mishamyrt/mysh/workflows/build/badge.svg
[actions]:        https://github.com/mishamyrt/mysh/actions?query=workflow%3A%22build%22
[goreport-badge]: https://goreportcard.com/badge/github.com/mishamyrt/mysh
[goreport]:       https://goreportcard.com/report/github.com/mishamyrt/mysh
