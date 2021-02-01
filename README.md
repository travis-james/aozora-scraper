# aozora-scraper
This is a small concurrent webscraper that downloads all works of an author from https://www.aozora.gr.jp/

## !! WARNING !!
This program was created so that I could learn a bit about go routines. If you use this program, do NOT use it 
maliciously or irresponsibly. Excessive requests could put unneccessary burden on the receiving website, and 
possibly get you in trouble. Be kind and responsible.

## How to use this program
It has the following command line flags:
- ap: author page, such as https://www.aozora.gr.jp/index_pages/person11.html The program accepts any link
provided to it to be of that format, where each author's url is distinguished by the person number ("/personxx.html")
- dn: directory name. This is the location of where the program will save all the files. Example, "works" would save
the files to a folder named works. Note, the directory name must not exist already, it must be a new directory.

An example of how one would run it:
```
go run ./cmd -ap=https://www.aozora.gr.jp/index_pages/person35.html -dn=osamu
```

### Assumptions
- As already stated, the provided link must be of the form https://www.aozora.gr.jp/index_pages/personxx.html
- This program only downloads works that are zip files (which have a .txt file inside of the zip), if no zip file is
found, the program will return an error.
- I've tested this on about 5 authors' web pages. It is possible there are other authors' pages that have
HTML that don't conform to how this program tokenizes web pages.
