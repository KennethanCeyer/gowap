<p align="center"><img width="150" src="https://www.pigno.se/static/assets/images/gowap-logo.svg" /></p>
<p align="center">
  <a href="https://github.com/KennethanCeyer/gowap"><img src="https://img.shields.io/github/release/KennethanCeyer/gowap.svg" alt="Github Release"></a>
</p>
<p align="center">
    <a href="https://travis-ci.org/KennethanCeyer/gowap"><img src="https://travis-ci.org/KennethanCeyer/gowap.svg?branch=master" alt="Build Status" /></a>
</p>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/KennethanCeyer/gowap"><img src="https://goreportcard.com/badge/github.com/KennethanCeyer/gowap" alt="Go Report Card"></a>
  <a href="https://codebeat.co/projects/github-com-kennethanceyer-gowap-master"><img src="https://codebeat.co/badges/8225c214-4c2b-49da-98d0-276cbccea9e8" alt="codebeat badge"></a>
<a href="https://app.fossa.io/projects/git%2Bgithub.com%2FKennethanCeyer%2Fgowap?ref=badge_shield" alt="FOSSA Status"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2FKennethanCeyer%2Fgowap.svg?type=shield"/></a>
  <a href="https://codeclimate.com/github/KennethanCeyer/gowap/maintainability"><img src="https://api.codeclimate.com/v1/badges/f8cb6d015910de13f018/maintainability" alt="Maintainability"></a>
  <a href="https://codeclimate.com/github/KennethanCeyer/gowap/test_coverage"><img src="https://api.codeclimate.com/v1/badges/f8cb6d015910de13f018/test_coverage" alt="Test Coverage"></a>
</p>

## :warning: Notice

This project is working in progress project

If you follow the below installation guide, you may not get the intended result

But you can express your interest by clicking star on this project.

The more star in this repository, the more nervous the maintainer will be and will focus more on this project than any other project.

## :clap:for those who are suffering from managing multiple ssh keys

#### the goal of this project is quite simple

- ssh-key changes must be simple and quick
- you can set an alias name for each of the ssh keys
- it could be nice, if this cli can support to register your ssh-key to github by using api

## :triangular_ruler: Blueprint

#### rules

- it must follows [cleancode rule](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29)
- proceed with TDD-based development
- minimize dependence
- support all popular os

#### usage

1. generate ssh key for test

```bash
$ ssh-keygen
```

2. add `~/.ssh/id_rsa` `~/.ssh/id_rsa` as `home`

```bash
$ gowap add home
enter file in which to add the private key (~/.ssh/id_rsa):
enter file in which to add the public key (~/.ssh/id_rsa.pub):
...
your ssh key is added as `home`
```

3. hiding current ssh key for generate new ssh key

```bash
$ gowap archive
ssh key is archived *home -> archive
```

4. generate new ssh key

```bash
$ ssh-keygen
```

5. add `~/.ssh/id_rsa` `~/.ssh/id_rsa` as `company`

```bash
$ gowap add company
enter file in which to add the private key (~/.ssh/id_rsa):
enter file in which to add the public key (~/.ssh/id_rsa.pub):
...
your ssh key is added as 'company'
```

6. list all keys

```bash
$ gowap list
* company
home
```

7. situation: you clone your company's private repository with company ssh-key

```bash
$ git clone git@github.com:nickname/some-awesome-project.git
```

8. you can edit and push because company ssh key has a permission for that

```bash
$ touch some-changes
$ git add some-changes
$ git commit -m "add some-changes"
$ git push origin master
```

9. but you need to update your personal repository

```bash
$ git clone git@github.com:personal-nickname/personal-project.git
```

10. you can't clone this repository bacause company ssh key hasn't any permissions for access

```bash
Permission denied (publickey).
```

11. you can checkout your personal ssh key `home`

```bash
$ gowap checkout home
your now key is 'home'
```

12. you can clone now

```bash
$ git clone git@github.com:personal-nickname/personal-project.git
```

## :package: Installation

#### :yellow_heart: go

```bash
$ go get github.com/KennethanCeyer/gowap
$ gowap -v
NAME:
   gowap - simple ssh swap tool

USAGE:
   gowap.exe [global options] command [command options] [arguments...]

VERSION:
   x.x.x

AUTHOR:
   kenneth ceyer <https://github.com/KennethanCeyer>

COMMANDS:
     init, i    initialize ssh path
     go, g      swap ssh profile
     remove, r  remove ssh profile
     change, c  change ssh profile
     list, l    show list ssh profiles
     search, s  search ssh profile
     help, h    Shows a list of commands or help for one command
     ...

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
   ...
```

#### :broken_heart: linux (still yet not supported)

```bash
$ apt-get install gowap
```

#### :broken_heart: Mac (still yet not supported)

```bash
$ brew install gowap
```

#### :broken_heart: Windows (still yet not supported)

```bash
$ choco install gowap
```

## :mag: License

gowap is under Apache 2.0 license

you can, of course. download it, use it, modify it


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FKennethanCeyer%2Fgowap.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FKennethanCeyer%2Fgowap?ref=badge_large)