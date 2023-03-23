# ac-profile
You can view atcoder profile information on the terminal.
## Install
```bash
go install github.com/cotora/ac-profile
```
## Usage
Running the command with a user name as an argument will display the user's profile information.
```shell-session
# ac-profile cotora
[Algorithm]
User           : cotora
Country/Region : Japan
Birth Year     : 2001
Twitter ID     : @cotora_kyopro
Rank           : 12993rd
Rating         : 934
Highest Rating : 934
Rated Matches  : 44
Last Competed  : 2023/03/19
```
Running with the `-h` flag as an option will display the heuristic information.
```shell-session
# ac-profile -h cotora
[Heuristic]
User           : cotora
Country/Region : Japan
Birth Year     : 2001
Twitter ID     : @cotora_kyopro
Rank           : 977th
Rating         : 1060
Highest Rating : 1060
Rated Matches  : 4
Last Competed  : 2023/02/18
```