# prman - bot to inform pull requests to slack created for MMMcorp

prman is slack slash command written in Go.
prman gives you information of pull request that is not WIP && review requested to you.

## Table of Contents

* [Installation](#installation)
* [Usage](#usage)
* [License](#license)

## Installation


```
$ go get github.com/ygnmhdtt/prman
```

### Setup

### 1. Create Slash command for your workspace

Create Slash command . [reference](https://api.slack.com/slash-commands).

Configs:

* `command name` : as you like
* `URL` : `your_server:7000`
* `Method` : `Post`
* `Name` `Icon` : as you like
* `escape` : `off`

And, `token` must be exported following configuration.

### 2. Environment Variables

prman requires 2 environment variables.

```
PR_GITHUB_TOKEN="123412341234xxxxyyyyzzzz"
PR_GITHUB_ORGANIZATION="exampleorg"
PR_VALID_TOKEN="11112222hhhhkkkk"
```

`PR_VALID_TOKEN` is token slack generated.

You can create `PR_GITHUB_TOKEN` at  [Personal API tokens](https://github.com/blog/1509-personal-api-tokens).
If you want to get information of private repositories, token must have permission to get them.

### 3. create json file

Create `prman-members.json` .
It must be placed at same directory as prman binary file.

Sample:

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

Then, please run prman on your server like:

```
$ $GOPATH/bin/prman &
```

## Usage

Get pull requests that requested review to you:

```
/gp
```

Get pull requests that requested review to `user`:

```
/gp user
```

`user` must be slack username, **not** github username.

## FAQ

* How prman decide whether pull request is wip?

If title of pr is starts with `WIP` or `(WIP)` or `[WIP]` or `【WIP】` , prman will think it is wip.

## License
MIT
