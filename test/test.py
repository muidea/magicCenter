'test.py'

from account import user
from account import group
from cache import cache
from cas import login
from content import article
from content import catalog
from content import link
from content import media
from fileregistry import file_registry


if __name__ == '__main__':
    print("execute test script")
    login.main()
    user.main()
    group.main()
    cache.main()
    article.main()
    catalog.main()
    link.main()
    media.main()
    file_registry.main()
