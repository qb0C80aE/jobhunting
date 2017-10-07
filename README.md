# Where is README_ja.md?

[Don't worry, sir great Google Translate helps you.](https://translate.googleusercontent.com/translate_c?act=url&depth=1&hl=ja&ie=UTF8&prev=_t&rurl=translate.google.co.jp&sl=en&sp=nmt4&tl=ja&u=https://github.com/qb0C80aE/jobhunting&usg=ALkJrhj_Zh2lSf55CLPo6nYAcc89YVt0EA)

# What's this?

Recently, I've heard that now recruiters are picking up skilled engineers by checking an activity graph as known as a green lawn on their GitHub profile page. So that I guess that if we write a job hunting message on the lawn as a canvas, recruiters will think that engineer is so cool, and contact you immediately.

Yes, just like this.

![what?](https://github.com/qb0C80aE/jobhunting/raw/master/example.png)

# How can I use this?

## Install

First, you need to install `git`. And then, download from the links below.

* [linux/amd64](https://github.com/qb0C80aE/jobhunting/releases/download/0.0.1/jobhunting-linux-amd64.tgz)
* [linux/386](https://github.com/qb0C80aE/jobhunting/releases/download/0.0.1/jobhunting-linux-386.tgz)
* [windows/amd64](https://github.com/qb0C80aE/jobhunting/releases/download/0.0.1/jobhunting-windows-amd64.zip)
* [windows/386](https://github.com/qb0C80aE/jobhunting/releases/download/0.0.1/jobhunting-windows-386.zip)

Please don't forget put these into the PATH enabled directory.  

FYI, I don't know how I can build the binaries for MacOS on Linux, for now. If you know, please tell me that.

## Create a new repository on your GitHub account

Go to `github.com` and create a repository used to draw the text. Any name is good for that.

## Clone the repository

```
$ git clone <your repo url>
```

Make sure that git config user.name and user.email are valid, and the current branch is the default branch of this repository.

## Create some files into the repository directory

* a file used to draw the text, the default is grass.txt
  * a text contained in this file must be expressed in 50x7 cells, using 0 and 1. 1 indicated the foreground value, and 0 is the background one.
* a file used to put commit messages. The default is message.txt
  * the file must contain at least one message.

See `githib.com/qb0C80aE/jobhunting/grass.txt` and `githib.com/qb0C80aE/jobhunting/message.txt` as samples.

## Just execute jobhunting in the repository directory

```
$ jobhunting
```

If you have already worked on GitHub and contributed to something, the lawn will be normalized.
In this case, you can use `-s` option to emphasize your text by committing given times.

```
$ jobhunting -s 50
```

It will commit 50 times per cell.

## Just push

```
$ git push origin master
```

Now check your GitHub profile.

# Note

## How can I change the job hunting message?

Just in case, I've prepared a special service for you.
Try this:

```
$ curl "https://texttobinary.herokuapp.com/proxyart?bg=0&fg=1&size=10&text=JOBHUNTING"
```

Then you can get an output like below.

```
00000000000000000000000000000000000000000000000000
00000000000000000000000000000000000000000000000000
00010011001110010010100101001011111001001001001100
00010100101001010010100101101000100001001101010010
00010100101001010010100101101000100001001101010000
00010100101110011110100101101000100001001101010110
10010100101001010010100101011000100001001011010010
10010100101001010010100101011000100001001011010010
01100011001110010010011001001000100001001001001110
00000000000000000000000000000000000000000000000000
00000000000000000000000000000000000000000000000000
```

Just remove head and tail two lines each, and copy the content left into your grass file.

## Do I have to destroy and create the text repository every day?

Perhaps, you could be rescued by using `hub` command to automate the operation.

# In the first place,

This is a crappy software, and give useless loads to GitHub. So, if you have enough common sense, it's supposed that you're aware of that you should not use this.

