# fn-extensions

Three example extensions for Fn server - check the article that explains these in more details: https://medium.com/@pjausovec/playing-with-the-fn-project-8c6939cfe5cc 

## Listener: Call counter extension (/callcount)
Counts a number of times an app has been called and outputs that value to stdout.

## Middleware: Cancel call middleware (/cancelcall)
Middleware that cancels the chain of calls in case a certain header is present in the call.

## Custom API endpoint: Call logs
Custom endpoint that returns the information about calls that were made to the app.

