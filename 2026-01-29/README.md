# Dispatch

This program receives a **JSON request** and calls the right function.

## What is dispatch?

Dispatch means "send to the right place".

When you send a command, `dispatch.py` reads it and calls the correct function in `users.py`.

## How to use

Send a JSON string as an argument:

```bash
python dispatch.py '{"command": "list"}'
```

## Commands

### List all users

```bash
python dispatch.py '{"command": "list"}'
```

### Login

```bash
python dispatch.py '{"command": "login", "email": "madoka@example.com", "password": "madoka123"}'
```

### Add a user

```bash
python dispatch.py '{"command": "add", "username": "nagisa", "email": "nagisa@example.com", "password": "nagisa123"}'
```

### Delete a user

```bash
python dispatch.py '{"command": "delete", "email": "nagisa@example.com"}'
```

### Deactivate a user

```bash
python dispatch.py '{"command": "deactivate", "email": "sayaka@example.com"}'
```

### Change email

```bash
python dispatch.py '{"command": "change-email", "email": "kyoko@example.com", "new_email": "kyoko2@example.com"}'
```

### Change password

```bash
python dispatch.py '{"command": "change-password", "email": "homura@example.com", "old_password": "homura123", "new_password": "newpass"}'
```

## Files

- `dispatch.py` - Receives JSON and calls functions
- `users.py` - User functions (CRUD)
- `accounts.json` - User data
