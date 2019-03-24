#!/bin/sh

./eagle &
# jobs
# disown
echo moved graphql to background
nginx -g 'daemon off;'