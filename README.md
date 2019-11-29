# Remoty

Simple application that allows you to control your computer remotely, currently we will receive commands in order to controls **internet download manager** but soon we will support more commands.

## Install

Installation process is a straight forward `go install`.

- server: `go install ./cmd/server`
- cli: `go install ./cmd/remoty`

## Server

Remoty server is kinda a deamon that should run on your host machine and will listen for incoming commands

**Environment variables**:

- `RT_PORT`: server port (`1995`)
- `RT_IDM`: path to **internet download manager** exe (`c:/IDMan.exe`)

## Cli

Simple command line interface to interact with remoty server

**Environment variables**:

- `RT_ADDR`: remoty server address (`localhost:1995`)
