# qlik-hackathon-2019-12

Hackathon project to implement a basic webhook that listens for Qlik events

build & run container
```
docker build --tag qlik-hackathon-2019-12-image .
docker run --name qlik-hackathon-2019-12 -p 8080:8080 -d qlik-hackathon-2019-12-image
```

clean up
```
docker rm -f qlik-hackathon-2019-12
docker rmi qlik-hackathon-2019-12-image
```

add to nginx conf
```
vi /etc/nginx/nginx.conf (add)

        location /postreceive {
            proxy_pass http://localhost:8080;
        }

sudo service nginx stop
sudo service nginx start
```