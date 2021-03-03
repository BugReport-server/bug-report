This is a server that simply prints the URL from a GET request directed at it.

For example, if it is deployed at bug-report-vercel.app, a GET request to https://bug-report-vercel.app/api/test?hello should print "https://bug-report-vercel.app/api/test?hello"

Before February 23rd, a get request to https://bug-report-vercel.app/api/test?♡ would print "https://bug-report-vercel.app/api/test?♡", but after that day it is printing "https://bug-report-vercel.app/api/test?%E2%99%A1". While this may or may not be expected behaviour, it was not like this for at least a year. Additionally, any redirects (such as a domain pointed to now.sh redirecting to vercel.app) will cause this encoding to happen twice, first encoding ♡ to %E2%99%A1 then encoding that to %C3%A2%E2%84%A2%C2%A1. 
