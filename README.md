# DALE

Email alias service (backend API)

## Setup
### Database
Set up PostgreSQL database. In `./.env` enter the connection string:
```
C_STRING="postgresql://<username>:<pw>@<server>/<db>?sslmode=require"
```

### Redis
Setup Redis server. In `./.env` enter the redis address and password:
```
REDIS_ADDR="addr:port"
REDIS_PWD="password"
```

## Usage
cd into the root directory.

1. `go mod tidy` to update, get dependencies, etc.
2. For development, install [air](https://github.com/air-verse/air?tab=readme-ov-file#installation) if not already: `go install github.com/air-verse/air@latest`
3. Start live server by `air`

## Roles
A role `x` inherits permissions from roles lower than it.  
- `0` - User
  - Access to most endpoints, access their own resources
- `1` - Premium
  - `0` permissions
  - Access to extra endpoints, relaxed limits
- `2` - Admin
  - Access to all endpoints, access to all resources

## Endpoints
Base url for API: `/`

## /
## GET /ping
Returns `{"message": "pong"}`

## /users
**Role**: *ADMIN* required  

### GET /
**Role**: *ADMIN* required  
Gets all users and their information  
Returns `[{user}, ...]`

### POST /
**Role**: *ADMIN* required  
Create user with provided information  
Body:
```json
{
  "email": "email@email.com",
  "password": "password",
  "role": 0,
}
```

### GET /:id
**Role**: *ADMIN* required  
Gets specific user information  
Returns `{user}`

## /aliases
**Role:** *USER* required

### GET /
**Role:** *USER* required
Gets all aliases by user  
Returns:
```json
[
    {
        "ID": 3,
        "CreatedAt": "2024-08-27T04:53:34.476623+01:00",
        "UpdatedAt": "2024-08-27T04:53:34.476623+01:00",
        "DeletedAt": null,
        "userID": 2,
        "aliasPrefix": "MnKRCvh",
        "isActive": true,
        "isDeleted": false
    }
]
```

### POST /
**Role:** *USER* required
Creates an alias for the user, no parameters  
Returns:
```json
{
    "ID": 3,
    "CreatedAt": "2024-08-27T04:53:34.4766236+01:00",
    "UpdatedAt": "2024-08-27T04:53:34.4766236+01:00",
    "DeletedAt": null,
    "userID": 2,
    "aliasPrefix": "MnKRCvh",
    "isActive": true,
    "isDeleted": false
}
```

### DELETE /:id
**Role:** *USER* required
Deletes a user's alias  
Returns:
```json
{
    "message": "Alias deleted successfully"
}
```

### POST /toggle/:id
**Role:** *USER* required
Toggles the user's alias active status  
Returns:
```json
{
    "ID": 4,
    "CreatedAt": "2024-08-27T05:16:58.345478+01:00",
    "UpdatedAt": "2024-08-27T05:17:03.5071936+01:00",
    "DeletedAt": null,
    "userID": 2,
    "aliasPrefix": "xypLS48",
    "isActive": false,
    "isDeleted": false
}
```

#### GET /admin
**Role**: *ADMIN* required  
Gets all aliases  
Returns:
```json
[
    {
        "ID": 1,
        "CreatedAt": "2024-08-27T04:52:55.035347+01:00",
        "UpdatedAt": "2024-08-27T04:52:55.035347+01:00",
        "DeletedAt": null,
        "userID": 1,
        "aliasPrefix": "ecJYewh",
        "isActive": true,
        "isDeleted": false
    },
    ...
]
```

#### GET /admin/:id
**Role**: *ADMIN* required  
Gets a specific alias
Returns:
```json
{
    "ID": 1,
    "CreatedAt": "2024-08-27T04:52:55.035347+01:00",
    "UpdatedAt": "2024-08-27T04:52:55.035347+01:00",
    "DeletedAt": null,
    "userID": 1,
    "aliasPrefix": "ecJYewh",
    "isActive": true,
    "isDeleted": false
}
```

#### GET /admin/user/:userID
**Role**: *ADMIN* required  
Gets all aliases by a user  
Returns:
```json
[
    {
        "ID": 1,
        "CreatedAt": "2024-08-27T04:52:55.035347+01:00",
        "UpdatedAt": "2024-08-27T04:52:55.035347+01:00",
        "DeletedAt": null,
        "userID": 1,
        "aliasPrefix": "ecJYewh",
        "isActive": true,
        "isDeleted": false
    }
]
```

## /auth

### POST /login
Logs into an account  
Parameters:
```json
{
  "email": "email",
  "password": "password"
}
```
Returns:
```json
{
    "message": "Login successful"
}
```
or
```json
{
    "error": "Invalid Credentials"
}
```

### POST /signup
Creates an account
Parameters:
```json
{
  "email": "email",
  "password": "password"
}
```
Returns:
```json
{
    "message": "User created"
}
```
or
```json
{
    "error": "..."
}
```

### POST /logout
Logs out of the current account
