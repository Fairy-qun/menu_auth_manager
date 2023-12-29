import { createRouter, createWebHashHistory } from 'vue-router'

// 默认路由，所有用户共享
const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../view/Login.vue'),
    meta: {
      name: '登录'
    }
  },
  {
    path: '/sys/home',
    name: 'index',
    component: () => import('../view/Index.vue'),
    meta: {
      name: '首页'
    }
  }
]

// 动态路由加载
const asyncRoutes = [
  {
    path: '/sys/home',
    name: '/sys/home',
    component: () => import('../view/Home.vue'),
    meta: {
      name: '主控台'
    }
  },
  {
    path: '/sys/user',
    name: '/sys/user',
    component: () => import('../view/User.vue'),
    meta: {
      name: '用户管理'
    }
  },
  {
    path: '/sys/role',
    name: '/sys/role',
    component: () => import('../view/Role.vue'),
    meta: {
      name: '角色管理'
    }
  },
  {
    path: '/sys/menu',
    name: '/sys/menu',
    component: () => import('../view/Menu.vue'),
    meta: {
      name: '菜单管理'
    }
  },
  {
    path: '/site/golang',
    name: '/site/golang',
    component: () => import('../view/Gsite.vue'),
    meta: {
      name: 'golang官网'
    }
  },
  {
    path: '/site/php',
    name: '/site/php',
    component: () => import('../view/Psite.vue'),
    meta: {
      name: 'php官网'
    }
  },
  {
    path: '/about/me',
    name: '/about/me',
    component: () => import('../view/About.vue'),
    meta: {
      name: '关于我们'
    }
  }
]

export const router = createRouter({
  routes,
  history: createWebHashHistory()
})

// 动态添加路由方法
const asyncAddRoute = menus => {
  // 判断是否有新的路由
  let hasNewRoutes = false
  // 根据传进来的菜单进行遍历，添加路由
  const findAndAddRoutesByMenus = arr => {
    arr.forEach(item => {
      let items = asyncRoutes.find(o => o.path == item.path)
      if (items && !router.hasRoute(items.path)) {
        router.addRoute('index', items)
        hasNewRoutes = true
      }
      if (item.children && item.children.length > 0) {
        findAndAddRoutesByMenus(item.children)
      }
    })
  }
  findAndAddRoutesByMenus(menus)
  return hasNewRoutes
}

// 前置路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (!token && to.path !== '/login') {
    router.push('/login')
  } else if (token) {
    const menus = JSON.parse(localStorage.getItem('userInfo')).menus
    const hasNewRouter = asyncAddRoute(menus)
    hasNewRouter ? next(to.path) : next()
  } else {
    next()
  }
  const title = to.meta.name
  document.title = title + '-菜单权限管理'
})
