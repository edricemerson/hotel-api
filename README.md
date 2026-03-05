# Hotel Booking REST API

## Overview

This project is a **Hotel Booking REST API** that allows users to register, log in, browse hotels and rooms, and make room bookings.
The system is built using a RESTful architecture and stores data in a **PostgreSQL database**.

The API manages several core resources such as **users, hotels, rooms, bookings, and payments**.

---


## Database Structure

The system uses the following main tables:

* **users** – stores registered users
* **hotels** – stores hotel information
* **rooms** – stores rooms belonging to hotels
* **bookings** – stores booking transactions
* **booking_payments** – stores payment information for bookings

A **database view (`booking_report`)** is used to generate booking reports by joining users, hotels, rooms, and bookings.

---

## REST API Endpoints

### Authentication

| Method | Endpoint  | Description         |
| ------ | --------- | ------------------- |
| POST   | /register | Register a new user |
| POST   | /login    | Login user          |

### Hotels

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| GET    | /hotels     | Get all hotels    |
| GET    | /hotels/:id | Get hotel details |
| POST   | /hotels     | Create a hotel    |
| PUT    | /hotels/:id | Update hotel      |
| DELETE | /hotels/:id | Delete hotel      |

### Rooms

| Method | Endpoint          | Description        |
| ------ | ----------------- | ------------------ |
| GET    | /rooms            | Get all rooms      |
| GET    | /rooms/:id        | Get room details   |
| GET    | /hotels/:id/rooms | Get rooms by hotel |
| POST   | /rooms            | Create room        |
| PUT    | /rooms/:id        | Update room        |
| DELETE | /rooms/:id        | Delete room        |

### Bookings

| Method | Endpoint            | Description              |
| ------ | ------------------- | ------------------------ |
| POST   | /bookings           | Create booking           |
| GET    | /bookings           | Get all bookings         |
| GET    | /bookings/:id       | Get booking details      |
| GET    | /users/:id/bookings | Get user booking history |
| DELETE | /bookings/:id       | Cancel booking           |

### Reports

| Method | Endpoint          | Description        |
| ------ | ----------------- | ------------------ |
| GET    | /reports/bookings | Get booking report |

---

## Technology Stack

* **Backend Language:** Golang
* **Database:** PostgreSQL

---
