'test.py'

from account import user
from account import group
from cache import cache
from cas import cas
from content import article
from content import catalog
from content import link
from content import media


if __name__ == '__main__':
    print("execute test script")
    cas.main()
    user.main()
    group.main()
    cache.main()
    article.main()
    catalog.main()
    link.main()
    media.main()
