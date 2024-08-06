const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = function(event) {
    const message = JSON.parse(event.data);
    displayMessage(message);
};

function sendMessage(content) {
    const message = {
        sender: 'user1',
        content: content,
        timestamp: new Date().toISOString()
    };
    socket.send(JSON.stringify(message));
}

function displayMessage(message) {
    const messageElement = document.createElement('div');
    messageElement.textContent = `[${message.timestamp}] ${message.sender}: ${message.content}`;
    document.getElementById('chat').appendChild(messageElement);
}
