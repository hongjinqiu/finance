#!/bin/bash
#encoding=utf8

import os,re

def deal():
    os.chdir('.')
    dirLi = os.walk('.')
    result = []
    for item in dirLi:
        for subItem in item[2]:
            filePath = os.path.join(item[0], subItem)
            fIn = open(filePath, 'r')
            content = fIn.read()
            fIn.close()
            
            li = re.findall(r'(?msi)\.use.*?\{', content)
            for useItem in li:
                result.append(useItem)

            li = re.findall(r'(?msi)requires(.*?)\]', content)
            for useItem in li:
                result.append(useItem)
    #print '\n'.join(result)
    result2 = []
    for item in result:
        li = re.findall(r'"(.*?)"', item)
        for useItem in li:
            if useItem not in result2:
                result2.append(useItem)
        li = re.findall(r'\'(.*?)\'', item)
        for useItem in li:
            if useItem not in result2:
                result2.append(useItem)
    print ','.join(['"%s"' % (item) for item in result2])


if __name__ == '__main__':
    deal()
