from os import EX_OSFILE
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

def gen_dto():
    path_in = path_project / 'app' / 'model' / 'internal'
    for p in path_in.glob('*.go'):
        print(p)

        s = file.read_all_text(p)
        lines = s.split('\n')

        structStart = -1
        for i, line in enumerate(lines):
            if ' struct {' in line:
                structStart = i
                break

        lines = lines[structStart:]

        lines = [line for line in lines if (not 'CreatedAt' in line and not 'UpdatedAt' in line and not 'DeletedAt' in line)]

        response = '\n'.join(lines)
        response = response.replace(' struct', 'ApiResponse struct')

        lines = [line for line in lines if 'Id' not in line]

        request = '\n'.join(lines)
        request = request.replace(' struct', 'ApiRequest struct')


        text = file.read_all_text(path_project/'app'/'model'/ p.name)
        if 'ApiRequest' in text:
            print('ApiRequest  has already exists, skipped.')
        else:
            text += request
            print('ApiRequest gen successful :)')
            
        if 'ApiResponse' in text:
            print('ApiResponse  has already exists, skipped.')
        else:
            text += response
            print('ApiResponse gen successful :)')

        file.write_all_text(path_project/'app'/'model'/ p.name,text)

def gen_router():
    template = file.read_all_text(path_project/'script'/'crud_gen'/'template'/'route.txt')
    dir_cfg = path_project/'script'/'crud_gen'/'cfg'

    route = file.read_all_text(path_project/'router'/'router.go')
    for file_cfg in dir_cfg.glob('*.json'):
        cfg = json.loads(file.read_all_text(file_cfg))

        name = cfg['小写'] 

        if name in route:
            print(f'{name} has already exists, skipped.')
        else:
            content = fill_template(template,cfg)
            route = route.replace('// crud_gen will insert here',content)
            print(f'{name} gen successful :)')
    file.write_all_text(path_project/'router'/'router.go',route)
def main():
    gen_api()
    gen_dto()
    gen_router()


if __name__ == '__main__':
    main()