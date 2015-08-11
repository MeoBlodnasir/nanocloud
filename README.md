# Webapp #

## Requirements ##
[io.js](https://iojs.org) or [node.js](https://nodejs.org)

## Install ##
- Clone this repository to a local directory.
- Go into the directory and run `npm install`.
- Start to code by running `npm start`.

## npm Scripts ##
- `npm install`
 - Gets all node modules and dependencies.
 - Gets all Javascript, CSS and Typescript libraries with bower.
 - Gets all Typescript definitions.
 - Concatenates js and css files of the libraries.
 - Builds ts and less files.
- `npm start`
 - Builds ts and less files.
 - Starts to watch the file modifications to build again.
 - Starts an HTTP server at `http://localhost:8080/`.