# Choreia: A Static Analyzer to Generate Choreography Automata from Go Source Code

## Abstract 
Choreographies are an emerging paradigm for describing concurrent systems
which have been gaining momentum in the last few years. The main purpose is to provide the developer with a tool that allows an immediate understanding of the interaction occuring between the participants within the system and how they interact with each other. Starting from individual participants, and their local views, it is possible to recompose in a bottom-up manner the entire Choreography (or global view) of the system. An added advantage of the Choreographies is that, when they respect well defined properties, they give guarantees on the absence of typical concurrency problems such as Deadlocks, Liveness and Race Conditions. There are various formal models of Choreographies, this thesis deals specifically with Choreography Automata, based on Finite State Automata (FSA).  
This thesis presents Choreia: a static analysis tool that, starting from a Go source code, derives the
Choreography Automata of the concurrent system in a bottom-up manner.

## Credits and licensing
This work is licensed under a [Creative Commons Zero v1.0 Universal](https://creativecommons.org/share-your-work/public-domain/cc0).  
All credits to the respective authors of the cited papers, special thanks the my relator of this thesis: Professor [Ivan Lanese](https://dblp.org/pid/56/3713)