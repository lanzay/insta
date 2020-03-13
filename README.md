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

```
https://www.instagram.com/explore/locations/359545221/moscow/
https://www.instagram.com/explore/locations/17326249/moscow-russia/
https://www.instagram.com/explore/locations/1206739199510403/
https://www.instagram.com/explore/locations/327275377680277/
https://www.instagram.com/explore/locations/1/
```