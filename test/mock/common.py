'common'

import random

RAW_CONTENT = 'abcdefghijklmnopqrstuvwxyz'

# 随机生成一个词
def word():
    'word'
    ret = ""
    return ret.join(random.sample(RAW_CONTENT, random.randint(3, 12)))

# 随机生成一个名字，首字母大写
def name():
    'name'
    return word().title()

# 随机生成一个句子
def sentence():
    'sentence'
    val_array = []
    count = random.randint(5, 13)
    index = 0
    while index < count:
        val_array.append(word())
        index = index + 1

    ret = " "
    return ret.join(val_array) + ". "

# 随机生成一个标题
def title():
    'title'
    val_array = []
    count = random.randint(5, 10)
    index = 0
    while index < count:
        val_array.append(word())
        index = index + 1

    ret = " "
    return ret.join(val_array)

# 生成一段话
def paragraph():
    'paragraph'
    val_array = []
    count = random.randint(5, 10)
    index = 0
    while index < count:
        val_array.append(sentence())
        index = index + 1

    ret = ""
    return ret.join(val_array)

def content():
    'content'
    val_array = []
    count = random.randint(3, 60)
    index = 0
    while index < count:
        val_array.append(paragraph())
        index = index + 1

    ret = "\n"
    return ret.join(val_array)

def email():
    'email'
    url_array = []
    index = 0
    while index < 2:
        url_array.append(word())
        index = index + 1
    domain = '.'
    return '%s@%s'%(word(), domain.join(url_array))

def picker_list(data_list, num):
    'picker_list'
    ret = []
    data_len = len(data_list)
    if data_len <= num:
        return data_list.copy()

    tmp_list = data_list[:]
    while True:
        if len(ret) >= num:
            break
        offset = random.randint(0, len(tmp_list)-1)
        ret.append(tmp_list[offset])
        del tmp_list[offset]
    return ret.copy()

def picker_dict(data_dict, num):
    'picker_dict'
    ret = {}
    data_len = len(data_dict)
    if data_len <= num:
        return data_dict.copy()

    tmp_dict = data_dict.copy()
    while True:
        if len(ret) >= num:
            break
        offset = random.randint(0, len(tmp_dict)-1)
        for key in tmp_dict.keys():
            if offset == 0:
                ret[key] = tmp_dict[key]
                tmp_dict.pop(key)
                break
            else:
                offset = offset -1

    return ret

if __name__ == '__main__':
    print(word())
    print(name())
    print(sentence())
    print(title())
    print(paragraph())
    print(content())
    print(email())
