# Presentation Site

Uses [mdbook](https://rust-lang.github.io/mdBook/) to build a documentation site and [marp-cli](https://github.com/marp-team/marp-cli) to build the presentation slides to go along with the site.

Feel free to rip this folder out, the book.toml in the parent directory, and the `.github/workflows/mdbook.yml` if you want to use this repo as a starting point and don't need the site.

## Local development

```bash
mdbook build
npx @marp-team/marp-cli@latest docs/workshop/slides.md --allow-local-files --html --template bespoke -o build/slides.html
open build/index.html
```

## GitHub workflow

`.github/workflows/mdbook.yml` is setup to build and deploy the site and the workshop slides together.

mdbook is installed by binary and runs a quick `mdbook build` command that relies on the book.toml

`slides.md` is turned into an html file using [marp-cli](https://github.com/marp-team/marp-cli) and are added to the `build/` folder that is deployed to GitHub pages
