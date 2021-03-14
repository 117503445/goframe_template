#!/usr/bin/env python3
# coding: utf-8

import subprocess
from pathlib import Path


def run_shell_script(script: str):
    lines = script.split('\n')
    for line in lines:
        subprocess.call(args=line, shell=True)


def text_to_lf():
    for path in Path('.').glob('**/*'):
        # print(path)
        if str(path)[0] != '.' and not path.is_dir():
            try:
                with open(path, 'r', encoding='utf-8')as f:
                    lines = f.readlines()
                    text = ''.join(lines)
                with open(path, 'w', newline='\n', encoding='utf-8')as f:
                    f.write(text)
            except UnicodeDecodeError:
                # bin or gbk ..
                print('skip!!!')


def main():
    run_shell_script('''gf gen dao
gf swagger --pack -y
gofmt -l -s -w ./''')
    text_to_lf()


if __name__ == '__main__':
    main()
