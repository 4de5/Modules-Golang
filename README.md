# PENGENALAN GO MODULES
Go-Module mulai 1.11 dan fix di 1.12 || Nonton video-nya PZM aja

## MEMBUAT MODULE
1. go mod init <nama-module>
2. Go akan auto membuat file go.mod
3. Udah ga perlu main GOPATH

## MERILIS MODULE
1. Punya akun GitHub dulu, karena ini bisa berinteraksi dengan Git
2. Buat projek di GitHub (ex: <go-say-hello>) ga pake README
3. copy URL (ex: <github.com/asepdwisaputra/go-say-hello>)
4. buat folder sesuai nama Repo (ex: <go-say-hello>)
5. buka IDE VSC masuk folder tadi >> Terminal
6. buat init module (<go mod init github.com/asepdwisaputra/go-say-hello>)
7. akan muncul file go.mod
8. buat file baru di go-modules (ex: say_hello.go)
9. buat program SayHello >> git init >> git add go.mod >> git add say_hello.go >> git status >> git commit -m 'membuat module say-hello'
	- git untuk folder go-say-hello
10. buka GitHub copy lokasi(samping https/ssh <github.com/asepdwisaputra/go-modules.git>)
11. git remote add origin https://github.com/asepdwisaputra/go-modules.git
    - sebenarnya ada juga di bagian tengah tabel
12. git push origin master >> refresh web github >> refresh web
13. Membuat TAG :
	- git tag v1.0.0
	- git push origin v1.0.0 >> Refresh Web

##  MENAMBAH DEPENDENCY
1. buat dan buka (ex: app-say-hello) di VSC
2. buat repo (ex: app-say-hello) di Web GitHub
3. copas lokasi >> go mod init github.com/asepdwisaputra/app-say-hello
4. buat main file (main.go)
5. get module(bisa cek go.mod folder go-say-hello) yg tadi biar bisa impor (ex: go get github.com/asepdwisaputra/go-say-hello)
 6. tunggu download >> muncul go.sum >> dan ada tambahan baris di go.mod (require.....)
7. coba panggil func SayHello
8. di tutor gada git init >> git add .  >> git status >> git commit -m 'membuat module say-hello' >> git remote add ... >> git push ....
9. tapi tak coba keluar dari TUTORIAL

## UPDATE MODULE
- digunakan untuk upgrade aplikasi(menambah tag baru)
	1. edit file go-say-hello (helloworld)
	2. git add namafile.go >> git commit -m 'pesan' [disini saya pake git commit -am 'update hello world']
		// Kenapa pake git commit -am?? cek catatan git, pernah nulis
	3. git push origin master
	4. git tag v1.1.0
		// saat melakukan perubahan harus menggunakan tag berbeda, ini penting
	5. git push origin v1.1.0
	6. go-say-hello punya 2 tag di GH

## UPDATE DEPENDENCY
1. bisa ubah go.mod versinya(app-say-hello) bagian v1.0.0 ganti v1.1.0
2. go get
3. atau hapus bagian baris required di go.mod >> go get
4. PEMULA GA USAH BIKIN REPO app-say-hello

## Major Upgrade
- merubah secara besar penulisan kode, sehingga tidak backward compatible. Biasanya ini SANGAT DIHINDARI, namun apabila ini jalan terakhir karena ada bug/kondisi keamanan maka perlu di lakukan. modul berubah, tag v2.0.0
	1. ubah modul go-say-hello, ubah func SayHello dengan wajib menggunakan parameter nama string
		- ini jadi major karena otomatis kode SayHello() eror semua --> SayHello("Asep")
	2. pergi ke go.mod >> pada bagian module tambahi /v2 di akhir, agar saat go get tidak langsung mengambil versi v2.0.0(tidak misscom.)
	3. git add go.mod say-hello.go
	4. git commit -m 'update say-hello.go with parameter'
	5. git push origin master >> git tag v2.0.0 >> git push origin v2.0.0
	6. cek GH ada v2
	7. kita coba go get di app-say-hello
		- buka go.mod >> hapus baris require
		- go get lokais v2(github.com/asepdwisaputra/go-say-hello/v2)
		- main.go >> edit&tambah parameter nama
