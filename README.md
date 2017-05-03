# Scout [![Build Status](https://travis-ci.org/project-regista/scout.svg?branch=master)](https://travis-ci.org/project-regista/scout)

The Scout repository contains a list of functions to handle HTTP requests to Soccerama API.

The currently supported resources are:

[Competitions](https://soccerama.pro/docs/1.2/competitions):
  - CompetitionCountry                : get information about a single competition including country
  - CompetitionsCountry               : get all the competitions including country

[Countries](https://soccerama.pro/docs/1.2/countries):
  - Country                           : get information about a single country
  - Countries                         : get all the countries

[Matches](https://soccerama.pro/docs/1.2/matches):
  - MatchSeason                       : get information about a single match including season
  - MatchesSeason                     : get all the matches including season

[Seasons](https://soccerama.pro/docs/1.2/seasons):
  - SeasonCompetition                 : get information about a single season including competition
  - SeasonsCompetition                : get all the season including competition


## HOWTO:

```
------------------------------------------------------------------------
regista - scout
------------------------------------------------------------------------
build                          build test container
test                           run unit and benchmark tests
```
