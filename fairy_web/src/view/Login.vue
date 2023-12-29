<template>
  <div class="container">
    <!-- 标题 -->
    <div class="title">欢迎回来</div>
    <!-- 表单 -->
    <el-form :model="userInfo" label-width="60px">
      <el-form-item label="用户名">
        <el-input v-model="userInfo.username" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="userInfo.password" />
      </el-form-item>
    </el-form>
    <!-- 登录按钮 -->
    <el-button
      type="primary"
      size="small"
      style="width: 100%"
      @click="loginHandler"
      >登录</el-button
    >
    <!-- 注册按钮 -->
    <div class="register">还没有账号吗?<span>去注册</span></div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()
const userInfo = reactive({
  username: '',
  password: ''
})

const loginHandler = () => {
  if (userInfo.username === '' || userInfo.password === '') {
    alert('用户名及密码不能为空')
    return
  }
  console.log(userInfo)
  axios.post('http://127.0.0.1:5001/login', userInfo).then(res => {
    console.log(res)
    if (res.data.code === 200) {
      // 存储用户菜单
      localStorage.setItem('userInfo', JSON.stringify(res.data.data.userInfo))
      localStorage.setItem('token', res.data.data.token)
      router.replace('/sys/home')
    } else {
      alert(res.data.msg)
    }
  })
}
</script>

<style lang="scss" scoped>
.container {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 5px 10px;
  width: 400px;
  height: 250px;
  background-color: pink;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  border-radius: 5px;

  .title {
    font-size: 20px;
    font-weight: 700;
  }

  .register {
    position: absolute;
    bottom: 0;
    right: 20px;
    font-size: 14px;
    span {
      font-size: 16px;
      color: '#ffbb00';
    }
  }
}
</style>
