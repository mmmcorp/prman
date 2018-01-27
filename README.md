# prman - inform pull requests to slack for MMM

## Table of Contents

* [Installation](#installation)
* [Usage](#usage)
* [License](#license)

## Installation

This library doesn't depends on any other packages.

```
$ go get github.com/ygnmhdtt/prman
```

## Usage

You need 3 configurations.

### 1. Environment Variables

prman requires 2 environment variables.

```
PR_GITHUB_TOKEN="123412341234xxxxyyyyzzzz"
PR_GITHUB_ORGANIZATION="exampleorg"
```

You can create `PR_GITHUB_TOKEN` by [Personal API tokens](https://github.com/blog/1509-personal-api-tokens).

### 2. create json file

You need to create `prman-members.json` .
It must be placed same directory as prman binary file.

Here is sample:

```
{
  "members": [
    "yagi:ygnmhdtt",
    "sla:git",
    "daru:farid"
  ]
}
```

left side of `:` is slack username. The other is github username.

When 1 and 2 config has done, please run prman on your server like:

```
$ $GOPATH/bin/prman &
```

### 3. Create Slash command for your workspace

Create Slash command by [here](https://api.slack.com/slash-commands).

Here is configs:

* `command name` : as you like
* `URL` : your_server:7000
* `Method` : `Post`
* `Name` `Icon` : as you like
* `escape` : `off`

## License
MIT
