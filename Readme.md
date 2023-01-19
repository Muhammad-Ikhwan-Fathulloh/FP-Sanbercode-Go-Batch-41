# Final Project Sanbercode Go Batch 41

<ul>
<li>Nama : Muhammad Ikhwan Fathulloh</li>
<li>Final Project Bootcamp Golang Sanbercode Batch 41</li>
</ul>

### Resource :
<ul>
<li>Github : https://github.com/Muhammad-Ikhwan-Fathulloh/FP-Sanbercode-Go-Batch-41</li>
<li>Deploy Railway : https://collabs.up.railway.app/</li>
<li>Documentation Postman : https://documenter.getpostman.com/view/19977900/2s8ZDX2Mdu</li>
<li>PPT : https://docs.google.com/presentation/d/1MpGq3PKj6LE9tWqqjf8vFbIAWCWz11K6rglFTk5otg8/edit?usp=sharing</li>
</ul>

### ERD

<img src="documentation/ERD-FP Golang.drawio.png">

### Endpoint Auth/User
Endpoint ini bertanggung jawab membuat generate token untuk token ke endpoint pengelolaan.
Method | Path | Keterangan
------------- | ------------- | -------------
***POST*** | *`/login`* | Men-generate token untuk mengakses endpoint yang berfungsi sebagai pengelolaan. Token akan didapatkan setelah pengguna mengirim json berupa "*email dan password*" yang terdaftar
***GET*** | *`/users`* | Mengakses data pengguna | token
***GET*** | *`/users/:user_id`* | Mengakses data pengguna berdasakan id pengguna | token
***POST*** | *`/users`* | Membuat data pengguna baru | token
***PUT*** | *`/users/:user_id`* | Mengubah data pengguna berdasakan id pengguna | token
***DELETE*** | *`/users/:user_id`* | Menghapus data pengguna berdasakan id pengguna | token

### Endpoint Community
Endpoint ini bertanggung jawab mengelola data komunitas yang berkolaborasi.
Method | Path | Keterangan
------------- | ------------- | -------------
***GET*** | *`/communities`* | Mengakses data komunitas 
***GET*** | *`/communities/:community_id`* | Mengakses data komunitas berdasakan id komunitas 
***POST*** | *`/communities`* | Membuat data komunitas baru | token
***PUT*** | *`/communities/:community_id`* | Mengubah data komunitas berdasakan id komunitas | token
***DELETE*** | *`/communities/:community_id`* | Menghapus data komunitas berdasakan id komunitas | token

### Endpoint Kategori Acara
Endpoint ini bertanggung jawab mengelola data kategori acara.
Method | Path | Keterangan
------------- | ------------- | -------------
***GET*** | *`/event-categories`* | Mengakses data kategori acara 
***GET*** | *`/event-categories/:event_category_id`* | Mengakses data kategori acara berdasakan id kategori acara 
***POST*** | *`/event-categories`* | Membuat data kategori acara baru | token
***PUT*** | *`/event-categories/:event_category_id`* | Mengubah data kategori acara berdasakan id kategori acara | token
***DELETE*** | *`/event-categories/:event_category_id`* | Menghapus data kategori acara berdasakan id kategori acara | token

### Endpoint Acara
Endpoint ini bertanggung jawab mengelola data acara yang diadakan komunitas.
Method | Path | Keterangan
------------- | ------------- | -------------
***GET*** | *`/events`* | Mengakses data acara 
***GET*** | *`/events/:event_id`* | Mengakses data acara berdasakan id acara 
***GET*** | *`/events-by-category/:event_category_id`* | Mengakses data acara berdasakan id kategori acara 
***GET*** | *`/events-by-community/:community_id`* | Mengakses data acara berdasakan id komunitas 
***POST*** | *`/events`* | Membuat data acara baru | token
***PUT*** | *`/events/:event_id`* | Mengubah data acara berdasakan id acara | token
***DELETE*** | *`/events/:event_id`* | Menghapus data acara berdasakan id acara | token


### Endpoint Partisipan Acara
Endpoint ini bertanggung jawab mengelola data partisipan acara yang diadakan komunitas dan memberikan akses generate QR Code untuk akses masuk acara.
Method | Path | Keterangan
------------- | ------------- | -------------
***GET*** | *`/event-participants`* | Mengakses data partisipan acara 
***GET*** | *`/event-participants/:event_participant_id`* | Mengakses data partisipan acara berdasakan id partisipan acara | token
***POST*** | *`/event-participants`* | Membuat data partisipan acara baru | token
***PUT*** | *`/event-participants/:event_participant_id`* | Mengubah data partisipan acara berdasakan id partisipan acara | token
***DELETE*** | *`/event-participants/:event_participant_id`* | Menghapus data partisipan acara berdasakan id partisipan acara | token