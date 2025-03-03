Go backend for Booking (Multi-User)
==================

### How to start

```bash
go build
./booking-go
```

### Config 

```yml
db:
  path: db.sqlite    # path to the database
  resetonstart: true # reset data on server restart
server:
  url: "http://localhost:3000"
  port: ":3000"
  cors:
    - "*"
```
