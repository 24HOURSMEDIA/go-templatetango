# Example with json in an env variable

You can pass json in an environment variable, and use it
to create objects with defaults, or iterate over multiple items.

Example:


```twig
{# get the hosts as json from an ENV variable if you wish, and apply the json_decode filter. #}
{# for this example we provide it as an hardcoded string #}
{% set hosts = '[{"host": "localhost"}, {"host": "example.com", "scheme": "https", "port": 443}]' | json_decode %}

{# elements of hosts are merged in a defaultHost with default value if they are not specified #}
{% set defaultHost = {"host": null, "scheme": "http", "port": 80} %}

{% for host in hosts %}
    {% set host = defaultHost | merge(host)  %}
    <server>
        listen {{ host.host }}:{{ host.port }}
        {% if host.scheme == "https" %}ssl on{% else %}ssl off{% endif %}
    </server>
{% endfor %}
```

Output:

```text
    <server>
        listen localhost:80
        ssl off
    </server>
    
    <server>
        listen example.com:443
        ssl on
    </server>
```