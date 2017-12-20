#!/bin/sh

# MIT License
# 
# Copyright (c) 2017 Rob Baruch
# 
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
# 
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
# 
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
# 

## is search on
if [ `cat config.toml | grep algolia_search | awk ' { print $NF } '` != 'true' ]
then
	echo "no searching requested, exiting"
	exit 1
fi

APP_ID=`cat config.toml | grep -i algolia_appId | awk ' { print $NF } ' | sed -e 's/"//g'`
# this is the public key -- API_KEY=`cat config.toml | grep -i algolia_apiKey | awk ' { print $NF } '` | sed -e 's/"//g'
INDEX=`cat config.toml | grep -i algolia_indexName | awk ' { print $NF } ' | sed -e 's/"//g'`
API_KEY='** your_update_key_goes_here **'

if [ "$APP_ID" == "" ]
then
	echo "no app id set, exiting"
	exit 1
fi

if [ "$API_KEY" == "" ]
then
	echo "no app key set, exiting"
	exit 1
fi

if [ "$INDEX" == "" ]
then
	echo "no index set, exiting"
fi

echo "performing algolia update with:"
echo "\t$APP_ID"
echo "\t$API_KEY"
echo "\t$INDEX"
echo ""
./algtool -id "$APP_ID" -index "$INDEX" -key "$API_KEY" && echo "\nupdated successfully."
echo ""
echo "done."

