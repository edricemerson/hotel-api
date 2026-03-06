# Hotel Booking REST API

## Overview

This project is a **Hotel Booking REST API** that allows users to register, log in, browse hotels rooms, and make room bookings.
The system is built using a RESTful architecture and stores data in a **PostgreSQL database**.

The API manages several core resources such as **users, rooms, bookings**.

---


## Database Structure

The system uses the following main tables:

* **users** – stores registered users
* **rooms** – stores rooms belonging to hotels
* **bookings** – stores booking transactions

---

## REST API Endpoints

### Authentication

| Method | Endpoint  | Description         |
| ------ | --------- | ------------------- |
| POST   | /register | Register a new user |
| POST   | /login    | Login user          |

### Rooms

| Method | Endpoint          | Description        |
| ------ | ----------------- | ------------------ |
| GET    | /rooms            | Get all rooms      |
| GET    | /rooms/:id        | Get room by id     |
| POST   | /rooms            | Create room        |
| PUT    | /rooms/:id        | Update room        |
| DELETE | /rooms/:id        | Delete room        |

### Bookings

| Method | Endpoint            | Description              |
| ------ | ------------------- | ------------------------ |
| POST   | /bookings           | Create booking           |
| GET    | /bookings           | Get booking based on jwt |
| GET    | /bookings/:id       | Get booking details      |
| PUT    | /bookings/:id       | Update booking details   |
| DELETE | /bookings/:id       | Cancel booking           |

---

## Technology Stack

* **Backend Language:** Golang
* **Database:** PostgreSQL

---

