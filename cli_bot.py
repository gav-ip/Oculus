import os
import typer
import requests
import threading
import time
from web3 import Web3
from dotenv import load_dotenv 

#Loading in API Keys
load_dotenv()

#connect to an ethereum node using infura API
url = f"https://mainnet.infura.io/v3/{os.getenv('INFURA_API_KEY')}"
web3 = Web3(Web3.HTTPProvider(url))

def clear_terminal():
    if os.name == 'nt':
        os.system('cls')

    # For macOS and Linux
    else:
        os.system('clear')

def get_coin_price(coin):

    coin = coin.upper()

    # CoinAPI URL for Ethereum data
    url = f"https://rest.coinapi.io/v1/exchangerate/{coin}/USD"
    
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
        coin_price = data["rate"]
        return coin_price
    else:
        return None
def live_server(coin):
    print(f"Fetching live price for {coin.upper()}. Press 'Ctrl + C' to quit.")
    try:
        while True:
            coin_price = get_coin_price(coin)
            if coin_price:
                print(f"Current {coin.upper()} price: ${coin_price:.2f}")
            else:
                print("Unable to fetch", coin ,"price.")
            time.sleep(5)
    except KeyboardInterrupt:
        print("\nServer stopped by user.")

# Main function
if __name__ == "__main__":
    #checking if connection is successful
    if web3.is_connected():
        print("connected to Ethereum Mainnet")
        
    else:
        print("Failed to connect")
    
    while True:
        coin = input("Please choose coin price to track: BTC, ETH, SOL\n")
        if coin == "BTC" or coin == "ETH" or coin == "SOL":
            break
        else:
            coin = input("Please choose coin price to track: BTC, ETH, SOL\n")

    live_server(coin)