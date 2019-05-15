#!/bin/bash
USER_COUNT=`cat /etc/passwd | grep '^$1' -c`
USER_NAME=$1
if [ $USER_COUNT -ne 1 ]
 then
 useradd $USER_NAME
 echo $2 | passwd $USER_NAME --stdin
 else
 echo 'user exits'
fi
