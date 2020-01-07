## 1.1docker-compose 安装过程中2.0~3.0存在的问题

![截图](./images/2019-12-20-1.png)

#### 1.1 可能问题 

```shell
Installing collected packages: texttable, subprocess32, urllib3, certifi, chardet, idna, requests, websocket-client, pycparser, cffi, bcrypt, pynacl, cryptography, paramiko, docker, backports.shutil-get-terminal-size, docker-compose
  Found existing installation: texttable 0.8.1
ERROR: Cannot uninstall 'texttable'. It is a distutils installed project and thus we cannot accurately determine which files belong to it which would lead to only a partial uninstall.
SENSETIME\taoshumin_vend
```

解决方案:

```shell
$ sudo pip install docker-compose --ignore -installed request
```



#### 1.2 可能存在问题

![截图](./images/2019-12-20-2.png)

```shell
, line 97, in <module>
    from pip._vendor.urllib3.contrib import pyopenssl
  File "/usr/local/lib/python2.7/dist-packages/pip/_vendor/urllib3/contrib/pyopenssl.py", line 46, in <module>
    import OpenSSL.SSL
  File "/usr/lib/python2.7/dist-packages/OpenSSL/__init__.py", line 8, in <module>
    from OpenSSL import rand, crypto, SSL
  File "/usr/lib/python2.7/dist-packages/OpenSSL/SSL.py", line 118, in <module>
    SSL_ST_INIT = _lib.SSL_ST_INIT
AttributeError: 'module' object has no attribute 'SSL_ST_INIT'
```

解决方案:

```shell
$ rm -rf /usr/lib/python2.7/dist-packages/OpenSSL
$ rm -rf /usr/lib/python2.7/dist-packages/pyOpenSSL-0.15.1.egg-info
$ sudo pip install pyopenssl
```

## 1.2 docker 制作镜像

```shell
# 制作镜像
$ docker build -t 镜像名:tag . 
# 运行镜像
$ docker run -itd --name 容器名 -p 宿主机port:容器port 镜像名
```

