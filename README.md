# Package Statistics Tool

## Table of Contents
- [Description](#description)
- [Solution](#solution)
  * [Algorithm](#algorithm)
  * [Performance](#performance)
  * [Tools](#tools)
- [Prerequisites](#prerequisites) 
- [Usage](#usage)


## Description 
The current `git` repository contains a `CLI` implementation for parsing a file from a `URL` containing compressed content indices associated with an architecture(amd64, arm64, mips etc). </br>
The architecture is passed via a cmd argument. The usage of our CLI will be developed in the following sections.</br>
For the aforementioned problem two solutions will be provided.</br>
The following solution was implemenented in approximately 8 hours.

## Solution 

### Algorithm
This problem from an algorithmic perspective assuming that the file's lines are parsed and read as supposed to, can be identified as a sorting problem. </br>
Our implementation language was in `Go` so the [sort](https://pkg.go.dev/sort) package of the `Go` was used to sort our files.</br>
The `sort` function needs a slice as input meaning that the appropriate data transformations need to occur in order to end up in the wanted data structure.</br>
Every line is parsed and every package in the line is added in a hashmap incrementing the counter for the associated package or initializing it if the package does not exist in the hashmap.</br>

### Performance
First a serial solution was implemented in order to check the correctness and to check the results (`size` of the file, `execution-time`, etc)</br>
The performance of a single threaded solution was quite quickly regarding the fact that each file has a size of 1,5 million lines to parse.</br>

Later, a concurrency solution was implemented in order to take advantage of `Go`'s easy applying `concurrency` mechanisms which are not other than `goroutines` and `channels`. </br> 
The synchronization schema implemenented consists of multiple `senders` and a single `receiver`.</br>
The `decision` to make arises as per the chunk size to be used in terms of lines that a `go routine` would be responsible for. </br>
This decision depends on multiple factors as per the host machine running the code. </br>
Spawning multiple goroutines (exceeding) host's CPU threads cores can affect the performance of our solution.<br>
The host machine where the solution was implemented a `chunk size` of `176171` lines was used.</br>
The reasoning behind our decision is that our `host machine consists of 10` cores.</br>
We wanted `9` goroutines to act as the senders and `1` goroutine to act as the receiver.</br>
This decision was taken as the result of the `lscpu` command on our host machine.</br>
The handy part of the algorithm is to extract the file's size.</br>
In our case this was difficult without having to parse the whole file as our executable downlads and unzips a compressed file, meaning we do not have a direct view of the file's size.</br>
We did not want to affect the performance so some tests were run offline and seen that the average file consists of lines of 69 bytes in size and a total of 1.5 millions lines.</br>

```matlab
 chunk size = (1.5 * 10^6 * 69) / 9 = 176171 
```

Each goroutine is responsible to parse 176151 lines associate the packages in an array, sending the aforementioned batch of strings to the receiver via a common `buffered` channel.</br>
Indeed having tested `bigger` or `less` chunk size was not as optimal.

The `execution time` in the concurrent implementation was 1 second better than the non concurrent one.

### Tools
The tools that were used to develop the CLI were 
1. the [Cobra](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md) tool, along with
2. [Viper](https://github.com/spf13/viper)

## Prerequisites
The solution was developed using the `go1.23.2` version.</br>
The following command needs to be executed in order to build the `binary` file from the [package-statistics-chunks](/package-statistics-chunks/) directory.

```bash
go build -o package-statistics
```
The same command can be executed in the [package-statistics-one-routine](/package-statistics-one-routine/) if wished.</br>

## Usage
To execute the cli the following cmd needs to be executed

```bash
 ./package_statistics --config [path to the config file] [arch name]
```
 In our case the [.cobra.yaml](/package-statistics-chunks/.cobra.yaml).<br> 
 The following are declared in the config file: 

 ```yaml
 debianURL: "http://ftp.uk.debian.org/debian/dists/stable/main/"
 outputFile: "./files"
 chunkSize: 176171
 ```




















