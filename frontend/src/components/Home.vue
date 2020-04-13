<template>
  <el-container class="home-container">
    <!-- 头部区域 -->
    <el-header>
      <div>
        <span>电商后台管理系统</span>
      </div>
      <el-button type="info" @click="logout">退出</el-button>
    </el-header>
    <!-- 页面主体区域 -->
    <el-container>
      <el-aside :width="asideSize">
        <!-- 侧边栏 -->
        <div class="toggle-button" @click="toggle">|||</div>
        <el-menu
          :unique-opened="true"
          background-color="#333744"
          text-color="#fff"
          active-text-color="#409eff"
          :collapse="isCollapse"
          :collapse-transition="false"
          :router="true"
          :default-active="highlight_path"
        >
          <el-submenu :index="item.id.toString()" v-for="item in menulist" :key="item.id">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>{{item.authName}}</span>
            </template>
            <el-menu-item
              :index="'/'+child.path"
              v-for="child in item.children"
              :key="child.id"
              @click="highlight('/'+child.path)"
            >
              <i class="el-icon-menu"></i>
              <span>{{child.authName}}</span>
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <!-- 右侧主体内容 -->
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      menulist: [],
      isCollapse: false,
      asideSize: '200px',
      highlight_path: ''
    }
  },
  created() {
    this.getMenuList()
    this.highlight_path = window.sessionStorage.getItem('highlight')
  },
  methods: {
    toggle() {
      this.isCollapse = !this.isCollapse
      if (this.isCollapse) {
        this.asideSize = '64px'
      } else {
        this.asideSize = '200px'
      }
    },
    logout() {
      window.sessionStorage.clear()
      this.$router.push('/login')
    },
    // 获取所有的菜单
    async getMenuList() {
      const { data: src } = await this.$http.get('menus')
      if (src.meta.status !== 200) return this.$message.error(src.meta.msg)
      this.menulist = src.data
      console.log(src)
    },
    highlight(value) {
      // console.log('hi')
      window.sessionStorage.setItem('highlight', value)
      this.highlight_path = value
    }
  }
}
</script>

<style lang= "less" scoped>
.el-header {
  background-color: #373d41;
  display: flex;
  justify-content: space-between;
  padding-left: 0px;
  align-items: center;
  color: white;
  font-size: 20px;
}
.el-aside {
  background-color: #333744;
}
el-main {
  background-color: #eaedf1;
}
.home-container {
  height: 100%;
}
.iconfont {
  margin-right: 10px;
}
.el-menu {
  border-right: 0px;
}
.toggle-button {
  background-color: #4a5064;
  font-size: 10px;
  line-height: 24px;
  color: white;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
