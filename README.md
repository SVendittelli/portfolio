# portfolio

Live reloading can be enabled by using [cosmtrek/air](https://github.com/cosmtrek/air)

Use pandoc to automatucally convert markdown files to html:

```sh
docker run --rm --volume "$(pwd):/data" --entrypoint "/data/convert.sh" pandoc/latex:2.6
```
