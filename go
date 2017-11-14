
server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

    location /fw/ {
        alias /opt/www/fw/;
        autoindex on;
    }

    location /map/ {
        alias /opt/nord-ffmap-backend/outjson/;
        autoindex on;
    }

   location /fw/stable/factory {
             proxy_pass http://nord.freifunk.net/firmware/stable/factory/;
             proxy_connect_timeout 6s;
    }

  location /firmware/ {
        alias /opt/firmware/;
        autoindex on;

#Bar_Freifunk_841
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe7f:8fdc;

#NDS-MailBoxes
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc02;

#FF-Hanstedt-20
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bfa6;

#Lambda3-by-Gyrdon
allow 2a03:2267:4e6f:7264:227:19ff:fec4:b698;

#Sonnenau
2allow a03:2267:4e6f:7264:56e6:fcff:fef1:254a;

#Freifunk-OpenWLAN
allow 2a03:2267:4e6f:7264:ee08:6bff:fe73:72ac;

#FF-BU-Burger-Grill_BA3a
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:d27c;

#FF-HB-050
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:6bf6;

  ####### Unterhalb keine Änderungen durchführen! ###############

deny all;
 }

  root /opt/www;
  autoindex on;
  index index.html index.htm;
  server_name localhost;
}
