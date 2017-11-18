
server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

 #   location /fw/ {
 #       alias /opt/www/fw/;
 #       autoindex on;
 #   }

    location /map/ {
        alias /opt/nord-ffmap-backend/outjson/;
        autoindex on;
    }

   location /fw/stable/ {
             proxy_pass http://nord.freifunk.net/firmware/stable/;
             proxy_connect_timeout 6s;
    }

  #location /firmware/ {
  #      alias /opt/firmware/;
  #      autoindex on;

  ####### Unterhalb keine Änderungen durchführen! ###############
  
 #}

  root /opt/www;
  autoindex on;

deny 2a03:2267:4e6f:7264:f6f2:6dff:fe52:b854;
deny 2a03:2267:4e6f:7264:f6f2:6dff:fe5f:1462;
deny 2a03:2267:4e6f:7264:6a72:51ff:fe62:b5ba;
deny 2a03:2267:4e6f:7264:6a72:51ff:fe62:b7a8;
deny 2a03:2267:4e6f:7264:6a72:51ff:fe48:bf9c;
deny 2a03:2267:4e6f:7264:ee08:6bff:fe2b:bdf2;
deny 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc22;
deny 2a03:2267:4e6f:7264:ee08:6bff:fea4:ca6e;
deny 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc64;
deny 2a03:2267:4e6f:7264:ee08:6bff:fe78:67c4;
deny 2a03:2267:4e6f:7264:ee08:6bff:fe78:721c;

index index.html index.htm;
  server_name localhost;
}
