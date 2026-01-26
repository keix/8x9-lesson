# User CRUD Simulation (JSON-based)

This repository contains a **simple user CRUD simulation** implemented in Python.

The goal of this project is **not** to teach frameworks or UI, but to help learners understand **how programs actually work**.

## What this is

- A small Python program
- Users are stored in a local JSON file
- Each operation is implemented as a **small function**
- No database
- No framework
- No magic

This simulates the **core logic of a Web application**.

## What you can do

The program supports basic user operations:

- Create a user
- Read user data
- Update user data
- Deactivate a user (delete-like behavior)

In other words: **CRUD**.

All data is stored in `accounts.json`.

## Why JSON?

JSON is just data.

- Easy to read
- Easy to edit
- Used everywhere (Web APIs, configs, databases)

A real Web app does the same thing,
but stores this data in a database instead of a file.

## Why no classes?

Classes are useful, but they add concepts that beginners do not need at first.

This project uses:

- `dict`
- `list`
- `function`

This makes **data flow and behavior visible**.

## Local or server?

This code runs locally.

A server-based Web app does the **same thing**, but calls these functions through HTTP. The difference is **how functions are called**,
not **what they do**.

## Files

- `users.py`  
  User-related logic (CRUD functions)

- `accounts.json`  
  User data (acts like a database)

## How to run

```bash
python users.py
```
