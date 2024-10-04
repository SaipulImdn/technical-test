````markdown
# User Management System

This project is a backend service designed to manage user registration, login, and CRUD operations. Built using Golang with the Gin framework, it also incorporates dynamic Docker configuration for easy deployment.

## Features

- **User Registration**: Allows users to create an account with unique username and email validation.
- **User Login**: Authenticates users with their email and password.
- **CRUD Operations**: Provides endpoints for creating, reading, updating, and deleting user data.
- **Dynamic Docker Configuration**: Utilizes environment variables for flexible Docker deployments.

## Technologies Used

- **Golang**: Primary programming language
- **Gin**: HTTP web framework for Golang
- **GORM**: ORM library for Golang
- **MySQL**: Database for storing user data
- **Docker**: Containerization for simplified deployment

## Getting Started

### Prerequisites

- Go 1.21+
- Docker
- MySQL

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   ```
````

2. **Set Up Environment Variables:**

   Create a `.env` file with the following contents. Ensure this file is added to `.gitignore`:

   ```plaintext
   APP_ENV=development
   APP_PORT=5000
   DB_HOST=localhost
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=yourdatabase
   DB_PORT=3306
   JWT_SECRET_KEY=your-secret-key
   ```

3. **Run the Application Locally:**

   ```bash
   go run main.go
   ```

### Docker Setup

1. **Build the Docker Image:**

   ```bash
   docker build -t user-management .
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -p 5000:5000 --env-file .env user-management
   ```

   The application should now be accessible at `http://localhost:5000`.

## API Endpoints

The following endpoints are available in this API:

- **POST** `/api/v1/register`: Register a new user.
- **POST** `/api/v1/login`: Log in a user.
- **GET** `/api/v1/users`: Retrieve all users.
- **GET** `/api/v1/users/:username`: Retrieve a specific user by username.
- **PUT** `/api/v1/users/:username`: Update a user by username.
- **DELETE** `/api/v1/users/:username`: Delete a user by username.

### Example Requests

- **Register User**:

  ```bash
  curl -X POST http://localhost:5000/api/v1/register -d '{"username": "newuser", "email": "newuser@example.com", "password": "password123"}'
  ```

- **Login User**:

  ```bash
  curl -X POST http://localhost:5000/api/v1/login -d '{"email": "newuser@example.com", "password": "password123"}'
  ```

- **Get All Users**:

  ```bash
  curl -X GET http://localhost:5000/api/v1/users
  ```

- **Get User by Username**:

  ```bash
  curl -X GET http://localhost:5000/api/v1/users/newuser
  ```

- **Update User**:

  ```bash
  curl -X PUT http://localhost:5000/api/v1/users/newuser -d '{"email": "updateduser@example.com", "username": "updateduser"}'
  ```

- **Delete User**:

  ```bash
  curl -X DELETE http://localhost:5000/api/v1/users/newuser
  ```

## Dynamic Docker Configuration

The Dockerfile is set up to dynamically use environment variables for configuration. This allows you to modify settings without altering the application code directly.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any proposed changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

```

Make sure to replace `your-username` and `your-repository` with your actual GitHub username and repository name. This `README.md` provides a comprehensive overview of your project, including setup instructions and examples of how to use the API. You can copy and paste it directly into your `README.md` file.
```
