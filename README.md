# Collection of opinionated services with pluggable implementations

This project is build from two assumptions: 
1. Majority of applications don't need full potential of relational databases, distributed event store platforms, etc.
2. It's possible to build an opiniated and simple set of APIs to satisfy needs of majority applications.

Benefits that such generic backbone services can provide:
* Plugabble implementations: use provided services if you are running in a cloud or use self-hosted backends for on-prem installations.
* Much easier transitoin between clouds aka "no vendor lock".
