# Postman - Alterra AGMC (Day 8) - Deployment

> [See content in English](./en-postman.md)

## Navigasi Utama

- [Membuat Linode Compute Instance](./id-linode-setup.md)
- [Setup SSH dengan key](./id-setup-ssh-key.md)
- [Setup Docker](./id-setup-docker.md)
- [Deploy aplikasi dengan image dari Docker Registery](./id-deploy.md)
- ➡️ Test akses API dari Postman

## Daftar Isi

- [Postman](#postman)
- [Route Status](#route-status)
- [Route Auth](#route-auth)
  - [Signup](#signup-post-authsignup)
  - [Login](#login-post-authlogin)
- [Route Books](#route-books)
  - [Get all books](#get-all-books-get-apibooks)
  - [Create book](#create-book-post-apibooks)
  - [Get book by ID](#get-book-by-id-get-apibooksid)
  - [Update book](#update-book-put-apibooksid)
  - [Delete book](#delete-book-delete-apibooksid)
- [Route User](#route-user)
  - [Get all users](#get-all-users-get-apiusers)
  - [Get user by ID](#get-user-by-id-get-apiusersid)
  - [Update user](#update-user-put-apiusersid)
  - [Delete user](#delete-user-delete-apiusersid)

## Postman

Mari kita coba test akses API menggunakan Postman.

Berikut collection variable yang saya gunakan.

![image-20220926081015419](./assets/image-20220926081015419.png)

Jika ingin menggunakan collection yang sama bisa download di [sini](./Alterra%20AGMC%20Day%208.postman_collection.json).

## Route Status

![image-20220926081137816](./assets/image-20220926081137816.png)

## Route Auth

### Signup (POST `/auth/signup`)

![image-20220926081341641](./assets/image-20220926081341641.png)

### Login (POST `/auth/login`)

![image-20220926081813709](./assets/image-20220926081813709.png)

## Route Books

### Get all books (GET `/api/books`)

![image-20220926081837406](./assets/image-20220926081837406.png)

### Create book (POST `/api/books`)

![image-20220926081919642](./assets/image-20220926081919642.png)

### Get book by ID (GET `/api/books/:id`)

![image-20220926081956267](./assets/image-20220926081956267.png)

### Update book (PUT `/api/books/:id`)

![image-20220926082020741](./assets/image-20220926082020741.png)

### Delete book (DELETE `/api/books/:id`)

![image-20220926082043821](./assets/image-20220926082043821.png)

## Route User

### Get all users (GET `/api/users`)

![image-20220926082111487](./assets/image-20220926082111487.png)

### Get user by ID (GET `/api/users/:id`)

![image-20220926082126902](./assets/image-20220926082126902.png)

### Update user (PUT `/api/users/:id`)

![image-20220926082144619](./assets/image-20220926082144619.png)

### Delete user (DELETE `/api/users/:id`)

![image-20220926082201191](./assets/image-20220926082201191.png)
