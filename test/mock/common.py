import random
import array

rawContent = 'abcdefghijklmnopqrstuvwxyz'

# 随机生成一个词
def word():
    ret = ""
    return ret.join(random.sample(rawContent, random.randint(5,12)))

# 随机生成一个名字，首字母大写
def name():
    return word().title()

# 随机生成一个句子
def sentence():
    valArray = []
    count = random.randint(5,13)
    ii = 0
    while(ii < count):
        valArray.append(word())
        ii = ii + 1

    ret = " "
    return ret.join(valArray) + ". "

# 随机生成一个标题
def title():
    valArray = []
    count = random.randint(5,10)
    ii = 0
    while(ii < count):
        valArray.append(word())
        ii = ii + 1

    ret = " "
    return ret.join(valArray)

# 生成一段话
def paragraph():
    valArray = []
    count = random.randint(5,10)
    ii = 0
    while(ii < count):
        valArray.append(sentence())
        ii = ii + 1

    ret = ""
    return ret.join(valArray)

def content():
    valArray = []
    count = random.randint(3,60)
    ii = 0
    while(ii < count):
        valArray.append(paragraph())
        ii = ii + 1

    ret = "\n"
    return ret.join(valArray)    

if __name__ == '__main__':
    print(word())
    print(name())
    print(sentence())
    print(title())
    print(paragraph())
    print(content())