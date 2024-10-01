# traefik-crawler-user-agents

This repository contains a list of HTTP user-agents used by robots, crawlers, and spiders as in single JSON file.

Adapted for Yaegi interpreter and able to use in Traefik plugins.

Forked from Go package: <https://pkg.go.dev/github.com/monperrus/crawler-user-agents>

Each `pattern` is a regular expression. It should work out-of-the-box wih your favorite regex library.

If you use this project in a commercial product, [please sponsor it](https://github.com/sponsors/monperrus).

## Install

### Direct download

Download the [`crawler-user-agents.json` file](https://raw.githubusercontent.com/monperrus/crawler-user-agents/master/crawler-user-agents.json) from this repository directly.

### Go

Go: use [this package](https://pkg.go.dev/github.com/monperrus/crawler-user-agents),
it provides global variable `Crawlers` (it is synchronized with `crawler-user-agents.json`),
functions `IsCrawler` and `MatchingCrawlers`.

Example of Go program:

```go
package main

import (
  "fmt"

  "github.com/stape-io/traefik-crawler-user-agents"
)

func main() {
  userAgent := "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)"

  isCrawler := agents.IsCrawler(userAgent)
  fmt.Println("isCrawler:", isCrawler)

  indices := agents.MatchingCrawlers(userAgent)
  fmt.Println("crawlers' indices:", indices)
  fmt.Println("crawler's URL:", agents.Crawlers[indices[0]].URL)
}
```

Output:

```
isCrawler: true
crawlers' indices: [237]
crawler' URL: https://discordapp.com
```

## Contributing

I do welcome additions contributed as pull requests.

The pull requests should:

* contain a single addition
* specify a discriminant relevant syntactic fragment (for example "totobot" and not "Mozilla/5 totobot v20131212.alpha1")
* contain the pattern (generic regular expression), the discovery date (year/month/day) and the official url of the robot
* result in a valid JSON file (don't forget the comma between items)

Example:

    {
      "pattern": "rogerbot",
      "addition_date": "2014/02/28",
      "url": "http://moz.com/help/pro/what-is-rogerbot-",
      "instances" : ["rogerbot/2.3 example UA"]
    }

## License

The list is under a [MIT License](https://opensource.org/licenses/MIT). The versions prior to Nov 7, 2016 were under a [CC-SA](http://creativecommons.org/licenses/by-sa/3.0/) license.

## Related work

There are a few wrapper libraries that use this data to detect bots:

* [Voight-Kampff](https://github.com/biola/Voight-Kampff) (Ruby)
* [isbot](https://github.com/Hentioe/isbot) (Ruby)
* [crawlers](https://github.com/Olical/crawlers) (Clojure)
* [isBot](https://github.com/omrilotan/isbot) (Node.JS)

Other systems for spotting robots, crawlers, and spiders that you may want to consider are:

* [Crawler-Detect](https://github.com/JayBizzle/Crawler-Detect) (PHP)
* [BrowserDetector](https://github.com/mimmi20/BrowserDetector) (PHP)
* [browscap](https://github.com/browscap/browscap) (JSON files)
