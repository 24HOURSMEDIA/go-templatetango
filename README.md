# Template Tango

[Full documentation](https://24hoursmedia.github.io/go-templatetango)

A command line template parser, intended for parsing configuration files.

* Environment variables can be used in the template, as well as conditions.
* The parser is a standalone app with no dependencies.
* Templates use a templating language based on a simplified Twig template language.

## Example

```
tango parse:file nginx_server.conf.twig >> nginx_server.conf
```



