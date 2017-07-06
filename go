
server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

    location /fw/ {
        alias /opt/www/fw/;
        autoindex on;
    }

   location /fw/stable/factory {
             proxy_pass http://nord.freifunk.net/firmware/stable/factory/;
             proxy_connect_timeout 6s;
    }

  location /firmware/ {
        alias /opt/firmware/;
        autoindex on;

#FF-BU-Buchholz_Gemeinde_2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:3128;

#ff-nord-computerschuster
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:ea84;

#1.8.2017
#ff-nord-computerschuster
allow 	2a03:2267:4e6f:7264:f6f2:6dff:fedc:ea84;

#Lambda3-by-Gyrdon
allow 2a03:2267:4e6f:7264:227:19ff:fec4:b698;
# und mehr:


allow 2a01:4f8:162:74a2::11;
allow 2001:41d0:401:3100::416c;

allow 2a03:2267:4e6f:7264:6666:b3ff:fec2:509a;

#### Trage deine Knoten hier ein ######
#NDS-MailBoxes
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc02;

#Fremdenzimmer_Westphal
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe6b:dbc7;

#freifunk-Windbergen_A
allow 2a03:2267:4e6f:7264:62e3:27ff:febe:1c78;

#Sonnenau
allow 2a03:2267:4e6f:7264:56e6:fcff:fef1:254a;

#Meerkabarett-Eventhalle-1
allow 2a03:2267:4e6f:7264:62e3:27ff:fefb:4bcd;




#FF-Elpersb-HausDithmarschen-00
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:acee;

#FF-BU-Woe-Chstr12
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e4b8;

#STD-01
allow 2a03:2267:4e6f:7264:c66e:1fff:fe2c:79c2;
#RR
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:b462;

#behrens_reisen
allow 2a03:2267:4e6f:7264:9ade:d0ff:feab:5114;

#Lebaekee
allow 2a03:2267:4e6f:7264:ea94:f6ff:fea1:e6bc;

#28052017
allow 2a03:2267:4e6f:7264:cabe:19ff:fe6f:3cd9;
#FF-BU-WH-Heide_Haus_2-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9f04;

#NDS-FFNH-0114-Handeloh-AWCV
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e27e;
#NDS-FFNH-0517-Handeloh-AWCV
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:7aac;

#Moevenberg-5
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:35c8;
#Moevenberg-6
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3702;

#freifunk_hagen_sk_04
allow 2a03:2267:4e6f:7264:16cc:20ff:feb1:f38;
#freifunk_hagen_sk_09
allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:6536;
#freifunk_hagen_sk_06
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:be20;


#24052017

#Piraten-Stade
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe3b:6772;

#Schuh-Friedrich
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:7d60;

#freifunk_hagen_hi_2a
allow 2a03:2267:4e6f:7264:9ade:d0ff:fec1:f762;
#freifunk_hagen_ds_3_vorne
allow 2a03:2267:4e6f:7264:62e3:27ff:feec:e12a;
#freifunk_hagen_ds_9a
allow 2a03:2267:4e6f:7264:9ade:d0ff:fec2:138;
#freifunk_hagen_ds_11
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:79d8;

#freifunk_hagen_ds_3_dach
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe6f:340;
#freifunk_hagen_hi_2a
allow 2a03:2267:4e6f:7264:9ade:d0ff:fec1:f762;
#freifunk_hagen_hi_01a
allow 2a03:2267:4e6f:7264:16cc:20ff:fe3b:5864;
#freifunk_hagen_ha_26
allow 2a03:2267:4e6f:7264:ee08:6bff:fe56:b922;
#freifunk_hagen_ha_28
allow 2a03:2267:4e6f:7264:32b5:c2ff:fee0:b2;

#freifunk_hagen_hi_22
allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:5a6;

#freifunk_hagen_hi_18_v2
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:4b1e;
#freifunk_hagen_hi_18_v2
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:4b1e;
#freifunk_hagen_hi_18_Dach_v2
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:be1e;


#NDS-Hue-Jugendfreizeitheim-0290
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:c930;
#NDS-Hue-Jugendfreizeitheim-1106
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:eefe;
#NDS-Hue-Jugendfreizeitheim-1105
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:d120;


#MJ1-Freifunk
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:68a0;

#FF-HB-GrWi-Muehl-8
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:6b0e;
#FF-HB-GrWi-Muehl-8-1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:40be;

#drestedt_asyl
allow 2a03:2267:4e6f:7264:6a72:51ff:fe2e:3bd7;
#drestedt_offloader
allow 2a03:2267:4e6f:7264:218:71ff:fe4d:4b38;

#FF-STD-BUX-Altstadt1
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:458e;

#GeWo317
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:be9a;

#FitnessparkFriesland
allow 2a03:2267:4e6f:7264:ee08:6bff:fe82:9cae;

#FF-BU-Wilster_Eiscafe-Rialto_2
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e54c;

#drestedt_asyl_jupiter
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:b16e;
#https://mesh.nord.freifunk.net/#!v:m;n:60e327bdb16e

#Rosengasse
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:fffa;

#FF-BU-Woe-CNS6
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef32;

#FF_SL_Husumer_Str_001
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;

#FF-HB-Bist-Muehl-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:c988;

#freifunk-nazar-kebabhaus
allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;

#FreifunkGUWarder
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:b0a2;

#Prontera
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8640;

#FFNH-Undeloh-HEZ1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe5f:1496;

#FF-BU-St_Annen-Parkplatz1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:36e0;

#FF-BU-St_Annen-Kirche
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:49c4;

#FF_SL_Husumer_Str_001
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;
#https://mesh.nord.freifunk.net/#!v:m;n:e8de27fd2cce

#FF-RD_BI-26B
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1d0;
#FF-RD_BI-26A
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:c6b0;

#FFNH-Deutsches-Haus-02
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:7bb6;
#FFNH-Deutsches-Haus
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe30:2274;

#FFNH-Undeloh-HH01
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:ba42;

#MED-Rathausplatz-OFF-01
allow 2a03:2267:4e6f:7264:219:99ff:fe7a:9115;
#MED-Rathausplatz-FF-01
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:e06c;
#MED-Rathausplatz-FF-02
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d1d4;
#Hartmann_Moden
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:5526;
#Tee-Speicher
allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:602;

#Meinsbur_03
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefb:5e38;
#Meinsbur_04
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:7db4;
#Meinsbur_02
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefb:5cd6;
#Freifunk-Meinsbur
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:1226;

#Lotto_Exter
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:f87a;

##offline
#Elbmarsch_507
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a34;
#https://mesh.nord.freifunk.net/#!v:m;n:60e327c70a34

#FF-RD-012
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:a068;
#https://mesh.nord.freifunk.net/#!v:m;n:60e327b7a068



#SID6581
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:7b20;

#freifunk_hagen_sk_09
allow fdda:fee6:187:0:62e3:27ff:fe60:6536;
#freifunk_hagen_sk_06
allow fdda:fee6:187:0:62e3:27ff:febd:be20;

# FB-3
allow fdda:fee6:187:0:f6f2:6dff:fedc:388c;
#FB-1
allow fdda:fee6:187:0:f6f2:6dff:fedc:3d6e;
#spaeter: 
#FB-2
#allow fdda:fee6:187:0:f6f2:6dff:fedc:3c1e;
# was ist mit dem test router?
#allow fdda:fee6:187:0:62e3:27ff:fece:b9aa;


#GeWo316
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c0dc;

#Elbmarsch_203
allow 2a03:2267:4e6f:7264:62e3:27ff:fea1:7f74;
#Elbmarsch_202
allow 2a03:2267:4e6f:7264:62e3:27ff:fed6:1580;
#Elbmarsch_011_Rot_2
allow 2a03:2267:4e6f:7264:32b5:c2ff:fec2:6198;
#Elbmarsch_205_Gelb
allow 2a03:2267:4e6f:7264:c66e:1fff:fefe:c35a;


#NDS-FFNH-170-Hanstedt-012
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6854;

#GeWo316
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c0dc;

#freifunk14cc205a76d6
allow 2a03:2267:4e6f:7264:16cc:20ff:fe5a:76d6;

#GeWo313
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c3b2;

#GeWo1821
allow 2a03:2267:4e6f:7264:8616:f9ff:fe5b:8800;

#FFNH-Juz3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:4600;

#GeWo1825
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:b864;

#Elbmarsch_204_Rot_1
allow 2a03:2267:4e6f:7264:12fe:edff:fec2:3540;


#23.05.2017
#FF-HB-Asch-Schu-6
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a91e;
#FF-HB-Asch-Schu-6-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d164;

#SY01-GH-Tinnum
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:c594;

#Faehrhaus_Hodorf1
allow 2a03:2267:4e6f:7264:16cc:20ff:fec0:177e;

#NDS-FFNH-0113-Handeloh-HSAR
allow	2a03:2267:4e6f:7264:ee08:6bff:fea4:c97a;

#KelliNet-008
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe3f:6394;

#FF-PflegeDiakonie
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:6d78;

#FFNH-Origene-Praxis
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ae9e;

#Doerpstedt-FFW-Sport
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe33:3e4e;

#NDS-FFNH-241-Hanstedt-Heidekrug
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c278;

#Freifunk_Meldorf10
allow 2a03:2267:4e6f:7264:cad3:a3ff:fe5c:c166;

#Freifunk-Tafel-Meldorf
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:6b7e;

#NDS-FFNH-123-NeuWu-Courage
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:80ae;

#Seevextender
allow 2a03:2267:4e6f:7264:eade:27ff:fe21:dc4d;

#GeWo314
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:b768;

#GeWo312
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:b252;


#ferienhof_moeller_5
allow 2a03:2267:4e6f:7264:c6e9:84ff:fea1:e910;
#ferienhof_moeller_3
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:b498;

#Cuxhaven_Nordseewind_1
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:a732;

#FF-Hanstedt-17
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea2:335c;

#NDS-FFNH-160-Hanstedt-016
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:5e28;
#NDS-FFNH-500-Hanstedt-002
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:872e;

#Elbmarsch_506
allow 2a03:2267:4e6f:7264:62e3:27ff:fec6:fa04;

#FF-HB-Asch-Schu-4-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:4580;
#FF-HB-Asch-Schu-6-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fef2:ca62;


#FF-Nord-amaism34894ns
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:73f2;

#Salon-Cordes
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed2:e250;

#GeWo1827
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fb02;

#SY02-01-GH-Morsum
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:c5cc;

#hle-c7
allow 2a03:2267:4e6f:7264:c66e:1fff:feea:ee72;

#NDS-FFNH-254-Schaeferstieg31
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c1e8;

#FF-DIT-OFF-Albersd-Freizeitbad1
allow 2a03:2267:4e6f:7264:219:99ff:fe51:25ef;

#SLN
allow 2a03:2267:4e6f:7264:eade:27ff:fe59:213;

#diekkieker
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:a2b2;

#GeWo309
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bc4a;

#FF-WES-Bruhs
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe8b:7abc;

#FF-BU-EW10_1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:305c;

#SY04-Beratung-Geschwister-Scholl
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:4e9a;

#Freifunk-Horneburg
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe57:d8f4;

#GeWo321
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:b772;

#FF-DIT-AOEZA1
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e4d2;

#Route66
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:a832;

#22.05.2017
#FF-Nord-mq84bn56x830x
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4fa8;



#freifunk14cc20b0f096
allow 2a03:2267:4e6f:7264:16cc:20ff:feb0:f096;


#Cuxhaven_Nordseewind_2
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:b09e;

#FFNH-Handeloh-009
allow 2a03:2267:4e6f:7264:62e3:27ff:fed6:126c;

#FFNH-NeuWu-Marktplatz-Offloader
allow 2a03:2267:4e6f:7264:fa1a:67ff:fea5:fc24;
#NDS-FFNH-126-NeuWu-Rathaus-SW
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8138;



#FF-HB-Asch-Schu-2-1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:2a96;
#FF-HB-Asch-Schu-6-3
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:569a;

#ff-nf-gar-vermittlung-hoffmann
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:82ee;

#Fleestedt_Hand_In_Hand_0001
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe80:eb36;
#Fleestedt_Hand_In_Hand_0008
allow 2a03:2267:4e6f:7264:62e3:27ff:fe91:bade;

#NDS-FFNH-0531-Erikastrasse01
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:7ad2;

#FFNH-Undeloh-HEZ1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe48:bfa8;
#FFNH-Undeloh-HEZ2
allow 2a03:2267:4e6f:7264:6a72:51ff:fe48:bfa8;
#FFNH-VVUndeloh-002
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8d28;
#FFNH-Undeloh-HRP1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe48:c100;
#FFNH-Undeloh-VVMP01
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4e:f631;

#NDS-FFNH-167-Buchholz-Diakonie
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:67aa;
#NDS-FFNH-227-Buchholz-Cafe-Inter
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:c95e;

#FFNH-Origene-Praxis
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ae9e;

#FFNH-Juz1
allow 2a03:2267:4e6f:7264:a62b:b0ff:fec7:b1ba;
#FFNH-Juz2
allow 2a03:2267:4e6f:7264:a62b:b0ff:fec7:b0ea;

#Elbmarsch_Haus5_neu
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe7a:20be;
#Elbmarsch_504
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a2e;
#Elbmarsch_501
allow 2a03:2267:4e6f:7264:62e3:27ff:fec6:fc1c;
#Elbmarsch_505
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ac8;
#Elbmarsch_005
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a70;
#Elbmarsch_512
allow 2a03:2267:4e6f:7264:62e3:27ff:fed6:1114;
#Elbmarsch_507
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a34;
#Elbmarsch_502
allow 2a03:2267:4e6f:7264:62e3:27ff:fec6:fbe4;

#Cafe_Saimi
allow 2a03:2267:4e6f:7264:62e3:27ff:febe:208a;

#Cafe_Saimi_Terrasse
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:25d6;
#Provinzial_Nord_Meldorf
allow 2a03:2267:4e6f:7264:62e3:27ff:feb8:49e;
#MED-Rathausplatz-FF-04
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:e0ea;
#MED-Rathausplatz-FF-03
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:dfea;

#Elbmarsch_101_refugees
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:3efa;
#Elbmarsch_204_Rot_1
allow 2a03:2267:4e6f:7264:12fe:edff:fec2:3540;
#Elbmarsch_002_blau
allow 2a03:2267:4e6f:7264:eade:27ff:fef9:1dba;
#Elbmarsch_013_Rot_3
allow 2a03:2267:4e6f:7264:32b5:c2ff:fec2:6002;
#Elbmarsch_006
allow 2a03:2267:4e6f:7264:32b5:c2ff:fee2:ed9c;
#Elbmarsch_001
allow 2a03:2267:4e6f:7264:227:19ff:fecc:721c;
#Elbmarsch_003
allow 2a03:2267:4e6f:7264:16cc:20ff:fe6f:e58e;

#NDS-FFWL001-Hoopte-Clubheim1
allow 2a03:2267:4e6f:7264:8616:f9ff:fe49:6b6;
#NDS-FFWL001-HOOPTE01
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:b612;
#NDS-FFWL005-Winsen02
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:a910;

#allow FF-BU-EW10_2
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:74b0;

#FFNH-Egestorf-AL-U1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6c:64c4;


#freifunk-helmstorf-1
allow 2a03:2267:4e6f:7264:12fe:edff:fe08:f09c;

#Andre
allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:4446;



#NDS-FFNH-483-Hanstedt-027
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:86de;
#NDS-FFNH-507-Hanstedt-021
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:832c;
#FF-Hanstedt-26
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe4b:97ae;

#FFNH-Deutsches-Haus-02
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:7bb6;

#NDS-FFNH-506-Schaeferstieg27
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:83f4;
#NDS-FNH-255-Schaeferstieg27
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c726;

#NDS-FFNH-168-Winsen-001
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:609a;
#NDS-FFNH-166-Winsen-002
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73fa;
#AMD-Nordermarkt-Touri
allow 2a03:2267:4e6f:7264:eade:27ff:fef6:d4e2;
#Ristorante_Mama_Leone
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:dcfe;

#NDS-FFNH-507-Hanstedt-021
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:832c;
#NDS-FFNH-490-Hanstedt-003
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8fd4;

#FF-Hanstedt-04
allow 2a03:2267:4e6f:7264:c6e9:84ff:fef0:705a;
#NDS-FFNH-508-Hanstedt-008
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:861a;

#NDS-FFNH-217-Brackel-Tischlein01
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:7402;
#NDS-FFNH-218-Brackel-Tischlein02
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:67f0;

#FF-Egestorf-Nord-03
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:ab82;
#NDS-FFNH-0535-Egestorf-Nord-01
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe88:6db0;
#NDS-FFNH-0536-Egestorf-Nord-02
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:84fa;

#NDS-FFNH-164-Hanstedt-Buchhandlu
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:678a;
#NDS-FFNH-163-Hanstedt-Buchhandlu
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:686e;

#FFNH-Egestorf-AL-P1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe62:b5a5;

#NDS-FFNH-1764-Bispingen-RH01
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fb70;
#NDS-FFNH-0202-Bispingen-JK
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6620;
#NDS-FFNH-0144-Bispingen-AKR1
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc90;

#FFNH-Undeloh-DS02
allow a03:2267:4e6f:7264:6a72:51ff:fe62:b5ba;

#NDS-FFNH-0237-Undeloh-SH01
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cb7e;
#NDS-FFNH-0238-Undeloh-SH02
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cb96;

#FFNH-Undeloh-LC01
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e0f4;
#FFNH-Undeloh-LC02
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8dfe;

#FFNH-Handeloh-RH01
allow 2a03:2267:4e6f:7264:c66e:1fff:fefe:4184;
#NDS-FFNH-1711-Handeloh-RHD1
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:ab22;
#FFNH-Handeloh-BVV03
allow 2a03:2267:4e6f:7264:62e3:27ff:feec:e5ca;

#NDS-FFNH-0113-Handeloh-HSAR
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:c97a;
#NDS-FFNH-1393-Handeloh-HK01
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:dc84;
#NDS-FFNH-0112-Handeloh-KS09
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:c9f8;
#NDS-FFNH-0098-Handeloh-AW16
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cab6;


#NDS-FFNH-1705-1602-Hoeckel-OL2
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:aa78;
#FFNH-Handeloh-007
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8dc0;
#FFNH-Hoeckel-H16VG1
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8e78;
#NDS-FFNH-0099-Handeloh-HSRO
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e35c;
#NDS-FFNH-1708-Hoeckel-KG1
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:a8c0;

#spasibasarebu
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:7f34;
## 21.05.2017
#FF-Hanstedt-20
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bfa6;


#HBJensen_Sylt_7
allow 2a03:2267:4e6f:7264:92f6:52ff:feb5:2c8a;
#Freifunk-SYLT
allow 2a03:2267:4e6f:7264:16cc:20ff:fee1:5654;
#DIAVOLO
allow 2a03:2267:4e6f:7264:16cc:20ff:febb:d1ae;

#NDS-FFNH-0532-Erikastrasse02
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee1:79a8;

#FF-HB-Borg-Ferienhaus_Britta
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d194;
#FF-HB-Borg-Ferienhaus_Tina
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:64b2;

#FF-HB-Borg-Trae-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1a6;
#FF-HB-Borg-Trae-1a
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d102;

#NDS-FFNH-242-Hanstedt-009
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c75a;

#NDS-FFNH-243-Hanstedt-010
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c152;
#FF-Hanstedt-22
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea2:3384;

#drestedt_asyl_jupiter
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:b16e;
#drestedt_asyl_saturn
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:9be2;
#mushreq
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe9f:1bce;
#drestedt_asyl_mars
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:9a9a;

#NDS-FFNH-131-NeuWu-Turek
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8010;
#NDS-FFNH-1401-NeuWu-Uhu
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:cac8;
#NDS-FFNH-139-NeuWu-Plover
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc8a;
#NDS-FFNH-212-NeuWu-Sterna
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:679a;
##weiter mit https://mesh.nord.freifunk.net/#!v:m;n:0019999ae5b3

#NDS-FFNH-213-Moisburg-Rathaus
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6884;
#NDS-FFNH-211-Moisburg-Rathaus
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:678c;

#Bodderbrod_mit_Kees
allow 2a03:2267:4e6f:7264:c6e9:84ff:fefd:db0;
#Leuchte_des_Nordens
allow 2a03:2267:4e6f:7264:6a72:51ff:fe34:afc2;

#Wiep4
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:4018;
#Wiep5
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:262c;
#Wiep2
allow 2a03:2267:4e6f:7264:62e3:27ff:fec6:fc5a;
#hechthausen-001
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe6e:c514;
#hechthausen-002
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe92:4de8;

#Privilegierte_Stadtapotheke
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:ff4c;
#Ristorante_Mama_Leone
#allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:dcfe;
#freifunk_Meld_Suederm
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:5406;
#Provinzial_Nord_Meldorf
#allow 2a03:2267:4e6f:7264:62e3:27ff:feb8:49e;
#MED-Rathausplatz-FF-03
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:dfea;
#MED-Rathausplatz-FF-02
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d1d4;
#MED-Rathausplatz-FF-04
#allow a03:2267:4e6f:7264:f6f2:6dff:fe85:e0ea;

#HV-Stelle-1043-v2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3b62;
#HV-Stelle-1043
allow 2a03:2267:4e6f:7264:8616:f9ff:fe66:5f92;
#HV-Stelle
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:8caa;

#ff-nf-spo-fewo-carstens-1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73c8;

#InuYasha
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe69:4277;
#Kikyo
allow 2a03:2267:4e6f:7264:6a72:51ff:fe06:5e95;
#Kaede
allow 2a03:2267:4e6f:7264:c66e:1fff:feb9:13f8;
#Kiara
allow 2a03:2267:4e6f:7264:eade:27ff:fe65:cae4;
#Kagome
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe6e:9cba;

#FF-NMS-CB-01
allow 2a03:2267:4e6f:7264:12fe:edff:fe92:eb24;
#FF-NMS-CB-04
allow 2a03:2267:4e6f:7264:ee08:6bff:fe56:d9a0;
#FF-NMS-NSM-01
allow 2a03:2267:4e6f:7264:6a72:51ff:fe16:6376;
#FF-NMS-AB
allow 2a03:2267:4e6f:7264:12fe:edff:fe7a:38fa;

#tangogolf
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:d83e;
#tangogolf_2
allow 2a03:2267:4e6f:7264:a62b:b0ff:fec7:cdb8;

#FF-Stolpe-DGH-0
allow 2a03:2267:4e6f:7264:219:99ff:fe68:dece;
#FF-Stolpe-DGH-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe22:840c;
#FF-Stolpe-DGH-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe22:8092;

#eae-segeberg-uplink
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:39ca;
#eae-badsegeberg-01
allow 2a03:2267:4e6f:7264:46d9:e7ff:fe74:5a5b;
#eae-segeberg-ap2
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a998;

#KelliNet-009
allow 2a03:2267:4e6f:7264:32b5:c2ff:fed5:c96c;
#KelliNet-005
allow 2a03:2267:4e6f:7264:16cc:20ff:fe2b:7568;

#KelliNet-002
allow 2a03:2267:4e6f:7264:c66e:1fff:feb3:967e;
#KelliNet-001
allow 2a03:2267:4e6f:7264:6a72:51ff:fe50:a341;
#KelliNet-006
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe3d:f346;
#KelliNet-008
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe3f:6394;

#MBD_Freifunk_Ehndorf
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:8aa4;
#MBD_Freifunk_Ehndorf2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:a3de;

#Ferienhof-Struve-11-c
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e9c6;
#Ferienhof-Struve-14
allow 2a03:2267:4e6f:7264:92f6:52ff:fec4:b47f;
#Ferienhof-Struve-11-d
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:21aa;
#Ferienhof-Struve-13a
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe3e:8dd6;
#Ferienhof-Struve-11-d
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:21aa;
#Ferienhof-Struve-11b
allow 2a03:2267:4e6f:7264:92f6:52ff:fec4:a9f6;
#Ferienhof-Struve-16
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe3e:8bd4;
#Ferienhof-Struve-15
allow 2a03:2267:4e6f:7264:92f6:52ff:fec4:b2c6;
#Ferienhof-Struve-11a
allow 2a03:2267:4e6f:7264:6666:b3ff:fe60:6f23;

#FF-DIT-Albersdorf-Freizeitbad2
allow 2a03:2267:4e6f:7264:6a72:51ff:fe44:da41;
#FF-DIT-Albersdorf-Freizeitbad1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4e:28ba;

#FNF_Corcuscant
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:7146;
#FNF_Alderaan
allow 2a03:2267:4e6f:7264:46d9:e7ff:feb7:e715;

#schiri012
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:dd34;
#freifunk_hagen_sk_09
#allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:6536;
#freifunk_hagen_sk_06
#allow 2a03:2267:4e6f:7264:62e3:27ff:febd:be20;
#freifunk_hagen_sk_04
#allow 2a03:2267:4e6f:7264:16cc:20ff:feb1:f38;

#FF-Elpersb-HausDithmarschen
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:acee;
#FF-Elpersb-HausDithmarschen-02
allow 2a03:2267:4e6f:7264:c66e:1fff:fea7:52f2;
#FF-Elpersb-HausDithmarschen-01
allow 2a03:2267:4e6f:7264:32b5:c2ff:fef0:987a;

#ff-nf-spo-fewo-carstens-1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73c8;
#ff-nf-spo-fewo-carstens-2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bb9c;
#FF-RD_BI-26B
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1d0;

#FF-HB-Asch-Huett-10
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1cc; 

#Meldorfer_Buecherstube
#weiter mit https://mesh.nord.freifunk.net/#!v:m;n:c8d3a35cc166
allow 2a03:2267:4e6f:7264:f6f2:6dff:fef6:66fe;

#Schuhhaus_Carstensen
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:db02;

#Eiscafe-Boethern
allow 2a03:2267:4e6f:7264:c66e:1fff:fed6:6f7c;
#VHS_Meldorf_1
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe3e:8e34;
#FF_Meldorf_DerKulturonkel
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:aa9c;

#FF-DIT-Stadtbuecherei2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d3d8;
#FF-DIT-Stadtbuecherei1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d542;
#VHS_Meldorf_2
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:8fc0;

#FF-BU-St_Annen-Kirche
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:49c4;

#FF-HB-029
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:6f8c;

#FF-HB-Asch-ses-iwi
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4110;
#FF-HB-Asch-ses-iwi-0
allow 2a03:2267:4e6f:7264:6666:b3ff:fe2e:fa5;
#FF-HB-Asch-Schu-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:dfec;

#MTV-von-1860-ev-Heide
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:a5c4;
#MTV-von-1860-ev-Heide_Gaststube
allow 2a03:2267:4e6f:7264:ee08:6bff:feec:273e;

#FF-BU-Woe-Ring11a
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e470;

#FF-WesthofWoehrden-2
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e296;

#FF-WesthofWoehrden-1
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e498;

#FF-BU-Woe-Chstr2
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e494;
#FF-BU-Woe-Chstr8
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9fc6;

#FF-BU-Woe-CNS6
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef32;

#NDS-Hue-Rathaus-0292
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc00;
#NDS-Hue-Rathaus-0291
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e11c;

#Am-Exer-01
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe6e:6922;
#SeasideAppartment02
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe6e:3ed2;

#Eckernfoerde02
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bf12;
#Eckernfoerde01
allow 2a03:2267:4e6f:7264:c66e:1fff:fee8:eec4;

#Doerpstedt-FFW-Sport
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe33:3e4e;

#FF-BU-GB8
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:c71c;

#FF-BU-223
allow 2a03:2267:4e6f:7264:a62b:b0ff:feaa:dca;

#FFNH-Deutsches-Haus-02
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:7bb6;
#FFNH-Deutsches-Haus-01
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:79ce;

#FFNH-Origene-Praxis
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ae9e;

#Freifunk_Geestlandhalle
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:9190;

#Elbmarsch_OL3
allow 2a03:2267:4e6f:7264:219:99ff:fe5b:45cc;
#Elbmarsch_RS1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6f:68a4;
#Elbmarsch_RS2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:8cac;
#Elbmarsch_RS3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6f:6c08;

#FB-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3c1e;
#FB-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3d6e;
#FB-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:388c;

#behard
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed7:1eb2;

#FF_SL_Husumer_Str_001
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;

#FF-DIT-Kappeln1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:9710;
#FF-DIT-Kappeln2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15c4;
#FF-DIT-Kappeln3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15be;

#Etoile-V
allow 2a03:2267:4e6f:7264:46d9:e7ff:fedb:fb24;
#Napapijri
allow 2a03:2267:4e6f:7264:ea94:f6ff:fea1:d60e;
#Tobacco
allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:2276;


#RBZ10
allow 2a03:2267:4e6f:7264:6670:2ff:fed3:9fc;

#FF-BU-THW_BA35_1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4b60;

#HollenbeckFF_0x06
allow 2a03:2267:4e6f:7264:a2f3:c1ff:feac:26c8;
#HollenbeckFF_0x02
allow 2a03:2267:4e6f:7264:12fe:edff:feb7:5dcc;
#HollenbeckFF_0x04
allow 2a03:2267:4e6f:7264:6666:b3ff:feaf:dfca;

#FF-BU-223
allow 2a03:2267:4e6f:7264:a62b:b0ff:feaa:dca;

#FF-BU-LS30
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae0a;

#Christian_Sylt-Futro
allow 2a03:2267:4e6f:7264:219:99ff:fe7a:704e;

# Norderstedt-Harksheide01 
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e75c;

  #18052017
  #FF_Windbergen_Wodansberg
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe54:c26e;
  #FF_Windbergen_Wodansberg_2
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe54:c28e;
  #FF_Windbergen_Wodansberg_3
  allow 2a03:2267:4e6f:7264:c66e:1fff:fe2d:a208;

  #Pension-Galerie_Vierhoefen_03
  allow 2a03:2267:4e6f:7264:c66e:1fff:fe63:4724;

  #DieMuehle
  allow 2a03:2267:4e6f:7264:a62b:b0ff:fed1:e178;

#NDS-FFNH-124-NeuWu-Courage
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8048;
#NDS-FFNH-123-NeuWu-Courage
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:80ae;
#NDS-FFNH-125-NeuWu-Rathaus-NW
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:814e;
  #NDS-FFNH-120-NeuWu-Rathaus-SO
  allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8178;

#Faehrhaus_Hodorf1
allow 2a03:2267:4e6f:7264:16cc:20ff:fec0:177e;
  #Faehrhaus_Hodorf2
  allow 2a03:2267:4e6f:7264:6a72:51ff:fe36:c0c;

  #FF-Hanstedt-12
  allow 2a03:2267:4e6f:7264:62e3:27ff:feed:86fc;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #NDS-FFNH-0509-Bispingen-CPE1
  allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8514;

#ferienhof_moeller
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:77c;
#ferienhof_moeller_2
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:e3ee;

  #ferienhof_moeller_4
  allow 2a03:2267:4e6f:7264:eade:27ff:fe37:e090;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #freifunk-nazar-kebabhaus
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;

  #FF-HB-161
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:43a8;

#Norbi
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d6ba;
  #FressflashRuftDasKaenguru
  allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d73e;

  ####### Unterhalb keine Änderungen durchführen! ###############

deny all;
 }

  root /opt/www;
  autoindex on;
  index index.html index.htm;
  server_name localhost;
}
