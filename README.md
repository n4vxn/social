# Social Backend API

Welcome to the **Social Backend API**, a fully-featured backend service designed for social networking platforms. This API, built using the Chi router in Go, integrates robust user authentication, efficient post and comment management, and seamless integration with PostgreSQL.

The backend uses JWT for secure session management and leverages GORM for database migrations, ensuring data consistency and ease of development. With modular routing and well-structured services, the API provides a scalable foundation for building social applications.

---

## ✨ Features

- **User Authentication:** Secure JWT-based cookie authentication, user account activation, and password management.
- **Post Management:** Endpoints for creating, retrieving, updating, and deleting posts. Includes support for likes, dislikes, and pagination.
- **Comment Management:** Full CRUD operations on comments with like/dislike functionalities and nested replies.
- **Health Monitoring:** Dedicated endpoint to check API availability and health.

---

## 🚀 Key Services

### 1. Health Monitoring
- A simple endpoint to verify that the API and router are functioning as expected.

### 2. Authentication
- Comprehensive user registration, login, logout, password reset, and OAuth flows.

### 3. Posts
- Manage user-generated content with robust CRUD endpoints and additional support for engagement actions like likes and dislikes.

### 4. Comments
- Enable dynamic discussions with threaded comments. Supports likes/dislikes for better user engagement.

---

## 🛠 API Endpoints

### Health Routes
- **GET `/health/router`:** Returns a status message indicating the health of the API.

### Authentication Routes
- **POST `/auth/register`:** Register a new user.
- **GET `/auth/activate/{token}`:** Activate a user account via token.
- **POST `/auth/login`:** Login and initiate a secure JWT session.
- **POST `/auth/logout`:** Logout and invalidate the JWT cookie.
- **POST `/auth/forgot-password`:** Request a password reset email.
- **POST `/auth/reset-password/{token}`:** Reset the user password using the token.
- **GET `/auth/google/login`:** Initiate login with Google OAuth.
- **GET `/auth/google/callback`:** Handle Google OAuth callback.
- **GET `/auth/github/login`:** Initiate login with GitHub OAuth.
- **GET `/auth/github/callback`:** Handle GitHub OAuth callback.

### Post Routes
- **GET `/api/v1/posts`:** Retrieve all posts with optional pagination.
- **GET `/api/v1/posts/{postID}`:** Get details of a specific post.
- **POST `/api/v1/posts`:** Create a new post.
- **PATCH `/api/v1/posts/{postID}`:** Update an existing post.
- **DELETE `/api/v1/posts/{postID}`:** Delete a post.
- **POST `/api/v1/posts/{postID}/like`:** Like a post.
- **DELETE `/api/v1/posts/{postID}/like`:** Remove like from a post.
- **POST `/api/v1/posts/{postID}/dislike`:** Dislike a post.
- **DELETE `/api/v1/posts/{postID}/dislike`:** Remove dislike from a post.

### Comment Routes
- **GET `/api/v1/comments/{commentID}`:** Retrieve a specific comment.
- **GET `/api/v1/posts/{postID}/comments`:** Retrieve comments for a specific post.
- **POST `/api/v1/posts/{postID}/comments`:** Add a new comment to a post.
- **PUT `/api/v1/comments/{commentID}`:** Update a comment.
- **DELETE `/api/v1/comments/{commentID}`:** Delete a comment.
- **POST `/api/v1/comments/{commentID}/like`:** Like a comment.
- **DELETE `/api/v1/comments/{commentID}/like`:** Remove like from a comment.
