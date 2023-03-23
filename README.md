# ac-profile
ターミナル上でatcoderのプロフィール情報を表示することができます.
## Install
```bash
go install github.com/cotora/ac-profile
```
## Usage
ユーザー名を引数として実行するとそのユーザーのプロフィール情報が表示されます.
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
オプションとして`-h`フラグをつけて実行すると、ヒューリスティックの情報が表示されます
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