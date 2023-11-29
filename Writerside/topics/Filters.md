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

| Filter name | Description                     |
|-------------|---------------------------------|
| `json_value` | Encodes a value as a json value |
|             |                                 |
|             |                                 |
|             |                                 |
|             |                                 |
|             |                                 |

### `json_value`

Example (VAR can be a string, number, boolean, null).
Note that environment variables are always strings, so you need to cast them to the correct type.

```json
{
  "foo": {{ VAR | json_value }}
}
```

## Default filters

Some of the default filters provided with the template engine:

| Filter name         | Description |
|---------------------|-------------|
| `abs`               |             |
| `batch`             |             |
| `capitalize`        |             |
| `convert_encoding`  |             |
| `date`              |             |
| `date_modify`       |             |
| `default`           |             |
| `first`             |             |
| `format`            |             |
| `join`              |             |
| `json_encode`       |             |
| `keys`              |             |
| `last`              |             |
| `length`            |             |
| `lower`             |             |
| `merge`             |             |
| `nl2br`             |             |
| `number_format`     |             |
| `raw`               |             |
| `replace`           |             |
| `reverse`           |             |
| `round`             |             |
| `say_hello`         |             |
| `slice`             |             |
| `sort`              |             |
| `split`             |             |
| `striptags`         |             |
| `title`             |             |
| `trim`              |             |
| `upper`             |             |
| `url_encode`        |             |