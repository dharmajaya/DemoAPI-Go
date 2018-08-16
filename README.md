# Instructions for the Demo API - Go

*Follow the Steps to make it Run Successfully.

*Operating systems can be Ubuntu 16.0 or Above , CentOS 6.0 or Above

*I have configured with Ubuntu 18.04

Steps
-----

* Please use to install with sudo or root user privileges


Please make sure with the following
====================================

  * [Ubuntu VM or OS 14.0 or Above](http://www.ubuntu.com)
  * [Go 1.6+](https://golang.org/dl)
  * [Redis](http://redis.io)
  * [MySQL](https://www.mysql.com)
  * [Gmail Account](https://accounts.google.com/SignUp)

* **Requirements for the Demo API**
* User Interaction Logic Flow
* User Login & Logout
* User send verification code through sms code
* Feature to prevent the sms api with spam
* User Reset password
* User send Reference link through email
* User can join with Reference Link
* Count the joined members through reference link
* Check how many clicks with the reference link
* News Center
* Promote News
* News
* Board News show with list
* Go Language and Beego Orm Framework

Choosen the Frameworks
=======================
* Choosen the Language & Framework as GO Lang & Beego(Mvc) Orm Framework
* Session Management with Redis (Handling the Session through Redis Server)
* Security Used Capcha Security

1. Setup and configure required packages for the ubuntu os or ubuntu Virtual Machine
2. sudo apt-get update
3. Check the mysql server is installed or not. If not install a new mysql server sudo apt-get install mysql-server-5.6 or more
4. Check the mysql is working or not. mysql -u username -p  password (mysql > )
5. Create a sample database for the demo api (create database api-demo)

mysql> show databases;
      mysql> show databases;
      +--------------------+
      | Database           |
      +--------------------+
      | information_schema |
      | api_demo           |
      +--------------------+

6. Install make tool -- sudo apt-get install make gcc tcl
7. Downlaod and Install Redis Server (if wget is not working then install the wget )
  wget http://download.redis.io/releases/latest-release.tar.gz
8. Move to /opt directory using sudo tar -C /opt -xzf redis-latest-relase.tar.gz
9. Go to the /opt/redis-latest-relase directry and check the redis server installation

  make
  -----
  make test
  ----------
10. If everyting is working fine then redis server is ready to use it.
11. Start the Redis Server  /opt/redis-3.2.0/src/redis-server &
 if everything is fine then you will see as follows

 11949:M 16 Aug 2018 08:49:26.575 # Current maximum open files is 4096. maxclients has been reduced to 4064 to compensate for low ulimit. If you need higher maxclients increase 'ulimit -n'.
                 _._
            _.-``__ ''-._
       _.-``    `d.  `_.  ''-._           Redis 4.9.104 (00000000/0) 64 bit
   .-`` .-```.  ```\/    _.,_ ''-._
  (    '      ,       .-`  | `,    )     Running in standalone mode
  |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
  |    `-._   `._    /     _.-'    |     PID: 11949
   `-._    `-._  `-./  _.-'    _.-'
  |`-._`-._    `-.__.-'    _.-'_.-'|
  |    `-._`-._        _.-'_.-'    |           http://redis.io
   `-._    `-._`-.__.-'_.-'    _.-'
  |`-._`-._    `-.__.-'    _.-'_.-'|
  |    `-._`-._        _.-'_.-'    |
   `-._    `-._`-.__.-'_.-'    _.-'
       `-._    `-.__.-'    _.-'
           `-._        _.-'
               `-.__.-'

 11949:M 16 Aug 2018 08:49:26.576 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.

 Port will be 6379

 12. Download the golanguage latest version and install
 wget https://storage.googleapis.com/golang/go(latest relase).linux-amd64.tar.gz

 13. Extract the golangauge sudo tar -C /usr/local -xzf go(latest release).linux-amd64.tar.gz

 14. Update the GOlocation with profile

 echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile

 source /etc/profile

 15. Setup the Go Environment path

 mkdir usergolanguage directory/gocode

 echo 'export GOPATH=$golanguage directory path/gocode' >> .bashrc
 echo 'export PATH=$PATH:$GOPATH/bin' >> .bashrc

 source .bashrc

 16. Almost the Project is set up is ready
 17 . The configuration of the project directory structuse as follows

  * **Main configurations with the project**
  * ** conf - app.conf -> All the information about the database , url, max idle , connection etc. **
  * **models --> entity relationship configurations , userobjects. **
  * **controllers --> All ther controllers for the requirments . ex. login and logout **
  * **views -->  All the front end templates configurations **
  * ** static --> Default static image, jquery file configruations **
  * ** main.go --> Start or Initiate the Project**


  * **ProjectName ---Demoapi**
    |----- assest -- user requirments image and images
    |----- conf
            |--- app.conf
    |----- controllers
            |--- main.go
            |--- user_delete.go
            |--- user_encryption.go
            |--- user_erro_handling_pages.go
            |--- user_index.go
            |--- user_joined_reference.go
            |--- user_news.go
            |--- user_phone_verification.go
            |--- user_profile.go
            |--- user_reference_generate.go
            |--- user_resetpassword.go
            |--- user_security.go
            |--- user_signin.go
            |--- user_signup.go
            |--- user_verify.go
    |----- models
            |--- account.go
            |--- forms.go
    |----- routers
            |--- router.go
    |----- static
            |--- css
            |--- img
            |--- jquery
            |--- js
            |---twbs
    |----- test
            |--- default_test.go
    |----- vendor
            |--- github.com
            |--- golang.com
            |--- vendor.json
    |----- views
            |--- deletecontroller
                 |--- get.tpl
                 |--- post.tpl
            |--- indexcontroller
                 |--- get.tpl
            |--- newscontroller
                 |--- footer.tpl
                 |--- get.tpl
                 |--- post.tpl
            |--- profilecontroller
                 |--- footer.tpl
                 |--- get.tpl
                 |--- post.tpl
            |--- referencecontroller
                 |--- footer.tpl
                 |--- get.tpl
                 |--- post.tpl
            |--- resetpasswordcontroller
                 |--- get.tpl
                 |--- post.tpl
            |--- securitycontroller
                 |--- footer.tpl
                 |--- get.tpl
                 |--- post.tpl
            |--- shared
                 |--- footer.tpl
                 |--- header.tpl
                 |--- layout.tpl
            |--- signincontroller
                 |--- get.tpl
                 |--- post.tpl
            |--- singupcontroller
                 |--- footer.tpl
                 |--- get.tpl
                 |--- post.tpl
            |--- error_page.tpl
    |--- main.go
    |--- README.md

  18. update the import packages
  19. if its not working use with go get -u packagename
   ex: go get -u github.com/astaxie/beego

  20. Before run the project please check the redis server and mysql server is running on the operating system or virtual machine

  21. If there is not compilation error then run with main.go

  22. When we run main.go it will create the tables automatically. No need to create externally

   * **Build the project go build main.go**
   * **Run the Project go run main.go**


  mysql> show tables;
  +--------------------+
  | Tables_in_api_demo |
  +--------------------+
  | market             |
  | profile            |
  | user               |
  +--------------------+
  3 rows in set (0.00 sec)


mysql> desc market;
+--------------+--------------+------+-----+---------+-------+
| Field        | Type         | Null | Key | Default | Extra |
+--------------+--------------+------+-----+---------+-------+
| uid          | varchar(255) | NO   | PRI | NULL    |       |
| asset_name   | varchar(255) | NO   |     |         |       |
| new_price    | double       | NO   |     | 0       |       |
| board_news   | varchar(255) | NO   |     |         |       |
| news         | varchar(255) | NO   |     |         |       |
| promote_news | varchar(255) | NO   |     |         |       |
+--------------+--------------+------+-----+---------+-------+

mysql> desc profile;
+---------------------+--------------+------+-----+---------+-------+
| Field               | Type         | Null | Key | Default | Extra |
+---------------------+--------------+------+-----+---------+-------+
| uid                 | varchar(255) | NO   | PRI | NULL    |       |
| username            | varchar(255) | NO   |     |         |       |
| reference_id        | varchar(255) | NO   |     |         |       |
| joined_date         | datetime     | NO   |     | NULL    |       |
| registers_count     | int(11)      | NO   |     | 0       |       |
| click_count         | int(11)      | NO   |     | 0       |       |
| password_reset_date | datetime     | NO   |     | NULL    |       |
+---------------------+--------------+------+-----+---------+-------+
7 rows in set (0.00 sec)



Check it the images from /assest/

Description :
------------
During the shortime i had learned and did R&D.

This project can run with out any bug. Need to check with all the api for save and update etc..

I Can improve the go lanaguage with as soon as possible. I believe in myself .

But i can implement it. Might be need some time. ( i can adopt very quickly)

Check Screenshots Manually
---------------------------
Refer the /asset/ folder

ScreenShots Links
-----------------
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/database2018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/database12018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/database22018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/code12018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/code2018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/Runpage2018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/DemoAPIGO.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/homepage12018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/homepage2018-08-16.png)
![](https://github.com/dharmajaya/DemoAPI-Go/blob/master/assets/singuppage2018-08-16.png)

* ** Thanks for the Great Opportunity**