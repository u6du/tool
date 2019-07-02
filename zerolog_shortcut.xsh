#!/usr/bin/env xonsh
from os.path import join,abspath,dirname
from mako.template import Template

ROOT = join(dirname(dirname(abspath(__file__))),"zerolog")

def code(method, func_li):
    txt = Template("""package ${method.lower()}

import (
    "fmt"
    "net"
    "time"

    "github.com/u6du/zerolog"
    "github.com/u6du/zerolog/log"
)

func Msg(msg string) {
    log.${method}().Out(msg)
}

func Msgf(format string, v ...interface{}) {
    log.${method}().Out(fmt.Sprintf(format, v...))
}

%for func, args in li:
func ${func} *zerolog.Event {
    return log.${method}().${func.split("(")[0]}(${args})
}

%endfor
""").render(method=method, li=func_li)
    
    method = method.lower()
    cd @(ROOT)
    mkdir -p @(method)
    echo @(txt) > @(method)/@(method).go


def args(line):
    li = []
    for i in line.split("(")[1].split(")")[0].split(","):
        li.append( i.strip().split(" ")[0] )
    return ", ".join(li)


def func(txt):
    for i in txt.split("(")[1].split(")")[0].split(","):
        param = i.strip().split(" ").pop()
        if param.startswith("Log"):
            txt = txt.replace(param, "zerolog."+param)
    return txt.replace("Event","zerolog.Event")


def main():
    prefix = "func (e *Event) "
    li = []
    with open(join(ROOT,"event.go")) as f:
        for i in f:
            i = i.strip()
            if i.startswith(prefix):
                if i.find("*Event {") < 0:
                    continue
                i = i[len(prefix):-9]
                if i[0].isupper() and "()" not in i:
                    li.append((func(i), args(i)))
    
    for word in "Info Warn Debug".split():
        code(word,li)

if __name__ == "__main__":
    main()
