import requests

cookies = {
    'FatSecret.AntiForgery': 'CfDJ8BPT8Y-3o7FDkRiB4ECfYY_Yh3GRWzocfapchO2HQwloaTlRqspYczsUTd0SEjkJCz30igyxp3DxE6qTXOF7iAtOGwVXezflZG37xgVXnSdRSHspZLqgRDShco7Tmny1Yke7ZNlpmE3iRnjF84G7fXY',
    'AWSALB': 'f4ajqeI4YEX0tlTYPu7rsd1lyty9UhatspAxiyMXa7wOYQZXEb8/QeVhhQ1WkSlEvfxVuLLLoda7UePkQYPVhxowUliHrWw0B1LZDn+qBgLSYjFg3ajAwHzpfkKt',
    'AWSALBCORS': 'f4ajqeI4YEX0tlTYPu7rsd1lyty9UhatspAxiyMXa7wOYQZXEb8/QeVhhQ1WkSlEvfxVuLLLoda7UePkQYPVhxowUliHrWw0B1LZDn+qBgLSYjFg3ajAwHzpfkKt',
    'DemoToken': 'CfDJ8BPT8Y-3o7FDkRiB4ECfYY_xoBNj6GBRHGhN4bqnuHq0RGpkwDhJjOAGVtT0D4pSE0DaT2ftNZyc_l5xIj77xqbIljTFQ5Q3klMSBX_fy_hWnCcM77Gq7ciQNrE2YyTbW6tHlvutmjpqMuWPq9b0hoPchD38HxwlHEFpSZSlJ_QbxzTGZXwfk42CJfCead372Q',
}

headers = {
    'accept': '*/*',
    'accept-language': 'ru-RU,ru;q=0.9',
    'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
    # 'cookie': 'FatSecret.AntiForgery=CfDJ8BPT8Y-3o7FDkRiB4ECfYY_Yh3GRWzocfapchO2HQwloaTlRqspYczsUTd0SEjkJCz30igyxp3DxE6qTXOF7iAtOGwVXezflZG37xgVXnSdRSHspZLqgRDShco7Tmny1Yke7ZNlpmE3iRnjF84G7fXY; AWSALB=f4ajqeI4YEX0tlTYPu7rsd1lyty9UhatspAxiyMXa7wOYQZXEb8/QeVhhQ1WkSlEvfxVuLLLoda7UePkQYPVhxowUliHrWw0B1LZDn+qBgLSYjFg3ajAwHzpfkKt; AWSALBCORS=f4ajqeI4YEX0tlTYPu7rsd1lyty9UhatspAxiyMXa7wOYQZXEb8/QeVhhQ1WkSlEvfxVuLLLoda7UePkQYPVhxowUliHrWw0B1LZDn+qBgLSYjFg3ajAwHzpfkKt; DemoToken=CfDJ8BPT8Y-3o7FDkRiB4ECfYY_xoBNj6GBRHGhN4bqnuHq0RGpkwDhJjOAGVtT0D4pSE0DaT2ftNZyc_l5xIj77xqbIljTFQ5Q3klMSBX_fy_hWnCcM77Gq7ciQNrE2YyTbW6tHlvutmjpqMuWPq9b0hoPchD38HxwlHEFpSZSlJ_QbxzTGZXwfk42CJfCead372Q',
    'origin': 'https://platform.fatsecret.com',
    'priority': 'u=1, i',
    'referer': 'https://platform.fatsecret.com/api-demo',
    'sec-ch-ua': '"Brave";v="131", "Chromium";v="131", "Not_A Brand";v="24"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-platform': '"Windows"',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'same-origin',
    'sec-gpc': '1',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36',
    'x-requested-with': 'XMLHttpRequest',
}

data = {
    'MarketLocale': 'RU',
    'LanguageLocale': 'ru',
    'SearchTerm': 'вкусно и ',
    #'Token': 'CfDJ8BPT8Y-3o7FDkRiB4ECfYY_xoBNj6GBRHGhN4bqnuHq0RGpkwDhJjOAGVtT0D4pSE0DaT2ftNZyc_l5xIj77xqbIljTFQ5Q3klMSBX_fy_hWnCcM77Gq7ciQNrE2YyTbW6tHlvutmjpqMuWPq9b0hoPchD38HxwlHEFpSZSlJ_QbxzTGZXwfk42CJfCead372Q',
    #'FatSecret.AntiForgery': 'CfDJ8BPT8Y-3o7FDkRiB4ECfYY_vL0Pd2sSR1IDgjU9SWL4VDem1kuSwpofCkVfKvmdWVM2uL8A3nFsuNNKiV3Zw4rklDt7ryaoWDSXRvhw446EffAvil0WWwj2SfPi2T5oDWRa_Q1BLgWwgjtm8RzdSmUI',
}

response = requests.post('https://platform.fatsecret.com/api-demo/foods-search', cookies=cookies, headers=headers, data=data)

print(response)
print(response.text)

with open('a', 'w') as file:
    file.write(response.text)