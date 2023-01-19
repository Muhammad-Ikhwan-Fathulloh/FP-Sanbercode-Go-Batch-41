# Final Project Sanbercode Go Batch 41

### Endpoint Auth
Endpoint ini bertanggung jawab membuat generate token untuk token ke endpoint pengelolaan.
Method | Path | Keterangan
------------- | ------------- | -------------
***POST*** | *`/login`* | Men-generate token untuk mengakses endpoint yang berfungsi sebagai pengelolaan. Token akan didapatkan setelah pengguna mengirim json berupa "*email dan password*" yang terdaftar