Takes tatoeba.org files and creates a filtered laguange pair file, for ex:

    wget http://downloads.tatoeba.org/exports/sentences.tar.bz2
    wget http://downloads.tatoeba.org/exports/links.tar.bz2

    # On Mac: wget https://github.com/slava-vishnyakov/tatoeba-pair/releases/download/0.1/tatoeba-pair-darwin-amd64
    # on rest - compile from source

    ./tatoeba-pair eng cmn > eng-cmn.tsv

The output is tab-separated:

    head -3 eng-cmn.tsv

    Let's try something.	我們試試看！
    Let's do it.	我們試試看！
    Let's try it.	我們試試看！
