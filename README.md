# kjbc
The king James Bible concordance for your terminal.

## Usage:

`kjbc [options] word [word]`
This will search the concordance to find the entry number for an English word.
If only one entry exists for a word, kjbc will return the entry definition as well.

`kjbc [options] entry [number]`
This will return the information for a given entry number.

### Options:
`-l [hebrew/greek]` The target language to search within. Defaults to greek when omitted.

`-v` Display the verses that correspond with the concordance entry number. Defaults to false when omitted.

`-u` Display the usage information when set to true. Defaults to false when omitted.

`-help` Display the list of options. Defaults to false when omitted.

## Support The Project:
There are several ways to support this project. Visit [swab.dev/support](https://swab.dev/support) for more information.

---

P.S. - I did not type up the entire Strong's concordance. That was already done and the text files may be found at http://www.ebibleconcordances.com/

I simply formatted the information to how I needed it in order to create the CLI.