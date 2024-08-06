Project Structure

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
├── frontend/
│   ├── src/
│   │   ├── index.ts
│   │   ├── auth.ts
│   │   ├── posts.ts
│   │   ├── comments.ts
│   │   ├── chat.ts
│   │   ├── utils.ts
│   │   └── styles/
│   │       └── styles.css
│   ├── dist/
│   │   ├── index.js
│   │   ├── auth.js
│   │   ├── posts.js
│   │   ├── comments.js
│   │   ├── chat.js
│   │   └── utils.js
│   ├── index.html
│   └── tsconfig.json
├── .gitignore
├── README.md
└── LICENSE
```
