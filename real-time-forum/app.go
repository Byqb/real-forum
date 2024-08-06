package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    _ "github.com/mattn/go-sqlite3"
)

var upgrader = websocket.Upgrader{}
var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite3", "./forum.db")
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", serveIndex)
    http.HandleFunc("/ws", handleWebSocket)
    http.HandleFunc("/register", handleRegister)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/posts", handlePosts)
    http.HandleFunc("/post", handlePost)
    http.HandleFunc("/comment", handleComment)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Serve the single HTML file
func serveIndex(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Forum</title>
    <style>
        body { font-family: Arial, sans-serif; }
        #navbar { background-color: #333; color: white; padding: 1em; display: flex; justify-content: space-around; }
        #navbar a { color: white; text-decoration: none; }
        .form-container, .content-container { margin: 2em; }
        #posts-feed, #post-details, #messages { display: none; }
        #posts-feed.active, #post-details.active, #messages.active { display: block; }
        textarea { width: 100%; }
    </style>
</head>
<body>
    <div id="app">
        <nav id="navbar">
            <a href="#" id="home-link">Home</a>
            <a href="#" id="messages-link">Messages</a>
            <a href="#" id="logout-link">Logout</a>
        </nav>

        <div id="registration-form" class="form-container">
            <h2>Register</h2>
            <form id="register-form">
                <label for="nickname">Nickname:</label>
                <input type="text" id="nickname" name="nickname" required>
                <label for="age">Age:</label>
                <input type="number" id="age" name="age" required>
                <label for="gender">Gender:</label>
                <select id="gender" name="gender" required>
                    <option value="male">Male</option>
                    <option value="female">Female</option>
                    <option value="other">Other</option>
                </select>
                <label for="first-name">First Name:</label>
                <input type="text" id="first-name" name="first-name" required>
                <label for="last-name">Last Name:</label>
                <input type="text" id="last-name" name="last-name" required>
                <label for="email">Email:</label>
                <input type="email" id="email" name="email" required>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
                <button type="submit">Register</button>
            </form>
        </div>

        <div id="login-form" class="form-container">
            <h2>Login</h2>
            <form id="login-form">
                <label for="login-nickname">Nickname or Email:</label>
                <input type="text" id="login-nickname" name="login-nickname" required>
                <label for="login-password">Password:</label>
                <input type="password" id="login-password" name="login-password" required>
                <button type="submit">Login</button>
            </form>
        </div>

        <div id="posts-feed" class="content-container">
            <h2>Posts</h2>
            <div id="posts-list"></div>
        </div>

        <div id="post-details" class="content-container">
            <h2>Post Details</h2>
            <div id="post-content"></div>
            <div id="comments-section">
                <h3>Comments</h3>
                <div id="comments-list"></div>
                <form id="comment-form">
                    <label for="comment-content">Add a Comment:</label>
                    <textarea id="comment-content" name="comment-content" required></textarea>
                    <button type="submit">Submit Comment</button>
                </form>
            </div>
        </div>

        <div id="messages" class="content-container">
            <h2>Private Messages</h2>
            <div id="online-users">
                <h3>Online Users</h3>
                <ul id="online-users-list"></ul>
            </div>
            <div id="chat-container">
                <div id="chat-header">
                    <h3>Chat with: <span id="chat-user-name"></span></h3>
                </div>
                <div id="chat"></div>
                <form id="message-form">
                    <textarea id="message-content" name="message-content" required></textarea>
                    <button type="submit">Send</button>
                </form>
            </div>
        </div>

        <div id="loading-message" class="hidden">Loading...</div>
        <div id="error-message" class="hidden"></div>
    </div>

    <script>
        const socket = new WebSocket('ws://localhost:8080/ws');

        socket.onopen = () => {
            console.log('WebSocket connection opened');
        };

        socket.onmessage = (event) => {
            const message = JSON.parse(event.data);
            displayMessage(message);
        };

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
        });

        function displayMessage(message) {
            const chat = document.getElementById('chat');
            const messageElement = document.createElement('div');
            messageElement.textContent = `${message.sender}: ${message.content}`;
            chat.appendChild(messageElement);
        }
    </script>
</body>
</html>
`

    fmt.Fprintf(w, tmpl)
}

// WebSocket handler
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error while connecting:", err)
        return
    }
    // Handle WebSocket messages
}

// Handle registration
func handleRegister(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    r.ParseForm()
    nickname := r.FormValue("nickname")
    age := r.FormValue("age")
    gender := r.FormValue("gender")
    firstName := r.FormValue("first-name")
    lastName := r.FormValue("last-name")
    email := r.FormValue("email")
    password := r.FormValue("password")
    // Password hashing and DB insertion
}

// Handle login
func handleLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    r.ParseForm()
    nicknameOrEmail := r.FormValue("login-nickname")
    password := r.FormValue("login-password")
    // Authentication logic
}

// Handle posts
func handlePosts(w http.ResponseWriter, r *http.Request) {
    // Retrieve and return posts
}

// Handle a specific post
func handlePost(w http.ResponseWriter, r *http.Request) {
    // Retrieve and return a specific post
}

// Handle comments
func handleComment(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    r.ParseForm()
    postID := r.FormValue("post_id")
    content := r.FormValue("content")
    // Insert comment into database
}
