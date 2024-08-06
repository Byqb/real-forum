# Forum Application

## Overview

This project is an upgraded forum application with real-time features, including private messaging, live updates, and more. It integrates multiple technologies to create a seamless user experience:

- **SQLite** for database management
- **Golang** for backend services and WebSocket handling
- **JavaScript** for frontend interactions and WebSocket communication
- **HTML** for page structure
- **CSS** for styling

## Features

1. **Registration and Login**
   - User registration with nickname, age, gender, first name, last name, email, and password.
   - Login using either nickname or email combined with password.
   - Logout functionality from any page.

2. **Posts and Comments**
   - Creation of posts with categories.
   - Commenting on posts.
   - Viewing posts in a feed.
   - Viewing comments only when a post is clicked.

3. **Private Messages**
   - Real-time private messaging between users.
   - Display online/offline users organized by recent activity.
   - Sending messages to online users.
   - Viewing past conversations with a user.
   - Loading additional messages on scroll (throttled/debounced).

## Technologies Used

- **Backend:**
  - **Golang**: Handles server logic, WebSocket communication, and database interactions.
  - **Gorilla WebSocket**: Manages WebSocket connections.
  - **SQLite**: Stores user data, posts, and messages.
  - **bcrypt**: Handles password hashing.

- **Frontend:**
  - **JavaScript**: Manages real-time updates and client-side interactions.
  - **HTML**: Structures the web pages.
  - **CSS**: Styles the application.

### Prerequisites

- Go 
- SQLite

## Serve the Frontend

- The frontend is served by the backend application itself. Ensure that the backend is running, and it will handle serving the HTML, CSS, and JavaScript files.

## Access the Application

- Open your web browser and go to http://localhost:8080 (or the port specified in your backend configuration).

## Usage
- Register: Navigate to the registration page and fill out the required details.
- Login: Use your credentials to log in.
- Create Posts: Access the post creation form and submit posts.
- Comment on Posts: View a post and add comments.
- Send Private Messages: Use the chat interface to send and receive messages in real-time.

### Steps

1. **Clone the Repository**
```bash
   git clone https://github.com/Byqb/real-forum.git
   cd real-forum
```


# Project Structure

```forum-application/
├── backend/
│   ├── main.go
│   ├── routes.go
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── posts.go
│   │   ├── comments.go
│   │   └── chat.go
│   ├── models/
│   │   ├── user.go
│   │   ├── post.go
│   │   ├── comment.go
│   │   └── message.go
│   ├── database/
│   │   ├── db.go
│   │   ├── migrations/
│   │   │   ├── init.sql
│   │   │   └── schema.sql
│   ├── utils/
│   │   └── auth.go
│   └── go.mod
|── frontend/
    ├── src/
    │   ├── index.ts
    │   ├── auth.ts
    │   ├── posts.ts
    │   ├── comments.ts
    │   ├── chat.ts
    │   ├── utils.ts
    │   └── styles/
    │       └── styles.css
    ├── dist/
    │   ├── index.js
    │   ├── auth.js
    │   ├── posts.js
    │   ├── comments.js
    │   ├── chat.js
    │   └── utils.js
    ├── index.html
    └── tsconfig.json
```
