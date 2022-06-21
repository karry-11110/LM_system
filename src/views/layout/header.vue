<template>
    <el-row type="flex" justify="center" align="middle">
      <el-col :span="8">
        <el-row>
          <el-col :span="8">
           <div class="ginx_title"><a href="/" id="home-link" title="ginxcast">ginxcast</a></div>
          </el-col>
          <el-col class="" :span="8" :offset="8">
            <div class="header_name"> <router-link class="to-link" :to="{path: '/create'}"><a title="ginxcast">我要提问</a></router-link> | wangkun | <a href="/user/logout" title="ginxcast">登出</a> </div>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
</template>

<script>
export default {
  methods: {
  }
}
</script>

<style scoped>
.ginx_title {
  float: left;
  vertical-align: top;
  font-size: 24px;
  font-weight: 400;
  margin: 0 15px 0 0;
  line-height: 34px;
  padding-top: 10px;
}

.header_name {
  float: right;
  font-size: 13px;
  font-weight: 400;
  margin: 0 15px 0 0;
  line-height: 34px;
  padding-top: 10px;
  padding-top: 10px;
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}
.header_name > a {
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}

.ginx_title > a {
  background-color: transparent;
  color: #486e9b;
  text-decoration: none;
}
.el-header {
  background-color: #409EFF;
  text-align: center;
  line-height: 80px;
}

</style>
