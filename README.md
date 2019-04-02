# gssh

gssh is meant to make using the gcloud ssh command easier by eliminating the need to
know the complete name of a given VM and it's zone.

![Preview](terminalizer/preview.gif)

gssh was built for my needs specifically which is why flags such as `--tunnel-through-iap`
are enabled by default. No reason that cannot change if others have an opinion.

## Badges

[![CircleCI](https://img.shields.io/circleci/project/github/regner/gssh.svg?style=for-the-badge)](https://circleci.com/gh/regner/gssh)
[![Codecov branch](https://img.shields.io/codecov/c/github/regner/gssh/master.svg?style=for-the-badge)](https://codecov.io/gh/regner/gssh)
[![Release](https://img.shields.io/github/release/regner/gssh.svg?style=for-the-badge)](https://github.com/regner/gssh/releases/latest)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)

## Installation

You're welcome to download the source and compile from that. There is no brew release
yet, that may come in the future if people want it. You can also download compiled
releases from the [releases](https://github.com/regner/gssh/releases/latest) page.

## TODO

* [ ] Add ability to disable iap tunnel
* [ ] Create a preview video/gif/something to demo gssh
* [ ] Add release to brew
* [ ] Add test coverage for main.go
* [ ] Add test coverage for survey prompts
* [ ] Improve filtering
