import requests
import os
import json


# def get_all_txs(last_block=None):
#     # This query gets all txs in which 0x91e6..38 was mentioned and a contract was created (0xc3c3..c3)
#     url = "https://app.hexagate.com/api/v1/ethereum/mainnet/concord/transactions/?addresses=0x91e677b07f7af907ec9a428aafa9fc14a0d3a338&addresses=0xc3c3c3c3c3c3c3c3c3c3c3c3c3c3c3c3c3c3c3c3&reverse=true"

#     headers = {"X-Hexagate-Api-Key": "5EtYscGfQxK8TUjgYqGfvg=="}
#     if last_block is not None:
#         url += f"&to_block={last_block}"
#     r = requests.get(url, headers=headers)
#     return r.json()


# pod_creation_topic = (
#     "0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a"
# )
# pods = set()

# txs = get_all_txs()
# last_block = txs[0]["block_number"]
# while txs:
#     for tx in txs:
#         last_block = min(last_block, tx["block_number"])
#         for log in tx["logs"]:
#             if log["topics"][0] == pod_creation_topic:
#                 pods.add(log["topics"][1])
#     # print("Last block", last_block)
#     txs = get_all_txs(last_block=last_block - 1)

# print(pods)


# def get_all_monitors():
#     url = "https://api.hexagate.com/api/v1/monitoring/user_monitors"
#     headers = {"X-Hexagate-Api-Key": "5EtYscGfQxK8TUjgYqGfvg=="}

#     r = requests.get(url, headers=headers)
#     return r.json()

# if __name__ == '__main__':
#     print(get_all_monitors())  


url = "https://api.hexagate.com/api/v1/monitoring/user_monitors"

headers = {
  'Content-Type': 'application/json',
  'X-Hexagate-Api-Key': '5EtYscGfQxK8TUjgYqGfvg=='
}

response = requests.request("GET", url, headers=headers)
print(response.text)