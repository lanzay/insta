# Instagram scraper
![Go](https://github.com/lanzay/insta/workflows/Go/badge.svg)

Scrape photos from Instagram without authorisation

Scrap by: 
* Users names
* List of users names
* Tag

Output:
BY: username, list of user name, tag
```
insta -u nasa,kyliejenner -l list.txt -t topmodel foto
```

+Webhook
```
insta -u nasa,kyliejenner -l list.txt -t topmodel -w https://myserver.ru/json foto
```

tag:
* https://www.instagram.com/explore/tags/aaa/?__a=1

PS
Собранные бинарники тут https://github.com/lanzay/insta/releases/
