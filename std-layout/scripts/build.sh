#!/usr/bin/env bash
# 为go的二进制包注入版本信息
VersionDir=cmd
# gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
BuildDate=$(date +'%Y-%m-%d_%H:%M:%S(ts:%s)')
CurBranch=$(git symbolic-ref --short -q HEAD) # 在哪个分支上编译的
GitCommit=$(git log --pretty=format:'%H' -n 1)
LastCommitter=$(git log -1 --pretty=format:'%an') # 最后一次提交的作者

# go build flags
ldflags="-X $VersionDir.buildDate=$BuildDate -X $VersionDir.curBranch=$CurBranch -X $VersionDir.gitCommit=$GitCommit -X $VersionDir.lastCommitter=$LastCommitter"

export GOOS=linux && cd cmd && go build -ldflags "$ldflags"
