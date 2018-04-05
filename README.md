# Go wrapper for the official PUBG API

[![GitHub license](https://img.shields.io/github/license/driquet/gopubg.svg)](https://github.com/driquet/gopubg)
[![Build Status](https://travis-ci.org/driquet/gopubg.svg?branch=master)](https://travis-ci.org/driquet/gopubg)

The goal of this project is to wrap the official PUBG API ([learn more about
the API](https://documentation.playbattlegrounds.com/en/introduction.html). To
use this wrapper, you need an API key.

## Installation

To install the wrapper:
```
go get -u github.com/driquet/gopubg
```

## Usage

### Create an API instance

```
import "github.com/driquet/gopubg"

api := gopubg.NewAPI("<your key here>")
```

### Matches
### Players
### Telemetry
