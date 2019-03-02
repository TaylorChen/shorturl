> ## This is a short url solution
> 1. This solution uses redis as a temporary storage to store long-url and short-url mapping relationships. If you want to persist, you can use the database as a permanent solution.
> 2. Solved the problem of small probability hash conflicts.
> 3. You can use nginx to build a short url domain proxy, such as my.cn. When you visit http://m.cn/SlnkP, you can directly jump to the original address.

> ## How to run
```powershell
git clone git@github.com:TaylorChen/shorturl.git
cd shorturl
go run main.go -c conf/shorturl.conf
```

> ## Support Functions
> 1. Generate a short url
```powershell
curl -d "longurl=https://en.wikipedia.org/wiki/Google" "http://127.0.0.1:8088/gen_short_url"
http://m.cn/SlnkP
````
> 2. Get a long url
```powershell
curl "http://127.0.0.1:8088/get_long_url?short=SlnkP"
https://en.wikipedia.org/wiki/Google
````
