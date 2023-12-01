# Include, embed

You can include and embed other files.
The file locations are relative to the working directory.

`include` just includes a file and by default passes all variables to 
the include (you can override that behavior, see the example).

`embed` basically does the same, but allows to override blocks
in the embedded file with alternative content.

main.twig:
```twig
{% set name = 'Foo' %}
{% include 'some_include.twig' with {name: name} only %}

{% embed 'some_embed.twig' %}
   {% block message %}Hello {{ name }}{% endblock %}
{% endembed %}
```

some_include.twig:
```twig

Hello {{ name }} from include!

```

some_embed.twig:

```twig
Some text line.
{% block message %}
Default message
{% endblock %}
Another text line
```