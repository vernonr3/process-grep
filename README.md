# process-grep


This is an early version... say 0.21

It is designed to extract key information from huge text files (e.g. program log files)

The program allows a progressive sequence of grep searches and regular expressions to be run.

The steps are configured in json files.

It's easiest to explain by considering the following problem

Imagine you have a log file of **300,000 lines**. It's been generated from a highly concurrent program - so lines may not be strictly in order!
You're interested in those lines from the log where Error code XX occurs
Each of these line also contain identifiers.
The identifiers help indicate what the preceding and following events for Error XX were.

So you want
1. to find the lines where Error code XX occurs (grep)
2. use a regular expression to find the other indentifier in this line
3. use grep for each identifier to find the other set of lines

The input is one json configuration file and log file. 
The output from c can be a set of short extracts from the log file.

This radically reduces the effort required to establish what may have caused Error XX and what resulted.

## Limitations

1. The code doesn't yet cope with requests to capture multiple groups from the regular expression
2. It's only been used for the use case above - so no more than 3 steps - and other possibilities haven't been explored.
3. Unit test coverage is moderate. The logic for executing multiple steps could be broken down and tested more thoroughly.
4. total development time so far < 1 day...
