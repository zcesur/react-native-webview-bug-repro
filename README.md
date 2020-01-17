# react-native-webview-bug-repro
## Prerequisites

- `go`
- `yarn`/`npm`

## Running the app

1. Run HTTP server with `go run server.go`
2. Install the dependencies with `yarn install`
3. Replace `localhost` in App.js with the address of the machine running the server, if necessary.
4. Start the metro bundler with `yarn start`
5. Run the app with `yarn android` or `yarn ios`
