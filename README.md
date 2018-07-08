# wail 😱

Wail is a little go tool that listens for pipe input and then streams it to a websocket.

It's useful for sharing the contents of a log to other places on the interwebs.

## Usage

### On the server

```sh
./myServer | wail
```

### In the client

```js
const ws = new WebSocket("ws://localhost:80", "protocolOne");
ws.onmessage = function(event) {
  console.log(event.data);
};
```
