#!/usr/bin/env python3
# coding: utf-8

import subprocess
from pathlib import Path
import pysftp
import yaml
from htutil import file
import paramiko
project_path = str(Path('.').absolute())
project_name = Path('.').absolute().name

image_name = f'{project_name}_dev'
container_name = f'{project_name}_dev'

config_path = Path(__file__).parent / 'config.yaml'
config = yaml.load(file.read_all_text(config_path), Loader=yaml.FullLoader)


def run_shell_script(script: str):
    lines = script.split('\n')
    for line in lines:
        print(line)
        subprocess.call(args=line, shell=True)


def build_image():
    #run_shell_script('docker pull golang:1.16.0-alpine3.13\ndocker pull alpine:3.13')

    # run_shell_script(f'docker rm -f {container_name}')

    run_shell_script(f'docker build -t {image_name} -f Dockerfile_dev .')
    # run_shell_script(f'docker run --name {container_name} --restart=always -d --network host -v {project_path}/config:/root/config -v {project_path}/tmp/log:/root/log {image_name}')
    run_shell_script(
        f'docker save --output ./tmp/{image_name}.tar {image_name}')


def upload_image():
    with pysftp.Connection(config['server']['host'], username='root', private_key='~/.ssh/id_rsa') as sftp:
        with sftp.cd('/root'):
            sftp.put(f'./tmp/{image_name}.tar')


def run_server_container():
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(
        paramiko.AutoAddPolicy())  # 把要连接的机器添加到known_hosts文件中

    ssh.connect(hostname=config['server']['host'], username='root')
    ssh.exec_command('cd /root')
    ssh.exec_command(f'docker rm -f {container_name}')
    ssh.exec_command(f'docker rmi -f {image_name}')
    ssh.exec_command(f'docker load --input {image_name}.tar')
    # ssh.exec_command(f'docker run --input {image_name}.tar')
    # todo RUN!!!


def main():
    # build_image()
    # upload_image()
    run_server_container()


if __name__ == '__main__':
    main()
