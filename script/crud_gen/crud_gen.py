from htutil import file
from pathlib import Path

import json

path_project = Path(__file__).parent.parent.parent
project_name = path_project.name

def path_to_abs_str(path:Path) -> str:
    return str(path.absolute())

def fill_template(template: str, value: dict) -> str:
    for d in value:
        template = template.replace('${'+d+'}', value[d])
    return template


def gen_api():
    template = file.read_all_text(path_to_abs_str(path_project/'script'/'crud_gen'/'template'/'api.txt'))

    dir_cfg = path_project/'script'/'crud_gen'/'cfg'
    for file_cfg in dir_cfg.glob('*.json'):
        cfg = json.loads(file.read_all_text(file_cfg))
        file_code = path_project / 'app' / 'api' / (cfg['小写'] + '.go')
        if file_code.exists():
            print(f'{file_code} has already exists, skipped.')
        else:
            content = fill_template(template,cfg) 
            file.write_all_text(file_code,content)
            print(f'{file_code} gen successful :)')


def main():
    # print(path_project)
    gen_api()
    


if __name__ == '__main__':
    main()