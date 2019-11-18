#!/usr/bin/env python3

import yaml
import sys
import os
import os.path
from os import path
from datetime import date
from glob import glob 
import requests

DIVE_FOLDER = os.getenv('HOME') + '/.local/share/dive'
GLOBAL_CONFIG_FILE = DIVE_FOLDER + '/global.yaml'
REMOTES_FOLDER = DIVE_FOLDER + '/remotes'
REMOTES_CONFIG = DIVE_FOLDER + '/remotes.yaml'
LOCAL_CONFIG_FILE = '.dive.yaml'
SELF_PATH = ''
VERBOSE = False

def update_configs():
  remotes_config_file = open(REMOTES_CONFIG, 'r+')
  remotes_config = yaml.load(remotes_config_file, Loader=yaml.FullLoader)
  for config_url in remotes_config:
    get_config(config_url)

def load_config(file_path):
  with open(file_path, 'r') as file:
    config = yaml.load(file, Loader=yaml.FullLoader)
    file.close()
    return config

def get_config(url):
  f = requests.get(url)
  config_text = f.text
  config = yaml.load(config_text, Loader=yaml.FullLoader)
  config_name = config['namespace']
  f = open(REMOTES_FOLDER + '/' + config_name + '.yaml', 'w')
  del config['namespace']
  yaml.dump(config, f)

def add_config(url):
  get_config(url)
  if path.exists(REMOTES_CONFIG):
    remotes_config_file = open(REMOTES_CONFIG, 'r+')
    remotes_config = yaml.load(remotes_config_file, Loader=yaml.FullLoader)
    if url not in remotes_config:
      remotes_config.append(url)
  else:
    remotes_config_file = open(REMOTES_CONFIG, 'x')
    remotes_config = []
    remotes_config.append(url)
    yaml.dump(remotes_config, remotes_config_file)
  remotes_config_file.close()
  
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
  elif sys.argv[i] == 'get':
    i += 1
    url = sys.argv[i]
    add_config(url)
    quit()
  elif sys.argv[i] == 'update':
    i += 1
    update_configs()
    quit()
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

files = glob(REMOTES_FOLDER + '/*')
for remote_config in files:
  config.update(load_config(remote_config))

host_config = {}
if host is not None and user is None:
  if 'aliases' in config and host in config['aliases']:
    real_name = config['aliases'][host]
    host_config = config['hosts'][real_name]
  elif host in config['hosts']:
    host_config = config['hosts'][host]

if 'user' not in host_config:
  if 'user' in config:
    host_config['user'] = config['user']
  else:
    host_config['port'] = os.getenv('USER')
if 'host' not in host_config:
  host_config['host'] = '127.0.0.1'
if 'port' not in host_config:
  if port is not None:
    host_config['port'] = port
  elif 'port' in config:
    host_config['port'] = config['port']
  else:
    host_config['port'] = '22'

connect(host_config)
