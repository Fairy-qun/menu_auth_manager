<template>
  <div class="main">
    <div class="aside-menu">
      <el-menu
        :default-active="defaultActive"
        :collapse-transition="false"
        :unique-opened="true"
        @select="handleSelect"
      >
        <template v-for="(item, index) in asideMenu" :key="index">
          <el-sub-menu
            v-if="item.children && item.children.length > 0"
            :index="item.name"
          >
            <template #title>
              <el-icon>
                <component :is="item.icon"></component>
              </el-icon>
              <span>{{ item.name }}</span>
            </template>
            <el-menu-item
              v-for="(item2, index2) in item.children"
              ::key="index2"
              :index="item2.path"
            >
              <el-icon>
                <component :is="item2.icon"></component>
              </el-icon>
              <span>{{ item2.name }}</span>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item v-else :index="item.path">
            <el-icon>
              <component :is="item.icon"></component>
            </el-icon>
            <span>{{ item.name }}</span>
          </el-menu-item>
        </template>
      </el-menu>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useRouter, useRoute, onBeforeRouteUpdate } from 'vue-router'

const router = useRouter()
const route = useRoute()
// 获取菜单
const asideMenu = JSON.parse(localStorage.getItem('userInfo')).menus

// 绑定激活的index
const defaultActive = ref(route.path)
// 监听路由
onBeforeRouteUpdate((to, from) => {
  defaultActive.value = to.path
})
const handleSelect = e => {
  // e = e == '/' ? '/home' : e
  //   console.log(e)
  router.push(e)
}
</script>

<style lang="scss" scoped>
.main {
  display: flex;
  justify-content: space-between;
}
.content {
  margin-left: 200px;
  padding: 5px 20px;
}
.aside-menu {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  overflow-y: auto;
  overflow-x: hidden;
  transition: all 0.2s;
  background-color: rgba(255, 255, 255, 0.807);
  box-shadow: 2px 0px 20px #ccc;
  width: 200px;
  .el-menu {
    border: none;
  }
}
</style>
