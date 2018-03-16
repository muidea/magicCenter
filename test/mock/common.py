'common'

import random

RAW_CONTENT = 'abcdefghijklmnopqrstuvwxyz'

# 随机生成一个词
def word():
    'word'
    ret = ""
    return ret.join(random.sample(RAW_CONTENT, random.randint(5, 12)))

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

if __name__ == '__main__':
    print(word())
    print(name())
    print(sentence())
    print(title())
    print(paragraph())
    print(content())
