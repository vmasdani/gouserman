#!./venv/bin/python3
import subprocess
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('action')

args = parser.parse_args()

print('Generating proto Go')
subprocess.run('./protoc -I=$PWD/protos --go_out=$PWD/backend $PWD/protos/*.proto', shell=True, cwd='.')
print('Generating proto TS')


if args.action == 'run':
    steps = [
        ('docker build -t gouserman .', '.'),
        ('docker compose up', '.'),
    ]

    for (s, cwd) in steps:
        subprocess.run(s, cwd=cwd, shell=True)

elif args.action == 'build':
    steps = [
        ('docker build -t gouserman .', '.')
    ]

    for (s, cwd) in steps:
        subprocess.run(s, cwd=cwd, shell=True)
