# %product% extra Filters

Specially crafted filters:

| Filter name           | Description                                                                    |
|-----------------------|--------------------------------------------------------------------------------|
| `json_value`          | Encodes a value as a json value                                                |
| `json_casted_value`   | Encodes a value as a json value but first tries to cast to int, etc            |
| `json_escape`         | Escapes a json string                                                          |
| `json_decode`         | Decode a json string for further processing                                    |
|                       |                                                                                |
| `rawurlencode`        | Escapes spaces (and other special chars) for url with a % symbol               |
| `boolify`             | Convert 'on', 'off', 'true', '1', 1 etc to boolean values                      |
| `bool_switch`         | Boolifies the value and return the first or the second argument                |
| `exists`              | Check if a variable with a name exists in the current scope                    |
| `value`               | Retrieve a value by name from the current context, or a default,or null        |
| `apply_mapping`       | Maps variables by name from the current context or an object to another object |
| `extract_objects`     | Create objects from numbered variable names in the current scope               |
| `objects_from_prefix` | Create objects from variables in the current scope sharing the same prefix     |
| `tidy`                | Clean up extra newlines in texts and attempts to fix indentation               |

## `json_value`

Example (VAR can be a string, number, boolean, null).
Note that environment variables are always strings, so you need to cast them to the correct type.

```twig
{
  "foo": {{ VAR | json_value }}
}
```

## `json_casted_value`

json_casted_value tries to cast the value to a boolean, number or null before encoding it as a json value.

```twig
[
  {{ 'True' | json_casted_value }},
  {{ 'false' | json_casted_value }},
  {{ 'null' | json_casted_value }},
  {{ '124.5' | json_casted_value }},
  {{ 'null' | json_casted_value }},
  {{ 'foo' | json_casted_value }}
]
```

Result:

```json
[
  true,
  false,
  null,
  124.5,
  null,
  "foo"
]
```

## `json_escape`

```twig
[
  "-- {{ '"To be or not to be" - Shakespeare' | json_escape }} --"
]
```

Result:

```json
[
  "-- \"To be or not to be\" - Shakespeare --"
]
```

## `json_decode`

```twig
{% set json = '{"foo": "bar", "foobar": 2}' %}
{% for key, value in json | json_decode %}
  {{ key }}: {{ value }}
{% endfor %
```

## `rawurlencode`

Raw url encode is also called path encoding.
It encodes spaces and other special characters with a percent symbol.

```twig
{{ 'To be or not to be' | rawurlencode }}
```

Result:

```text
To%20be%20or%20not%20to%20be
```

## `boolify`

Convert various formats to a boolean.

Integer and float values (or strings that convert to an int or float) that are not `0` return true.
Integer and float values (or strings that convert to an int or float) that equal `0` return false.
null returns false
The following strings return true: 'true', 'on', 'yes', 'y', 'enable', 'enabled'
The following strings return false: 'false', 'off', 'no', 'n', 'disable', 'disabled'

Examples:

```twig
{% if 'on' | boolify  %}
true
{% endif %}
{% if 'On' | boolify  %}
true
{% endif %}
{% if 'EnAbLed' | boolify  %}
true
{% endif %}
{% if 0.1 | boolify  %}
true
{% endif %}
{% if 'false' | boolify  %}
false
{% endif %}
```

## `bool_switch`

Return one of the arguments dependent upon the boolified value.

Example:

```twig
{{ 'enabled' | bool_switch('the switch is enabled', 'the switch is disabled' }}
{{ 'off' | bool_switch('the switch is enabled', 'the switch is disabled' }}
```

Results:
```
the switch is enabled
the switch is disabled
```

## `exists`

```twig
{% set foo = 'bar' %}
Foo: {% if "foo" | exists %}{{ foo }}{% elseif %}does not exist{% endif %}
Xyz: {% if "xyz" | exists %}{{ xyz }}{% elseif %}does not exist{% endif %}
```

## `value`

Return a value of a variable by name in the current scope,
or a default value if specified.

```twig
{% set foo = "foo_value" %}
{{ "foo" | value("default_foo") }}
{{ "bar" | value("default_bar") }}
```

Output:
```
foo
default_bar
```

## `apply_mapping` 

A powerful filter that maps variables by name from the current context or an object,
to another object.

Useful to create uniform objects from .env parameters like:
- HOST1=.., HOST2=..
- PORT1=.., PORT2=..

Examples:

```
{# map properties from one object to another #}
{% set source = {"foo": "foo_value", "bar": "bar_value"} %}
{% set mapping = {"attr1": "foo", "attr2": "bar", "attr3": "notexist"} %}
{% set obj = mapping|apply_mapping(source) %}
{{ obj.attr1 }}-{{ obj.attr2 }}-{{ obj.attr3 }}
{# prints: `foo_value-bar_value-` }
```

```
{# create an object from variables in the current scope, with a default if they do not exists #}
{% set default = "default_value" %}
{% set foo = "foo_value" %}
{% set bar = "bar_value" %}
{% set obj = {"attr1": "foo", "attr2": "bar", "attr3": "notexist"} | apply_mapping(nil, nil, default) %}
{{ obj.attr1 }}-{{ obj.attr2 }}-{{ obj.attr3 }}
{# prints: `foo_value-bar_value-default_value` #}
```

```
{# create an object from variables in the current scope with a suffix, useful to create arrays of objects #}
{# from environment variables like VARIABLE1=, VARIABLE2= #}

{% set HOST1 = "example.com" %}
{% set PORT1 = "443" %}
{% set HOST2 = "foobar.com" %}
{% set PORT2 = "80" %}
{% set mapping = {"host": "HOST", "port": "PORT"} %}
{% set obj1 = mapping|apply_mapping(nil, "1") %}
{% set obj2 = mapping|apply_mapping(nil, "2") %}
{{ obj1.host }}:{{ obj1.port }}
{{ obj2.host }}:{{ obj2.port }}
{# prints: #}
{# example.com:443 #}
{# foobar.com:80 #}
```

## `extract_objects`

`extract_objects` is a powerful filter that can be used to extract objects from the current scope.

Example:

```
{% set HOST = "foo" %}
{% set PORT = "80" %}
{% set HOST_0 = "bar" %}
{% set PORT_0 = "81" %}
{% set HOST_1 = "qux" %}
{% set PORT_1 = "82" %}
{# included because it is the first key in the mapping #}
{% set HOST_70 = "fooqux" %}
{# NOT included becausethe first key in the mapping is missing #}
{% set PORT_71 = "fooqux" %}

{# try to map from 0...100 #}
{% set params = {"host": "HOST", "port": "PORT"} | extract_objects(100, "HOST", "undefined") %}

{{ params | json_encode }}
{# output: [{"host":"foo","port":"80"},{"host":"bar","port":"81"},{"host":"qux","port":"82"},{"host":"fooqux","port":"undefined"}]} #}

{% for paramObj in params %}{{ paramObj.host }}:{{ paramObj.port }}
{% endfor %}
```

## `objects_from_prefix`

This filter returns a list of objects from variables with the same prefix.
The prefix is the first part of the variable name before the first underscore.
A number may be added after the prefix so multiple objects can be returned.
You can specify an optional default value.

```

{% set proxy_host = "host" %}
{% set proxy_port = "port" %}
{% set proxy1_host = "host1" %}
{% set proxy1_port = "port1" %}
{% set proxy4_host = "host4" %}
{% set proxy4_port = "port4" %}

{% set objects = "proxy" | objects_from_prefix(10) %}
{{ objects | json_encode }}

Expect:
[{"host":"host","port":"port"},{"host":"host1","port":"port1"},{"host":"host4","port":"port4"}]

.. with default values

{% set default = {"generator": "tango", "host": "default_host"} %}
{% set proxy5_port = "port5" %}
{% set objects = "proxy" | objects_from_prefix(10, default) %}
{{ objects | json_encode }}

Expect:
[{"generator":"tango","host":"host","port":"port"},{"generator":"tango","host":"host1","port":"port1"},{"generator":"tango","host":"host4","port":"port4"},{"generator":"tango","host":"default_host","port":"port5"}]
```

## `tidy`

Clean up multiple newlines and attempts to fix standardization.
Converts tabs to spaces.
Very useful in templates.
Assumes indentation of 4 spaces.

```
{% set contents %}

   foo
	   baz
	                   bar
                        qwe




qux
		test
{% endset %}
======
{{ contents|tidy }}
======
```

Output:

```
======
    foo
        baz
                bar
                    qwe

qux
            test

======
```






