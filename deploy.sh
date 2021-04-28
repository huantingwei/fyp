# run this script in base directory

#!/bin/bash

true || rm fyp
export PATH=$PATH:/usr/local/go/bin
cd backend
go build .
mv fyp ../
cd ../
./fyp