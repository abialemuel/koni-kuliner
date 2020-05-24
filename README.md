# Koni-kuliner

Application for my mom ❤️ ❤️ ❤️ 

## Setup

### Prerequisite

- Go
- MySQL

### Installation

1. Copy `env.sample` to `.env`
  ```
  cp env.sample .env
  ```
2. Create database
  ```
  mysql> create database kuekoni_db;
  ```
3. Run Database Migration
  ```
  make migrate
  ```
4. Build & run the service
  ```
  make run
  ```

5. Try this
  ```
  localhost:5000
  ```

  