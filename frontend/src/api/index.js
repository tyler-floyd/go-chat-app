// api/index.js

var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
    console.log("Connecting to server...");

    socket.onopen = () => {
        console.log("Successfully connected to server");
    };

    socket.onmessage = (msg) => {
        console.log(msg);
    };

    socket.onclose = (event) => {
        console.log('Socket closed connection: ', event);
    };

    socket.onerror = (error) => {
        console.log('Socket error: ', error);
    };  
};

let sendMsg = (msg) => {
    console.log("Sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };
