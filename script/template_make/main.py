from htutil import file
import json
from pathlib import Path
import toml
import os

def template_make(raw: str, cfg: dict) -> str:
    for d in cfg:
        raw = raw.replace(cfg[d], '${'+d+'}')
    return raw


def main():
    for dir_in in os.listdir('in'): 
        path_in = Path('in') / dir_in

        raw = file.read_all_text(path_in / 'raw.txt')
        cfg = json.loads(file.read_all_text(path_in / 'cfg.json'))

        template = template_make(raw, cfg)

        path_out = Path('out')
        file.write_all_text(path_out / dir_in /'template.txt', template)


if __name__ == '__main__':
    main()
