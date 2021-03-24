from htutil import file
from pathlib import Path

project_path = Path(__file__).parent.parent.parent

def path_to_abs_str(path:Path) -> str:
    return str(path.absolute())

project_path_str = path_to_abs_str(project_path)
project_name = project_path.name

def main():
    print(project_path)
    print(project_path_str)
    print(project_name)


if __name__ == '__main__':
    main()