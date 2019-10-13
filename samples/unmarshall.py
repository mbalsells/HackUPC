import re
import os

res = open("regex.txt", "r").read();
txt = open("drama.txt", "r").read();

print(len(re.findall(res, txt)))
