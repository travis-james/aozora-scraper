# aozora-scraper
This is a small CLI concurrent webscraper that downloads all works of an author from https://www.aozora.gr.jp/

## !! WARNING !!
This program is purely for educational purposes.
If you use this program, do NOT use it maliciously or irresponsibly. Excessive requests could put unneccessary burden on the receiving website, and possibly get you in trouble. Be kind and responsible.

## How to use this program
It has the following command line flags:
- ap: author page, such as https://www.aozora.gr.jp/index_pages/person11.html Each author's url is distinguished by the person number ("/personxx.html")
- dn: directory name. This is the location of where the program will save all the files. Example, "osamu" would save
the files to a folder named osamu. Note, the directory name must not exist already, it must be a new directory.

An example of how one would run it:
```
go run ./cmd -ap=https://www.aozora.gr.jp/index_pages/person35.html -dn=osamu
```

## Why did I make this?
I wanted to learn more about web scraping, Go routines, and concurrency.

### Assumptions
- As already stated, the provided link must be of the form https://www.aozora.gr.jp/index_pages/personxx.html
- All links to works on an author page start with "../cards" and therefore all links of a given
work are:
https://www.aozora.gr.jp + "/cards"
- This program only downloads works that are zip files (which have a .txt file inside of the work), if no zip file is
found, the program will return an error. Some works only have a pdf or html link rather than a zip, in that case this
program returns an error.
- I've tested this on about 5 authors' web pages. It is possible there are other authors' pages that have
HTML that don't conform to how this program tokenizes web pages.
