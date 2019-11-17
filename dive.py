#!/usr/bin/env python3

import yaml
import sys
import os
import os.path
from os import path

GLOBAL_CONFIG_FILE = os.getenv('HOME') + '/.dive.yaml'
LOCAL_CONFIG_FILE = './.dive.yaml'
SELF_PATH = ''
VERBOSE = False

def load_config(file_path):
  with open(file_path, 'r') as file:
    config = yaml.load(file, Loader=yaml.FullLoader)
    file.close()
    return config

def connect(host_config):
  command = ''
  if 'user' in host_config:
    command += host_config['user'] + '@'
  command += host_config['host']
  if 'port' in host_config:
    command += ' -p ' + host_config['port']
  if (VERBOSE):
    print('Host: ' + host_config['host'])
    print('User: ' + host_config['user'])
    print('Port: ' + host_config['port'])
    print('Connecting...')
  return os.system('ssh ' + command)

i = 0
count = len(sys.argv)
host = None
user = None
port = None
save = False
while i < count:
  if i == 0:
    SELF_PATH = sys.argv[i]
  elif sys.argv[i] == '-p':
    i += 1
    port = sys.argv[i]
  elif sys.argv[i] == '-v':
    VERBOSE = True
  elif sys.argv[i] == '-s':
    save = True
  else:
    if ('@' in sys.argv[i]):
      connection_args = sys.argv[i].split('@')
      user = connection_args[0]
      host = connection_args[1]
    else:
      host = sys.argv[i]
  i += 1

config = {}
if path.exists(GLOBAL_CONFIG_FILE):
  config = load_config(GLOBAL_CONFIG_FILE)
if path.exists(LOCAL_CONFIG_FILE):
  config.update(load_config(LOCAL_CONFIG_FILE))

host_config = None
if host is not None and user is None:
  if (host in config['aliases']):
    real_name = config['aliases'][host]
    host_config = config['hosts'][real_name]
  elif (host in config['hosts']):
    host_config = config['hosts'][host]

if host_config is None:
  connect({
    'host': host or 'localhost',
    'user': user or os.getenv('USER'),
    'port': port or '22'
  })
else:
  connect(host_config)
