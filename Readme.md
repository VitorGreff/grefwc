## wc CLI
Command line tool to fetch info about files.
Currently, this repo supports the following operations:
- get number of bytes -> -c
- get number of lines -> -l
- get number of words -> -w
- get number of characters -> -m
## How to use
Build the source go code:
```console
go build gfwc.go
```
Then, run the binary by typing:
```console
./gfwc <flags-list> <filepath>
./gfwc -c -w example.txt
```
As all flags are optional, when you dont provide any of the given options, the program will execute all of them. Also, this tool supports stdin data, so this:
```console
cat example.txt | ./gfwc -c -w 
```
will work just fine. Notice that, when no filepath is provided, the program will search data on the stdin channel.