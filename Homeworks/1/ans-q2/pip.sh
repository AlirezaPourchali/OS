#!/bin/bash

name=$(pip list --outdated | awk '{ print $1 }')

for i in $name; do
    pip install --upgrade "$i"
done
