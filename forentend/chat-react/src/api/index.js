var socket = new WebSocket("ws://192.168.10.27:8080/ws");
// const token="2QDx3iWPVmW99GMsb0G6Aiy75Ie_4HsLpP4cC5s2Lj2JuNRX6"
// var socket = new WebSocket(`wss://your-ngrok-url-here.ngrok.io/ws?token=${token}`);

let connect = cb => {
    console.log("Attempting Connection....");


    socket.onopen = () => {
        console.log("Connect"); 
    }
    socket.onmessage = msg => {
        console.log(msg);
        cb(msg)
    }
    
    socket.onclose = event => {
        console.log("Scoket cloased ", event);
    }
    socket.onerror = error => {
        console.log("socket error ", error);
    }
}
let sendMsg = msg => {
    console.log("sending msg ", msg);
    socket.send(msg);
}

export {
    connect, sendMsg
}