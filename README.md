# process-grep

This is an early version... say 0.2

It is designed to extract key information from huge text files (e.g. program log files)

The program allows a progressive sequence of grep searches and regular expressions to be run.

The steps are configured in json files.

It's easiest to explain by considering the following problem

Imagine you have a log file of 300,000 lines. It's been generated from a highly concurrent program - so lines may not be strictly in order!
You're interested in those lines from the log where Error code XX occurs
Each of these line also contain identifiers.
The identifiers help indicate what the preceding and following events for Error XX were.

So you want
a) to find the lines where Error code XX occurs (grep)
b) use a regular expression to find the other indentifier in this line
c) use grep for each identifier to find the other set of lines

The input is one json configuration file and log file. 
The output from c can be a set of short extracts from the log file.

This radically reduces the effort required to establish what may have caused Error XX and what resulted.

Limitations

a) The code doesn't yet cope with requests to capture multiple groups from the regular expression
b) It's only been used for the use case above - so no more than 3 steps - and other possibilities haven't been explored.
c) whilst unit tests exists - the match between the tests and the latest version of the code isn't great. They are in need to tidying.
