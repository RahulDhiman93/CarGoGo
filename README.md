# CarGoGo - A Carpooling Service
Welcome to CarGoGo, a carpooling service designed to connect drivers and passengers for shared rides. This project focuses on providing a backend service written in GoLang, with PostgreSQL as the underlying database.

# Introduction
CarGoGo is a carpooling service that aims to reduce traffic congestion, lower transportation costs, and promote a more sustainable and efficient way of commuting. The project is currently in its backend-only stage, implemented in GoLang with a PostgreSQL database.

# Features
**User Authentication:** Secure user registration and authentication system.
**Ride Management:** Create, view, and manage ride listings for drivers and passengers.
**Matching Algorithm:** Efficient algorithm to match passengers with drivers based on preferences and routes.
**Review System:** Allow users to provide and view reviews for each other to build a trustworthy community.

# Getting Started
You can follow these instructions to get a copy of the project up and running on your local machine for development and testing.

# Prerequisites
1. GoLang is installed on your machine. Visit https://golang.org/dl/ for installation instructions.
2. PostgreSQL database installed and running. Visit https://www.postgresql.org/download/ for installation instructions.

# Installation
Clone the repository:
`git clone https://github.com/RahulDhiman93/CarGoGo.git`

Change into the project directory:
`cd CarGoGo`

Install dependencies:
1. `go mod download`
2. `go tidy`
3. `brew install postgresql@14`
4. `brew install gobuffalo/tap/pop`
5. `Install DBeaver for GUI`

# Usage
1. Run `soda migrate` for up migration.
2. Make sure PostgreSQL is running.
3. Run the application: `./run.sh`.
4. Visit `http://localhost:8080` with endpoints in Postman to access the CarGoGo API.

# Contributing
Contributions are welcome! Please check the whole project for details on our code of conduct and the process for submitting pull requests.



