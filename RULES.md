# Project Structure and Guidelines

This document outlines the structure and conventions for the Golang API template project.

## Root Directory
- `main.go`: Application entry point
- `go.mod`: Go module definition
- `RULES.md`: This file, containing project structure and guidelines

## Internal Packages

### `/internal`
Contains all application-specific code that shouldn't be imported by other applications.

#### `/internal/entity`
- **Purpose**: Domain models and business logic entities
- **Rules**:
  - Define core business objects
  - Should not contain any database or API specific code
  - Use clear, descriptive names that reflect business domain

#### `/internal/handlers`
- **Purpose**: HTTP request handlers
- **Rules**:
  - One file per resource/endpoint group
  - Handle request/response formatting
  - No business logic - delegate to services
  - Name format: `{resource}_handler.go`

#### `/internal/middlewares`
- **Purpose**: HTTP middleware components
- **Rules**:
  - One file per middleware
  - Handle cross-cutting concerns (auth, logging, etc.)
  - Name format: `{purpose}_middleware.go`

#### `/internal/repositories`
- **Purpose**: Data access layer
- **Rules**:
  - Interact with databases/external services
  - Implement repository interfaces defined in services
  - Name format: `{entity}_repository.go`

#### `/internal/request`
- **Purpose**: Request/response DTOs (Data Transfer Objects)
- **Rules**:
  - Define request/response structures
  - Include validation tags
  - Group related DTOs in the same file

#### `/internal/route`
- **Purpose**: Route definitions
- **Rules**:
  - Define API routes and their handlers
  - Group related routes together
  - Apply middlewares at the route level

#### `/internal/services`
- **Purpose**: Business logic layer
- **Rules**:
  - Implement business use cases
  - Define service interfaces
  - Call repositories for data access
  - Name format: `{domain}_service.go`

## Public Packages (`/pkg`)
Contains packages that can be imported by other applications.

### `/pkg/common`
- **Purpose**: Shared utilities and constants
- **Rules**:
  - Generic utilities used across the application
  - Constants and enums
  - Small, focused utility functions

### `/pkg/configs`
- **Purpose**: Configuration management
- **Rules**:
  - Load and validate configuration
  - Environment variable handling
  - Configuration structs and defaults

### `/pkg/database`
- **Purpose**: Database connections and migrations
- **Rules**:
  - Database connection setup
  - Migration files
  - Database utility functions

### `/pkg/util`
- **Purpose**: General utility functions
- **Rules**:
  - Pure utility functions
  - No business logic
  - Well-documented and tested

## Other Directories

### `/scripts`
- **Purpose**: Build and deployment scripts
- **Rules**:
  - One script per specific task
  - Include usage instructions in comments
  - Make scripts executable

### `/server`
- **Purpose**: Server configuration and startup
- **Rules**:
  - Server initialization
  - Middleware setup
  - Graceful shutdown handling

## General Guidelines

### Naming Conventions
- Use `camelCase` for variables and functions
- Use `PascalCase` for types, interfaces, and exported functions
- Use `UPPER_SNAKE_CASE` for constants
- Use `snake_case` for file names

### Error Handling
- Always handle errors explicitly
- Use custom error types for domain-specific errors
- Provide meaningful error messages
- Log errors with context

### Testing
- Place test files next to the code they test
- Use `_test.go` suffix for test files
- Table-driven tests for multiple test cases
- Mock dependencies for unit tests

### Code Organization
- Keep files small and focused
- Group related functionality
- Use interfaces to define behavior
- Document exported functions and types

### Dependencies
- Keep dependencies to a minimum
- Use Go modules for dependency management
- Document any non-standard dependencies

## Best Practices

1. **Keep it Simple**: Favor simple, readable code over clever solutions
2. **Single Responsibility**: Each function/struct should have one responsibility
3. **Dependency Injection**: Pass dependencies explicitly
4. **Error Handling**: Handle errors where they occur
5. **Logging**: Log important events and errors
6. **Documentation**: Document public APIs and complex logic
7. **Testing**: Write tests for critical paths
8. **Performance**: Optimize only after measuring
9. **Security**: Follow security best practices
10. **Error Messages**: Make error messages helpful and actionable
