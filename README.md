# xjsonl

xjsonl is a tool to convert xSV to JSON Lines.

## Usage

```
Usage of ./xjsonl:
  -keys string
    	JSON object keys
  -sep string
    	line separator. not split if empty (default ",")
  -version
    	Print version and exit
```

```
$ printf "foo,bar\nbar,zoo" | xjsonl
["foo","bar"]
["bar","zoo"]

$ printf "foo\tbar\nbar\tzoo" | xjsonl -sep '\t'
["foo","bar"]
["bar","zoo"]

$ printf "foo,bar\nbar,zoo" | xjsonl -sep ""
["foo,bar"]
["bar,zoo"]

$ printf "foo,bar\nbar,zoo" > data.csv
$ xjsonl data.csv
["foo","bar"]
["bar","zoo"]

$ printf "foo,bar\nbar,zoo" | xjsonl -keys a,b
{"a":"foo","b":"bar"}
{"a":"bar","b":"zoo"}
```
