twitterbot
====================

Twitterbot is a wrapper around the Anaconda Go package for accessing the Twitter API.

### Config

After generating the appropriate tokens through the Twitter developer portal copy them into a local file somewhere that follows the format in sample/sample_config.

### Creating a Bot

All bots are created by passing in the location of the keys config to `New`.

````go
bot, err := twitterbot.New("path/to/config")
````

Once run, the bot will loop forever posting and sleeping for 10 minutes. The content of each tweet is defined by the argument to `Run`, which is a struct with a `NextTweet` method. The output of this method is automatically de-duped within `Run`, so the same content is not posted consecutively.

A very simple example that just tries to post "Hello" over and over again would look like:

````go

// Read config location and create bot above...

type Hello struct {}

func (h Hello) NextTweet() string {
    return "Hello"
}

tweetMaker := Hello{}
bot.Run(tweetMaker)
````

An example of a full script that tweets out the top post on the Reddit front page can be found in samples.
