# %product%

<include from="library.md" element-id="app_urls"/>

%product% parses templates written in a twig-like templating language.

Environment variables and variables loaded from a .env file
are made available in the templates.

The application is available as a stand-alone binary with NO dependencies,
so you can easily include it in your docker image (by copying it from the %product% image)


## Intended usage

It's primary usage if parsing configuration templates in docker container boot scripts.
These templates can have conditional sections.

## Example

### Parsing a template file:

```
%command% parse:file nginx_server.conf.twig >> nginx_server.conf
```

### Example template:

nginx.conf:

```
server { # simple reverse-proxy
    listen       80;
    
    {% if HTTPS_ENABLED == 1 %}
    listen 443; 
    ssl_certificate     {{ SSL_CERT }};
    ssl_certificate_key {{ SSL_CERT_KEY }};
    {% endif %}

    # pass requests
    location / {
      proxy_pass      http://127.0.0.1:{{ BACKEND_PORT }};
    }
  }
```

.env
```
HTTPS_ENABLED=0
SSL_CERT=www.example.com.crt
SSL_CERT_KEY=www.example.com.key
BACKEND_PORT=80
```