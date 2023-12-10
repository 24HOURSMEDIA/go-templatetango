# Filters

You can list all available filters with the command 'stick:list-filters':

```
%command% stick:list-filters
```

## Using filters in templates

Use filters in templates just like in twig:

```
Hello {{ name | upper }}.
```

## %product% filters

Specially crafted filters:

| Filter name         | Description                                                         |
|---------------------|---------------------------------------------------------------------|
| `json_value`        | Encodes a value as a json value                                     |
| `json_casted_value` | Encodes a value as a json value but first tries to cast to int, etc |
| `json_escape`       | Escapes a json string                                               |
| `json_decode`       | Decode a json string for further processing                         |
|                     |                                                                     |
| `rawurlencode`      | Escapes spaces (and other special chars) for url with a % symbol    |
| `boolify`           | Convert 'on', 'off', 'true', '1', 1 etc to boolean values           |
| `bool_switch`       | Boolifies the value and return the first or the second argument     |

### `json_value`

Example (VAR can be a string, number, boolean, null).
Note that environment variables are always strings, so you need to cast them to the correct type.

```twig
{
  "foo": {{ VAR | json_value }}
}
```

### `json_casted_value`

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

### `json_escape`

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

### `json_decode`

```twig
{% set json = '{"foo": "bar", "foobar": 2}' %}
{% for key, value in json | json_decode %}
  {{ key }}: {{ value }}
{% endfor %
```

### `rawurlencode`

Raw url encode is also called path encoding. 
It encodes spaces and other special characters with a percent symbol.

```twig
{{ 'To be or not to be' | rawurlencode }}
```

Result:

```text
To%20be%20or%20not%20to%20be
```

### `boolify`

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

### `bool_switch`

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

## Default filters

Some of the default filters provided with the template engine:

| Filter name        | Description |
|--------------------|-------------|
| `abs`              |             |
| `batch`            |             |
| `capitalize`       |             |
| `convert_encoding` |             |
| `date`             |             |
| `date_modify`      |             |
| `default`          |             |
| `first`            |             |
| `format`           |             |
| `join`             |             |
| `json_encode`      |             |
| `keys`             |             |
| `last`             |             |
| `length`           |             |
| `lower`            |             |
| `merge`            |             |
| `nl2br`            |             |
| `number_format`    |             |
| `raw`              |             |
| `replace`          |             |
| `reverse`          |             |
| `round`            |             |
| `slice`            |             |
| `sort`             |             |
| `split`            |             |
| `striptags`        |             |
| `title`            |             |
| `trim`             |             |
| `upper`            |             |
| `url_encode`       |             |