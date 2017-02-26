# Scout [![Build Status](https://travis-ci.org/project-regista/scout.svg?branch=master)](https://travis-ci.org/project-regista/scout)


## Howto:

Start a fresh new instance of Neo4j running in a Docker container

>`make neo4/start`

Change the default Neo4j password:

>`make neo4j/dev-pass`

Note that the first command will create the folder `$HOME/neo4j-regista/data` if this isn't already present.

### Available Commands
```
------------------------------------------------------------------------
regista
------------------------------------------------------------------------
build                          Build unit test container
clean                          Remove images and containers
neo4j/dev-pass                 Set new password for LOCAL development instance
neo4j/start                    Start Neo4j instance
neo4j/stop                     Stop Neo4j instance
neo4j/volume/purge             Remove Neo4j data volume @ $HOME/neo4j-regista/data
neo4j/volume                   Create Neo4j data volume @ $HOME/neo4j-regista/data
test                           Run unit tests
```
