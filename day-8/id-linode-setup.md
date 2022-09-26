# Membuat Linode Compute Instance - Alterra AGMC (Day 8) - Deployment

> [See content in English](./en-linode-setup.md)

## Navigasi Utama

- ➡️ Membuat Linode Compute Instance
- [Setup SSH dengan key](./id-setup-ssh-key.md)
- [Setup Docker](./id-setup-docker.md)
- [Deploy aplikasi dengan image dari Docker Registery](./id-deploy.md)
- [Test akses API dari Postman](./id-postman.md)

## Membuat Linode Compute Instance

Pada Dashboard Linode, klik **Create** dan pilih **Linode** untuk membuat linode compute instance.

![image-20220924152130631](./assets/image-20220924152130631.png)

Pada **Distribution Images** pilih **Ubuntu 22.04 LTS** dan pada **Region** pilih region misalkan Singapore.

![image-20220924152251123](./assets/image-20220924152251123.png)

Pada **Linode Plan** pilih tab **Shared CPU** kemudian pilih **Linode 2 GB**.

![image-20220924152429346](./assets/image-20220924152429346.png)

Pada **Linode Label** masukkan nama yang diinginkan, label ini akan ditampilkan pada list linodes yang kita miliki.

![image-20220924152600162](./assets/image-20220924152600162.png)

Isi Root Password.

![image-20220924152709811](./assets/image-20220924152709811.png)

Untuk SSH Keys kita skip saja dulu, karena nanti kita akan [setup manual](./id-setup-ssh-key.md) lewat terminal saja.

Lihat summary, pastikan sudah sesuai.

![image-20220924152833141](./assets/image-20220924152833141.png)

Kemudian klik **Create Linode**.

Setelah itu kita akan dibawa ke halaman detail dari server yang baru saja kita buat.

![image-20220926060530467](./assets/image-20220926060530467.png)

Tunggu hingga status **PROVISIONING** berubah menjadi **RUNNING**.

![image-20220926060855569](./assets/image-20220926060855569.png)

Server sudah siap digunakan, lanjut ke [Setup SSH dengan key](./id-setup-ssh-key.md).