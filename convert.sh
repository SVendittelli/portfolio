#!/bin/sh
pandoc ./src/about-me.md -f markdown-smart -t html -o ./templates/files/about-me.hbs
