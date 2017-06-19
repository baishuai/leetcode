# --coding: utf-8 ---
# wget https://leetcode.com/api/problems/algorithms/
# 需要 cookies， 直接通过浏览器下载吧
import json
import os
import shutil

# url = "https://leetcode.com/api/problems/algorithms/"

SUFFIXES = {'.go': 'Go', '.py': 'Python', '.java': 'Java'}


def type_and_file(idx):
    '''
    return lang type and file path of problems idx
    if find nothing, return None
    '''
    res = {}
    subdir_name = 'p' + str(idx).zfill(3)
    print('find idx', './algorithms/' + subdir_name)
    path = './algorithms/' + subdir_name
    if os.path.exists(path) and os.path.isdir(path):
        for file in os.listdir(path):
            if 'test' in file:
                continue
            if os.path.isfile(os.path.join(path, file)):
                name, extension = os.path.splitext(file)
                print(name, extension)
                if extension in SUFFIXES:
                    res[SUFFIXES[extension]] = os.path.join(path, file)
    return res


def parse(filename, out):
    '''
    parse items in filename, and write to out
    '''
    levels = {1: "Easy", 2: "Medium", 3: "Hard"}
    with open(filename) as file:
        j = json.load(file)

        out.write(
            "完成进度 {0}/{1}\n\n".format(j['num_solved'], j['num_total']))

        out.write("| # | Title | Solution | Difficulty | Finshed |\n")
        out.write("|---| ----- | -------- | ---------- | ---- |\n")
        for ssp in j['stat_status_pairs']:
            status = ssp['status'] == 'ac'
            idx = int(ssp['stat']['question_id'])
            title = ssp['stat']['question__title']
            level = levels[ssp['difficulty']['level']]
            slug = ssp['stat']['question__title_slug']

            if status:
                out.write(
                    "|{0} |[{1}](https://leetcode.com/problems/{2})|".format(idx, title, slug))
                find = ' 404 '
                for lang, path in type_and_file(idx).items():
                    out.write("[{0}]({1})  ".format(lang, path))
                    find = ':heavy_check_mark:'
                out.write("|{0} |{1} |\n".format(level, find))
            else:
                out.write("|{0} |[{1}](https://leetcode.com/problems/{2})| | {3} | |\n".format(
                    idx, title, slug, level))


def main():
    '''main function'''
    filename = "leetcode.com.json"
    shutil.copyfile("README_HEAD.md", "README.md")
    with open("README.md", 'a') as out:
        parse(filename, out)


if __name__ == '__main__':
    main()
