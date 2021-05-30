# kjbc
The king James Bible concordance for your terminal.

## Usage:

`kjbc [options] word [word]`
This will search the concordance to find the entry number for an English word.

`kjbc [options] entry [number]`
This will return the information for a given entry number.

### Options:
`-l [hebrew/greek]` The target language to search within. Defaults to greek when omitted.

`-v` Display the verses that correspond with the concordance entry number. Defaults to false when omitted.

`-u` Display the usage information when set to true. Defaults to false when omitted.

`-help` Display the list of options. Defaults to false when omitted.