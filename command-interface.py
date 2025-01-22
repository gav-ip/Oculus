import typer
from web3 import Web3
import requests

#connect to an ethereum node using infura API
infura_url = "https://mainnet.infura.io/v3/834f5b297d404c70a7f920c99450dc29"
web3 = Web3(Web3.HTTPProvider(infura_url))

#checking if connection is successful
if web3.is_connected():
    print("connected to Ethereum Mainnet")
else:
    print("Failed to connect")
