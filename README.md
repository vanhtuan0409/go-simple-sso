# Simple Single-Sign-On

### Features

* [x] Cross domain support
* [x] In-mem datastore for user and session
* [ ] Role-based access restriction
* [ ] Logout across devices

### Summary

Single sign-on (SSO) is a session and user authentication service that permits a user to use one set of login credentials (e.g., name and password) to access multiple applications (or website). The service authenticates the end user for all the applications the user has been given rights to and eliminates further prompts when the user switches applications during the same session. On the back end, SSO is helpful for logging user activities as well as monitoring user accounts.

Take Google as an example. You only need to login once but you can access to all of google services like Gmail, Calendar, Drive, ...

![diagram](https://drive.google.com/file/d/1Fo2HtYKDA5tdZXKGBhtiZuNhq_kPkuAW/view)

### Installation

```
git clone https://github.com/vanhtuan0409/go-simple-sso.git
cd go-simple-sso
docker-compose up
```

Edit your `hosts` file to add to following lines.

* For windows, it is `C:\Windows\System32\Drivers\etc\hosts`
* For mac, it is `/etc/hosts`

```
127.0.0.1 web1.com
127.0.0.1 web2.com
127.0.0.1 login.com
```

If you can go to **web1.com** or **web2.com** now, you will be redirect to **login.com**. After login, you can access both of these 2 website.
