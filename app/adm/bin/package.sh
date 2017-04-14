#!/bin/sh

cd /app/gopath/src/gopkg.in/goyy/goyy.v0 && svn update
cd /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm && go build adm.go

mkdir -p /app/webapps/gyadm/bin
mkdir -p /app/assets/gyadm
mkdir -p /app/assets/hndev/ui

rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/sys/static/adm/ /app/assets/gyadm/sys
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/static/ /app/assets/gyadm
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/ui/static/ /app/assets/hndev/ui

rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/sys/templates/adm/ /app/webapps/gyadm/templates/sys
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/ui/templates/ /app/webapps/gyadm/templates/ui
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/templates/ /app/webapps/gyadm/templates

rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/adm  /app/webapps/gyadm
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/conf /app/webapps/gyadm

rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/bin/startup.sh /app/webapps/gyadm/bin
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/bin/shutdown.sh /app/webapps/gyadm/bin
rsync --exclude=.svn -r /app/gopath/src/gopkg.in/goyy/goyy.v0/app/adm/bin/restart.sh /app/webapps/gyadm/bin
