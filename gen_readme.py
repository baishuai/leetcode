# --coding: utf-8 ---
# wget https://leetcode.com/api/problems/algorithms/
# 需要 cookies， 直接通过浏览器下载吧
import json
import os
import shutil

url = "https://leetcode.com/api/problems/algorithms/"

levels = {1: "Easy", 2: "Medium", 3: "Hard"}

py_ans = [297, 278, 202, 191, 160, 151, 138, 117, 116, 190]

filename = "leetcode.com.html"

shutil.copyfile("README_HEAD.md", "README.md")

with open(filename) as f:
    with open("README.md", 'a') as out:
        j = json.load(f)

        out.write("完成进度 {0}/{1}\n\n".format(j['num_solved'], j['num_total']))

        out.write("| # | Title | Solution | Difficulty | Finshed |\n")
        out.write("|---| ----- | -------- | ---------- | ---- |\n")
        for ssp in j['stat_status_pairs']:
            status = ':heavy_check_mark:' if ssp['status'] == 'ac' else ''
            idx = int(ssp['stat']['question_id'])
            title = ssp['stat']['question__title']
            level = levels[ssp['difficulty']['level']]
            slug = ssp['stat']['question__title_slug']

            lang = "Go"
            suffix = "go"
            if idx in py_ans:
                lang = "Python"
                suffix = "py"

            if status:
                out.write("|{0} |[{1}](https://leetcode.com/problems/{2})|[{3}](./algorithms/p{4}/{4}.{5}) | {6} | {7} |\n".format(
                    idx, title, slug, lang, str(idx).zfill(3), suffix, level, status))
            else:
                out.write("|{0} |[{1}](https://leetcode.com/problems/{2})|  | {3} |  |\n".format(
                    idx, title, slug, level))
