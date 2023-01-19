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
Sistem Informasi Acara Komunitas dan Tiket ini memiliki 5 tabel :
<ul>
<li>Tabel users</li>
<li>Tabel communities</li>
<li>Tabel event_categories</li>
<li>Tabel events</li>
<li>Tabel event_participants</li>
</ul>
<img src="documentation/ERD-FP Golang.drawio.png">

### API
Sistem Informasi Acara Komunitas dan Tiket ini memiliki 5 endpoint :
## Endpoint Main
Endpoint utama.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/`* | Menampilkan identitas pembuat api

## Endpoint Auth/User
Endpoint ini bertanggung jawab membuat generate token untuk token ke endpoint pengelolaan.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***POST*** | *`/api/login`* | Men-generate token untuk mengakses endpoint yang berfungsi sebagai pengelolaan. Token akan didapatkan setelah pengguna mengirim request json berupa email dan password yang sudah terdaftar
***GET*** | *`/api/secured/users`* | Mengakses data pengguna | token
***GET*** | *`/api/secured/users/:user_id`* | Mengakses data pengguna berdasakan id pengguna | token
***POST*** | *`/api/secured/users`* | Membuat data pengguna baru | token
***PUT*** | *`/api/secured/users/:user_id`* | Mengubah data pengguna berdasakan id pengguna | token
***DELETE*** | *`/api/secured/users/:user_id`* | Menghapus data pengguna berdasakan id pengguna | token

## Endpoint Community
Endpoint ini bertanggung jawab mengelola data komunitas yang berkolaborasi.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/api/communities`* | Mengakses data komunitas 
***GET*** | *`/api/communities/:community_id`* | Mengakses data komunitas berdasakan id komunitas 
***POST*** | *`/api/secured/communities`* | Membuat data komunitas baru | token
***PUT*** | *`/api/secured/communities/:community_id`* | Mengubah data komunitas berdasakan id komunitas | token
***DELETE*** | *`/api/secured/communities/:community_id`* | Menghapus data komunitas berdasakan id komunitas | token

## Endpoint Kategori Acara
Endpoint ini bertanggung jawab mengelola data kategori acara.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/api/event-categories`* | Mengakses data kategori acara 
***GET*** | *`/api/event-categories/:event_category_id`* | Mengakses data kategori acara berdasakan id kategori acara 
***POST*** | *`/api/secured/event-categories`* | Membuat data kategori acara baru | token
***PUT*** | *`/api/secured/event-categories/:event_category_id`* | Mengubah data kategori acara berdasakan id kategori acara | token
***DELETE*** | *`/api/secured/event-categories/:event_category_id`* | Menghapus data kategori acara berdasakan id kategori acara | token

## Endpoint Acara
Endpoint ini bertanggung jawab mengelola data acara yang diadakan komunitas.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/api/events`* | Mengakses data acara 
***GET*** | *`/api/events/:event_id`* | Mengakses data acara berdasakan id acara 
***GET*** | *`/api/events-by-category/:event_category_id`* | Mengakses data acara berdasakan id kategori acara 
***GET*** | *`/api/events-by-community/:community_id`* | Mengakses data acara berdasakan id komunitas 
***POST*** | *`/api/secured/events`* | Membuat data acara baru | token
***PUT*** | *`/api/secured/events/:event_id`* | Mengubah data acara berdasakan id acara | token
***DELETE*** | *`/api/secured/events/:event_id`* | Menghapus data acara berdasakan id acara | token


## Endpoint Partisipan Acara
Endpoint ini bertanggung jawab mengelola data partisipan acara yang diadakan komunitas dan memberikan akses generate QR Code untuk akses masuk acara.
Method | Path | Keterangan | Auth
------------- | ------------- | ------------- | -------------
***GET*** | *`/api/event-participants`* | Mengakses data partisipan acara 
***GET*** | *`/api/secured/event-participants/:event_participant_id`* | Mengakses data partisipan acara berdasakan id partisipan acara | token
***POST*** | *`/api/secured/event-participants`* | Membuat data partisipan acara baru | token
***PUT*** | *`/api/secured/event-participants/:event_participant_id`* | Mengubah data partisipan acara berdasakan id partisipan acara | token
***DELETE*** | *`/api/secured/event-participants/:event_participant_id`* | Menghapus data partisipan acara berdasakan id partisipan acara | token