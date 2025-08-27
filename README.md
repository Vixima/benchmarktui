# Aim Benchmark TUI

*god i need a better name for this shit.*

## A simple and lightweight way to track your progress in various aim trainer benchmarks.
### Created using Go and various Charm libraries.

# TODO
- [ ] actually implement the tui using bubbletea v2 (i wiped the code because it was shit, the other todo items are remnants of said code)
- [ ] save sheet
- [ ] colour code ranks
- [ ] pull ranks from csv instead of hardcoding
- [ ] make the progress bar work
- [ ] pull latest scores from the kovaaks folder
- [ ] pull benchmarks from the mastersheets instead of forcing the user to download a csv

_could maybe have a program like wget, curl, or aria2 as a dependency in conjunction with a csv output link for a mastersheet, like uh
`aria2c https://docs.google.com/spreadsheets/d/1k3XRRwn-mAbvb3rXIODFPrsnkgSEfh0xg1EWJICcSPo/gviz/tq\?tqx\=out:csv\&sheet\=Intermediate -o viscose_int_mastersheet.csv`_

_seems stupid and i could probably just do it with some sorta go package._

- [ ] add more benchmarks
- [ ] save data directly to a google sheets
- [ ] cool score graphs maybe??

### other cool projects
#### [sens](https://github.com/junapur/sens)
[sens](https://github.com/junapur/sens) is a cli sensitivity converter created by fellow aimer and nerd [junapur](https://github.com/junapur), it's really nice to use and is a great lightweight alternative to the bigger websites.
