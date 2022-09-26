

# Alterra AGMC (Day 8) - Deployment

## Task

**Deployment EC2**

1. Membuat VM di EC2, dan Implementasi Security Group EC2.
2. Melakukan SSH Remote ke VM di AWS EC2 (Dengan Key) serta dijelaskan key dan password.
3. Deploy your Program to EC2.

**Deployment RDS**

1. Membuat DB di RDS.
2. Migrate Your Local Data to RDS.
3. Connect Your Application to RDS.

## Deployment

Disini saya tidak menggunakan **AWS** tetapi menggunakan [Linode](https://www.linode.com/) sebagai target deployment.

Aplikasi sudah ter-deploy dan bisa diakses di http://172.104.44.239:15000/status

Postman collection tersedia di [sini](./Alterra%20AGMC%20Day%208.postman_collection.json).

Berikut langkah-langkah yang gunakan saat melakukan deploy:

- [Membuat Linode Compute Instance](./id-linode-setup.md)
- [Setup SSH dengan key](./id-setup-ssh-key.md)
- [Setup Docker](./id-setup-docker.md)
- [Deploy aplikasi dengan image dari Docker Registery](./id-deploy.md)
- [Test akses API dari Postman](./id-postman.md)

## Server Spec

**Linode 2 GB**

- Ubuntu 22.04 LTS
- 1 CPU, 50 GB Storage, 2 GB RAM
- 2 TB Transfer
- 40 Gbps In / 2 Gbps Out