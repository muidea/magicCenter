
from account import User
from account import Group
from cache import Cache
from cas import Login
from content import Article
from content import Catalog
from content import Link
from content import Media
from fileregistry import FileRegistry


if __name__ == '__main__':
    print("execute test script")
    Login.main()
    User.main()
    Group.main()
    Cache.main()
    Article.main()
    Catalog.main()
    Link.main()
    Media.main()
    FileRegistry.main()
    
