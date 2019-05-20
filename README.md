# Building An App with Go Modules and Athens :tada:

Hey Gophers! We're gonna build an awesome webapp with [Gin](https://github.com/gin-gonic/gin). Well, it's pretty basic but it shows cat and dog pictures so it's still pretty awesome :grinning:.

The cool part though? Instead of pulling all my webapp's dependencies directly from version control systems like GitHub (which we've always been doing in the past), we're gonna build it using [Athens](https://docs.gomods.io).

In fact, we're actually gonna do the build in _three_ different ways! And we'll talk about each one.

>I gave this demo at [GopherCon EU 2019](https://www.gophercon.es/#aaron-schlesinger-anchor) in the session titled "The Athens Project - A Proxy Server for Go Modules".

![athens banner](https://d33wubrfki0l68.cloudfront.net/517b1a47a4e75a20cbddd06e936d73837d3c2297/c84eb/banner.png)

# About the Web App

The web application we're going to build is pretty standard. I built a little server using [gin](https://github.com/gin-gonic/gin) as the framework, and it shows some HTML pages with cat pictures on them. 

The cool thing is that the whole codebase is modules-aware. That means it has a `go.mod` file that keeps track of all my code's dependencies. The built-in [Go Modules](https://github.com/golang/go/wiki/Modules) system then uses that file to download all of my dependencies from Athens (or, it could use any other module proxy server!) so it can build my code.

# Run The Demo

Below is how how to do the demo yourself. The instructions are for Linux/Mac OS X systems.

## First Way: Build Using a Demo Athens

Let's first build the code by telling our `go` tool to use a public, demo Athens server:

```console
$ export GOPROXY=https://athens.azurefd.net
```

>Note: the proxy at this URL is experimental and should not be used in production settings. It's not supported by me or anyone else :smile:

### Clear Local Module Cache

The `go` toolchain keeps a local on-disk cache of every module & version it's ever downloaded. That means if we really want to test Athens out when we build, we should empty the cache so that every module is a miss! Here's how to do it:

```console
$ sudo rm -r $GOPATH/pkg/mod
```

### Build & Run The Server!

No need to change your familiar tools to do this:

```console
$ go run .
```

## Second Way: Run Your Own Athens

Since the public URL I used is for demos, you shouldn't rely on it for your production code. That's all good, though - Athens is meant for _you_ to run inside your own organization to serve up your own code. Even private code!

It's pretty easy to run your own Athens too. It's a free (as in beer) open source project and we provide [instructions](https://docs.gomods.io/install) for running the server in a few different configurations. Today, we're going to use [Docker](https://www.docker.com/) to run ours.

First, run this to start Athens up:

```console
$ docker run -p 3000:3000 -e GO_ENV=development gomods/athens:v0.4.0
```

And then to set your environment variable to point to the local server, similar to last time:

```console
$ export GOPROXY=http://localhost:3000
```

Also don't remember to clear your cache just like last time:

```console
$ sudo rm -rf $GOPATH/pkg/mod
```

And then start up the server again too! (Don't forget to shut down the old one):

```console
$ go run .
```

## Third Way: Use Your Athens While Offline :scream:

Like I mentioned in the last step, Athens is meant to run inside of organizations to serve up their private code. To do that, it keeps its own database of _all_ the code you request from it - public and private. That means that if you have an Athens server running and you do a build against it, you can shut down your internet connection and do the same build against it without an internet connection!

>Athens can be used inside of organizations that have completely firewalled off access to the internet

So, let's get started! The first step is the easiest - keep the Athens server from last time running.

Next, clear out your cache again:

```console
$ sudo rm -rf $GOPATH/pkg/mod
```

And then, **shut down your internet connection** :see_no_evil:.

And finally, do the build & run:

```console
$ go run .
```

And you're done!

## Finally

Thanks for following along! Now you're both a Gopher _and_ an Athenian :green_heart:.

If you want to learn more, check out [docs.gomods.io](https://docs.gomods.io)! We'd also love for you to get involved - here are some ways to do so:

- Come star [our repo](https://github.com/gomods/athens)
- Come say hello on the `#athens` channel in the [Gophers Slack](https://invite.slack.golangbridge.org/)
  - This is a great place to come ask for help getting started and ask questions too :smile:
- Come to one of our [developer meetings](https://docs.gomods.io/contributing/community/developer-meetings/). Absolutely everybody is welcome, regardless of experience, background or anything else
- And of course, file [bug reports](https://github.com/gomods/athens/issues/new/choose) or [feature requests](https://github.com/gomods/athens/issues/new/choose) and [contribute code](https://docs.gomods.io/contributing/new/development/)!

# Keep on rockin', Gophers!

![athens gopher](./athens-gopher.png)
