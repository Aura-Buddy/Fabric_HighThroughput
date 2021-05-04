#!/bin/sh

declare -i a=0
declare -i b=1
while [ $a -lt $1 ]
do 
    echo $(date) >> data.txt
    ./attemptingConcurrency $b
    b=$((b+1))
    # increment the value
    a=`expr $a + 1`
done
