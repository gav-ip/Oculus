import os
import typer
import requests
from web3 import Web3
from dotenv import load_dotenv 

#Loading in API Keys
load_dotenv()

#connect to an ethereum node using infura API
url = f"https://mainnet.infura.io/v3/{os.getenv('INFURA_API_KEY')}"
web3 = Web3(Web3.HTTPProvider(url))

def get_eth_price():
    # CoinAPI URL for Ethereum data
    url = "https://rest.coinapi.io/v1/exchangerate/ETH/USD"
    
    # Headers include the CoinAPI key
    headers = {
        "X-CoinAPI-Key": os.getenv("COIN_API_KEY")  # Load API key from .env file
    }
    
    # Make the API request
    response = requests.get(url, headers=headers)
    
    # Check if the request was successful
    if response.status_code == 200:
        data = response.json()
        # Extract the Ethereum price in USD
        eth_price = data["rate"]
        return eth_price
    else:
        print(f"Failed to fetch data. HTTP Status Code: {response.status_code}")
        print(response.text)  # Print error details
        return None
    
def main():
    #checking if connection is successful
    if web3.is_connected():
        print("connected to Ethereum Mainnet")
    else:
        print("Failed to connect")

    #testing if eth price can be printed 
    eth_price = get_eth_price()
    if eth_price:
        print(f"ETH price in USD: ${eth_price:.2f}")
    else:
        print("Failed to retrieve ETH price")
main()