# xjsonl

xjsonl is a tool to convert xSV to JSON Lines.

[![Build Status](https://travis-ci.org/winebarrel/xjsonl.svg?branch=master)](https://travis-ci.org/winebarrel/xjsonl)

## Usage

```
Usage of xjsonl:
  -keys string
    	json object keys
  -sep string
    	line separator. not split if empty (default ",")
  -version
    	print version and exit
  -with-header
    	consider the first line as a header
```

```
$ printf 'foo,bar\nbar,zoo' | xjsonl
["foo","bar"]
["bar","zoo"]

$ printf 'foo\tbar\nbar\tzoo' | xjsonl -sep '\t'
["foo","bar"]
["bar","zoo"]

$ printf 'foo,bar\nbar,zoo' | xjsonl -sep ''
["foo,bar"]
["bar,zoo"]

$ printf 'foo,bar\nbar,zoo' > data.csv
$ xjsonl data.csv
["foo","bar"]
["bar","zoo"]

$ printf 'foo,bar\nbar,zoo' | xjsonl -keys a,b
{"a":"foo","b":"bar"}
{"a":"bar","b":"zoo"}

$ printf "foo,bar\nzoo,baz\n1,2" | xjsonl -with-header
{"foo":"zoo","bar":"baz"}
{"foo":"1","bar":"2"}
```
