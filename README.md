# process-grep


This is an early version... say 0.22

It is designed to extract key information from huge text files (e.g. log files)

The program allows a progressive sequence of grep searches and regular expressions to be run.

The steps are configured in json files. The example.json files are provided to give the idea - the others are real files.

It's easiest to explain by considering the following problem

Imagine you have a log file of **300,000 lines**. It's been generated from a highly concurrent program - so lines may not be strictly in order!
You're interested in those lines from the log where Error code XX occurs
Each of these line also contain network identifiers.
The network identifiers help indicate what the preceding and following events for Error XX were.

So you want
1. to find the lines where Error code XX occurs (grep)
2. use a regular expression to find the other indentifier in this line
3. use grep for each identifier to find the other set of lines

The input is one json configuration file and log file. 
The output from c can be a set of short extracts from the log file.

This radically reduces the effort required to establish what may have caused Error XX and what resulted.

Whilst it is plausible to do this using some combination of bash, grep and awk it wouldn't be easy to debug - or configure as things evolved.

## Limitations

1. The code doesn't yet cope with requests to capture multiple groups from the regular expression. Not sure whether it makes sense though - probably easier/better to do each as separate sequence.
2. It's only been used so far for the use case above - so no more than 3 steps - and other possibilities haven't been explored.
3. total development time so far < 1 day...
