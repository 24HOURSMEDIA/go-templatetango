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
|                     |                                                                     |
| `rawurlencode`       | Escapes spaces (and other special chars) for url with a % symbol    |
|                     |                                                                     |

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
| `say_hello`        |             |
| `slice`            |             |
| `sort`             |             |
| `split`            |             |
| `striptags`        |             |
| `title`            |             |
| `trim`             |             |
| `upper`            |             |
| `url_encode`       |             |