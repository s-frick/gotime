# Gotime

A cli(maybe later more) to easily track your project times right in your terminal. This project is in a very early stage, things may change quickly.

This project is inspired by [Watson](https://github.com/TailorDev/Watson), which currently has an uncertain future and is not being maintained. However, since I want to use some features of Watson that are a bit buggy at the moment I started to develop my own time tracking CLI in Go. Maybe someone will find it useful.

## What's working

A constantly changing list of things that work or should work in future.

- [x] Start a frame
  - [x] with tags, you'll be asked for confirmation if it's the first time you using a tag
  - [x] with start time, `--at "15:04"`
  - [x] with stopping of current running frame, start time also works here
- [x] Stop a frame
  - [x] with stop time, `--at "15:04"`
- [ ] configuration
  - [ ] tag confirmation
  - [ ] stop on start
  - [ ] allow start/stop times in future
  - [ ] ...
- [ ] Frame Log, as of now you can inspect `~/.gotime/frames`
- [ ] Aggregations
- [ ] Reports

## Install/Usage

Currently there is no release. So you have to build it yourself. Clone this repo and run:

```bash
make build
```

Usage:

```bash
./bin/gotime start [TAGS] [OPTIONS]
./bin/gotime stop [OPTIONS]
```
