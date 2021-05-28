# kjbc
The king James Bible concordance for your terminal.

## Usage:

`kjbc [options] find [word]`
This will search the concordance to find the entry number for an English word.

`kjbc [options] get [number]`
This will return the information for a given entry number.

### Options:
`-l [hebrew/greek]` The target language to search within. Defaults to greek when omitted.

`-v [true/false]` Display the verses that correspond with the concordance entry number. Defaults to false when omitted.

`-u [true/false]` Display the usage information when set to true. Defaults to false when omitted.

`-help [true/false]` Display the list of options. Defaults to false when omitted.