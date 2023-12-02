# An Example For Implementing Http Server Application

## Feature

- I'm trying to implement modular monolith in clean architecture.
- This repository is just experimental. 
- Please open issue or pull request if you are interested in.

## Technical Stack

| Category | Selected | URL |
| ---- | ---- | ---- |
| Frame Work | echo | <https://echo.labstack.com/> |
| Database Client | bun | <https://bun.uptrace.dev/> |
| Dependency Injection | wire | <https://github.com/google/wire> |
| Mock | gomock | <https://pkg.go.dev/go.uber.org/mock/gomock> |
| I/F | oapi-codegen | <https://github.com/deepmap/oapi-codegen> |

## Test Policy

- tests of repository use db container without using mock, the others use mock.





