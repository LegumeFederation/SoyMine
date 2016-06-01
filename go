#!/bin/sh
date > started
cd integrate
./go
cd ..
cd postprocess
./go
cd ..
date > ended
