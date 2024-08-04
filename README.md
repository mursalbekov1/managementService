# Project Management API

## Overview

This project provides a RESTful API for managing users, tasks, and projects. The API allows for CRUD operations and includes search functionalities.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
    - [Users](#users)
    - [Tasks](#tasks)
    - [Projects](#projects)
- [Health Check](#health-check)
- [Contributing](#contributing)

## Installation 

1. **Clone the Repository**
```bash
git clone https://github.com/mursalbekov1/managementService.git
```

### Run and build the service:

```bash
make run
```

## API Endpoints

### Users

- **GET** `/v1/user/users`
- **POST** `/v1/user/users`
- **GET** `/v1/user/users/{id}`
- **PUT** `/v1/user/users/{id}`
- **DELETE** `/v1/user/users/{id}`
- **GET** `/v1/user/users/search?name={name}`
- **GET** `/v1/user/users/search?email={email}`
- **GET** `/v1/user/users/{id}/tasks`

### Tasks

- **GET** `/v1/task/tasks`
- **POST** `/v1/task/tasks`
- **GET** `/v1/task/tasks/{id}`
- **PUT** `/v1/task/tasks/{id}`
- **DELETE** `/v1/task/tasks/{id}`
- **GET** `/v1/task/tasks/search?title={title}`
- **GET** `/v1/task/tasks/search?status={status}`
- **GET** `/v1/task/tasks/search?priority={priority}`
- **GET** `/v1/task/tasks/search?assignee={userId}`
- **GET** `/v1/task/tasks/search?project={projectId}`

### Projects

- **GET** `/v1/project/projects`
- **POST** `/v1/project/projects`
- **GET** `/v1/project/projects/{id}`
- **PUT** `/v1/project/projects/{id}`
- **DELETE** `/v1/project/projects/{id}`
- **GET** `/v1/project/projects/search?title={title}`
- **GET** `/v1/project/projects/search?manager={userId}`

## Health Check

**GET** `/v1/healthCheck`