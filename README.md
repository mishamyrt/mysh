<p align="center">
    <img src="./docs/logo@2x.png" width="208px">
</p>

Dive — wrapper over SSH, which helps not to clog your head with unnecessary things. In Dive, you can specify a remote repository with SSH hosts and connect to it by knowing only the name. If the host address changes, you don't have to edit the configuration manually, just update from the repository.

<img src="https://github.com/mishamyrt/dive/workflows/Build%20binaries/badge.svg" alt="Build status">

## Features

* Hosts repository
* SSH compatible syntax (partly)
* Local aliases

## How to use

First of all, add the host repository to the dive.

```sh
$ dive get https://yourcompany.com/hosts/dive.yaml
'yourcompany' config successfully added
```

Now use the `hosts` command to see a list of all available hosts.

```sh
$ dive hosts
- yourcompany:mercury
- yourcompany:may
- yourcompany:deacon
```

As you can see, all the hosts are prefixed with the namespace. You can enter a hostname with or without a namespace. Usually you need a namespace to avoid collisions.

```sh
$ dive mercury
freddie@mercury:~# 
```

To update all your remotes you can use `update` command.

```sh
$ dive update
'yourcompany' is updated
'petproject' is updated
```

And now my favorite: you can put the `.dive.yaml` file in the project folder, prescribe host aliases there.

```yaml
aliases:
    test: mercury
```

Now, being in the folder with this file, you can connect to the host using the alias.

```sh
$ dive test
freddie@mercury:~# 
```
