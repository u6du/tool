#!/usr/bin/env xonsh

from os import listdir
from os.path import dirname,abspath,exists,join
from colorama import Fore,init
init()

ROOT = dirname(dirname(abspath(__file__)))


def main():
    for i in listdir(ROOT):
        dirpath = join(ROOT,i)
        if exists(join(dirpath,"go.mod")):
            print(Fore.GREEN+">>", dirpath+Fore.RESET)
            cd @(dirpath)
            go get -u
            go mod tidy
            git commit -m "go get -u"
            git pull
            git push


main()
