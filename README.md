# go-beginner

Golang Project for Onboarding Fresher

## Dựng 1 api service

### Các tính năng:

1. Tạo user mới: POST /api/user
2. Login: POST /api/auth/login
3. Refresh token: POST /api/auth/refresh-token
4. Get user info: GET /api/user
5. Update user info: PUT /api/user
6. Update password: PUT /api/user/password

- Các API 1, 2, 3 không cần authen
- Các API còn lại authen bằng JWT token

### Database schema

Table `user`:

- id
- created_at
- updated_at
- deleted_at
- name: unique
- email
- hashed_password

## Build docker image

Viết dockerfile và build thành image, run bằng docker/docker compose
