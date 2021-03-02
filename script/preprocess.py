import subprocess

def main():
    run_shell_script('''gf gen dao
gf swagger --pack -y
gofmt -l -s -w ./''')


def run_shell_script(script:str):
    lines = script.split('\n')
    for line in lines:
        subprocess.call(args=line, shell=True)


if __name__ == '__main__':
    main()
