// app.js
document.addEventListener('DOMContentLoaded', () => {
    const registerForm = document.getElementById('register-form');
    registerForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        const formData = new FormData(registerForm);
        const response = await fetch('/register', {
            method: 'POST',
            body: new URLSearchParams(formData),
        });
        if (response.ok) {
            alert('Registration successful');
        } else {
            alert('Registration failed');
        }
    });

    const loginForm = document.getElementById('login-form');
    loginForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        const formData = new FormData(loginForm);
        const response = await fetch('/login', {
            method: 'POST',
            body: new URLSearchParams(formData),
        });
        if (response.ok) {
            alert('Login successful');
            // Redirect or update UI
        } else {
            alert('Login failed');
        }
    });

    const messageForm = document.getElementById('message-form');
    messageForm.addEventListener('submit', (event) => {
        event.preventDefault();
        const content = document.getElementById('message-content').value;
        socket.send(JSON.stringify({ content }));
        document.getElementById('message-content').value = '';
    });

    // WebSocket setup
    const socket = new WebSocket('ws://localhost:8080/ws');
    socket.onopen = () => {
        console.log('WebSocket connection opened');
    };

    socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        displayMessage(message);
    };

    function displayMessage(message) {
        const chat = document.getElementById('chat');
        const messageElement = document.createElement('div');
        messageElement.textContent = `${message.sender}: ${message.content}`;
        chat.appendChild(messageElement);
    }
});
